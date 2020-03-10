package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/flavor"
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
	client, clientErr := openstack.NewECSV1(provider, gophercloud.EndpointOpts{})
	if clientErr != nil {
		fmt.Println("Failed to get the NewECSV1 client: ", clientErr)
		return
	}
	ListProjectFlavors(client)

	fmt.Println("main end...")

}

//ListProjectFlavors requests to query project flavors.
func ListProjectFlavors(sc *gophercloud.ServiceClient) {
	allPages, err := flavor.List(sc, &flavor.ListOpts{AvailabilityZone: "XXXXXXXX"}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	flavors, err := flavor.ExtractFlavors(allPages)
	if err != nil {
		fmt.Println("Test get flavor list error:", err)
		return
	}

	b, _ := json.MarshalIndent(flavors, "", "   ")
	fmt.Println(string(b))
}
