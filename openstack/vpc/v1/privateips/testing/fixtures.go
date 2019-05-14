package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/privateips"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var CreateOutput = `
{
  "privateips": [{
    "status": "DOWN",
    "id": "ea274524-f1cc-4078-8e67-c002be25c413",
    "subnet_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_owner": "",
    "ip_address": "192.168.0.12"
  }]
}
`

var CreateResponse = []privateips.PrivateIp{
	{
		Status:      "DOWN",
		ID:          "ea274524-f1cc-4078-8e67-c002be25c413",
		SubnetId:    "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:    "57e98940a77f4bb988a21a7d0603a626",
		DeviceOwner: "",
		IpAddress:   "192.168.0.12",
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/privateips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, CreateOutput)
	})
}

var GetOutput = `
{
  "privateip": {
    "status": "DOWN",
    "id": "ea274524-f1cc-4078-8e67-c002be25c413",
    "subnet_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_owner": "",
    "ip_address": "192.168.0.12"
  }
}
`
var GetResponse = privateips.PrivateIp{
	Status:      "DOWN",
	ID:          "ea274524-f1cc-4078-8e67-c002be25c413",
	SubnetId:    "5ae24488-454f-499c-86c4-c0355704005d",
	TenantId:    "57e98940a77f4bb988a21a7d0603a626",
	DeviceOwner: "",
	IpAddress:   "192.168.0.12",
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/privateips/ea274524-f1cc-4078-8e67-c002be25c413", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "privateips": [{
    "status": "DOWN",
    "id": "0740752d-0b2b-44e2-bf2a-edbe51c36227",
    "subnet_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_owner": "neutron:VIP_PORT",
    "ip_address": "192.168.0.196"
  }, {
    "status": "DOWN",
    "id": "127ce534-6b04-4523-9e6d-dd17fd25bf13",
    "subnet_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_owner": "network:router_interface_distributed",
    "ip_address": "192.168.0.1"
  }]
}
`

var ListResponse = []privateips.PrivateIp{
	{

		Status:      "DOWN",
		ID:          "0740752d-0b2b-44e2-bf2a-edbe51c36227",
		SubnetId:    "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:    "57e98940a77f4bb988a21a7d0603a626",
		DeviceOwner: "neutron:VIP_PORT",
		IpAddress:   "192.168.0.196",
	},
	{
		Status:      "DOWN",
		ID:          "127ce534-6b04-4523-9e6d-dd17fd25bf13",
		SubnetId:    "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:    "57e98940a77f4bb988a21a7d0603a626",
		DeviceOwner: "network:router_interface_distributed",
		IpAddress:   "192.168.0.1",
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/subnets/5ae24488-454f-499c-86c4-c0355704005d/privateips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/privateips/ea274524-f1cc-4078-8e67-c002be25c413 ", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
