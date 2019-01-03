package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/vpcs"
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
		fmt.Println("get vpc v1 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//TestCreateVPC(sc)
	TestGetVPC(sc)
	TestUpdateVPC(sc)
	//TestListVPC(sc)
	//TestDeleteVPC(sc)

	fmt.Println("main end...")
}

func TestCreateVPC(sc *gophercloud.ServiceClient) {

	resp, err := vpcs.Create(sc, vpcs.CreateOpts{
		Name: "ABC",
		Cidr: "192.168.0.0/16",
	}).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("vpc name is:", resp.Name)
	fmt.Println("vpc Id is:", resp.ID)
	fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("vpc Status is:", resp.Status)
	fmt.Println("vpc Cidr is:", resp.Cidr)
	fmt.Println("vpc Routes is:", resp.Routes)

}

func TestUpdateVPC(sc *gophercloud.ServiceClient) {

	resp, err := vpcs.Update(sc, "463497ec-7a31-4c82-91e7-360243e54be0", vpcs.UpdateOpts{
		Name: "ABC-back",
		Cidr: "192.168.0.0/24",
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("vpc name is:", resp.Name)
	fmt.Println("vpc Id is:", resp.ID)
	fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("vpc Status is:", resp.Status)
	fmt.Println("vpc Cidr is:", resp.Cidr)
	fmt.Println("vpc Routes is:", resp.Routes)

}

func TestGetVPC(sc *gophercloud.ServiceClient) {
	resp, err := vpcs.Get(sc, "463497ec-7a31-4c82-91e7-360243e54be0").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("vpc name is:", resp.Name)
	fmt.Println("vpc Id is:", resp.ID)
	fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("vpc Status is:", resp.Status)
	fmt.Println("vpc Cidr is:", resp.Cidr)
	fmt.Println("vpc Routes is:", resp.Routes)

}

func TestListVPC(sc *gophercloud.ServiceClient) {

	allpages, err := vpcs.List(sc, vpcs.ListOpts{
		//Limit: 2,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	vpcList,err := vpcs.ExtractVpcs(allpages)
	for _, resp := range vpcList {

		fmt.Println("vpc name is:", resp.Name)
		fmt.Println("vpc Id is:", resp.ID)
		fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
		fmt.Println("vpc Status is:", resp.Status)
		fmt.Println("vpc Cidr is:", resp.Cidr)
		fmt.Println("vpc Routes is:", resp.Routes)
	}

}

func TestDeleteVPC(sc *gophercloud.ServiceClient) {

	resp := vpcs.Delete(sc, "5aaaf1cc-9138-4958-955d-4cd6193ff9ff")
	if resp.Err != nil {
		fmt.Println(resp.Err)
		if ue, ok := resp.Err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete VPC success!")
}
