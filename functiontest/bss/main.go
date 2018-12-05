package main

import (
	"encoding/json"
	"fmt"
	"math"
	
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/account"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/resource"
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

	TestListDetail(sc)

    TestResourceDailyALL(sc)

	fmt.Println("main end...")
}

func TestListDetail(sc *gophercloud.ServiceClient) {
	var allRes resource.Resources

	optsTmp := resource.ListOpts{
		CustomerId: "3b011b89b2f64fb68782a43380e2a78f",
		//ResourceIds: "2251f59c-b1ef-4398-bfa8-321782f670a5",
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
	//此处为适配api bug代码，bss接口失败也只会返回http 200.
	if resTmp.ErrorCode == "" || resTmp.ErrorCode == "CBC.0999" {
		fmt.Println("List Detail faild!")
		return
	}

	allRes.ErrorCode = resTmp.ErrorCode
	allRes.ErrorMsg = resTmp.ErrorMsg
	allRes.TotalCount = resTmp.TotalCount
	//allRes.Data = make(resource.ResourceInstance, resTmp.TotalCount)

	//一次查10条
	onceCnt := 10
	totalCnt := resTmp.TotalCount
	queryTimes := int(math.Ceil(float64(totalCnt) / float64(onceCnt)))
	lastCnt := totalCnt - (queryTimes-1)*onceCnt
	fmt.Println("queryTimes:", queryTimes)

	opts := resource.ListOpts{
		CustomerId: "3b011b89b2f64fb68782a43380e2a78f",
		//ResourceIds: "2251f59c-b1ef-4398-bfa8-321782f670a5",
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
		//此处为适配api bug代码，bss接口失败也只会返回http 200.
		if res.ErrorCode == "" || res.ErrorCode == "CBC.0999" {
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
