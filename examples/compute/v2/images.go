package main

import (
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
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
	imageId := "{imageId}"
	ImagesList(client)
	ImageDelete(client, imageId)
	ImageGet(client, imageId)
	fmt.Println("main end...")
}

//Query images list
func ImagesList(client *gophercloud.ServiceClient) {
	listOpts := images.ListOpts{
		Status: "active",
	}
	// Query all images list information
	allPages, allPagesErr := images.ListDetail(client, listOpts).AllPages()
	if allPagesErr != nil {
		fmt.Println("allPagesErr:", allPagesErr)
		if ue, ok := allPagesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	// Transform images structure
	allImages, allImagesErr := images.ExtractImages(allPages)
	if allImagesErr != nil {
		fmt.Println("allImagesErr:", allImagesErr)
		if ue, ok := allImagesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("images list is : ")
	for _, image := range allImages {
		imageJson, _ := json.MarshalIndent(image, "", " ")
		fmt.Println(string(imageJson))
	}
}

//Delete image
func ImageDelete(client *gophercloud.ServiceClient, imageId string) {
	deleteErr := images.Delete(client, imageId).ExtractErr()
	if deleteErr != nil {
		fmt.Println("deleteErr:", deleteErr)
		if ue, ok := deleteErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("image delete success!")
}

//Get image details
func ImageGet(client *gophercloud.ServiceClient, imageId string) {
	image, imagesGetErr := images.Get(client, imageId).Extract()
	if imagesGetErr != nil {
		fmt.Println("imagesGetErr:", imagesGetErr)
		if ue, ok := imagesGetErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	imageJson, _ := json.MarshalIndent(image, "", " ")
	fmt.Println("image detail is : " + string(imageJson))
}
