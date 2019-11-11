package testing

import (
	"fmt"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

const StorageTypeResp = `
{ 
	"storage_type": [{
		"name": "COMMON",
		"az_status": {
			"az1": "normal",
			"az2": "normal"
		}
	},
	{
		"name": "ULTRAHIGH",
		"az_status": {
			"az1": "unsupported",
			"az2": "unsupported"
		}
	}]
}
`

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/storage-type/mysql", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, StorageTypeResp)
		params := r.URL.Query()
		versionname := params.Get("version_name")
		th.CheckEquals(t, "5.7", versionname)
	})
}
