package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/gophercloud/gophercloud/testhelper/client"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/job"
)

func TestGetJobResult(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	jobId := "2c9eb2c5544cbf6101544f0602af2b4f"

	th.Mux.HandleFunc(fmt.Sprintf("/jobs/%s",jobId), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetJobResultResponse)
	})

	jobrs,err := job.GetJobResult(fake.ServiceClient(),jobId)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, jobrs.Status, "SUCCESS")
	th.AssertEquals(t, jobrs.Id, jobId)
	th.AssertEquals(t, jobrs.Type, "createServer")
	th.AssertEquals(t, jobrs.BeginTime, "2016-04-25T20:04:34.604Z")
	th.AssertEquals(t, jobrs.EndTime, "2016-04-25T20:08:41.593Z")
	th.AssertEquals(t, jobrs.ErrorCode, "")
	th.AssertEquals(t, jobrs.FailReason, "")
	th.AssertEquals(t, jobrs.Entities.SubJobsTotal, 1)
}