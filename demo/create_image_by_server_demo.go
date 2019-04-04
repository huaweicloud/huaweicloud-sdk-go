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
	//
	createOpts := &cloudimages.CreateByServerOpts{
		Name:            "test_image_by_server",
		InstanceId:      "83822ddc-a6e1-41e0-9073-c2a0c7309fa9",
	}

	job, err_create := cloudimages.CreateImageByServer(client, createOpts).ExtractJob()

	if err_create != nil {
		if ue, ok := err_create.(*gophercloud.UnifiedError); ok {
			fmt.Println("Failed to create image from the server.")
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Succeed to create image!")
	fmt.Println("jobID:", job.Id)

	//ff80808265a3d24b0165a90b7f3953c0

	jr, err := cloudimages.GetJobResult(client, "ff80808265a3d24b0165a90b7f3953c0").ExtractJobResult()
	if err != nil {
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	fmt.Println("Id:", jr.Id)
	fmt.Println("Type:", jr.Type)
	fmt.Println("Status:", jr.Status)
	fmt.Println("BeginTime:", jr.BeginTime)
	fmt.Println("EndTime:", jr.EndTime)
	fmt.Println("ErrorCode:", jr.ErrorCode)
	fmt.Println("FailReason:", jr.FailReason)
	fmt.Println("Entities:", jr.Entities)
	fmt.Println("Entities.ImageId:", jr.Entities.ImageId)

}
