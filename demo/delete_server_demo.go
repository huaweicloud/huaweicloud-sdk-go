package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func main() {

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	var server_id = "bd99d40b-f9b8-4b35-b44b-44dfe248114f"

	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Fail to get the provider: ", err_auth)
		return
	}

	client, err_client := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	if err_client != nil {
		fmt.Println("Fail to get the computer client: ", err_client)
		return
	}

	err_delete := servers.Delete(client, server_id ).ExtractErr()

	if err_delete != nil {
		if se, ok := err_delete.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		} else{
			fmt.Println("Error:", err_delete)
		}
		return
	}

	fmt.Println("Start to delete server!")

}


