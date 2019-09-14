package cloudserversext

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/resource"
)

const ECSCLOUDSERVICETYPECODE = "hws.service.type.ec2"

func GetServerExt(client *gophercloud.ServiceClient, serverId string) (CloudServerExt, error) {

	//1, get ecs v1 data
	cloudServerResp, err := getServerInfo(client, serverId)
	if err != nil {
		return CloudServerExt{}, err
	}
	bssClient, err := openstack.NewBSSV1(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return CloudServerExt{}, err
	}

	volumeClient, err := openstack.NewBlockStorageV2(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return CloudServerExt{}, err
	}

	//2,获取相关bss信息
	resources, err := getChargingInfo(bssClient, serverId, "")
	if err != nil {
		return CloudServerExt{}, err
	}

	//3,获取相关volumes信息
	volumeInfo, err := getVolumeInfo(volumeClient, cloudServerResp)
	if err != nil {
		return CloudServerExt{}, err
	}

	//4,组装返回
	return CloudServerExt{
		CloudServer:    cloudServerResp,
		VolumeAttached: volumeInfo,
		Charging:       fmtChargingInfo(resources),
	}, nil

}

func GetPrepaidServerDetailByOrderId(client *gophercloud.ServiceClient, orderId string) ([]CloudServerExt, error) {
	var cloudServerExt []CloudServerExt

	//init bss client
	bssClient, err := openstack.NewBSSV1(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return cloudServerExt, err
	}

	//init volume client
	volumeClient, err := openstack.NewBlockStorageV2(client.ProviderClient, gophercloud.EndpointOpts{})
	if err != nil {
		return cloudServerExt, err
	}

	//get bss info
	resources, err := getChargingInfo(bssClient, "", orderId)
	if err != nil {
		return cloudServerExt, err
	}

	//get server info
	for _, v := range resources.Data {
		//遍历 cloud_service_type_code 类型为"hws.service.type.ec2"。
		if v.CloudServiceTypeCode == ECSCLOUDSERVICETYPECODE {
			cloudServerResp, err := getServerInfo(client, v.ResourceId)
			if err != nil {
				return cloudServerExt, err
			}

			vaExt, err := getVolumeInfo(volumeClient, cloudServerResp)
			if err != nil {
				return cloudServerExt, err
			}

			cloudServerExt = append(cloudServerExt, CloudServerExt{
				CloudServer:    cloudServerResp,
				VolumeAttached: vaExt,
				Charging:       fmtChargingInfo(resources),
			})
		}
	}
	return cloudServerExt, nil
}

func getServerInfo(client *gophercloud.ServiceClient, serverId string) (*cloudservers.CloudServer, error) {
	return cloudservers.Get(client, serverId).Extract()
}

func getChargingInfo(client *gophercloud.ServiceClient, serverId, orderId string) (resource.Resources, error) {
	var resources resource.Resources

	domainID := client.ProviderClient.DomainID
	resOpts := resource.ListOpts{
		CustomerId: domainID,
	}
	if serverId != "" {
		resOpts.ResourceIds = serverId
	}
	if orderId != "" {
		resOpts.OrderId = orderId
	}
	//bss接口只能查询到包周期的资源,查询错误时特殊处理,不返回
	resources, err := resource.ListDetail(client, resOpts).Extract()
	if err != nil {
		return resources, err
	}
	//此处为适配api bug代码，bss接口失败也只会返回http 200.
	//查询错误时特殊处理,不返回
	if resources.ErrorCode == "" || resources.ErrorCode != "CBC.0000" {
		return resources, gophercloud.NewSystemCommonError(resources.ErrorCode, resources.ErrorMsg)
	}
	return resources, nil
}

func fmtChargingInfo(res resource.Resources) (Charging) {
	var charging Charging
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
	return charging
}

func getVolumeInfo(client *gophercloud.ServiceClient, cloudServerResp *cloudservers.CloudServer) ([]VolumeInfo, error) {
	vaExt := make([]VolumeInfo, 0)

	if len(cloudServerResp.VolumeAttached) > 0 {
		for _, va := range cloudServerResp.VolumeAttached {
			volume, err := volumes.Get(client, va.ID).Extract()
			if err != nil { //错误时特殊处理，不返回。
				vaExt = append(vaExt, VolumeInfo{
					ID:         va.ID,
				})
				continue
			}
			vaExt = append(vaExt, VolumeInfo{
				ID:         va.ID,
				VolumeType: volume.VolumeType,
				Size:       volume.Size,
			})
		}
	}
	return vaExt, nil
}
