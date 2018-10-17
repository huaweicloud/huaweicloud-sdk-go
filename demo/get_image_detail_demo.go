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
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		ProjectID:        "f9b60643bb8e44349b75da40923cbcd3",
		AccessKey:        "HYO2CHUIHR5SBMLJQVXK",
		SecretKey:        "y5e0TNThIzb0TbsgWAcYFVcK4ejjBGZecCutoZbw",
		Domain:           "myhuaweicloud.com",
		Region:           "cn-north-1",
        DomainID:         "0986aafba48049a6b9457b89968eeabf",
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

	image, err_get := images.Get(client, "0825b367-0200-4b2e-9aae-9b7cf629c285").Extract()

	if err_get != nil {
		if ue, ok := err_get.(*gophercloud.UnifiedError); ok {
			fmt.Println("Failed to get the image detail.")
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Succeed to get the image detail!")
	fmt.Println("id:", image.ID)
	fmt.Println("createdAt:", image.CreatedAt)
	fmt.Println("name:", image.Name)
}


