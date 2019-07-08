package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/peerings"
)

func main() {

	fmt.Println("main start...")

	//AKSK authentication, initialization authentication parameters
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	//Initialization provider client
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		fmt.Println(err)
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//Initialization service client
	sc, err := openstack.NewVPCV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get network vpc v2 client failed")
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	CreatePeering(sc)
	GetPeering(sc)
	ListPeering(sc)
	DeletePeering(sc)
	UpdatePeering(sc)
	AcceptPeering(sc)
	RejectPeering(sc)

	fmt.Println("main end...")
}

// Create a peering
func CreatePeering(client *gophercloud.ServiceClient) {

	opts := peerings.CreateOpts{
		Name: "xxxxxx",
		RequestVpcInfo: peerings.VPCInfo{
			VpcID:    "xxxxxx",
			TenantID: "xxxxxx",
		},
		AcceptVpcInfo: peerings.VPCInfo{
			VpcID:    "xxxxxx",
			TenantID: "xxxxxx",
		},
	}

	result, err := peerings.Create(client, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("peering: %+v\r\n", result)

	fmt.Println("peering Id is:", result.ID)
	fmt.Println("peering Name is:", result.Name)
	fmt.Println("peering Status is:", result.Status)
	fmt.Println("peering Description is:", result.Description)
	fmt.Println("peering AcceptVpcInfo is:", result.AcceptVpcInfo)
	fmt.Println("peering RequestVpcInfo is:", result.RequestVpcInfo)

	fmt.Println(" Create success!")
}

// Get a peering
func GetPeering(client *gophercloud.ServiceClient) {

	id := "xxxxxx"

	result, err := peerings.Get(client, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("peering: %+v\r\n", result)

	fmt.Println("peering Id is:", result.ID)
	fmt.Println("peering Name is:", result.Name)
	fmt.Println("peering Status is:", result.Status)
	fmt.Println("peering Description is:", result.Description)
	fmt.Println("peering AcceptVpcInfo is:", result.AcceptVpcInfo)
	fmt.Println("peering RequestVpcInfo is:", result.RequestVpcInfo)

	fmt.Println(" Get success!")
}

// List peering
func ListPeering(client *gophercloud.ServiceClient) {

	result, err := peerings.List(client, peerings.ListOpts{
		Limit: 30,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	peerList, err := peerings.ExtractPeerings(result)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("peerings: %+v\r\n", peerList)

	for _, resp := range peerList {

		fmt.Println("peering Id is:", resp.ID)
		fmt.Println("peering Name is:", resp.Name)
		fmt.Println("peering Status is:", resp.Status)
		fmt.Println("peering Description is:", resp.Description)
		fmt.Println("peering AcceptVpcInfo is:", resp.AcceptVpcInfo)
		fmt.Println("peering RequestVpcInfo is:", resp.RequestVpcInfo)
	}
	fmt.Println(" List success!")
}

// Delete a peering
func DeletePeering(client *gophercloud.ServiceClient) {
	id := "xxxxxx"
	err := peerings.Delete(client, id).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Delete success!")
}

// Update a peering
func UpdatePeering(client *gophercloud.ServiceClient) {

	id := "xxxxxx"
	opts := peerings.UpdateOpts{
		Name: "xxxxxx",
	}

	result, err := peerings.Update(client, id, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("peering: %+v\r\n", result)

	fmt.Println("peering Id is:", result.ID)
	fmt.Println("peering Name is:", result.Name)
	fmt.Println("peering Status is:", result.Status)
	fmt.Println("peering Description is:", result.Description)
	fmt.Println("peering AcceptVpcInfo is:", result.AcceptVpcInfo)
	fmt.Println("peering RequestVpcInfo is:", result.RequestVpcInfo)

	fmt.Println("Update success!")
}

// Accept peering request
func AcceptPeering(client *gophercloud.ServiceClient) {

	id := "xxxxxx"

	result, err := peerings.Accept(client, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("peering: %+v\r\n", result)

	fmt.Println("peering Id is:", result.ID)
	fmt.Println("peering Name is:", result.Name)
	fmt.Println("peering Status is:", result.Status)
	fmt.Println("peering Description is:", result.Description)
	fmt.Println("peering AcceptVpcInfo is:", result.AcceptVpcInfo)
	fmt.Println("peering RequestVpcInfo is:", result.RequestVpcInfo)

	fmt.Println("AcceptPeering success!")
}

// Reject peering request
func RejectPeering(client *gophercloud.ServiceClient) {

	id := "xxxxxx"

	result, err := peerings.Reject(client, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("peering: %+v\r\n", result)

	fmt.Println("peering Id is:", result.ID)
	fmt.Println("peering Name is:", result.Name)
	fmt.Println("peering Status is:", result.Status)
	fmt.Println("peering Description is:", result.Description)
	fmt.Println("peering AcceptVpcInfo is:", result.AcceptVpcInfo)
	fmt.Println("peering RequestVpcInfo is:", result.RequestVpcInfo)

	fmt.Println("RejectPeering success!")
}
