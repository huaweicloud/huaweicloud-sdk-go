package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bssintl/v1/enquiry"
)

func main() {
	fmt.Println("enquiry start...")

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

	TestQueryRating(sc)
	fmt.Println("enquiry end...")
}

func TestQueryRating(client *gophercloud.ServiceClient) {
	var a = 0
	opts := enquiry.QueryRatingOpts{
		TenantId:                   "74610f3a5ad941998e91f076297ecf27",
		RegionId:                   "cn-north-1",
		AvaliableZoneId:            "cn-north-1",
		ChargingMode:               &a,
		PeriodType:                 1,
		PeriodNum:                  10,
		PeriodEndDate:              "",
		RelativeResourceId:         "546568dsdcsc",
		RelativeResourcePeriodType: 1,
		SubscriptionNum:            10,
		ProductInfo:                nil,
		InquiryTime:                "",
	}

	enquiry.QueryRating(client, opts)
}
