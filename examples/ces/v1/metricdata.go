package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/metricdata"
	"github.com/gophercloud/gophercloud/auth/aksk"
)

func main() {
	fmt.Println("main start...")

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
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

	BatchQueryMetricDatas(sc)
	fmt.Println("main end...")
}

func BatchQueryMetricDatas(sc *gophercloud.ServiceClient) {
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

	for _,data:=range metricdatas{

		fmt.Println("metric data Datapoints",data.Datapoints)
		fmt.Println("metric data Dimensions",data.Dimensions)
		fmt.Println("metric data MetricName",data.MetricName)
		fmt.Println("metric data Namespace",data.Namespace)
	}

	p, _ := json.MarshalIndent(metricdatas, "", " ")
	fmt.Println(string(p))
}
