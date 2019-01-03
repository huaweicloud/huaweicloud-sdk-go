package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/imagedata"
)

var imageid string

func main() {
	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	//provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}

	//获取ims服务客户端
	sc, err := openstack.NewIMSV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ims client failed")
		fmt.Println(err.Error())
		return
	}

	TestList(sc)
	TestCreate(sc)
	TestUpdate(sc)
	TestGet(sc)
	TestUpload(sc)
	TestDownload(sc)
	TestDelete(sc)

	fmt.Println("main end...")
}

func TestGet(sc *gophercloud.ServiceClient) {
	image, err := images.Get(sc, imageid).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get image detail success!")
	fmt.Println("Self:", image.Self)
	fmt.Println("Deleted:", image.Deleted)
	fmt.Println("DeletedAt:", image.DeletedAt)
	fmt.Println("VirtualEnvType:", image.VirtualEnvType)
}

func TestList(sc *gophercloud.ServiceClient) {
	allPages, err := images.List(sc, images.ListOpts{}).AllPages()
	if err != nil {
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Test get image list success!")
	p, _ := json.MarshalIndent(allImages, "", " ")
	fmt.Println(string(p))
}

func TestUpdate(sc *gophercloud.ServiceClient) {
	opts := images.UpdateOpts{
		images.ReplaceImageName{NewName: "testupdateimage"},
	}

	image, err := images.Update(sc, imageid, opts).Extract()
	if err != nil {
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	fmt.Println("Test update image success!")
	p, _ := json.MarshalIndent(image, "", " ")
	fmt.Println(string(p))
}

func TestCreate(sc *gophercloud.ServiceClient) {
	opts := images.CreateOpts{
		Name:"testcreateimage",
	}

	image, err := images.Create(sc, opts).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	imageid = image.ID
	fmt.Println("Test create image success!")
	fmt.Println(image.ID)
}

func TestDelete(sc *gophercloud.ServiceClient) {
	err := images.Delete(sc, imageid).ExtractErr()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete image success!")
}

func TestUpload(sc *gophercloud.ServiceClient) {
	err := imagedata.Upload(sc, imageid,bytes.NewReader([]byte{5, 3, 7, 24})).ExtractErr()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test upload image success!")
}

func TestDownload(sc *gophercloud.ServiceClient) {
	irder,err := imagedata.Download(sc, imageid).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bs, err := ioutil.ReadAll(irder)
	err =ioutil.WriteFile("./testimages.qcow2",bs,0777)
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test download image success!")
	os.Remove("./testimages.qcow2")
}

