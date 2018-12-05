package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/metricdata"
)

func main() {

	fmt.Println("main start...")

	//provider, err := common.AuthAKSK()
	provider, err := common.AuthToken()
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

	TestBatchQueryMetricDatas(sc)
	fmt.Println("main end...")
}

func TestBatchQueryMetricDatas(sc *gophercloud.ServiceClient) {
	opts := metricdata.BatchQueryOpts{
		Metrics: []metricdata.Metric{
			{
				Namespace: "SYS.VPC",
				Dimensions: []map[string]string{
					{
						"name":  "bandwidth_id",
						"value": "ea31a911-dad7-4218-9036-77a7c3a16a45",
					},
				},
				MetricName: "downstream_bandwidth",
			},
		},
		From:   1540526925098,
		To:     1540537725098,
		Period: "1",
		Filter: "average",
	}
	metricdatas, err := metricdata.BatchQuery(sc, opts).ExtractMetricDatas()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test batch query metric data success！")
	p, _ := json.MarshalIndent(metricdatas, "", " ")
	fmt.Println(string(p))
}
