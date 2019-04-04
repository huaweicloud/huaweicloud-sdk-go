package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudservers"

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
		fmt.Println("Fail to get the provider: ", err_auth)
		return
	}

	client, err_client := openstack.NewECSV1_1(provider, gophercloud.EndpointOpts{})

	if err_client != nil {
		fmt.Println("Fail to get the ECSV1_1 client: ", err_client)
		return
	}

	nics := []cloudservers.Nic{
		cloudservers.Nic{
			SubnetId: "8cd34c55-8f59-458c-a602-7476fcd1475c",
		},
	}

	rv := cloudservers.RootVolume{
		VolumeType: "SSD",
	}

	dvs := []cloudservers.DataVolume{
		cloudservers.DataVolume{
			VolumeType: "SATA",
			Size:       10,
		},
		cloudservers.DataVolume{
			VolumeType: "SAS",
			Size:       20,
		},
	}

	md := &cloudservers.MetaData{
		OpSvcUserId:    "e5b8b8758ce94763a6e2f93b3b3ffcba",
	}

	sep := &cloudservers.ServerExtendParam{
		ChargingMode:       "prePaid",
		PeriodType:         "month",
		PeriodNum:          1,
		IsAutoPay:          "false",
	}

	serverCreateOpts := cloudservers.CreateOpts{
		Name:               "test_prepaid_server",
		FlavorRef:          "s2.small.1",
		ImageRef:           "1189efbf-d48b-46ad-a823-94b942e2a000",
		VpcId:              "f8c4770d-aeb8-4935-9b2f-d1e9d9aba3f7",
		Nics:               nics,
		RootVolume:         rv,
		DataVolumes:        dvs,
		AvailabilityZone:   "cn-north-1a",
		ExtendParam:        sep,
		KeyName:            "KeyPair-test",
		MetaData:           md,
		IsAutoRename:       gophercloud.Disabled,
	}

	_, orderId, err_create := cloudservers.Create(client, serverCreateOpts)

	if err_create != nil {
		if ue, ok := err_create.(*gophercloud.UnifiedError); ok {
			fmt.Println("创建虚拟机失败。")
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("create cloudservers success!")
	fmt.Println("orderId:", orderId)
}


