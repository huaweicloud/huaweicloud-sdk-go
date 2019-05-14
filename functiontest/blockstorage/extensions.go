package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/extensions/apiversions"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/extensions/extensions"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/extensions/availabilityzones"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/extensions/volumetypes"
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

	TestGetAvailabilityZone(sc)
	TestGetExtensions(sc)
	TestTypeList(sc)
	TestTypeGet(sc)
	TestGetAPIList(sc)
	TestGetAPIVersion(sc)
	fmt.Println("main end...")
}

func TestGetAvailabilityZone(sc *gophercloud.ServiceClient) {

	resp, err := availabilityzones.List(sc).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	availabilityZones, err := availabilityzones.ExtractAvailabilityZones(resp)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(availabilityZones, "", " ")
	fmt.Println(string(b))

	for _, d := range availabilityZones {
		fmt.Printf("%+v\n", d)
	}

	fmt.Println("Test TestGetAvailabilityZone success!")

}

func TestGetExtensions(sc *gophercloud.ServiceClient) {

	resp, err := extensions.List(sc).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	extension, err := extensions.ExtractExtensions(resp)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(extension, "", " ")
	fmt.Println(string(b))

	for _, d := range extension {
		fmt.Printf("%+v\n", d)
	}

	fmt.Println("Test TestGetExtensions success!")
}

func TestTypeList(sc *gophercloud.ServiceClient) {
	opts := volumetypes.ListOpts{

	}

	resp, err := volumetypes.List(sc, opts).AllPages()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	types, err := volumetypes.ExtractVolumeTypes(resp)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(types, "", " ")
	fmt.Println(string(b))

	for _, d := range types {
		fmt.Printf("%+v", d)
	}

	fmt.Println("Test TestTypeList success!")
}

func TestTypeGet(sc *gophercloud.ServiceClient) {

	typedata, err := volumetypes.Get(sc, "8256dff6-deb5-4fdd-a080-ce9185a39b16").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(typedata, "", " ")
	fmt.Println(string(b))

	fmt.Println("Test TestTypeGet success!")
}

func TestGetAPIList(sc *gophercloud.ServiceClient) {

	apiversionsPage, err := apiversions.List(sc).AllPages()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	apiversionList, err1 := apiversions.ExtractAPIVersions(apiversionsPage)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	for _, d := range apiversionList {
		fmt.Printf("%+v\n", d)
	}
	fmt.Println("get TestGetAPIList success!")

}

func TestGetAPIVersion(sc *gophercloud.ServiceClient) {
	id := "v3.0"
	apiversionData, err := apiversions.Get(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(apiversionData, "", " ")
	fmt.Println(string(b))

	fmt.Println("Test TestGetAPIVersion success!")

}
