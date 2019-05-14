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
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
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

	fmt.Println("main end...")
}

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
