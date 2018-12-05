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
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		ProjectID:        "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		AccessKey:        "XXXXXXXXXXXXXXXXXXXX",
		SecretKey:        "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		Domain:           "myhuaweicloud.com",
		Region:           "cn-north-1",
		DomainID:         "0986aafba48049a6b9457b89968eeabf",
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

	rebootOpts := &servers.RebootOpts{
		Type: servers.SoftReboot,
	}

	err_reboot := servers.Reboot(client, server_id ,rebootOpts).ExtractErr()

	if err_reboot != nil {
		if se, ok := err_reboot.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		} else{
			fmt.Println("Error:", err_reboot)
		}
		return
	}

	fmt.Println("Start to reboot server!")

}


