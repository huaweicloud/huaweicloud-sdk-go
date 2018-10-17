package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootwithscheduler"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/schedulerhints"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/serversext"

	tokens3 "github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
)

func main() {
	fmt.Println("main start...")
	//provider, err := common.AuthToken()
	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println(err.Error())
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	// 设置计算服务的client
	sc, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println(err.Error())
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//TestListBrief(sc)
	//TestList(sc)
	//TestDelete(sc)
	TestServerextGET(sc)
	//TestGet(sc)
	TestExtCreateServer(sc)
	//TestExtGet(sc)
	//TestExtListServers(sc)
	//TestEndPointAndSeriveVersion(provider)

	fmt.Println("main end...")
}

func TestServerextGET(sc *gophercloud.ServiceClient)  {

	ID:="93899eb0-f092-4e74-bef1-0a8cf0ee4e16"
	resp,err:=serversext.Get(sc,ID)

	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(resp)


}


func TestListBrief(sc *gophercloud.ServiceClient) {
	allPages, err := servers.ListBrief(sc, servers.ListOpts{Limit: 3}).AllPages()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	allServers, err := servers.ExtractServerBriefs(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("List brief servers  success!")
	for _, s := range allServers {
		fmt.Println(s)
	}
}

func TestList(sc *gophercloud.ServiceClient) {
	allPages, err := servers.List(sc, servers.ListOpts{Limit: 5}).AllPages()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("List Servers:")
	for _, s := range allServers {
		fmt.Println(s)
	}
}

func TestExtListServers(sc *gophercloud.ServiceClient) {
	demandServers, monthlySvrs, err := serversext.ListServers(sc)
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("ListServers success!")

	fmt.Println("demandServers:")
	for _, s := range demandServers {
		fmt.Println(s.ID)
	}

	fmt.Println("monthlySvrs:")
	for _, s := range monthlySvrs {
		fmt.Println(s.ID)
	}
}

func TestExtGet(sc *gophercloud.ServiceClient) {
	serverId := "73238006-4b4b-4b93-b93e-f78404495e3b" //"1778e268-b577-430b-90e4-b656cc997290"
	serverExt, err := serversext.Get(sc, serverId)
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("ext Get success!")

	fmt.Println("serverExt.ID:", serverExt.ID)

	fmt.Println("serverExt.VolumeAttached:", serverExt.VolumeAttached)
	fmt.Println("serverExt.Charging:", serverExt.Charging)
	fmt.Println("serverExt.Charging.ValidTime:", serverExt.Charging.ValidTime)
	fmt.Println("serverExt.Charging.ExpireTime:", serverExt.Charging.ExpireTime)
	fmt.Println("serverExt.VpcId:", serverExt.VpcId)

	//fmt.Println("serverExt.Addresses['addr']:", )
}

func TestDelete(sc *gophercloud.ServiceClient)  {

		serverId := "73238006-4b4b-4b93-b93e-f78404495e3b" //"1778e268-b577-430b-90e4-b656cc997290"
	err := servers.Delete(sc, serverId).ExtractErr()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("ext delete success!")

}


func TestGet(sc *gophercloud.ServiceClient) {
	server, err := servers.Get(sc, "2251f59c-b1ef-4398-bfa8-321782f670a5").Extract()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("get server success, server id:", server.ID)

	fmt.Println("server:", server)

	fmt.Println("VolumeAttached[0]['ID']:", server.VolumeAttached[0]["id"])

	//fmt.Println("ID:", server.VolumeAttached.ID)
}

func TestExtCreateServer(sc *gophercloud.ServiceClient) {
	fmt.Println("***BEGIN TEST CreateServer***")
	//	baseOpts := servers.CreateOpts{
	//		Name: "ECS_xx1",
	//		//FlavorRef: "normal1",
	//		FlavorRef: "s2.xlarge.4",
	//		ImageRef:  "53b2fbb5-ef2c-412a-bb0a-571436fa78ad",
	//		Networks: []servers.Network{
	//			servers.Network{UUID: "5689bda8-767d-4029-9c9e-460c3e05f46a"},
	//			//servers.Network{UUID: "5689bda8-767d-4029-9c9e-460c3e05f46a"},
	//			//servers.Network{UUID: "5689bda8-767d-4029-9c9e-460c3e05f46a"},
	//			//servers.Network{Port: "36000"},
	//			//servers.Network{UUID: "5689bda8-767d-4029-9c9e-460c3e05f46a"},
	//			//servers.Network{UUID: "5689bda8-767d-4029-9c9e-460c3e05f46a", Port: "36001"},
	//			//servers.Network{Port: "36001"},
	//		},
	//		Metadata:         map[string]string{"hello": "world"},
	//		SecurityGroups:   []string{"default"},
	//		UserData:         []byte("IyEvYmluL2Jhc2gKZWNobyAncm9vdDokNiRPRjMxdlo0cm1CWUpvZzBLJE1ldlVrS3dSYVI0SmM2QVRaSi9lT2s4Q0ZFWUo1NFVSOFlvc2xsZUd0RERIRHd4TWRuU3lJcUw0WS9jN0MvSFlRcmRVZG45WXJKQnlhRnlvZm5ybjYuJyB8IGNocGFzc3dkIC1l"),
	//		AvailabilityZone: "eu-de-02",
	//	}

	baseOpts := servers.CreateOpts{
		Name:      "ECS_xx2",
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
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("server:", server)
	fmt.Println("***END TEST CreateServer***")
}

func TestEndPointAndSeriveVersion(provider *gophercloud.ProviderClient) {
	fmt.Println("***BEGIN TEST EndPointAndSeriveVersion***")
	//生成url的函数
	eo := gophercloud.EndpointOpts{}
	//eo.ApplyDefaults("compute")
	eo.Type = "compute"
	eo.Availability = "public"
	url, _ := provider.EndpointLocator(eo)
	fmt.Println("url:=provider.EndpointLocator(eo):", url)

	//构造catalog
	catalog := tokens3.ServiceCatalog{
		Entries: []tokens3.CatalogEntry{
			tokens3.CatalogEntry{
				ID:   "106d583bbbf445d18da621eda615c50d",
				Name: "Workspace",
				Type: "wks",
				Endpoints: []tokens3.Endpoint{
					tokens3.Endpoint{
						ID:        "cef7ceece3e24b44bc4ca409d6c2c7f8",
						Region:    "eu-de",
						Interface: "public",
						URL:       "https://172.30.57.212/v1.0/054efa2069a64785a196efe56c05ee74"}}},
			tokens3.CatalogEntry{
				ID:   "5d67a8f14ea6418993810bfb20ba1a46",
				Name: "novav2.1",
				Type: "computev2.1",
				Endpoints: []tokens3.Endpoint{
					tokens3.Endpoint{
						ID:        "653723e8d676459c9e81020d9cf8c761",
						Region:    "eu-de",
						Interface: "public",
						URL:       "https://ecs.eu-de.otc.t-systems.com/v2.1/054efa2069a64785a196efe56c05ee74"}}},
			tokens3.CatalogEntry{
				ID:   "875ee02920614b36b8a6806a54dc453f",
				Name: "nova",
				Type: "compute",
				Endpoints: []tokens3.Endpoint{
					tokens3.Endpoint{
						ID:        "967e00662dec4d5da8f9ef77b19e6eed",
						Region:    "eu-de",
						Interface: "public",
						URL:       "https://ecs.eu-de.otc.t-systems.com/v2/054efa2069a64785a196efe56c05ee74"}}}},
	}

	opts := gophercloud.EndpointOpts{Type: "computev2.1", Availability: "public"}
	url, _ = openstack.V3EndpointURL(&catalog, opts)
	fmt.Println("url, _ := openstack.V3EndpointURL(&catalog, opts) url:", url)
}
