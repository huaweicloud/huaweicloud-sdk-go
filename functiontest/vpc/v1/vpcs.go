package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/vpcs"
)

func main() {

	fmt.Println("main start...")
	//AKSK authentication, initialization authentication parameters
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	//Initialization provider client
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//Initialization service client
	sc, err := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get vpc v1 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	CreateVPC(sc)
	GetVPC(sc)
	UpdateVPC(sc)
	ListVPC(sc)
	DeleteVPC(sc)

	fmt.Println("main end...")
}

func CreateVPC(sc *gophercloud.ServiceClient) {

	resp, err := vpcs.Create(sc, vpcs.CreateOpts{}).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("vpc: %+v\r\n", resp)
	fmt.Println("vpc name is:", resp.Name)
	fmt.Println("vpc Id is:", resp.ID)
	fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("vpc Status is:", resp.Status)
	fmt.Println("vpc Cidr is:", resp.Cidr)
	fmt.Println("vpc Routes is:", resp.Routes)
	fmt.Println("Create success!")

}

func UpdateVPC(sc *gophercloud.ServiceClient) {

	resp, err := vpcs.Update(sc, "1ce9a242-f01d-4fb0-9bfd-17bdc6b65853", vpcs.UpdateOpts{
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
	fmt.Printf("vpc: %+v\r\n", resp)
	fmt.Println("vpc name is:", resp.Name)
	fmt.Println("vpc Id is:", resp.ID)
	fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("vpc Status is:", resp.Status)
	fmt.Println("vpc Cidr is:", resp.Cidr)
	fmt.Println("vpc Routes is:", resp.Routes)
	fmt.Println("Update success!")

}

func GetVPC(sc *gophercloud.ServiceClient) {
	resp, err := vpcs.Get(sc, "1ce9a242-f01d-4fb0-9bfd-17bdc6b65853").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("vpc: %+v\r\n", resp)
	fmt.Println("vpc name is:", resp.Name)
	fmt.Println("vpc Id is:", resp.ID)
	fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("vpc Status is:", resp.Status)
	fmt.Println("vpc Cidr is:", resp.Cidr)
	fmt.Println("vpc Routes is:", resp.Routes)
	fmt.Println("Get success!")

}

func ListVPC(sc *gophercloud.ServiceClient) {

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

	vpcList, err1 := vpcs.ExtractVpcs(allpages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return
	}

	fmt.Printf("vpc: %+v\r\n", vpcList)
	for _, resp := range vpcList {

		fmt.Println("vpc name is:", resp.Name)
		fmt.Println("vpc Id is:", resp.ID)
		fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
		fmt.Println("vpc Status is:", resp.Status)
		fmt.Println("vpc Cidr is:", resp.Cidr)
		fmt.Println("vpc Routes is:", resp.Routes)
	}
	fmt.Println("List success!")

}

func DeleteVPC(sc *gophercloud.ServiceClient) {

	resp := vpcs.Delete(sc, "1ce9a242-f01d-4fb0-9bfd-17bdc6b65853")
	if resp.Err != nil {
		fmt.Println(resp.Err)
		if ue, ok := resp.Err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete success!")
}
