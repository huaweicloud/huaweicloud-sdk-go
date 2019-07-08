package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"

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

	VolumeCreate(sc)
	VolumesList(sc)
	VolumeGet(sc)
	VolumeUpdate(sc)
	VolumeDelete(sc)
	VolumesListBrief(sc)
	GetQuotaSet(sc)
	VolumesExportAsImage(sc)

	fmt.Println("main end...")
}

// create volume
func VolumeCreate(sc *gophercloud.ServiceClient) (volumeId string) {
	fmt.Println("start volume create...")
	createOpts := volumes.CreateOpts{
		AvailabilityZone: "az1.dc1",
		Description:      "volumeDescription",
		Size:             10,
		Name:             "volumeName",
		VolumeType:       "SATA",
	}
	volume, err := volumes.Create(sc, createOpts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("create volume success!")
	volumeId = volume.ID
	p, err := json.MarshalIndent(volume, "", "   ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(p))
	return volumeId
}

// get volume
func VolumeGet(sc *gophercloud.ServiceClient) {
	id := "xxx"
	volume, err := volumes.Get(sc, id).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("get volume detail success!")

	p, _ := json.MarshalIndent(volume, "", " ")
	fmt.Println(string(p))
}

// update volume
func VolumeUpdate(sc *gophercloud.ServiceClient) {
	id := "xxx"
	updatOpts := volumes.UpdateOpts{
		Name:        "volumeNameNew",
		Description: "volumeDescriptionNew",
	}

	resp, err := volumes.Update(sc, id, updatOpts).Extract()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("volume update success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

// delete volume
func VolumeDelete(sc *gophercloud.ServiceClient) {
	id := "xxx"
	err := volumes.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("volume delete success!")
}

// list volume detail
func VolumesList(sc *gophercloud.ServiceClient) {

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

// list volume brief
func VolumesListBrief(sc *gophercloud.ServiceClient) {
	fmt.Println("start volume list...")
	allPages, err := volumes.ListBrief(sc, volumes.ListBriefOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("list volume success!")

	allData, err1 := volumes.ExtractVolumesBrief(allPages)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	for _, v := range allData.VolumeList {
		b, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(b))
	}

}

// get tanant quota
func GetQuotaSet(sc *gophercloud.ServiceClient) {
	fmt.Println("start get quota set...")
	var vLimit, vInUse, gLimit, gInUse int
	projectId := "xxx"
	qs, err := volumes.GetQuotaSet(sc, projectId)
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	var bt volumes.BaseType

	for k, v := range qs.QuoSet {
		if k == "gigabytes" {
			if data, ok := v.(map[string]interface{}); ok {
				b, err := json.Marshal(data)

				if err != nil {
					return
				}
				err = json.Unmarshal(b, &bt)
				if err != nil {
					return
				}
				gLimit = bt.Limit
				gInUse = bt.InUse
			}
		}

		if k == "volumes" {
			if data, ok := v.(map[string]interface{}); ok {
				b, err := json.Marshal(data)

				if err != nil {
					return
				}
				err = json.Unmarshal(b, &bt)
				if err != nil {
					return
				}
				vLimit = bt.Limit
				vInUse = bt.InUse
			}
		}

	}

	fmt.Println(vLimit, vInUse, gLimit, gInUse)

	for k, v := range qs.QuoSet {
		fmt.Println("type is :", k)
		if data, ok := v.(map[string]interface{}); ok {
			for b1, b2 := range data {
				fmt.Println(b1, b2)
			}
		} else {
			fmt.Println("value is ", v)
		}
	}

	fmt.Println("get quota set success!")

}

// export volume to image
func VolumesExportAsImage(sc *gophercloud.ServiceClient) {
	fmt.Println("start volume export image...")
	opts := volumes.ExportVolumesOpts{
		ImageName: "xxx",
	}

	id := "xxx"
	resp, err := volumes.ExportVolumes(sc, id, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("volume export image success!")
	fmt.Println(resp)
}
