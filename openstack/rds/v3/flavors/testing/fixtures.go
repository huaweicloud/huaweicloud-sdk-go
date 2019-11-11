package testing

import (
	"fmt"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

const FlavorsMysqlResp = `
{
	"flavors": [{
		"vcpus": "1",
		"ram": 2,
		"spec_code": "rds.mysql.c2.medium.ha",
		"instance_mode": "ha"
	}, {
		"vcpus": "1",
		"ram": 2,
		"spec_code": "rds.mysql.c2.medium.rr",
		"instance_mode": "replica"
	}]
}
`
const FlavorsPostgreSQLResp = `
{
	"flavors": [{
		"vcpus": "2",
		"ram": 4,
		"spec_code": "rds.pg.c2.large",
		"instance_mode": "single",
		"az_status": {
			"az2xahz": "normal",
			"az1xahz": "normal",
			"az3xahz": "normal"
		}
	},{
		"vcpus": "48",
		"ram": 384,
		"spec_code": "rds.pg.i3.12xlarge.8",
		"instance_mode": "single",
		"az_status": {
			"az2xahz": "unsupported",
			"az1xahz": "unsupported",
			"az3xahz": "unsupported"
		}
	}, {
		"vcpus": "60",
		"ram": 512,
		"spec_code": "rds.pg.i3.15xlarge.8",
		"instance_mode": "single",
		"az_status": {
			"az2xahz": "unsupported",
			"az1xahz": "unsupported",
			"az3xahz": "unsupported"
		}
	}, {
		"vcpus": "48",
		"ram": 384,
		"spec_code": "rds.pg.i3.12xlarge.8.ha",
		"instance_mode": "ha",
		"az_status": {
			"az2xahz": "unsupported",
			"az1xahz": "unsupported",
			"az3xahz": "unsupported"
		}
	}]
}
`
func HandleListMysqlSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/flavors/MySQL", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, FlavorsMysqlResp)
		params := r.URL.Query()
		versionname := params.Get("version_name")
		th.CheckEquals(t, "5.7", versionname)
	})
}

func HandleListPgSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/flavors/PostgreSQL", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, FlavorsPostgreSQLResp)
		params := r.URL.Query()
		versionname := params.Get("version_name")
		th.CheckEquals(t, "9.6", versionname)
	})
}