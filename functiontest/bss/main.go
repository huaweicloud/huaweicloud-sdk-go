package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/resource"
	"math"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/account"
)

func main() {
	fmt.Println("main start...")

	//provider, err := common.AuthToken()
	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}

	// 初始化服务的client
	sc, err := openstack.NewBSSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	customerId := "477d81b6aab64cc3be86f15400ed2984"
	pageNo := 1
	pageSize := 100

	ListResourceDetailPage(sc,customerId,pageNo,pageSize)

	ListResourceDetailAll(sc,customerId)

    TestResourceDailyALL(sc)

	fmt.Println("main end...")
}

//查询客户包周期资源列表
func ListResourceDetailPage(client *gophercloud.ServiceClient, customerId string, pageNo int, pageSize int) {
	resourceListTmp := resource.ListOpts{
		CustomerId: customerId,
		//ResourceIds: "2251f59c-b1ef-4398-bfa8-321782f670a5",
		//OrderId: "xxxx",
		//OnlyMainResource: 1,
		//StatusList:"1,2",
		PageNo:   pageNo,
		PageSize: pageSize,
	}

	resResourceListTmp, err := resource.ListDetail(client, resourceListTmp).Extract()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	if resResourceListTmp.ErrorCode == "CBC.0000" {
		fmt.Println("ListResourceDetailPage success!")
		fmt.Println("ListResourceDetailPage.TotalCount:", resResourceListTmp.TotalCount)
		fmt.Println("ListResourceDetailPage.ErrorCode:", resResourceListTmp.ErrorCode)
		fmt.Println("ListResourceDetailPage.ErrorMsg:", resResourceListTmp.ErrorMsg)
		fmt.Println("ListResourceDetailPage.Data:", len(resResourceListTmp.Data))

		b,_:=json.MarshalIndent(resResourceListTmp,""," ")
		fmt.Println(string(b))

		return
	} else {
		fmt.Println("ListResourceDetailPage failed!")
	}
}


func ListResourceDetailAll(sc *gophercloud.ServiceClient, customerId string) {
	var allRes resource.Resources

	optsTmp := resource.ListOpts{
		CustomerId: customerId,
		PageNo:   1,
		PageSize: 1,
	}

	resTmp, err := resource.ListDetail(sc, optsTmp).Extract()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	if resTmp.ErrorCode != "CBC.0000" {
		fmt.Println("List Detail faild!")
		return
	}

	allRes.ErrorCode = resTmp.ErrorCode
	allRes.ErrorMsg = resTmp.ErrorMsg
	allRes.TotalCount = resTmp.TotalCount

	//一次查100条
	onceCnt := 100
	totalCnt := resTmp.TotalCount
	queryTimes := int(math.Ceil(float64(totalCnt) / float64(onceCnt)))
	lastCnt := totalCnt - (queryTimes-1)*onceCnt
	fmt.Println("queryTimes:", queryTimes)

	opts := resource.ListOpts{
		CustomerId: customerId,
		PageNo:   1,
		PageSize: onceCnt,
	}

	for i := 1; i <= queryTimes; i++ {
		//最后一次要
		if i == queryTimes {
			opts.PageSize = lastCnt
		}

		res, err := resource.ListDetail(sc, opts).Extract()
		if err != nil {
			fmt.Println("err:", err)
			if ue, ok := err.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return
		}

		if resTmp.ErrorCode != "CBC.0000" {
			fmt.Println("List Detail faild!")
			return
		}

		allRes.Data = append(allRes.Data, res.Data...)
	}

	fmt.Println("get resources success!")
	fmt.Println("allRes.TotalCount:", allRes.TotalCount)
	fmt.Println("allRes.ErrorCode:", allRes.ErrorCode)
	fmt.Println("allRes.ErrorMsg:", allRes.ErrorMsg)
	fmt.Println("allRes.Data:", len(allRes.Data))
}


func TestResourceDailyALL(sc *gophercloud.ServiceClient){

	reqTmp := account.ResourceDailyOpts{
		StartTime: "2018-06-01",
		EndTime: "2018-06-30",
		PayMethod: "0",
		CloudServiceType: "hws.service.type.ebs",
		RegionCode:"cn-xianhz-1",
		ResourceId:"",
		EnterpriseProjectId: "",
	}

	rspTmp,err := account.ListResourceDaily(sc, reqTmp)
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b,_:=json.MarshalIndent(rspTmp,""," ")
	fmt.Println(string(b))
}
