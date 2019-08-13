package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/metrics"
	"github.com/gophercloud/gophercloud/pagination"
)

func main() {

	fmt.Println("main start...")
	provider, err := common.AuthAKSK()
	//provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewCESV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ces client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	TestMetricsList(sc)
	fmt.Println("main end...")
}

func TestMetricsList(sc *gophercloud.ServiceClient) {
	limit := 10
	opts := metrics.ListOpts{
		Limit:     &limit,
		Namespace: "SYS.ELB",
		//Start:"SYS.ECS.inst_sys_status_error.instance_id:014f6ff1-4769-4f91-aab9-5e117092375a",
	}
	var metricsResp metrics.Metrics
	metricsResp.Metrics = make([]metrics.Metric, 0)
	var err error

	// 获取当前页数据
	err = metrics.List(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		metricsResp, err = metrics.ExtractMetrics(page)
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		return false, err
	})
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	bytes, _ := json.MarshalIndent(metricsResp, "", " ")
	fmt.Println(string(bytes))

	// 获取所有页数据
	allPages, err := metrics.List(sc, opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	metricsResp, err = metrics.ExtractAllPagesMetrics(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, _ = json.MarshalIndent(metricsResp, "", " ")
	fmt.Println(string(bytes))
}
