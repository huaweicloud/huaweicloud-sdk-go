package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	createOpts "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/serversext"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootwithscheduler"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/schedulerhints"
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

	client, err_client := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	if err_client != nil {
		fmt.Println("Fail to get the computer client: ", err_client)
		return
	}

	serverCreateOpts := servers.CreateOpts{
		Name:             "test",
		FlavorRef:        "s2.small.1",
		Networks:         []servers.Network{servers.Network{UUID: "fb510786-77f1-4534-8daa-2faa1b5b5a33"}},
		Metadata:         map[string]string{"hello": "world"},
		SecurityGroups:   []string{"fae12c64-6f57-42ff-9392-5fb40aee0f23"},
		AvailabilityZone: "cn-north-1a",
		KeyName:          "Keypair-test",
	}

	bdm_v2 := []createOpts.BlockDevice{
		createOpts.BlockDevice{
			BootIndex:           0,
			DeleteOnTermination: true,
			DestinationType:     "volume",
			SourceType:          "image",
			VolumeSize:          40,
			UUID:                "1189efbf-d48b-46ad-a823-94b942e2a000",
			VolumeType:          "SAS",
		},
		createOpts.BlockDevice{
			BootIndex:           -1,
			DeleteOnTermination: true,
			DestinationType:     "volume",
			SourceType:          "blank",
			VolumeSize:          10,
			VolumeType:          "SSD",
		},
	}

	sh := schedulerhints.SchedulerHints{
		CheckResources: "true",
	}

	vhOpts := bootwithscheduler.CreateOptsExt{
		CreateOptsBuilder: serverCreateOpts,
		BlockDevice:       bdm_v2,
		SchedulerHints:    sh,
	}

	server, err_create := serversext.CreateServer(client, vhOpts).Extract()
	if err_create != nil {
		if se, ok := err_create.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		} else {
			fmt.Println("Error:", err_create)
		}
		return
	}

	fmt.Println("server ID:", server.ID)
}


