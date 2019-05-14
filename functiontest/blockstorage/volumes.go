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
		return
	}

	//TestVolumesCreate(sc)
	//TestVolumesGet(sc)
	TestVolumesListBrief(sc)
	//TestVolumesUpdate(sc)
	//TestVolumesDelete(sc)
	//TestVolumesDeleteCascade(sc)
	TestVolumesList(sc)


	//TestVolumesMetadataCreate(sc)
	//TestVolumesMetadataGet(sc)
	//TestVolumesMetadataUpdate(sc)

	//TestVolumesMetadataGetKey(sc)
	//TestVolumesMetadataUpdateKey(sc)
	//TestVolumesMetadataDeleteKey(sc)
	//TestVolumesExtendSize(sc)
	//TestVolumesSetBootable(sc)
	//TestVolumesSetReadonly(sc)
	//TestVolumesExportAsImage(sc)
	TestVolumesGet(sc)

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
	volume, err := volumes.Get(sc, "0aefc7e2-207a-4e0b-8383-484997a6e59d").Extract()
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

func TestVolumesListBrief(sc *gophercloud.ServiceClient) {
	allPages, err := volumes.ListBrief(sc, volumes.ListBriefOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test volume List deatils success!")

	allData, _ := volumes.ExtractVolumesBrief(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData.VolumeList {
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

func TestVolumesList(sc *gophercloud.ServiceClient) {

	volumePage, err := volumes.List(sc, nil).AllPages()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	volumeList, err1 := volumes.ExtractVolumes(volumePage)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	for _, d := range volumeList {
		fmt.Println("volume id :", d.ID)
		fmt.Println("volume Name :", d.Name)
		fmt.Printf("%+v", d)
	}
	fmt.Println("get volume list  success!")

}

func TestVolumesMetadataCreate(sc *gophercloud.ServiceClient) {
	opts := volumes.MetadataOpts{
		Metadata: map[string]string{"kaka": "ads",},
	}

	volumeMetadata, err := volumes.CreateMetadata(sc, "e058f2ca-4338-461c-aa60-8e3718476fce", opts).ExtractMetadata()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println(" TestVolumesMetadataCreate success!")
	b, _ := json.MarshalIndent(volumeMetadata, "", " ")
	fmt.Println(string(b))

}

func TestVolumesMetadataGet(sc *gophercloud.ServiceClient) {
	volumeMetadata, err := volumes.GetMetadata(sc, "e058f2ca-4338-461c-aa60-8e3718476fce").ExtractMetadata()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestVolumesMetadataGet success!")
	b, _ := json.MarshalIndent(volumeMetadata, "", " ")
	fmt.Println(string(b))

}

func TestVolumesMetadataUpdate(sc *gophercloud.ServiceClient) {

	opts := volumes.MetadataOpts{
		Metadata: map[string]string{"kaka": "adass",},
	}

	volumeMetadata, err := volumes.UpdateMetadata(sc, "e058f2ca-4338-461c-aa60-8e3718476fce", opts).ExtractMetadata()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestVolumesMetadataUpdate success!")
	b, _ := json.MarshalIndent(volumeMetadata, "", " ")
	fmt.Println(string(b))

}

func TestVolumesMetadataGetKey(sc *gophercloud.ServiceClient) {
	volume, err := volumes.GetMetadataKey(sc, "e058f2ca-4338-461c-aa60-8e3718476fce", "kaka").ExtractMeta()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestVolumesMetadataGetKey success!")
	b, _ := json.MarshalIndent(volume, "", " ")
	fmt.Println(string(b))

}
func TestVolumesMetadataDeleteKey(sc *gophercloud.ServiceClient) {
	err := volumes.DeleteMetadataKey(sc, "e058f2ca-4338-461c-aa60-8e3718476fce", "kaka").ExtractErr()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestVolumesMetadataDeleteKeysuccess!")

}
func TestVolumesMetadataUpdateKey(sc *gophercloud.ServiceClient) {

	opts := volumes.MetaOpts{
		Meta: map[string]string{"kaka": "asdf"},
	}
	volume, err := volumes.UpdateMetadataKey(sc, "e058f2ca-4338-461c-aa60-8e3718476fce", "kaka", opts).ExtractMeta()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestVolumesMetadataUpdateKey success!")
	b, _ := json.MarshalIndent(volume, "", " ")
	fmt.Println(string(b))

}

func TestVolumesExtendSize(sc *gophercloud.ServiceClient) {
	opts := volumes.ExtendSizeOpts{
		NewSize: 50,
	}

	id := "e058f2ca-4338-461c-aa60-8e3718476fce"
	err := volumes.ExtendSize(sc, id, opts).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test volume TestVolumesExtendSize success!")
}

func TestVolumesSetBootable(sc *gophercloud.ServiceClient) {
	opts := volumes.SetBootableOpts{
		Bootable: true,
	}

	id := "e058f2ca-4338-461c-aa60-8e3718476fce"
	err := volumes.SetBootable(sc, id, opts).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test volume TestVolumesSetBootable success!")
}

func TestVolumesSetReadonly(sc *gophercloud.ServiceClient) {
	opts := volumes.SetReadOnlyOpts{
		ReadOnly: true,
	}

	id := "e058f2ca-4338-461c-aa60-8e3718476fce"
	err := volumes.SetReadOnly(sc, id, opts).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test volume TestVolumesSetReadonly success!")
}

func TestVolumesExportAsImage(sc *gophercloud.ServiceClient) {
	opts := volumes.ExportVolumesOpts{
		ImageName: "laa",
	}

	id := "e058f2ca-4338-461c-aa60-8e3718476fce"
	resp, err := volumes.ExportVolumes(sc, id, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println(resp)
	fmt.Println("Test volume TestVolumesExportAsImage success!")
}
