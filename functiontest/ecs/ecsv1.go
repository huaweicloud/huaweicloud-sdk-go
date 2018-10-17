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
	//TestGetEcs(sc)
	TestGetEcsExt(sc)
	fmt.Println("main end...")
}



func TestGetEcs(sc *gophercloud.ServiceClient) {
	//2c2cd6a9-c501-42a9-a679-53518e6757cc
	resp,err:=cloudservers.Get(sc,"d26b697b-3a74-4ec2-bd9d-5c3829f5d8a5").Extract()
	if err!=nil{
		fmt.Println(err)
	}
	b,errr:=json.MarshalIndent(*resp,""," ")

	if errr!=nil{

		fmt.Println(errr)
	}
	fmt.Println(string(b))

}





func TestGetEcsExt(sc *gophercloud.ServiceClient) {
	//2c2cd6a9-c501-42a9-a679-53518e6757cc
	//95b23c71-0016-4f80-b160-7c1e0341d205
	resp,err:=cloudserversext.GetServerExt(sc,"5925dbea-829c-4f4b-832e-ef61ea8b5677")
	if err!=nil{

		fmt.Println(err)
	}
	b,errr:=json.MarshalIndent(resp,""," ")

	if errr!=nil{

		fmt.Println(errr)
	}

	fmt.Println(string(b))

}
