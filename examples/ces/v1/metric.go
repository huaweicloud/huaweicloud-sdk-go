package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/metrics"
	"github.com/gophercloud/gophercloud/pagination"
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

	MetricsList(sc)
	fmt.Println("main end...")
}

func MetricsList(sc *gophercloud.ServiceClient) {
	limit := 10
	opts := metrics.ListOpts{
		Limit:     &limit,
		Namespace: "SYS.ELB",
		Start:"SYS.ECS.inst_sys_status_error.instance_id:014f6ff1-4769-4f91-aab9-5e117092375a",
	}
	var metricsOnePageResp metrics.Metrics
	var metricsAllPageResp metrics.Metrics
	metricsOnePageResp.Metrics = make([]metrics.Metric, 0)
	var err error

	// 获取当前页数据
	err = metrics.List(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		metricsOnePageResp, err = metrics.ExtractMetrics(page)
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

	res, marshalErr := json.MarshalIndent(metricsOnePageResp, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal metricsOnePageResp error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))

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

	metricsAllPageResp, err = metrics.ExtractAllPagesMetrics(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("metric metadata Count", metricsAllPageResp.MetaData.Count)
	fmt.Println("metric metadata Total", metricsAllPageResp.MetaData.Total)
	fmt.Println("metric metadata Marker", metricsAllPageResp.MetaData.Marker)
	fmt.Println("metric Metrics list", metricsAllPageResp.Metrics)

	for _, data := range metricsAllPageResp.Metrics {
		fmt.Println("metric data Unit", data.Unit)
		fmt.Println("metric data Dimensions", data.Dimensions)
		fmt.Println("metric data MetricName", data.MetricName)
		fmt.Println("metric data Namespace", data.Namespace)
	}

	res, marshalErr = json.MarshalIndent(metricsAllPageResp, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal metricsAllPageResp error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))
}
