package testing

import (
	"fmt"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/ports"
	"github.com/gophercloud/gophercloud/testhelper/client"

	"net/http"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var CreateOutput = `
{
  "port": {
    "id": "5e56a480-f337-4985-8ca4-98546cb4fdae",
    "name": "EricTestPort",
    "status": "DOWN",
    "admin_state_up": true,
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.208"
    }],
    "mac_address": "fa:16:3e:6b:b5:10",
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_id": "",
    "device_owner": "",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "extra_dhcp_opts": [],
    "allowed_address_pairs": [],
    "binding:vnic_type": "normal",
    "binding:vif_details": {},
    "binding:profile": {}
  }
}
`

var CreateResponse = ports.Port{
	ID:           "5e56a480-f337-4985-8ca4-98546cb4fdae",
	Name:         "EricTestPort",
	Status:       "DOWN",
	AdminStateUp: true,
	FixedIps: []ports.FixedIp{
		{
			SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
			IpAddress: "192.168.0.208",
		},
	},
	MacAddress:  "fa:16:3e:6b:b5:10",
	NetworkId:   "5ae24488-454f-499c-86c4-c0355704005d",
	TenantId:    "57e98940a77f4bb988a21a7d0603a626",
	DeviceId:    "",
	DeviceOwner: "",
	SecurityGroups: []string{
		"7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
	},
	ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
	AllowedAddressPairs: []ports.AllowedAddressPair{},
	BindingvnicType:     "normal",
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/ports", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "port": {
    "id": "5e56a480-f337-4985-8ca4-98546cb4fdae",
    "name": "ModifiedPort",
    "status": "DOWN",
    "admin_state_up": true,
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.208"
    }],
    "mac_address": "fa:16:3e:6b:b5:10",
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_id": "",
    "device_owner": "",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "extra_dhcp_opts": [],
    "allowed_address_pairs": [],
    "binding:vnic_type": "normal",
    "binding:vif_details": {},
    "binding:profile": {}
  }
}
`

var UpdateResponse = ports.Port{
	ID:           "5e56a480-f337-4985-8ca4-98546cb4fdae",
	Name:         "ModifiedPort",
	Status:       "DOWN",
	AdminStateUp: true,
	FixedIps: []ports.FixedIp{
		{
			SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
			IpAddress: "192.168.0.208",
		},
	},
	MacAddress:  "fa:16:3e:6b:b5:10",
	NetworkId:   "5ae24488-454f-499c-86c4-c0355704005d",
	TenantId:    "57e98940a77f4bb988a21a7d0603a626",
	DeviceId:    "",
	DeviceOwner: "",
	SecurityGroups: []string{
		"7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
	},
	ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
	AllowedAddressPairs: []ports.AllowedAddressPair{},
	BindingvnicType:     "normal",
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/ports/5e56a480-f337-4985-8ca4-98546cb4fdae", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "port": {
    "allowed_address_pairs": [],
    "extra_dhcp_opts": [],
    "updated_at": "2018-04-21T09:03:05",
    "device_owner": "",
    "binding:profile": {},
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.208"
    }],
    "id": "5e56a480-f337-4985-8ca4-98546cb4fdae",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "binding:vif_details": {},
    "binding:vif_type": "unbound",
    "qos_policy_id": null,
    "mac_address": "fa:16:3e:6b:b5:10",
    "status": "DOWN",
    "binding:host_id": "",
    "description": "",
    "tags": [],
    "device_id": "",
    "name": "ModifiedPort",
    "admin_state_up": true,
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "created_at": "2018-04-21T08:25:27",
    "binding:vnic_type": "normal"
  }
}
`

var GetResponse = ports.Port{
	ID:           "5e56a480-f337-4985-8ca4-98546cb4fdae",
	Name:         "ModifiedPort",
	Status:       "DOWN",
	AdminStateUp: true,
	FixedIps: []ports.FixedIp{
		{
			SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
			IpAddress: "192.168.0.208",
		},
	},
	MacAddress:  "fa:16:3e:6b:b5:10",
	NetworkId:   "5ae24488-454f-499c-86c4-c0355704005d",
	TenantId:    "57e98940a77f4bb988a21a7d0603a626",
	DeviceId:    "",
	DeviceOwner: "",
	SecurityGroups: []string{
		"7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
	},
	ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
	AllowedAddressPairs: []ports.AllowedAddressPair{},
	BindingvnicType:     "normal",
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/ports/5e56a480-f337-4985-8ca4-98546cb4fdae", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "ports": [{
    "id": "0740752d-0b2b-44e2-bf2a-edbe51c36227",
    "name": "",
    "status": "DOWN",
    "admin_state_up": true,
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.196"
    }],
    "mac_address": "fa:16:3e:74:e4:3a",
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_id": "",
    "device_owner": "neutron:VIP_PORT",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "extra_dhcp_opts": [],
    "allowed_address_pairs": [],
    "binding:vnic_type": "normal",
    "binding:vif_details": {},
    "binding:profile": {}
  }, {
    "id": "127ce534-6b04-4523-9e6d-dd17fd25bf13",
    "name": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
    "status": "DOWN",
    "admin_state_up": true,
    "fixed_ips": [{
      "subnet_id": "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
      "ip_address": "192.168.0.1"
    }],
    "mac_address": "fa:16:3e:75:43:f9",
    "network_id": "5ae24488-454f-499c-86c4-c0355704005d",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "device_id": "e23caa95-2599-4aa8-a2db-be3444450e78",
    "device_owner": "network:router_interface_distributed",
    "security_groups": ["7844d4b4-d78f-45dc-9465-2b4d1bca83a5"],
    "extra_dhcp_opts": [],
    "allowed_address_pairs": [],
    "binding:vnic_type": "normal",
    "binding:vif_details": {},
    "binding:profile": {}
  }],
  "ports_links": [{
    "href": "https://None/v2.0/ports?limit=3&marker=0740752d-0b2b-44e2-bf2a-edbe51c36227&page_reverse=True",
    "rel": "previous"
  }]
}
`

var ListResponse = []ports.Port{
	{
		ID:           "0740752d-0b2b-44e2-bf2a-edbe51c36227",
		Name:         "",
		Status:       "DOWN",
		AdminStateUp: true,
		FixedIps: []ports.FixedIp{
			{
				SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
				IpAddress: "192.168.0.196",
			},
		},
		MacAddress:          "fa:16:3e:74:e4:3a",
		NetworkId:           "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:            "57e98940a77f4bb988a21a7d0603a626",
		DeviceId:            "",
		DeviceOwner:         "neutron:VIP_PORT",
		SecurityGroups:      []string{"7844d4b4-d78f-45dc-9465-2b4d1bca83a5"},
		ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
		AllowedAddressPairs: []ports.AllowedAddressPair{},
		BindingvnicType:     "normal",
	}, {
		ID:           "127ce534-6b04-4523-9e6d-dd17fd25bf13",
		Name:         "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
		Status:       "DOWN",
		AdminStateUp: true,
		FixedIps: []ports.FixedIp{
			{
				SubnetId:  "7b4b101c-d5e2-4c52-9c6d-c6c7e1219919",
				IpAddress: "192.168.0.1",
			},
		},
		MacAddress:          "fa:16:3e:75:43:f9",
		NetworkId:           "5ae24488-454f-499c-86c4-c0355704005d",
		TenantId:            "57e98940a77f4bb988a21a7d0603a626",
		DeviceId:            "e23caa95-2599-4aa8-a2db-be3444450e78",
		DeviceOwner:         "network:router_interface_distributed",
		SecurityGroups:      []string{"7844d4b4-d78f-45dc-9465-2b4d1bca83a5"},
		ExtraDhcpOpts:       []ports.ExtraDHCPOpt{},
		AllowedAddressPairs: []ports.AllowedAddressPair{},
		BindingvnicType:     "normal",
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/ports", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/ports/5e56a480-f337-4985-8ca4-98546cb4fdae", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
