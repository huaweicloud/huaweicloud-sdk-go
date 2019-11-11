package main

import (
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"encoding/json"
)

func main() {
	fmt.Println("main start...")
	//Set authentication parameters
	tokenOpts := token.TokenOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		Username:         "{Username}",
		Password:         "{Password}",
		DomainID:         "{DomainID}",
		ProjectID:        "{ProjectID}",
	}
	//Init provider client
	provider, authErr := openstack.AuthenticatedClient(tokenOpts)
	if authErr != nil {
		fmt.Println("Failed to get the AuthenticatedClient: ", authErr)
		return
	}
	//Init service client
	client, clientErr := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if clientErr != nil {
		fmt.Println("Failed to get the NewComputeV2 client: ", clientErr)
		return
	}
	flavorsId := "{FlavorsId}"
	FlavorsList(client)
	FlavorsGet(client, flavorsId)
	fmt.Println("main end...")
}

//Query flavor list
func FlavorsList(client *gophercloud.ServiceClient) {
	listOpts := flavors.ListOpts{
		MinDisk: 20,
		MinRAM:  4096,
	}
	// Query all flavors list information
	allPages, allPagesErr := flavors.ListDetail(client, listOpts).AllPages()
	if allPagesErr != nil {
		fmt.Println("allPagesErr:", allPagesErr)
		if ue, ok := allPagesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	// Transform flavors structure
	allFlavors, allFlavorsErr := flavors.ExtractFlavors(allPages)
	if allFlavorsErr != nil {
		fmt.Println("allFlavorsErr:", allFlavorsErr)
		if ue, ok := allFlavorsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("flavors list is : ")
	for _, flavor := range allFlavors {
		flavorJson, _ := json.MarshalIndent(flavor, "", " ")
		fmt.Println(string(flavorJson))
	}
}

//Get flavor
func FlavorsGet(client *gophercloud.ServiceClient, flavorId string) {
	flavor, flavorsGetErr := flavors.Get(client, flavorId).Extract()
	if flavorsGetErr != nil {
		fmt.Println("flavorsGetErr", flavorsGetErr)
		if ue, ok := flavorsGetErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("flavors detail is : ")
	flavorJson, _ := json.MarshalIndent(flavor, "", " ")
	fmt.Println(string(flavorJson))
}
