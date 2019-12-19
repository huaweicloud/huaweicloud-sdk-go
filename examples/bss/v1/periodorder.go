package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/periodorder"
)

func main() {
	//AKSK auth，initial parameter.
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		DomainID:         "{domainID}",
	}

	//initial provider client。
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("get provider client failed")
		fmt.Println(errAuth.Error())
		return
	}

	// initial client
	sc, err := openstack.NewBSSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	QueryOrderList(sc)

	QueryOrderDetail(sc)

	PayPeriodOrder(sc)

	UnsubscribePeriodOrder(sc)

	QueryResourceStatusByOrderId(sc)

	CancelOrder(sc)

	QueryRefundOrderAmount(sc)
}

func QueryOrderList(client *gophercloud.ServiceClient) {
	var a = 10
	var b = 1
	opts := periodorder.QueryOrderListOpts{
		PageSize:  &a,
		PageIndex: &b,
	}

	periodorder.QueryOrderList(client, opts)
}

func QueryOrderDetail(client *gophercloud.ServiceClient) {
	opts := periodorder.QueryOrderDetailOpts{
		Offset: 1,
		Limit:  10,
	}

	orderId := "CS1908291151YAFGJ"
	periodorder.QueryOrderDetail(client, opts, orderId)
}

func PayPeriodOrder(client *gophercloud.ServiceClient) {
	var a = 1
	opts := periodorder.PayPeriodOrderOpts{
		OderId:         "CS1712271317IT8C4",
		PayAccountType:  &a,
		BpId:           "061013372c00d3410f3fc017ee1e8ac0",
		CouponIds:      nil,
	}

	periodorder.PayPeriodOrder(client, opts)
}

func UnsubscribePeriodOrder(client *gophercloud.ServiceClient) {
	opts := periodorder.UnsubscribePeriodOrderOpts{
		UnsubType:             4,
		UnsubscribeReasonType: 5,
		UnsubscribeReason:     " ",
	}

	orderId := "CS1908291151YAFGJ"
	periodorder.UnsubscribePeriodOrder(client, opts, orderId)
}

func QueryResourceStatusByOrderId(client *gophercloud.ServiceClient) {
	opts := periodorder.QueryResourceStatusByOrderIdOpts{
		Offset: 1,
		Limit:  10,
	}

	orderId := "CS1908291151YAFGJ"
	periodorder.QueryResourceStatusByOrderId(client, opts, orderId)
}

func CancelOrder(client *gophercloud.ServiceClient) {
	actionId := "cancel"
	opts := periodorder.CancelOrderOpts{
		OrderId:  "CS1908291151YAFGJ",
	}

	periodorder.CancelOrder(client, opts,actionId)
}

func QueryRefundOrderAmount(client *gophercloud.ServiceClient) {
	opts := periodorder.QueryRefundOrderAmountOpts{
		OrderId: "CS1908291151YAFGJ",
	}

	periodorder.QueryRefundOrderAmount(client, opts)
}
