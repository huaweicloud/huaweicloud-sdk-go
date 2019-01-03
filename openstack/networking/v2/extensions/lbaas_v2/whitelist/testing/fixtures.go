package testing

import (
	"testing"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/whitelist"
	th "github.com/gophercloud/gophercloud/testhelper"
	"net/http"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"fmt"
)

var CreateResp = `
{
    "whitelist": {
        "id": "eabfefa3fd1740a88a47ad98e132d238", 
        "listener_id": "eabfefa3fd1740a88a47ad98e132d238", 
        "tenant_id": "eabfefa3fd1740a88a47ad98e132d238", 
        "enable_whitelist": true, 
        "whitelist": "192.168.11.1,192.168.0.1/24,192.168.201.18/8,100.164.0.1/24"
    }
}`

var GetResp = `
{
    "whitelist": {
        "id": "eabfefa3fd1740a88a47ad98e132d238", 
        "listener_id": "eabfefa3fd1740a88a47ad98e132d238", 
        "tenant_id": "eabfefa3fd1740a88a47ad98e132d238", 
        "enable_whitelist": true, 
        "whitelist": "192.168.11.1,192.168.0.1/24,192.168.201.18/8,100.164.0.1/24"
    }
}`
var WhitelistBody = `
{
    "whitelists": [
        {
            "id": "eabfefa3fd1740a88a47ad98e132d238", 
            "listener_id": "eabfefa3fd1740a88a47ad98e132d238", 
            "tenant_id": "eabfefa3fd1740a88a47ad98e132d238", 
            "enable_whitelist": true, 
            "whitelist": "192.168.11.1,192.168.0.1/24,192.168.201.18/8,100.164.0.1/24"
        }, 
        {
            "id": "eabfefa3fd1740a88a47ad98e132d326", 
            "listener_id": "eabfefa3fd1740a88a47ad98e132d327", 
            "tenant_id": "eabfefa3fd1740a88a47ad98e132d436", 
            "enable_whitelist": true, 
            "whitelist": "192.168.12.1,192.168.1.1/24,192.168.203.18/8,100.164.5.1/24"
        }
    ]
}
`
var (
	WhitelistOne = whitelist.Whitelist{
		ID:              "eabfefa3fd1740a88a47ad98e132d238",
		TenantId:        "eabfefa3fd1740a88a47ad98e132d238",
		ListenerId:      "eabfefa3fd1740a88a47ad98e132d238",
		EnableWhitelist: true,
		Whitelist:       "192.168.11.1,192.168.0.1/24,192.168.201.18/8,100.164.0.1/24",
	}
	WhitelistTwo = whitelist.Whitelist{
		ID:              "eabfefa3fd1740a88a47ad98e132d326",
		TenantId:        "eabfefa3fd1740a88a47ad98e132d436",
		ListenerId:      "eabfefa3fd1740a88a47ad98e132d327",
		EnableWhitelist: true,
		Whitelist:       "192.168.12.1,192.168.1.1/24,192.168.203.18/8,100.164.5.1/24",
	}
)

func HandleWhitelistSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/whitelists", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()
		fmt.Println(r.Form)
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, WhitelistBody)
		case "eabfefa3fd1740a88a47ad98e132d326":
			fmt.Fprintf(w, `{ "pools": [] }`)
		default:
			t.Fatalf("/v2.0/lbaas/whitelists invoked with unexpected marker=[%s]", marker)
		}
	})
}

func HandleWhitelistALLSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/whitelists", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, WhitelistBody)

	})
}

func HandleWhitelistCreationSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/whitelists", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, CreateResp)
	})

}

func HandleWhitelistGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/whitelists/09e64049-2ab0-4763-a8c5-f4207875dc3e", func(w http.ResponseWriter, r *http.Request) {

		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetResp)
	})
}

func HandleWhitelistDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/whitelists/09e64049-2ab0-4763-a8c5-f4207875dc3e", func(w http.ResponseWriter, r *http.Request) {

		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleWhitelistUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/whitelists/09e64049-2ab0-4763-a8c5-f4207875dc3e", func(w http.ResponseWriter, r *http.Request) {

		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetResp)
	})
}
