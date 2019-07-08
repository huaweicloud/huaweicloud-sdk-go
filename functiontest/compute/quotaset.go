package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/quotasets"
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

	TestGetQuotaSetLimit(sc)
	TestGetQuotaSetDefault(sc)
	TestGetQuotaSet(sc)

	fmt.Println("main end...")
}

func TestGetQuotaSetLimit(sc *gophercloud.ServiceClient) {

	resp, err := quotasets.GetLimits(sc).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server get QuotaSetLimit success!")

	p, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(p))

}

func TestGetQuotaSet(sc *gophercloud.ServiceClient) {
	id := "128a7bf965154373a7b73c89eb6b65aa"

	resp, err := quotasets.Get(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server get QuotaSet success!")
	b,_:=json.MarshalIndent(*resp,""," ")
	fmt.Println(string(b))

}

func TestGetQuotaSetDefault(sc *gophercloud.ServiceClient) {
	id := "128a7bf965154373a7b73c89eb6b65aa"
	resp, err := quotasets.GetDefault(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test server get QuotaSetDefault success!")

	p, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(p))

}
