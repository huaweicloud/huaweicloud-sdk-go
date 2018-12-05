package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/policies"
)

var policyid string
var policyruleid string

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	//provider, err := common.AuthToken()
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

	TestPolicyList(sc)
	TestPolicyCreate(sc)
	TestPolicyGet(sc)
	TestPolicyUpdate(sc)
	TestPolicyRulesCreate(sc)
	TestPolicyRulesList(sc)
	TestPolicyRulesUpdate(sc)
	TestPolicyRulesGet(sc)
	TestPolicyRulesDelete(sc)
	TestPolicyDelete(sc)
	fmt.Println("main end...")
}

func TestPolicyCreate(sc *gophercloud.ServiceClient) {
	opts := policies.CreateOpts{
		Name:           "asd",
		RedirectPoolID: "a5a0c67e-a62c-4ffe-a094-795a879af926",
		ListenerID:     "cdc1a958-d9de-4d8f-b728-4a7867e723f9",
		Action:         "REDIRECT_TO_POOL",
	}

	resp, err := policies.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test create policy success!")
	policyid = resp.ID
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
}

func TestPolicyList(sc *gophercloud.ServiceClient) {
	allPages, err := policies.List(sc, policies.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get policy list success!")
	allData, err := policies.ExtractPolcies(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}
	p, _ := json.MarshalIndent(allData, "", " ")
	fmt.Println(string(p))
}

func TestPolicyGet(sc *gophercloud.ServiceClient) {
	resp, err := policies.Get(sc, policyid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	fmt.Println("Test get policy detail success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
}
func TestPolicyUpdate(sc *gophercloud.ServiceClient) {
	updatOpts := policies.UpdateOpts{
		Description: "asdddddddddddddd",
	}

	resp, err := policies.Update(sc, policyid, updatOpts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	fmt.Println("Test update policy success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
}

func TestPolicyDelete(sc *gophercloud.ServiceClient) {
	err := policies.Delete(sc, policyid).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete policy success!")
}

func TestPolicyRulesList(sc *gophercloud.ServiceClient) {
	allPages, err := policies.ListRules(sc, policies.RulesListOpts{}, policyid).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get policyrules List success!")
	policyrules, err := policies.ExtractPolicyRules(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	p, _ := json.MarshalIndent(policyrules, "", " ")
	fmt.Println(string(p))
}

func TestPolicyRulesGet(sc *gophercloud.ServiceClient) {
	policyrule, err := policies.GetRule(sc, policyid, policyruleid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get policyrule detail success!")
	p, _ := json.MarshalIndent(policyrule, "", " ")
	fmt.Println(string(p))
}

func TestPolicyRulesCreate(sc *gophercloud.ServiceClient) {
	opts := policies.CreateRuleOpts{
		CompareType: "EQUAL_TO",
		Type:        "HOST_NAME",
		Value:       "test.com",
	}
	policyrule, err := policies.CreateRule(sc, opts, policyid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test create policyrule success!")
	policyruleid = policyrule.Rule.ID
	p, _ := json.MarshalIndent(policyrule, "", " ")
	fmt.Println(string(p))
}

func TestPolicyRulesDelete(sc *gophercloud.ServiceClient) {
	err := policies.DeleteRule(sc, policyid, policyruleid).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete policyrule success!")
}

func TestPolicyRulesUpdate(sc *gophercloud.ServiceClient) {
	opts := policies.RuleUpdateOpts{
		CompareType: "EQUAL_TO",
		Value:       "cc.com",
	}

	policyrule, err := policies.UpdateRule(sc, opts, policyid, policyruleid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test update policyrule success!")
	p, _ := json.MarshalIndent(policyrule, "", " ")
	fmt.Println(string(p))
}
