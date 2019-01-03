package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
)

var secgroupid string

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

	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	TestSecGroupList(sc)
	TestSecGroupCreate(sc)
	TestSecGroupGet(sc)
	TestSecGroupUpdate(sc)
	TestSecGroupDelete(sc)

	fmt.Println("main end...")
}

func TestSecGroupList(sc *gophercloud.ServiceClient) {
	allpages, err := groups.List(sc, groups.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	secgroups, err := groups.ExtractGroups(allpages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get securitygroup list success!")
	p, _ := json.MarshalIndent(secgroups, "", " ")
	fmt.Println(string(p))
}

func TestSecGroupGet(sc *gophercloud.ServiceClient) {
	secgroup, err := groups.Get(sc, secgroupid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get securitygroup detail success!")
	p, _ := json.MarshalIndent(secgroup, "", " ")
	fmt.Println(string(p))
}

func TestSecGroupUpdate(sc *gophercloud.ServiceClient) {
	opts := groups.UpdateOpts{
		Name:"testsecgroup2",
		Description:"Functiontest of SecGroup Update",
	}

	secgroup, err := groups.Update(sc, secgroupid,opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test update securitygroup success!")
	p, _ := json.MarshalIndent(secgroup, "", " ")
	fmt.Println(string(p))
}

func TestSecGroupCreate(sc *gophercloud.ServiceClient) {
	opts := groups.CreateOpts{
		Name:"testsecgroup",
		Description:"Functiontest of SecGroup",
	}
	secgroup, err := groups.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test create securitygroup success!")
	secgroupid = secgroup.ID
	p, _ := json.MarshalIndent(secgroup, "", " ")
	fmt.Println(string(p))
}

func TestSecGroupDelete(sc *gophercloud.ServiceClient) {
	err := groups.Delete(sc, secgroupid).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete securitygroup success!")
}
