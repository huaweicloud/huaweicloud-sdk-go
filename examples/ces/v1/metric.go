package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/metrics"
	"github.com/gophercloud/gophercloud/pagination"
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

	MetricsList(sc)
	fmt.Println("main end...")
}

func MetricsList(sc *gophercloud.ServiceClient) {
	limit := 10
	opts := metrics.ListOpts{
		Limit:     &limit,
		Namespace: "SYS.ELB",
		//Start:"SYS.ECS.inst_sys_status_error.instance_id:014f6ff1-4769-4f91-aab9-5e117092375a",
	}
	var metricsresp metrics.Metrics
	metricsresp.Metrics = make([]metrics.Metric, 0)
	var err error

	// 获取当前页数据
	err = metrics.List(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		metricsresp, err = metrics.ExtractMetrics(page)
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
	bytes, _ := json.MarshalIndent(metricsresp, "", " ")
	fmt.Println(string(bytes))

	// 获取所有页数据
	allpages, err := metrics.List(sc, opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	metricsresp, err = metrics.ExtractAllPagesMetrics(allpages)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("metric metadata Count", metricsresp.MetaData.Count)
	fmt.Println("metric metadata Total", metricsresp.MetaData.Total)
	fmt.Println("metric metadata Marker", metricsresp.MetaData.Marker)
	fmt.Println("metric Metrics list", metricsresp.Metrics)

	for _, data := range metricsresp.Metrics {
		fmt.Println("metric data Unit", data.Unit)
		fmt.Println("metric data Dimensions", data.Dimensions)
		fmt.Println("metric data MetricName", data.MetricName)
		fmt.Println("metric data Namespace", data.Namespace)
	}

	bytes, _ = json.MarshalIndent(metricsresp, "", " ")
	fmt.Println(string(bytes))
}
