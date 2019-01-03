package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/secgroups"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get compute v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	GetServerSecurityGroupList(sc)
	fmt.Println("main end...")
}

func GetServerSecurityGroupList(sc *gophercloud.ServiceClient) {
	page, err := secgroups.ListByServer(sc, "81cfeb20-6c63-412c-81bf-4524d0e8e390").AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	PageData, err := secgroups.ExtractSecurityGroups(page)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("get Server Security Groups List success")

	for _, data := range PageData {

		fmt.Println("Server Security Groups name is ", data.Name)
		fmt.Println("Server Security Groups ID is ", data.ID)
		for _, rule := range data.Rules {
			fmt.Println("Server Security Groups Rule ID is", rule.ID)
			fmt.Println("Server Security Groups Rule FromPort is", rule.FromPort)
			fmt.Println("Server Security Groups Rule ToPort is", rule.ToPort)
			fmt.Println("Server Security Groups Rule IPProtocol is", rule.IPProtocol)
			fmt.Println("Server Security Groups Rule IPRange is", rule.IPRange)
			fmt.Println("Server Security Groups Rule ParentGroupID is ", rule.ParentGroupID)
		}
	}
}
