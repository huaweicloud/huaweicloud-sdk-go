package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/serversext"
	"github.com/gophercloud/gophercloud/auth/aksk"
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

	var server_id = "02057aa4-1e59-459b-814a-69b9a20a0c6a"

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

	server, err_get := serversext.Get(client, server_id)

	if err_get != nil {
		if se, ok := err_get.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		} else{
			fmt.Println("Error:", err_get)
		}
		return
	}

	fmt.Println("Server Name:", server.Name)
	fmt.Println("Server KeyName:", server.KeyName)
	fmt.Println("Server AvailbiltyZone:", server.AvailbiltyZone)
	fmt.Println("Server Status:", server.Status)
	fmt.Println("Server ChargingMode:", server.Charging.ChargingMode)
	fmt.Println("Server ValidTime:", server.Charging.ValidTime)
	fmt.Println("Server ExpireTime:", server.Charging.ExpireTime)
	fmt.Println("Server First Volume ID:", server.VolumeAttached[0].ID)
	fmt.Println("Server First VolumeType:", server.VolumeAttached[0].VolumeType)
	fmt.Println("Server Size:", server.VolumeAttached[0].Size)

}


