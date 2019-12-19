package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bssintl/v1/bill"
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
	sc, err := openstack.NewBSSIntlV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}
	QueryPartnerMonthlyBills(sc)
	QueryMonthlyExpenditureSummary(sc)
	QueryResourceUsageDetails(sc)
	QueryResourceUsageRecord(sc)
}

func QueryPartnerMonthlyBills(client *gophercloud.ServiceClient) {
	opts := bill.QueryPartnerMonthlyBillsOpts{
		ConsumeMonth: "2019-08",
	}
	detailRsp,err := bill.QueryPartnerMonthlyBills(client, opts).Extract()

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
	fmt.Println("PostPaidBill success")
}

func QueryMonthlyExpenditureSummary(client *gophercloud.ServiceClient) {
	opts := bill.QueryMonthlyExpenditureSummaryOpts{
		Cycle:                "2018-05",
	}
	detailRsp,err := bill.QueryMonthlyExpenditureSummary(client, opts).Extract()

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
	fmt.Println("CustomerSum success")
}

func QueryResourceUsageDetails(client *gophercloud.ServiceClient) {
	opts := bill.QueryResourceUsageDetailsOpts{
		Cycle:                "2018-05",
	}
	detailRsp,err := bill.QueryResourceUsageDetails(client, opts).Extract()

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
	fmt.Println("CustomerSum success")
}

func QueryResourceUsageRecord(client *gophercloud.ServiceClient) {
	opts := bill.QueryResourceUsageRecordOpts{
		StartTime:            "2019-08-01",
		EndTime:              "2019-08-31",
		CloudServiceTypeCode: "hws.service.type.ebs",
		RegionCode:           "cn-north-1",
		OrderId:              "orderId",
		PayMethod:            "0",
		Offset:               1,
		Limit:                10,
		ResourceId:           "hws.service.type.ebs",
		EnterpriseProjectId:  "pjzV4N9Uq1LWMUgh3fYAhJqUbtd6Ad8yALoRJeDoYGl0jWdZoS8UHbcGBqEin1Ia",
	}
	detailRsp,err := bill.QueryResourceUsageRecord(client, opts).Extract()

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
	fmt.Println("CustomerSum success")
}