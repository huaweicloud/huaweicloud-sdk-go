package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumetransfer"
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
		return
	}

	TestTransferCreate(sc)
	TestTransferAccept(sc)
	//TestTransferGet(sc)
	TestTransferListDetails(sc)
	//TestTransferList(sc)
	//TestTransferDelete(sc)
	fmt.Println("main end...")
}

func TestTransferCreate(sc *gophercloud.ServiceClient) {
	opts := volumetransfer.CreateOpts{
		VolumeID: "9a6f7d4d-63b8-4fb2-91cf-80934ca60664",
	}

	resp, err := volumetransfer.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

	fmt.Println("Test TestTransferCreate success!")

}

func TestTransferAccept(sc *gophercloud.ServiceClient) {
	opts := volumetransfer.AcceptOpts{
		AuthKey: "612fe80094d9a62f",
	}
	id := "efeff0a6-2d27-49dc-b340-15c247691fe1"
	resp, err := volumetransfer.Accept(sc, id, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

	fmt.Println("Test Transfer accept success!")
}

func TestTransferGet(sc *gophercloud.ServiceClient) {
	resp, err := volumetransfer.Get(sc, "b8df3db4-4597-4bdb-88da-af3d646802c7").Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

	fmt.Println("Test Transfer get success!")
}

func TestTransferListDetails(sc *gophercloud.ServiceClient) {
	allPages, err := volumetransfer.ListDetail(sc, volumetransfer.ListDetailOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	allData, _ := volumetransfer.ExtractTransferDetails(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData {
		fmt.Println(v.VolumeID)
		b, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("Test Transfer List details success!")
}

func TestTransferList(sc *gophercloud.ServiceClient) {

	volumePage, err := volumetransfer.List(sc, volumetransfer.ListDetailOpts{}).AllPages()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	volumeList, err1 := volumetransfer.ExtractTransferList(volumePage)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	for _, d := range volumeList {
		fmt.Println("volume id :", d.ID)
		fmt.Println("volume Name :", d.Name)
		fmt.Printf("%+v", d)
	}
	fmt.Println("get Transfer list  success!")

}

func TestTransferDelete(sc *gophercloud.ServiceClient) {
	id := "b8df3db4-4597-4bdb-88da-af3d646802c7"
	err := volumetransfer.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test TestTransferDelete success!")
}
