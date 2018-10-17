package cloudserversext

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootwithscheduler"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudservers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
)

func CreateServer(client *gophercloud.ServiceClient, opts cloudservers.CreateOptsBuilder) (jobId, orderId string, err error) {
	if bsOpts, ok := opts.(bootwithscheduler.CreateOptsExt); ok {
		//检查networks资源
		err = checkNetworks(client, bsOpts)
		if err != nil {
			return "", "", err
		}
	}

	return cloudservers.Create(client, opts)
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
	//fmt.Println(networks)

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

	//fmt.Println("sortedNetworks:", sortedNetworks)

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
