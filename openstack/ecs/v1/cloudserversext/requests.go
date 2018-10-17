package cloudserversext

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/resource"
)

func GetServerExt(client *gophercloud.ServiceClient, serverId string) (CloudServerExt, error) {

	//1, get ecs v1 data
	cloudeServerResp, err := cloudservers.Get(client, serverId).Extract()
	if err != nil {
		return CloudServerExt{}, err
	}

	//2,获取相关bss信息
	var charging Charging
	scBSS, err := openstack.NewBSSV1(client.ProviderClient, gophercloud.EndpointOpts{})

	if err != nil {
		return CloudServerExt{}, err
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
		return CloudServerExt{}, err
	}
	//此处为适配api bug代码，bss接口失败也只会返回http 200.
	//查询错误时特殊处理,不返回
	if res.ErrorCode == "" || res.ErrorCode != "CBC.0000" {
		return CloudServerExt{}, gophercloud.NewSystemCommonError(res.ErrorCode, res.ErrorMsg)
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

	//3,获取相关volumes信息
	scBS, err := openstack.NewBlockStorageV2(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return CloudServerExt{}, err
	}

	vaExt := make([]VolumeInfo, len(cloudeServerResp.VolumeAttached))

	if len(cloudeServerResp.VolumeAttached) > 0 {
		for i, va := range cloudeServerResp.VolumeAttached {
			volume, err := volumes.Get(scBS, va.ID).Extract()
			if err != nil { //错误时特殊处理，不返回。
				vaExt[i].ID = va.ID
				continue
			}
			vaExt[i].ID = va.ID
			vaExt[i].VolumeType = volume.VolumeType
			vaExt[i].Size = volume.Size
		}
	}

	//4,组装返回
	return CloudServerExt{
		CloudServer:    cloudeServerResp,
		VolumeAttached: vaExt,
		Charging:       charging,
	}, nil

}
