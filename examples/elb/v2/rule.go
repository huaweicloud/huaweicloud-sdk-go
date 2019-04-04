package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"encoding/json"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/policies"
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
		fmt.Println("Failed to get the provider: ", err_auth)
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

	policyid := "5cffadf0-ebda-431a-8e98-cdb218a1a221"

	ruleId := PolicyRulesCreate(sc, policyid)
	PolicyRulesList(sc, policyid)
	PolicyRulesGet(sc, policyid, ruleId)
	PolicyRulesUpdate(sc, policyid, ruleId)
	PolicyRulesDelete(sc, policyid, ruleId)

	fmt.Println("main end...")
}

func PolicyRulesList(sc *gophercloud.ServiceClient, policyid string) {
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

func PolicyRulesGet(sc *gophercloud.ServiceClient, policyid string, policyruleid string) {
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

func PolicyRulesCreate(sc *gophercloud.ServiceClient, policyid string) (policyruleid string) {
	opts := policies.CreateRuleOpts{
		CompareType: "EQUAL_TO",
		Type:        "PATH",
		Value:       "/test",
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
	return policyruleid
}

func PolicyRulesUpdate(sc *gophercloud.ServiceClient, policyid string, policyruleid string) {
	opts := policies.RuleUpdateOpts{
		CompareType: "EQUAL_TO",
		Value:       "/test-path",
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

func PolicyRulesDelete(sc *gophercloud.ServiceClient, policyid string, policyruleid string) {
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