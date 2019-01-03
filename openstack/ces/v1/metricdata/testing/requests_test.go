package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/metricdata"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestBatchQuery(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/batch-query-metric-data", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, BatchQueryRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, BatchQueryResponse)
	})

	options := metricdata.BatchQueryOpts{
		Metrics: []metricdata.Metric{
			{
				Namespace: "MINE.APP",
				Dimensions: []map[string]string{
					{
						"name":  "instance_id",
						"value": "33328f02-3814-422e-b688-bfdba93d4050",
					},
				},
				MetricName: "cpu_util",
			},
			{
				Namespace: "MINE.APP",
				Dimensions: []map[string]string{
					{
						"name":  "instance_id",
						"value": "33328f02-3814-422e-b688-bfdba93d4051",
					},
				},
				MetricName: "cpu_util",
			},
		},
		From:   1484153313000,
		To:     1484653313000,
		Period: "1",
		Filter: "average",
	}
	n, err := metricdata.BatchQuery(fake.ServiceClient(), options).ExtractMetricDatas()
	th.AssertNoErr(t, err)

	th.AssertEquals(t,n[0].Namespace , "MINE.APP")
	th.AssertEquals(t,n[0].MetricName , "cpu_util")
	th.AssertEquals(t,n[0].Unit , "request/s")
	th.AssertDeepEquals(t,n[0].Dimensions,[]map[string]interface{}{
		{
			"name": "instance_id",
			"value": "33328f02-3814-422e-b688-bfdba93d4050",
		},
	})
}
