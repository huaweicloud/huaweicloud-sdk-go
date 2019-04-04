package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
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

	client, err_client := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{})

	if err_client != nil {
		fmt.Println("Failed to get the NewImageServiceV2 client: ", err_client)
		return
	}

	err_del := images.Delete(client, "03a76032-604b-4edf-8049-418e7186bd4b").ExtractErr()

	if err_del != nil {
		if ue, ok := err_del.(*gophercloud.UnifiedError); ok {
			fmt.Println("Failed to delete the image.")
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Succeed to delete image!")
}


