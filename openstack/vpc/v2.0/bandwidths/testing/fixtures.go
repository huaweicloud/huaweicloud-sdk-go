package testing

import (
	"testing"
	th "github.com/gophercloud/gophercloud/testhelper"
	"net/http"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"fmt"
)



var respOrder = `
{
    "order_id": "dd0bdea6efe0"
}
`
var respBandwidth = `
{
    "bandwidth": {
        "id": "3fa5b383-5a73-4dcb-a314-c6128546d855",
        "name": "test",
        "size": 10,
        "share_type": "PER",
        "publicip_info": [
            {
                "publicip_id": "6285e7be-fd9f-497c-bc2d-dd0bdea6efe0",
                "publicip_address": "161.xx.xx.9",
                "publicip_type": "5_bgp",
                "ip_version": 4
            }
        ],
        "tenant_id": "8b7e35ad379141fc9df3e178bd64f55c",
        "bandwidth_type": "bgp"
    }
}
`

// HandleWithBSSInfoSuccessfully for testing
func HandleWithBSSInfoSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v2.0/128a7bf965154373a7b73c89eb6b65aa/bandwidths/3fa5b383-5a73-4dcb-a314-c6128546d855", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, respOrder)
	})

}

func HandleWithNameSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v2.0/128a7bf965154373a7b73c89eb6b65aa/bandwidths/3fa5b383-5a73-4dcb-a314-c6128546d855", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, respBandwidth)
	})

}




