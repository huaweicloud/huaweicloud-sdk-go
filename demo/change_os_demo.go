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

	client, err_client := openstack.NewECSV2(provider, gophercloud.EndpointOpts{})

	if err_client != nil {
		fmt.Println("Failed to get the NewECSV2 client: ", err_client)
		return
	}

	changeOpts := cloudserversV2.ChangeOpts{
		ImageID:       "2a50f694-b8e7-4a7a-8a51-0ff7f83d1345",
		KeyName:       "KeyPair-9ec0",
	}

	job, err_change := cloudserversV2.ChangeOS(client, "7ca89a44-0a72-444c-9032-907b68178575", changeOpts).ExtractJob()

	if err_change != nil {
		if ue, ok := err_change.(*gophercloud.UnifiedError); ok {
			fmt.Println("Failed to change the OS for server.")
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Succeed to change OS !")
	fmt.Println("jobID:", job.ID)
}


