package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

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

	TestGet(sc)
	//TestList(sc)

	fmt.Println("main end...")
}

func TestGet(sc *gophercloud.ServiceClient) {
	image, err := images.Get(sc, "48a6b63f-4e19-45c7-8cb8-113a791acaa8").Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("get image detail success!")
	fmt.Println("Self:", image.Self)
	fmt.Println("Deleted:", image.Deleted)
	fmt.Println("DeletedAt:", image.DeletedAt)
	fmt.Println("VirtualEnvType:", image.VirtualEnvType)
}

func TestList(sc *gophercloud.ServiceClient) {
	allPages, err := images.List(sc, images.ListOpts{Status: "ERROR"}).AllPages()
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

	fmt.Println(allImages)
}
