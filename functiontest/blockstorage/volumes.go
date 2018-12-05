package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
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

	//TestVolumesCreate(sc)
	//TestVolumesGet(sc)
	//TestVolumesListDetalis(sc)
	//TestVolumesUpdate(sc)
	TestVolumesDelete(sc)
	//TestVolumesDeleteCascade(sc)

	fmt.Println("main end...")
}
func TestVolumesCreate(sc *gophercloud.ServiceClient) {
	opts := volumes.CreateOpts{
		Size:             10,
		AvailabilityZone: "az1.dc1",
	}

	resp, err := volumes.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test volume create success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

func TestVolumesGet(sc *gophercloud.ServiceClient) {
	volume, err := volumes.Get(sc, "1a6d6c68-7ac6-4b08-8e87-089f5034c372").Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test volume get success!")
	b, _ := json.MarshalIndent(volume, "", " ")
	fmt.Println(string(b))

}

//func TestVolumesList(sc *gophercloud.ServiceClient) {
//	volume, err := volumes.List(sc).Extract()
//	if err != nil {
//		if ue, ok := err.(*gophercloud.UnifiedError); ok {
//			fmt.Println("ErrCode:", ue.ErrorCode())
//			fmt.Println("Message:", ue.Message())
//		}
//		return
//	}
//
//	fmt.Println("get volume success!")
//	fmt.Println("volume:", volume)
//}

func TestVolumesListDetalis(sc *gophercloud.ServiceClient) {
	allPages, err := volumes.List(sc, volumes.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test volume List success!")

	allData, _ := volumes.ExtractVolumes(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData {
		fmt.Println(v.UpdatedAt)
		b, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(b))
	}

}

func TestVolumesUpdate(sc *gophercloud.ServiceClient) {
	id := "1a6d6c68-7ac6-4b08-8e87-089f5034c372"

	updatOpts := volumes.UpdateOpts{
		Name:        "KAKAK EVS",
		Description: "new kaka EVS",
	}

	resp, err := volumes.Update(sc, id, updatOpts).Extract()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test volume update success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestVolumesDelete(sc *gophercloud.ServiceClient) {
	id := "1a6d6c68-7ac6-4b08-8e87-089f5034c372"
	err := volumes.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test volume delete success!")
}

func TestVolumesDeleteCascade(sc *gophercloud.ServiceClient) {
	id := "1a6d6c68-7ac6-4b08-8e87-089f5034c372"
	flag := false
	opts := volumes.DeleteOpts{
		Cascade: &flag,
	}
	err := volumes.DeleteCascade(sc, id, opts).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test volume delete success!")
}
