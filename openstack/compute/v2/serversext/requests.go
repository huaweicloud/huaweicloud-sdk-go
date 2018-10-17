package serversext

import (
	"math"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/resource"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootwithscheduler"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	//"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudserversext"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
	"fmt"
)

func CreateServer(client *gophercloud.ServiceClient, opts servers.CreateOptsBuilder) (r servers.CreateResult) {
	if client == nil {
		message := fmt.Sprintln(gophercloud.CE_NoClientProvidedMessage)
		err := gophercloud.NewSystemCommonError(gophercloud.CE_NoClientProvidedCode, message)
		r.Err = err
		return
	}
	if bsOpts, ok := opts.(bootwithscheduler.CreateOptsExt); ok {
		//检查networks资源
		err := checkNetworks(client, bsOpts)
		if err != nil {
			r.Err = err
			return
		}
		//检查volumes资源
		//		err = checkVolumes(client, bsOpts)
		//		if err != nil {
		//			r.Err = err
		//	        return
		//		}
	}

	r = servers.Create(client, opts)
	return
}

func ListServers(client *gophercloud.ServiceClient) (onDemandServers, monthlyServers []servers.ServerBrief, err error) {
	//step1:获取全量server,分页查询，每次查询500个。
	allPages, err := servers.ListBrief(client, servers.ListOpts{Limit: 500}).AllPages()
	if err != nil {
		return
	}

	allServerBriefs, err := servers.ExtractServerBriefs(allPages)
	if err != nil {
		return
	}

	//step2:通过bss接口获取全部monthly serverId
	scBSS, err := openstack.NewBSSV1(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return
	}

	domainID := client.ProviderClient.DomainID
	//多次查询，每次查300个
	resAll, err := listResourcesMultiTimes(scBSS, domainID, 300)
	if err != nil {
		return
	}

	//获取所有monthly serverId
	var monthlySvrIds []string
	for _, res := range resAll.Data {
		if "hws.resource.type.vm" == res.ResourceTypeCode {
			monthlySvrIds = append(monthlySvrIds, res.ResourceId)
		}
	}

	//step3:分离出monthly server 和 onDemand server
	for _, sb := range allServerBriefs {
		b := gophercloud.IsInStrSlice(monthlySvrIds, sb.ID)
		if b {
			monthlyServers = append(monthlyServers, sb)
		} else {
			onDemandServers = append(onDemandServers, sb)
		}
	}

	return
}

func listResourcesMultiTimes(sc *gophercloud.ServiceClient, domainId string, onceCnt int) (*resource.Resources, error) {
	var resAll resource.Resources

	optsTmp := resource.ListOpts{
		CustomerId: domainId,
		PageNo:     1,
		PageSize:   1,
	}

	resTmp, err := resource.ListDetail(sc, optsTmp).Extract()
	if err != nil {
		return &resAll, err
	}
	//此处为适配api bug代码，bss接口失败也只会返回http 200.
	if resTmp.ErrorCode == "" || resTmp.ErrorCode == "CBC.0999" {
		return &resAll, gophercloud.NewSystemCommonError("CBC.0999", resTmp.ErrorMsg)
	}

	resAll.ErrorCode = resTmp.ErrorCode
	resAll.ErrorMsg = resTmp.ErrorMsg
	resAll.TotalCount = resTmp.TotalCount

	//一次查10条
	totalCnt := resTmp.TotalCount
	queryTimes := int(math.Ceil(float64(totalCnt) / float64(onceCnt)))
	//lastCnt := totalCnt - (queryTimes-1)*onceCnt
	//fmt.Println("queryTimes:", queryTimes)

	opts := resource.ListOpts{
		CustomerId: domainId,
		PageNo:     1,
		PageSize:   onceCnt,
	}

	for i := 1; i <= queryTimes; i++ {
		opts.PageNo = i
		//最后一次要设置为实际次数
		//if i == queryTimes {
		//	opts.PageSize = lastCnt
		//}

		res, err := resource.ListDetail(sc, opts).Extract()
		if err != nil {
			return &resAll, err
		}
		//此处为适配api bug代码，bss接口失败也只会返回http 200.
		if res.ErrorCode == "" || res.ErrorCode == "CBC.0999" {
			return &resAll, gophercloud.NewSystemCommonError("CBC.0999", res.ErrorMsg)
		}

		resAll.Data = append(resAll.Data, res.Data...)
	}

	//fmt.Println("resAll.Data:", len(resAll.Data))
	return &resAll, nil
}

func Get(client *gophercloud.ServiceClient, serverId string) (ServerExt, error) {
	//获取基础server信息
	server, err := servers.Get(client, serverId).Extract()
	if err != nil {
		return ServerExt{}, err
	}

	//通过 ecs v1 接口 获取 server信息包括 vpc id
	ecsv1, err := openstack.NewECSV1(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return ServerExt{}, err
	}
	cloudeServerResp, err := cloudservers.Get(ecsv1, serverId).Extract()

	if err != nil {
		return ServerExt{}, err
	}

	//获取相关bss信息
	var charging Charging
	scBSS, err := openstack.NewBSSV1(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return ServerExt{}, err
	}

	domainID := client.ProviderClient.DomainID
	resOpts := resource.ListOpts{
		CustomerId:  domainID,
		ResourceIds: serverId,
	}

	//bss接口只能查询到包周期的资源,查询错误时特殊处理,不返回
	res, err := resource.ListDetail(scBSS, resOpts).Extract()
	if err != nil {
		//fmt.Println("Bss2")
		return ServerExt{}, err
	}
	//此处为适配api bug代码，bss接口失败也只会返回http 200.
	//查询错误时特殊处理,不返回
	if res.ErrorCode == "" || res.ErrorCode != "CBC.0000" {
		return ServerExt{}, gophercloud.NewSystemCommonError(res.ErrorCode, res.ErrorMsg)
	}

	data := res.Data
	//如果data没数据，则表示为按需server（非包周期server）
	if len(data) == 0 && res.ErrorCode == "CBC.0000" {
		charging.ChargingMode = "0"
		charging.ValidTime = ""
		charging.ExpireTime = ""
	}

	//如果data有数据，则表示为包周期server
	if len(data) > 0 {
		//不返回时间则是按需计费,返回时间则是包周期
		if data[0].ValidTime != "" && data[0].ExpireTime != "" {
			charging.ChargingMode = "1"
			charging.ValidTime = data[0].ValidTime
			charging.ExpireTime = data[0].ExpireTime
		}
	}

	//获取相关volumes信息
	scBS, err := openstack.NewBlockStorageV2(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return ServerExt{}, err
	}

	vaExt := make([]VolumeInfo, len(server.VolumeAttached))

	if len(server.VolumeAttached) > 0 {
		for i, va := range server.VolumeAttached {
			if len(va) > 0 {
				if id, ok := va["id"].(string); ok {
					volume, err := volumes.Get(scBS, id).Extract()
					if err != nil { //错误时特殊处理，不返回。
						vaExt[i].ID = id
						continue
					}

					vaExt[i].ID = id
					vaExt[i].VolumeType = volume.VolumeType
					vaExt[i].Size = volume.Size
				}
			}
		}
	}

	//vaExt := make([]VolumeInfo, len(cloudeServerResp.VolumeAttached))
	//
	//if len(cloudeServerResp.VolumeAttached) > 0 {
	//	for i, va := range cloudeServerResp.VolumeAttached {
	//		volume, err := volumes.Get(scBS, va.ID).Extract()
	//		if err != nil { //错误时特殊处理，不返回。
	//			vaExt[i].ID = va.ID
	//			continue
	//		}
	//		vaExt[i].ID = va.ID
	//		vaExt[i].VolumeType = volume.VolumeType
	//		vaExt[i].Size = volume.Size
	//	}
	//}
	//server := cloudeServerResp.CloudServer
	//
	////组装返回
	//	return &ServerExt{
	//	ID:             server.ID,
	//	TenantID:       server.TenantID,
	//	UserID:         server.UserID,
	//	Name:           server.Name,
	//	Status:         server.Status,
	//	Updated:        server.Updated,
	//	Created:        server.Created,
	//	Addresses:      server.Addresses,
	//	KeyName:        server.KeyName,
	//	TaskState:      server.TaskState,
	//	VMstate:        server.VMState,
	//	SecurityGroups: server.SecurityGroups,
	//	PowerState:     *server.PowerState,
	//	Metadata:       server.Metadata,
	//	Flavor:         server.Flavor,
	//	//Image:          server.Image,
	//	//Links:          server.Links,
	//	AvailbiltyZone: server.AvailabilityZone,
	//	VolumeAttached: vaExt,
	//	Charging:       charging,
	//	VpcId:          cloudeServerResp.CloudServer.Metadata.VpcID,
	//	ImageId : cloudeServerResp.CloudServer.Metadata.ImageID,
	//}, nil

	return ServerExt{
		ID:             server.ID,
		TenantID:       server.TenantID,
		UserID:         server.UserID,
		Name:           server.Name,
		Status:         server.Status,
		Updated:        server.Updated,
		Created:        server.Created,
		Addresses:      server.Addresses,
		KeyName:        server.KeyName,
		TaskState:      server.TaskState,
		VMstate:        server.VMstate,
		SecurityGroups: server.SecurityGroups,
		PowerState:     server.PowerState,
		Metadata:       server.Metadata,
		Flavor:         server.Flavor,
		Image:          server.Image,
		Links:          server.Links,
		AvailbiltyZone: server.AvailbiltyZone,
		VolumeAttached: vaExt,
		Charging:       charging,
		VpcId:          cloudeServerResp.Metadata.VpcID,
		ImageId:        cloudeServerResp.Metadata.ImageID,
	}, nil
}

//检测网络资源
func checkNetworks(client *gophercloud.ServiceClient, opts bootwithscheduler.CreateOptsExt) error {
	//创建network client
	networkClient, err := openstack.NewNetworkV2(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return err
	}
	base, err := opts.CreateOptsBuilder.ToServerCreateMap()
	if err != nil {
		return err
	}
	serverMap := base["server"].(map[string]interface{})
	nws := serverMap["networks"].([]map[string]interface{})

	//整理过的network,uuid和其数量
	sortedNetworks := make(map[string]int)
	for _, n := range nws {
		if v, ok := n["uuid"]; ok {
			uuid, _ := v.(string)
			if uuid != "" {
				if _, ok := sortedNetworks[uuid]; ok {
					sortedNetworks[uuid]++
				} else {
					sortedNetworks[uuid] = 1
				}
			}
		}
	}

	for uuid, needIps := range sortedNetworks {
		ia, err := networks.GetNetworkIpAvailabilities(networkClient, uuid)
		if err != nil {
			return err
		}

		availIps := ia.NetworkIpAvail.TotalIps - ia.NetworkIpAvail.UsedIps
		//fmt.Println("availIps:", availIps)
		//fmt.Println("needIps:", needIps)
		if availIps < needIps {
			return &gophercloud.UnifiedError{
				ErrCode:    "Ecs.1525",
				ErrMessage: "Insufficient IP addresses on the network.",
			}
		}

		continue
	}

	return nil
}

//检测存储资源
func checkVolumes(client *gophercloud.ServiceClient, opts bootwithscheduler.CreateOptsExt) error {
	//创建blockstorage client
	bsClient, err := openstack.NewBlockStorageV2(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return err
	}

	//计算需要的资源
	vNeed := 0
	gNeed := 0
	for _, bd := range opts.BlockDevice {
		if bd.SourceType != "volume" && bd.VolumeSize != 0 {
			vNeed++
			gNeed += bd.VolumeSize
		}
	}

	//计算现有的资源
	projectID := client.ProviderClient.GetProjectID()
	qs, err := volumes.GetQuotaSet(bsClient, projectID)
	if err != nil {
		return err
	}

	var vLimit, vInUse, gLimit, gInUse int
	if _, ok := qs.QuoSet.Volumes["limit"]; ok {
		vLimit = qs.QuoSet.Volumes["limit"]
	}
	if _, ok := qs.QuoSet.Volumes["in_use"]; ok {
		vInUse = qs.QuoSet.Volumes["in_use"]
	}
	if _, ok := qs.QuoSet.Gigabytes["limit"]; ok {
		gLimit = qs.QuoSet.Gigabytes["limit"]
	}
	if _, ok := qs.QuoSet.Gigabytes["in_use"]; ok {
		gInUse = qs.QuoSet.Gigabytes["in_use"]
	}

	//	fmt.Println("vNeed:", vNeed)
	//	fmt.Println("gNeed:", gNeed)
	//	fmt.Println("vLimit:", vLimit)
	//	fmt.Println("vInUse:", vInUse)
	//	fmt.Println("gLimit:", gLimit)
	//	fmt.Println("gInUse:", gInUse)

	//计算是否满足需要
	ue := &gophercloud.UnifiedError{
		ErrCode:    "Ecs.1524",
		ErrMessage: "Instance resource is temporarily sold out.",
	}

	if (vLimit != -1) && ((vLimit - vInUse) < vNeed) {
		return ue
	}

	if (gLimit != -1) && ((gLimit - gInUse) < gNeed) {
		return ue
	}

	return nil
}

/*
//判断资源是否已满，创建server前判断。
func checkHostAvailable(client *gophercloud.ServiceClient) (bool, error) {

	//查询所有status为error的servers
	listOpts := servers.ListOpts{
		Status: "ERROR",
	}

	allPages, err := servers.List(client, listOpts).AllPages()
	if nil != err {
		return true, err
	}

	allServers, err := servers.ExtractServers(allPages)
	if nil != err {
		return true, err
	}

	if 0 == len(allServers) {
		return true, err
	}

	for _, s := range allServers {
		b := strings.Contains(s.Fault.Message, "There are not enough hosts available")
		if b {
			return false, err
		}
	}

	return true, err
}
*/
