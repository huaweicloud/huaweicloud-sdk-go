package main

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootwithscheduler"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/schedulerhints"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/serversext"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

func main() {

	//test IAM
	providerIAM, err := TestIAM()
	if err != nil {
		fmt.Println("Test IAM failed!")
		PrintError(err)
		return
	} else {
		fmt.Println("Test IAM ok!")
	}

	//test AKSK
	providerAKSK, err := TestAKSK()
	if err != nil {
		fmt.Println("Test AKSK failed!")
		PrintError(err)
		return
	} else {
		fmt.Println("Test AKSK ok!")
	}

	//test ECS
	TestECS(providerIAM)
	if err != nil {
		fmt.Println("Test ECS failed!")
		PrintError(err)
		return
	} else {
		fmt.Println("Test ECS ok!")
	}

	//test IMS
	TestECS(providerAKSK)
	if err != nil {
		fmt.Println("Test IMS failed!")
		PrintError(err)
		return
	} else {
		fmt.Println("Test IMS ok!")
	}
}

func PrintError(err error) {
	if ue, ok := err.(*gophercloud.UnifiedError); ok {
		fmt.Println("ErrCode:", ue.ErrorCode())
		fmt.Println("Message:", ue.Message())
	} else {
		fmt.Println(err)
	}
}

/************ Test IAM **************/
func TestIAM() (*gophercloud.ProviderClient, error) {
	tokenOpts := token.TokenOptions{
		Username:         "",
		Password:         "",
		TenantID:         "128a7bf965154373a7b73c89eb6b65aa",
		DomainID:         "3b011b89b2f64fb68782a43380e2a78f",
		IdentityEndpoint: "",
		AllowReauth:      true,
	}

	return openstack.AuthenticatedClient(tokenOpts)
}

/************ Test AKSK **************/
func TestAKSK() (*gophercloud.ProviderClient, error) {
	akskOpts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "128a7bf965154373a7b73c89eb6b65aa",
		AccessKey:        "ETILDMOHPYAP0V0OREKD",
		SecretKey:        "O36HHeJ4Ol3VQbXWPcRUvDhF6iYyCsValpsMJSkU",
		Domain:           "",
		Region:           "xxx",
	}

	return openstack.AuthenticatedClient(akskOpts)
}

/************ Test ECS **************/
func TestECS(provider *gophercloud.ProviderClient) error {
	sc, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		return err
	}

	allPages, err := servers.List(sc, servers.ListOpts{}).AllPages()
	if err != nil {
		return err
	}

	_, err = servers.ExtractServers(allPages)
	if err != nil {
		return err
	}

	return err
}

/************ Test IMS **************/
func TestIMS(provider *gophercloud.ProviderClient) error {
	sc, err := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		return err
	}

	allPages, err := images.List(sc, images.ListOpts{}).AllPages()
	if err != nil {
		return err
	}

	_, err = images.ExtractImages(allPages)
	if err != nil {
		return err
	}

	return err
}

func TestGetServer(sc *gophercloud.ServiceClient) {
	server, err := servers.Get(sc, "931616e6-3d52-4974-b8b1-3a59b6e2c106").Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("get server success, server id:", server.ID)
}

func TestCreateServer(sc *gophercloud.ServiceClient) {
	fmt.Println("***BEGIN TEST CreateServer***")

	baseOpts := servers.CreateOpts{
		Name:      "ECS_xx1",
		FlavorRef: "c1.xlarge",
		//ImageRef:  "2a50f694-b8e7-4a7a-8a51-0ff7f83d1345",
		Networks: []servers.Network{
			servers.Network{UUID: "9a56640e-5503-4b8d-8231-963fc59ff91c"},
		},
		AvailabilityZone: "az1.dc1",
	}

	bd := []bootfromvolume.BlockDevice{
		bootfromvolume.BlockDevice{
			BootIndex:       0,
			DestinationType: "volume",
			SourceType:      "image",
			VolumeSize:      40,
			UUID:            "ee5c7dc8-acb8-4d93-8d47-b27610b3477d",
		},
	}

	sh := schedulerhints.SchedulerHints{
		CheckResources: "true",
	}

	bsOpts := bootwithscheduler.CreateOptsExt{
		CreateOptsBuilder: baseOpts,
		BlockDevice:       bd,
		SchedulerHints:    sh,
	}

	server, err := serversext.CreateServer(sc, bsOpts).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("server:", server)
	fmt.Println("***END TEST CreateServer***")
}
