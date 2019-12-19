package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/account"
)

func main() {
	//AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	//初始化provider client。
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
		return
	}

	// 初始化服务的client
	sc, err := openstack.NewBSSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	startTime := "2018-06-01"
	endTime := "2018-06-30"

	ListResourceDailyAll(sc, startTime, endTime)

	customerId := ""
	pageNo := 1
	pageSize := 100
	ListResourceDetailPage(sc, customerId, pageNo, pageSize)
}

func ListResourceDailyAll(client *gophercloud.ServiceClient, startTime string, endTime string) {
	//初始化查询参数
	reqTmp := account.ResourceDailyOpts{
		StartTime:           startTime,
		EndTime:             endTime,
		PayMethod:           "0",
		CloudServiceType:    "hws.service.type.ebs",
		RegionCode:          "cn-xianhz-1",
		ResourceId:          "",
		EnterpriseProjectId: "",
	}

	//根据查询参数获取消费汇总列表
	rspTmp, err := account.ListResourceDaily(client, reqTmp)
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//打印返回参数，如currency，totalRecord，totalAmount，dailyRecords等
	fmt.Println("Succeed to get the ResourceDaily List!")
	fmt.Println("totalRecord:", rspTmp.TotalRecord)
	fmt.Println("currency:", rspTmp.Currency)
	fmt.Println("totalAmount:", rspTmp.TotalAmount)
}
