package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/snapshots"
)

func main() {

	fmt.Println("main start...")

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

	//TestSnapshotsCreate(sc)
	//TestSnapshotsGet(sc)
	TestSnapshotsList(sc)
	//TestSnapshotsUpdateMetadata(sc)
	TestSnapshotsDelete(sc)

	fmt.Println("main end...")
}
func TestSnapshotsCreate(sc *gophercloud.ServiceClient) {
	opts := snapshots.CreateOpts{
		VolumeID: "bc9aef05-299f-4e87-bd7a-779780020690",
		Name:     "kaka",
	}

	resp, err := snapshots.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test Snapshots create success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

func TestSnapshotsGet(sc *gophercloud.ServiceClient) {
	resp, err := snapshots.Get(sc, "bd6f11c7-b06b-4868-bf98-81a4e714cbc4").Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Snapshots get success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

}

func TestSnapshotsList(sc *gophercloud.ServiceClient) {
	allPages, err := snapshots.List(sc, snapshots.ListOpts{Limit: 5}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Snapshots List success!")

	allData, _ := snapshots.ExtractSnapshots(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData.SnapshotsLinks {
		fmt.Println()
		b, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(b))
	}

	for _, v := range allData.Snapshots {
		b, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(b))
	}

}

func TestSnapshotsUpdateMetadata(sc *gophercloud.ServiceClient) {
	id := "bd6f11c7-b06b-4868-bf98-81a4e714cbc4"
	bb := make(map[string]interface{})
	bb["c"] = "sadfasdf"

	updatOpts := snapshots.UpdateMetadataOpts{
		Metadata: bb,
	}

	resp, err := snapshots.UpdateMetadata(sc, id, updatOpts).ExtractMetadata()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test Snapshots update success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

func TestSnapshotsDelete(sc *gophercloud.ServiceClient) {
	id := "bd6f11c7-b06b-4868-bf98-81a4e714cbc4"

	err := snapshots.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Snapshots delete success!")
}
