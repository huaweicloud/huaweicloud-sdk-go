package main

import (
	"github.com/gophercloud/gophercloud/openstack"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudservers"
)

func main() {
	fmt.Println("main start...")
	gophercloud.EnableDebug = true
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
	//Init  service client
	client, clientErr := openstack.NewECSV1_1(provider, gophercloud.EndpointOpts{})
	if clientErr != nil {
		fmt.Println("Failed to get the NewComputeV2 client: ", clientErr)
		return
	}
	ServerCreate(client)
	fmt.Println("main end...")

}

//Create a server (v1.1 version)
func ServerCreate(client *gophercloud.ServiceClient) {
	nics := []cloudservers.Nic{
		cloudservers.Nic{
			SubnetId: "cc7953b3-110f-4e87-b240-ff4915548875",
		},
	}
	rv := cloudservers.RootVolume{
		VolumeType: "SATA",
	}
	dvs := []cloudservers.DataVolume{
		cloudservers.DataVolume{
			VolumeType: "SATA",
			Size:       60,
		},
		cloudservers.DataVolume{
			VolumeType: "SATA",
			Size:       70,
		},
	}
	opts := cloudservers.CreateOpts{
		Name:             "ecs_cloud_xx2",
		FlavorRef:        "c1.xlarge",
		ImageRef:         "2a50f694-b8e7-4a7a-8a51-0ff7f83d1345",
		VpcId:            "b7ff7a9b-cc95-4dd0-b76a-f586c88e6556",
		Nics:             nics,
		RootVolume:       rv,
		DataVolumes:      dvs,
		AvailabilityZone: "az1.dc1",
	}
	jobId, orderId, createErr := cloudservers.Create(client, opts)
	if createErr != nil {
		fmt.Println("createErr:", createErr)
		if ue, ok := createErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("jobId is " + jobId)
	fmt.Println("orderId is " + orderId)
	fmt.Println("server create success!")
}
