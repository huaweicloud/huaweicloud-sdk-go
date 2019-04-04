package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudserversext"
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

	var orderID = "CS1811281510AF3Z2"

	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Fail to get the provider: ", errAuth)
		return
	}

	client, errClient := openstack.NewECSV1(provider, gophercloud.EndpointOpts{})

	if errClient != nil {
		fmt.Println("Fail to get the ecs v1 client: ", errClient)
		return
	}

	servers, errGet := cloudserversext.GetPrepaidServerDetailByOrderId(client, orderID)

	if errGet != nil {
		if se, ok := errGet.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		} else {
			fmt.Println("Error:", errGet)
		}
		return
	}

	if len(servers) == 0 {
		fmt.Println("there is no server info on this order")
	}

	for _, data := range servers {
		//charging info
		fmt.Println("Charging ChargingMode:", data.Charging.ChargingMode)
		fmt.Println("Charging ExpireTime:", data.Charging.ExpireTime)
		fmt.Println("Charging ValidTime:", data.Charging.ValidTime)

		//cloud server info
		fmt.Println("CloudServer Status:", data.CloudServer.Status)
		fmt.Println("CloudServer ID:", data.CloudServer.ID)
		fmt.Println("CloudServer Name:", data.CloudServer.Name)
		fmt.Println("CloudServer Metadata:", data.CloudServer.Metadata)
		fmt.Println("CloudServer LaunchedAt:", data.CloudServer.LaunchedAt)
		fmt.Println("CloudServer AvailabilityZone:", data.CloudServer.AvailabilityZone)

		//server volume info
		for _, volumeData := range data.VolumeAttached {
			fmt.Println("VolumeAttached ID:", volumeData.ID)
			fmt.Println("VolumeAttached Size:", volumeData.Size)
			fmt.Println("VolumeAttached VolumeType:", volumeData.VolumeType)

		}
	}

}
