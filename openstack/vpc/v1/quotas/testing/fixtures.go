package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/quotas"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var ListOutput = `
{
  "quotas": {
    "resources": [{
      "type": "vpc",
      "used": 3,
      "quota": 5,
      "min": 0
    }]
  }
}
`

var ListResponse = quotas.Quota{
	Resources: []quotas.Resource{
		{
			Type:  "vpc",
			Used:  3,
			Quota: 5,
			Min:   0,
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/quotas", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}
