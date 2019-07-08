package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumetransfer"

	"encoding/json"
)

func main() {

	fmt.Println("main start...")

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{DomainID}",
	}

	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("get provider client failed: ", err_auth)
		return
	}

	sc, err := openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get BlockStorage v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	VolumeTransferCreate(sc)
	VolumeTransferAccept(sc)
	VolumeTransferGet(sc)
	VolumeTransferListDetails(sc)
	VolumeTransferList(sc)
	VolumeTransferDelete(sc)
	fmt.Println("main end...")
}

// create volume transfer
func VolumeTransferCreate(sc *gophercloud.ServiceClient) (authKey string, id string) {
	fmt.Println("start volume transfer create...")
	opts := volumetransfer.CreateOpts{
		VolumeID: "xxx",
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

	authKey = resp.AuthKey
	id = resp.ID
	fmt.Println("create volume transfer success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

	return authKey, id

}

// accept volume transfer
func VolumeTransferAccept(sc *gophercloud.ServiceClient) {
	fmt.Println("start volume transfer accept...")
	opts := volumetransfer.AcceptOpts{
		AuthKey: "xxx",
	}
	id := "xxx"
	resp, err := volumetransfer.Accept(sc, id, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("accept volume transfer success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

// get volume transfer
func VolumeTransferGet(sc *gophercloud.ServiceClient) {
	fmt.Println("start volume transfer get...")
	id := "xxx"
	resp, err := volumetransfer.Get(sc, id).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("get volume transfer success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

}

// list volume transfer detail
func VolumeTransferListDetails(sc *gophercloud.ServiceClient) {
	fmt.Println("start list volume transfer details...")
	allPages, err := volumetransfer.ListDetail(sc, volumetransfer.ListDetailOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("list volume transfer details success!")
	allData, err1 := volumetransfer.ExtractTransferDetails(allPages)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	for _, v := range allData {
		fmt.Println(v.VolumeID)
		b, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(b))
	}

}

// list volume transfer brief
func VolumeTransferList(sc *gophercloud.ServiceClient) {
	fmt.Println("start list volume transfer...")
	volumePage, err := volumetransfer.List(sc, volumetransfer.ListDetailOpts{}).AllPages()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("list volume transfer success!")
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

}

// delete volume transfer
func VolumeTransferDelete(sc *gophercloud.ServiceClient) {
	fmt.Println("start volume transfer delete...")
	id := "xxx"
	err := volumetransfer.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete volume transfer success!")
}
