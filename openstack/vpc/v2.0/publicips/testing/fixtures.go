package testing

import (
	"testing"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/publicips"
	th "github.com/gophercloud/gophercloud/testhelper"
	"net/http"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"fmt"
)

func HandleOndemandSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v2.0/128a7bf965154373a7b73c89eb6b65aa/publicips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		//w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, OndemandResp)
	})

}

func HandleWithBSSInfoSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/v2.0/128a7bf965154373a7b73c89eb6b65aa/publicips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		fmt.Fprintf(w, BSSResp)
	})
}

var OndemandResp = `
{
    "publicip": {
        "id": "f588ccfa-8750-4d7c-bf5d-2ede24414706",
        "status": "PENDING_CREATE",
        "type": "5_bgp",
        "public_ip_address": "161.17.101.7",
        "tenant_id": "8b7e35ad379141fc9df3e178bd64f55c",
        "create_time": "2015-07-16 04:10:52",
        "bandwidth_size": 2,
        "ip_version": 4
    }
}
`
var BSSResp = `
{
    "order_id": "CS1802081410IMDRN",
    "publicip_id": "f588ccfa-8750-4d7c-bf5d-2ede24414706"
}
`
var (
	size   = 2
	pernum = 1
	ipver  = 4
)
var CreateResultForBSS = publicips.PrePaid{
	OrderID:    "CS1802081410IMDRN",
	PublicipID: "f588ccfa-8750-4d7c-bf5d-2ede24414706",
}

var OndmandOpts = publicips.CreateOpts{
	PublicIP: publicips.PublicIP{
		Type:      "5_bgp",
		IPVersion: ipver,},
	Bandwidth: publicips.Bandwidth{
		Name:       "test1",
		Size:       size,
		ShareType:  "PER",
		ChargeMode: "bandwidth",
	},
}

var BSSOpts = publicips.CreateOpts{

	PublicIP: publicips.PublicIP{
		Type: "5_bgp",
	},
	Bandwidth: publicips.Bandwidth{
		Name:       "test1",
		Size:       size,
		ShareType:  "PER",
		ChargeMode: "bandwidth",
	},
	ExtendParam: publicips.ExtendParam{
		ChargeMode:  "prePaid",
		PeriodNum:   pernum,
		PeriodType:  "month",
		IsAutoPay:   "true",
		IsAutoRenew: "false",
	},
}
