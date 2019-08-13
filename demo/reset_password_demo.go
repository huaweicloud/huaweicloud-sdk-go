package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/auth/aksk"
	cloudserversV2 "github.com/gophercloud/gophercloud/openstack/ecs/v2/cloudservers"
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

	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}

	client, err_client := openstack.NewECSV2(provider, gophercloud.EndpointOpts{})

	if err_client != nil {
		fmt.Println("Failed to get the NewECSV2 client: ", err_client)
		return
	}

	pwd := "XXXXXXX"

	err_reset := cloudserversV2.ResetPassword(client, "d2e1a23b-0844-4580-a2c4-b2ca3a5d4167", pwd).ExtractErr()

	if err_reset != nil {
		if ue, ok := err_reset.(*gophercloud.UnifiedError); ok {
			fmt.Println("Failed to reset password for the server.")
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Succeed to reset password for server!")
}


