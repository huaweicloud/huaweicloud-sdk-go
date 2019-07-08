package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/volumeattach"
)

func main() {

	fmt.Println("main start...")

	//provider, err := common.AuthAKSK()
	provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get compute v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//TestTestGetVolumeAttachList(sc)
	//TestGetVolumeAttached(sc)
	//TestDeleteVolumeAttached(sc)
	TestDetachVolumeWithFlag(sc)

	fmt.Println("main end...")
}

func TestTestGetVolumeAttachList(sc *gophercloud.ServiceClient) {
	serverId := "4a5e7286-8da5-4bf8-9658-88f5fc604d2e"
	page, err := volumeattach.List(sc, serverId).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server get VolumeAttached success!")

	volumeList, err := volumeattach.ExtractVolumeAttachments(page)

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	p, _ := json.MarshalIndent(volumeList, "", " ")
	fmt.Println(string(p))

}

func TestGetVolumeAttached(sc *gophercloud.ServiceClient) {
	id := "609bdc81-b131-4886-8d02-160efae60185"
	serverId := "4a5e7286-8da5-4bf8-9658-88f5fc604d2e"
	resp, err := volumeattach.Get(sc, serverId, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server get VolumeAttached success!")
	b, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(b))

}

func TestDeleteVolumeAttached(sc *gophercloud.ServiceClient) {
	id := "128a7bf965154373wea7b73c89eb6b65aa"
	serverId := "4a5e7286-8da5-4bf8-9658-88f5fc604d2e"
	err := volumeattach.Delete(sc, serverId, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server del VolumeAttached success!")

}

//TestDetachVolumeWithFlag tests detachment of volume with delete flag
func TestDetachVolumeWithFlag(sc *gophercloud.ServiceClient) {
	volumeID := "{volumeID}"
	serverID := "{serverID}"
	err := volumeattach.DeleteWithFlag(sc, serverID, volumeID, 0).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server detach volume with delete flag success!")

}
