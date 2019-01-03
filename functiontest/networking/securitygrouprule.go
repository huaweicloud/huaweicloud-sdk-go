package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
)

var secgroupruleid string

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
		fmt.Println("get Network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	TestSecGroupRuleList(sc)
	TestSecGroupRuleCreate(sc)
	TestSecGroupRuleGet(sc)
	TestSecGroupRuleDelete(sc)

	fmt.Println("main end...")
}


func TestSecGroupRuleList(sc *gophercloud.ServiceClient) {
	allpages, err := rules.List(sc, rules.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	secgrouprules, err := rules.ExtractRules(allpages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get securitygrouprule list success!")
	p, _ := json.MarshalIndent(secgrouprules, "", " ")
	fmt.Println(string(p))
}

func TestSecGroupRuleGet(sc *gophercloud.ServiceClient) {
	secgrouprule, err := rules.Get(sc, secgroupruleid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get securitygrouprule detail success!")
	p, _ := json.MarshalIndent(secgrouprule, "", " ")
	fmt.Println(string(p))
}

func TestSecGroupRuleCreate(sc *gophercloud.ServiceClient) {
	opts := rules.CreateOpts{
		Direction:"ingress",
		EtherType:"IPv4",
		SecGroupID:"ee995a60-ea20-42cb-b2e6-470c38cffb95",
	}
	secgrouprule, err := rules.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test create securitygrouprule success!")
	secgroupruleid = secgrouprule.ID
	p, _ := json.MarshalIndent(secgrouprule, "", " ")
	fmt.Println(string(p))
}

func TestSecGroupRuleDelete(sc *gophercloud.ServiceClient) {
	err := rules.Delete(sc, secgroupruleid).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete securitygrouprule success!")
}
