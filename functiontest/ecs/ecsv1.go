package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudserversext"
	"encoding/json"
)

func main() {
	fmt.Println("main start...")

	//provider, err := common.AuthToken()
	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}
	sc, err := openstack.NewECSV1(provider, gophercloud.EndpointOpts{})
	fmt.Println(sc.Token())

	provider.Reauthenticate("")
	//sc.ProviderClient.Reauthenticate("")
	TestGetEcsExtbyOrderId(sc)
	fmt.Println(sc.Token())

	if err != nil {
		fmt.Println("get ecs v1 client failed")
		fmt.Println(err.Error())
		return
	}
	//TestGetEcs(sc)
	//TestGetEcsExtbyServerId(sc)
	TestGetEcsExtbyOrderId(sc)
	fmt.Println("main end...")
}

func TestGetEcs(sc *gophercloud.ServiceClient) {
	//2c2cd6a9-c501-42a9-a679-53518e6757cc
	resp, err := cloudservers.Get(sc, "d26b697b-3a74-4ec2-bd9d-5c3829f5d8a5").Extract()
	if err != nil {
		fmt.Println(err)
	}
	b, errr := json.MarshalIndent(*resp, "", " ")

	if errr != nil {

		fmt.Println(errr)
	}
	fmt.Println(string(b))

}

func TestGetEcsExtbyServerId(sc *gophercloud.ServiceClient) {
	//2c2cd6a9-c501-42a9-a679-53518e6757cc
	//95b23c71-0016-4f80-b160-7c1e0341d205
	resp, err := cloudserversext.GetServerExt(sc, "2544b973-ba5b-4cbd-a060-771ba4ec73e2")
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println("CloudServer id is:", resp.CloudServer.ID)
	fmt.Println("CloudServer charging mode is:", resp.Charging.ChargingMode)
	volumeAttached, _ := json.MarshalIndent(resp.VolumeAttached, "", "    ")
	fmt.Println("CloudServer volume attached is:", string(volumeAttached))
}

func TestGetEcsExtbyOrderId(sc *gophercloud.ServiceClient) {
	resp, err := cloudserversext.GetServerExtbyOrderId(sc, "CS1811091456QYTEX")
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range resp {
		fmt.Println("CloudServer id is:", v.CloudServer.ID)
		fmt.Println("CloudServer charging mode is:", v.Charging.ChargingMode)
		volumeAttached, _ := json.MarshalIndent(v.VolumeAttached, "", "    ")
		fmt.Println("CloudServer volume attached is:", string(volumeAttached))
	}
}
