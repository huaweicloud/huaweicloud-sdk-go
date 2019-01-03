package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/ces/v1/metrics"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()
		start := r.Form.Get("start")
		switch start {
		case "":
			fmt.Fprintf(w, "%s", ListResponse)
		case "SYS.ECS.cpu_util.instance_id:d9112af5-6913-4f3b-bd0a-3f96711e004d":
			fmt.Fprintf(w, EndPageResponse)
		default:
			t.Fatalf("Unexpected start: [%s]", start)
		}
	})
	limit := 1
	opts := metrics.ListOpts{
		Limit: &limit,
		Start: "",
	}
	allpage, err := metrics.List(fake.ServiceClient(), opts).AllPages()
	th.AssertNoErr(t, err)
	metricsresp, err := metrics.ExtractAllPagesMetrics(allpage)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, metricsresp.MetaData.Count, 1)
	th.AssertEquals(t, metricsresp.MetaData.Marker, "SYS.ECS.cpu_util.instance_id:d9112af5-6913-4f3b-bd0a-3f96711e004d")
	th.AssertEquals(t, metricsresp.MetaData.Total, 1)
	th.AssertDeepEquals(t, metricsresp.Metrics, []metrics.Metric{
		{
			Namespace: "SYS.ECS",
			Dimensions: []metrics.Dimension{
				{
					Name:  "instance_id",
					Value: "d9112af5-6913-4f3b-bd0a-3f96711e004d",
				},
			},
			MetricName: "cpu_util",
			Unit:       "%",
		},
	})
}
