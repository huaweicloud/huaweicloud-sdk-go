package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	v1 "github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	v3 "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
)

func main() {

	fmt.Println("main start...")

	//provider, err := common.AuthToken()
	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}

	sc, err := openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get BlockStorage v2 client failed")
		fmt.Println(err.Error())
		return
	}

	TestGetVolumes(sc)

	fmt.Println("main end...")
}

func TestGetVolumes(sc *gophercloud.ServiceClient) {
	volume, err := volumes.Get(sc, "9547e93c-2180-4f95-8190-ea5140854c86").Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("get volume success!")
	fmt.Println("volume:", volume)
}

func TestGetQuotaSet(sc *gophercloud.ServiceClient) {
	qs, err := volumes.GetQuotaSet(sc, "054efa2069a64785a196efe56c05ee74")
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("get quota set success!")
	fmt.Println("id:", qs.QuoSet.Id)
	fmt.Println("volumes:", qs.QuoSet.Volumes)
	fmt.Println("gigabytes:", qs.QuoSet.Gigabytes)
	fmt.Println("gigabytes.limit:", qs.QuoSet.Gigabytes["limit"])
	fmt.Println("gigabytes.in_use:", qs.QuoSet.Gigabytes["in_use"])
}

func TestListVolumes(provider *gophercloud.ProviderClient) {
	//v1 client
	clientV1, err := openstack.NewBlockStorageV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get BlockStorage v1 client failed")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("V1 Endpoint:", clientV1.Endpoint)
	fmt.Println("V1 IdentityBase:", clientV1.IdentityBase)

	//v3 client
	clientV3, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get BlockStorage v3 client failed")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("V3 Endpoint:", clientV3.Endpoint)
	fmt.Println("V3 IdentityBase:", clientV3.IdentityBase)

	///////////////////////////////////
	//查询所有v3 Volumes
	listOptsV3 := v3.ListOpts{}

	allPagesV3, err := v3.List(clientV3, listOptsV3).AllPages()
	if err != nil {
		// 异常处理
		panic(err)
	}

	allVolumesV3, err := v3.ExtractVolumes(allPagesV3)
	if err != nil {
		// 异常处理
		panic(err)
	}

	fmt.Println(allVolumesV3)

	////////////////////////////////////
	//查询所有v1 Volumes
	listOptsV1 := v1.ListOpts{}

	allPagesV1, err := v1.List(clientV1, listOptsV1).AllPages()
	if err != nil {
		// 异常处理
		panic(err)
	}

	allVolumesV1, err := v1.ExtractVolumes(allPagesV1)
	if err != nil {
		// 异常处理
		panic(err)
	}

	fmt.Println(allVolumesV1)
}
