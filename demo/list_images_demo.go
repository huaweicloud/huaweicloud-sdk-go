package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack/ims/v2/cloudimages"
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

	client, err_client := openstack.NewIMSV2(provider, gophercloud.EndpointOpts{})

	if err_client != nil {
		fmt.Println("Failed to get the NewIMSV2 client: ", err_client)
		return
	}

	listOpts := cloudimages.ListOpts{
		Isregistered: "true",
	}

	allPages, err_list := cloudimages.List(client,listOpts ).AllPages()

	if err_list != nil {
		if ue, ok := err_list.(*gophercloud.UnifiedError); ok {
			fmt.Println("Failed to list images.")
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	allImages, err_extract := cloudimages.ExtractImages(allPages)

	if err_extract != nil {
		fmt.Println("Unable to extract images: ",err_extract)
	}

	fmt.Println("Succeed to list images!")
	fmt.Println("First image ID is:",allImages[0].ID)
}


