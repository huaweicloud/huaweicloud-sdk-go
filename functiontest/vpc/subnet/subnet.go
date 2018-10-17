package main


import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/subnets"
	"encoding/json"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	TestSubnetList(sc)
	fmt.Println("main end...")
}

func TestSubnetList(sc *gophercloud.ServiceClient)  {

	allPages, err := subnets.List(sc,subnets.ListOpts{VpcID:"1d79d5ce-bc4c-48c6-88cd-4a8619f6ad2c"}).AllPages()
		if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test TestSubnetList success!")

	allData,_:=subnets.ExtractSubnets(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _,v :=range allData{
		p,_:=json.MarshalIndent(v,""," ")
		fmt.Println(string(p))
	}

}




