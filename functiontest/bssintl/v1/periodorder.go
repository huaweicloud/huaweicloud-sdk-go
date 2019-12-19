package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bssintl/v1/periodorder"
)

func main() {
	fmt.Println("periodorder start...")

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
	TestQueryOrderList(sc)

	TestQueryOrderDetail(sc)

	TestPayPeriodOrder(sc)

	TestUnsubscribePeriodOrder(sc)

	TestQueryResourceStatusByOrderId(sc)

	TestCancelOrder(sc)

	TestQueryRefundOrderAmount(sc)
	fmt.Println("periodorder end...")
}

func TestQueryOrderList(client *gophercloud.ServiceClient) {
	var a = 10
	var b = 1
	opts := periodorder.QueryOrderListOpts{
		PageSize:  &a,
		PageIndex: &b,
	}

	periodorder.QueryOrderList(client, opts)
}

func TestQueryOrderDetail(client *gophercloud.ServiceClient) {
	opts := periodorder.QueryOrderDetailOpts{
		Offset: 1,
		Limit:  10,
	}

	orderId := "CS1908291151YAFGJ"
	periodorder.QueryOrderDetail(client, opts, orderId)
}

func TestPayPeriodOrder(client *gophercloud.ServiceClient) {
	var a = 1
	opts := periodorder.PayPeriodOrderOpts{
		OderId:         "CS1712271317IT8C4",
		PayAccountType: &a,
		BpId:           "061013372c00d3410f3fc017ee1e8ac0",
		CouponIds:      nil,
	}

	periodorder.PayPeriodOrder(client, opts)
}

func TestUnsubscribePeriodOrder(client *gophercloud.ServiceClient) {
	opts := periodorder.UnsubscribePeriodOrderOpts{
		UnsubType:             4,
		UnsubscribeReasonType: 5,
		UnsubscribeReason:     "test",
	}

	orderId := "CS1908291151YAFGJ"
	periodorder.UnsubscribePeriodOrder(client, opts, orderId)
}

func TestQueryResourceStatusByOrderId(client *gophercloud.ServiceClient) {
	opts := periodorder.QueryResourceStatusByOrderIdOpts{
		Offset: 1,
		Limit:  10,
	}

	orderId := "CS1908291151YAFGJ"
	periodorder.QueryResourceStatusByOrderId(client, opts, orderId)
}

func TestCancelOrder(client *gophercloud.ServiceClient) {
	actionId := "cancel"
	opts := periodorder.CancelOrderOpts{
		OrderId:  "CS1908291151YAFGJ",
	}

	periodorder.CancelOrder(client, opts,actionId)
}

func TestQueryRefundOrderAmount(client *gophercloud.ServiceClient) {
	opts := periodorder.QueryRefundOrderAmountOpts{
		OrderId: "CS1908291151YAFGJ",
	}

	periodorder.QueryRefundOrderAmount(client, opts)
}
