package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bssintl/v1/customercredit"
)

func main() {
	fmt.Println("customercredit start...")

	//打开debug日志
	gophercloud.EnableDebug = true

	provider, err := common.AuthToken()
	//provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("auth success!")

	// 初始化服务的client
	sc, err := openstack.NewBSSIntlV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	TestQueryCredit(sc)
	TestSetCredit(sc)

	fmt.Println("customercredit end...")
}

func TestQueryCredit(client *gophercloud.ServiceClient) {
	opts := customercredit.QueryCreditOpts{
		CustomerId: "123",
	}
	detailRsp,err := customercredit.QueryCredit(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(detailRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestQueryCredit success")
}

func TestSetCredit(client *gophercloud.ServiceClient) {
	var a = 1
	opts := customercredit.SetCreditOpts{
		CustomerId:       "123",
		AdjustmentAmount: 1,
		MeasureId:        &a,
	}
	detailRsp,err := customercredit.SetCredit(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(detailRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("SetCredit success")
}
