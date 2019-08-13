package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/metricdata"
)

func main() {
	fmt.Println("main start...")
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
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

	BatchQueryMetricData(sc)
	GetMetricData(sc)
	GetEventData(sc)
	AddMetricData(sc)
	fmt.Println("main end...")
}

func BatchQueryMetricData(sc *gophercloud.ServiceClient) {
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

	metricData, err := metricdata.BatchQuery(sc, opts).ExtractMetricDatas()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test batch query metric data success！")

	for _, data := range metricData {
		fmt.Println("metric data Datapoints", data.Datapoints)
		fmt.Println("metric data Dimensions", data.Dimensions)
		fmt.Println("metric data MetricName", data.MetricName)
		fmt.Println("metric data Namespace", data.Namespace)
	}

	res, marshalErr := json.MarshalIndent(metricData, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal metricData error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))
}

func GetMetricData(sc *gophercloud.ServiceClient) {
	opts := metricdata.GetOpts{
		Namespace:  "SYS.ECS",
		MetricName: "cpu_util",
		From:       "1548041969418",
		To:         "1548052769418",
		Period:     "3600",
		Filter:     "average",
		Dim0:       "instance_id,070c1ed3-176a-446e-8eff-b116b529b4b7",
	}

	metricData, err := metricdata.Get(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Get metric data success！")
	res, marshalErr := json.MarshalIndent(metricData, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal metricData error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))
}

func AddMetricData(sc *gophercloud.ServiceClient) {
	opts := metricdata.AddMetricDataOpts{
		{
			Metric: metricdata.MetricInfo{
				Namespace: "MINE.APP",
				Dimensions: []metricdata.MetricsDimension{
					{
						Name:  "instance_id",
						Value: "33328f02-3814-422e-b688-bfdba93d4050",
					},
				},
				MetricName: "cpu_util",
			},
			Ttl:         172800,
			CollectTime: 1463598260000,
			Value:       60,
			Unit:        "%",
			Type:        "int",
		},
	}

	err := metricdata.AddMetricData(sc, opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Add metric data success！")
}

func GetEventData(sc *gophercloud.ServiceClient) {
	opts := metricdata.GetEventDataOpts{
		Namespace: "SYS.ECS",
		From:      "1548041969418",
		To:        "1548052769418",
		Dim0:      "instance_id,070c1ed3-176a-446e-8eff-b116b529b4b7",
		Type:      "instance_host_info",
	}

	eventData, err := metricdata.GetEventData(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Get event data success！")
	res, marshalErr := json.MarshalIndent(eventData, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal eventData error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))
}
