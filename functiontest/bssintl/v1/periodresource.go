package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bssintl/v1/periodresource"
)

func main() {
	fmt.Println("periodresource start...")

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
	TestQueryCustomerPeriodResourcesList(sc)
	TestRenewSubscriptionByResourceId(sc)
	TestUnsubscribeByResourceId(sc)
	TestEnableAutoRenew(sc)
	TestDisableAutoRenew(sc)
	fmt.Println("periodresource end...")
}

func TestQueryCustomerPeriodResourcesList(client *gophercloud.ServiceClient) {
	var a = 2
	var b = 1
	var c = 0
	opts := periodresource.QueryCustomerPeriodResourcesListOpts{
		ResourceIds:      "abc",
		OrderId:          "abc",
		OnlyMainResource: &a,
		StatusList:       "2,3,4,5,6",
		PageNo:           &b,
		PageSize:         &c,
	}
	detailRsp,err := periodresource.QueryCustomerPeriodResourcesList(client, opts).Extract()

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
	fmt.Println("TestDetail success")
}

func TestRenewSubscriptionByResourceId(client *gophercloud.ServiceClient) {
	var a = 2
	var b = 1
	var c= 0
	opts := periodresource.RenewSubscriptionByResourceIdOpts{
		ResourceIds: []string{"123"},
		PeriodType:   &a,
		PeriodNum:   &b,
		ExpireMode:  &b,
		IsAutoPay:   &c,
	}
	detailRsp,err := periodresource.RenewSubscriptionByResourceId(client, opts).Extract()

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
	fmt.Println("TestRenew success")
}

func TestUnsubscribeByResourceId(client *gophercloud.ServiceClient) {
	var b = 1
	var c= 0
	opts := periodresource.UnsubscribeByResourceIdOpts{
		ResourceIds:           []string{"123"},
		UnSubType:             &b,
		UnsubscribeReasonType: &c,
		UnsubscribeReason:     "",
	}
	detailRsp,err := periodresource.UnsubscribeByResourceId(client, opts).Extract()

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
	fmt.Println("TestDelete success")
}

func TestEnableAutoRenew(client *gophercloud.ServiceClient) {
	resourceId := "123"
	opts := periodresource.EnableAutoRenewOpts{
		ActionId:   "123",
	}
	detailRsp,err := periodresource.EnableAutoRenew(client, opts,resourceId).Extract()

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
	fmt.Println("TestAutoRenew success")
}

func TestDisableAutoRenew(client *gophercloud.ServiceClient) {
	resourceId := "123"
	opts := periodresource.DisableAutoRenewOpts{
		ActionId:   "123",
	}
	detailRsp,err := periodresource.DisableAutoRenew(client, opts,resourceId).Extract()

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
	fmt.Println("TestDeleteAutoRenew success")
}

