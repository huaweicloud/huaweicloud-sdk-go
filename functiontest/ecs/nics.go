package main

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/nics"
)

func main() {
	fmt.Println("main start...")

	provider, err := common.AuthToken()
	//provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}
	sc, err := openstack.NewECSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ecs v1 client failed")
		fmt.Println(err.Error())
		return
	}

	TestAddNics(sc)
	TestDelNics(sc)
	TestBindNic(sc)
	TestUnBindNic(sc)

	fmt.Println("main end...")
}

func TestAddNics(sc *gophercloud.ServiceClient) {
	opts := nics.AddOpts{
		Nics: []nics.Nic{
			{
				SubnetId:  "9a56640e-5503-4b8d-8231-963fc59ff91c",
				IpAddress: "",
				SecurityGroups: []nics.SecurityGroup{
					{
						ID: "2579b3f4-c6c8-4c37-822d-c7b1fbe4c9f6",
					},
				},
			},
		},
	}

	jobId, err := nics.AddNics(sc, "1696c5a8-d432-45be-b419-e11cac3c30c6", opts)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test addnics success,jobId is: ", jobId)
}

func TestDelNics(sc *gophercloud.ServiceClient) {
	opts := nics.DelOpts{
		Nics: []nics.Nics{
			{
				ID: "a2de1b18-eb76-45af-a36b-fb76f8dc9180",
			},
		},
	}

	jobId, err := nics.DeleteNics(sc, "1cdd5621-93b7-4c49-be6a-500229c196f2", opts)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test DelNics success,jobId is: ", jobId)
}

func TestBindNic(sc *gophercloud.ServiceClient) {
	reversebinding := true
	opts := nics.BindOpts{
		SubnetId:       "9a56640e-5503-4b8d-8231-963fc59ff91c",
		IpAddress:      "192.168.1.152",
		ReverseBinding: &reversebinding,
	}

	portId, err := nics.BindNic(sc, "ef6684ac-f1d2-459e-8c56-de724a68a122", opts).ExtractPortId()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test bind nic success,portId is: ", portId)
}

func TestUnBindNic(sc *gophercloud.ServiceClient) {
	reversebinding := false
	opts := nics.UnBindOpts{
		SubnetId:       "",
		IpAddress:      "",
		ReverseBinding: &reversebinding,
	}

	portId, err := nics.UnBindNic(sc, "ef6684ac-f1d2-459e-8c56-de724a68a122", opts).ExtractPortId()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test unbind nic success,portId is: ", portId)
}
