package main

import (
	"github.com/gophercloud/gophercloud"
	"fmt"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	az "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/availabilityzones"
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

	GetAZList(client)
	fmt.Println("main end...")
}

//GetAZList requests to get availability zones list.
func GetAZList(client *gophercloud.ServiceClient) {

	allPages, err := az.List(client).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	azinfo, err := az.ExtractAvailabilityZones(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Get az info success")
	for _, data := range azinfo {
		fmt.Println("az hosts is ", data.Hosts)
		fmt.Println("az ZoneName is ", data.ZoneName)
		fmt.Println("az ZoneState is ", data.ZoneState)
	}
}