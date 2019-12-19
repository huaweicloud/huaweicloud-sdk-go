package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/periodresource"
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
	QueryCustomerPeriodResourcesList(sc)
	RenewSubscriptionByResourceId(sc)
	UnsubscribeByResourceId(sc)
	EnableAutoRenew(sc)
	DisableAutoRenew(sc)
}

func QueryCustomerPeriodResourcesList(client *gophercloud.ServiceClient) {
	var a = 10
	var b = 1
	var c = 0
	opts := periodresource.QueryCustomerPeriodResourcesListOpts{
		ResourceIds:      "abc",
		OrderId:          "abc",
		OnlyMainResource: &c,
		StatusList:       "2,3,4,5,6",
		PageNo:           &b,
		PageSize:         &a,
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
	fmt.Println("Detail success")
}

func RenewSubscriptionByResourceId(client *gophercloud.ServiceClient) {
	var a = 2
	var b = 1
	var c = 0
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
	fmt.Println("Renew success")
}

func UnsubscribeByResourceId(client *gophercloud.ServiceClient) {
	var b = 1
	var c = 0
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
	fmt.Println("Delete success")
}

func EnableAutoRenew(client *gophercloud.ServiceClient) {
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
	fmt.Println("AutoRenew success")
}

func DisableAutoRenew(client *gophercloud.ServiceClient) {
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
	fmt.Println("DeleteAutoRenew success")
}

