package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/securitygroups"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var CreateOutput = `
{
  "security_group": {
    "id": "f7616338-fa30-42b8-bf6b-754c0701aab8",
    "name": "EricSG",
    "description": "",
    "security_group_rules": [{
      "direction": "egress",
      "ethertype": "IPv4",
      "id": "5658ade3-e87c-4016-afc6-b8efd3c2e306",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "4f87f0ea-1927-4b81-a4bb-83943b10585e",
      "remote_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "1438172b-20b8-4454-945c-9b5508a8e4f7",
      "protocol": "tcp",
      "port_range_max": 22,
      "port_range_min": 22,
      "remote_ip_prefix": "0.0.0.0/0",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "43fcaf3d-5406-4fc1-ab42-a18cc2660f01",
      "protocol": "tcp",
      "port_range_max": 3389,
      "port_range_min": 3389,
      "remote_ip_prefix": "0.0.0.0/0",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }]
  }
}
`
var PortNumber_3389 = 3389
var PortNumber_22 = 22

var CreateResponse = securitygroups.SecurityGroup{
	ID:          "f7616338-fa30-42b8-bf6b-754c0701aab8",
	Name:        "EricSG",
	Description: "",
	SecurityGroupRules: []securitygroups.SecurityGroupRule{
		{
			Direction:       "egress",
			Ethertype:       "IPv4",
			ID:              "5658ade3-e87c-4016-afc6-b8efd3c2e306",
			SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
		},
		{
			Direction:       "ingress",
			Ethertype:       "IPv4",
			RemoteGroupId:   "f7616338-fa30-42b8-bf6b-754c0701aab8",
			ID:              "4f87f0ea-1927-4b81-a4bb-83943b10585e",
			SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
		},
		{
			Direction:       "ingress",
			Ethertype:       "IPv4",
			ID:              "1438172b-20b8-4454-945c-9b5508a8e4f7",
			Protocol:        "tcp",
			PortRangeMax:    &PortNumber_22,
			PortRangeMin:    &PortNumber_22,
			RemoteIpPrefix:  "0.0.0.0/0",
			SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
		},
		{
			Direction:       "ingress",
			Ethertype:       "IPv4",
			ID:              "43fcaf3d-5406-4fc1-ab42-a18cc2660f01",
			Protocol:        "tcp",
			PortRangeMax:    &PortNumber_3389,
			PortRangeMin:    &PortNumber_3389,
			RemoteIpPrefix:  "0.0.0.0/0",
			SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
		},
	},
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/security-groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, CreateOutput)
	})
}

//var UpdateOutput = `
//{
//  "subnet": {
//    "id": "c9aba52d-ec14-40cb-930f-c8153e93c2db",
//    "status": "ACTIVE"
//  }
//}
//`
//
//var UpdateResponse = subnets.UpdateResponse{
//    Subnet: struct {
//        Id     string `json:"id"`
//        Status string `json:"status"`
//    }{
//        Id:     "c9aba52d-ec14-40cb-930f-c8153e93c2db",
//        Status: "ACTIVE",
//    },
//}
//
//func HandleUpdateSuccessfully(t *testing.T) {
//    th.Mux.HandleFunc("/vpcs/ea3b0efe-0d6a-4288-8b16-753504a994b9/su
// bnets/c9aba52d-ec14-40cb-930f-c8153e93c2db", func(w http.ResponseWriter, r *http.Request)
// {
//        th.TestMethod(t, r, "PUT")
//        th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
//
//        w.Header().Add("Content-Type", "application/json")
//        fmt.Fprintf(w, UpdateOutput)
//    })
//}

var GetOutput = `
{
  "security_group": {
    "id": "f7616338-fa30-42b8-bf6b-754c0701aab8",
    "name": "EricSG",
    "description": "",
    "security_group_rules": [{
      "direction": "egress",
      "ethertype": "IPv4",
      "id": "5658ade3-e87c-4016-afc6-b8efd3c2e306",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "4f87f0ea-1927-4b81-a4bb-83943b10585e",
      "remote_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "1438172b-20b8-4454-945c-9b5508a8e4f7",
      "protocol": "tcp",
      "port_range_max": 22,
      "port_range_min": 22,
      "remote_ip_prefix": "0.0.0.0/0",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "43fcaf3d-5406-4fc1-ab42-a18cc2660f01",
      "protocol": "tcp",
      "port_range_max": 3389,
      "port_range_min": 3389,
      "remote_ip_prefix": "0.0.0.0/0",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }]
  }
}
`

var GetResponse = securitygroups.SecurityGroup{
	ID:          "f7616338-fa30-42b8-bf6b-754c0701aab8",
	Name:        "EricSG",
	Description: "",
	SecurityGroupRules: []securitygroups.SecurityGroupRule{
		{
			Direction:       "egress",
			Ethertype:       "IPv4",
			ID:              "5658ade3-e87c-4016-afc6-b8efd3c2e306",
			SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
		},
		{
			Direction:       "ingress",
			Ethertype:       "IPv4",
			RemoteGroupId:   "f7616338-fa30-42b8-bf6b-754c0701aab8",
			ID:              "4f87f0ea-1927-4b81-a4bb-83943b10585e",
			SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
		},
		{
			Direction:       "ingress",
			Ethertype:       "IPv4",
			ID:              "1438172b-20b8-4454-945c-9b5508a8e4f7",
			Protocol:        "tcp",
			PortRangeMax:    &PortNumber_22,
			PortRangeMin:    &PortNumber_22,
			RemoteIpPrefix:  "0.0.0.0/0",
			SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
		},
		{
			Direction:       "ingress",
			Ethertype:       "IPv4",
			ID:              "43fcaf3d-5406-4fc1-ab42-a18cc2660f01",
			Protocol:        "tcp",
			PortRangeMax:    &PortNumber_3389,
			PortRangeMin:    &PortNumber_3389,
			RemoteIpPrefix:  "0.0.0.0/0",
			SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
		},
	},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/security-groups/f7616338-fa30-42b8-bf6b-754c0701aab8", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "security_groups": [{
    "id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
    "name": "Sys-default",
    "description": "default",
    "vpc_id": "default",
    "security_group_rules": [{
      "direction": "egress",
      "ethertype": "IPv4",
      "id": "2aafd879-c90f-49e0-87ee-0e4f0721c40c",
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }, {
      "direction": "egress",
      "ethertype": "IPv6",
      "id": "bf04230c-208a-4864-b0ff-876c21b33d0c",
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }, {
      "direction": "ingress",
      "ethertype": "IPv6",
      "id": "ee07e0d4-b4ec-4c68-a754-ff01652d4e11",
      "remote_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "164bcf26-c9c0-417a-abc5-c129ad0c1854",
      "protocol": "tcp",
      "port_range_max": 3389,
      "port_range_min": 3389,
      "remote_ip_prefix": "0.0.0.0/0",
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "ff0eb13a-9ea7-432c-9032-9316a298d43a",
      "protocol": "tcp",
      "port_range_max": 22,
      "port_range_min": 22,
      "remote_ip_prefix": "0.0.0.0/0",
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "d534b6e7-d001-4c36-b34d-ae70aa48f988",
      "remote_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
      "security_group_id": "7844d4b4-d78f-45dc-9465-2b4d1bca83a5"
    }]
  }, {
    "id": "f7616338-fa30-42b8-bf6b-754c0701aab8",
    "name": "EricSG",
    "description": "",
    "security_group_rules": [{
      "direction": "egress",
      "ethertype": "IPv4",
      "id": "5658ade3-e87c-4016-afc6-b8efd3c2e306",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "4f87f0ea-1927-4b81-a4bb-83943b10585e",
      "remote_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "1438172b-20b8-4454-945c-9b5508a8e4f7",
      "protocol": "tcp",
      "port_range_max": 22,
      "port_range_min": 22,
      "remote_ip_prefix": "0.0.0.0/0",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }, {
      "direction": "ingress",
      "ethertype": "IPv4",
      "id": "43fcaf3d-5406-4fc1-ab42-a18cc2660f01",
      "protocol": "tcp",
      "port_range_max": 3389,
      "port_range_min": 3389,
      "remote_ip_prefix": "0.0.0.0/0",
      "security_group_id": "f7616338-fa30-42b8-bf6b-754c0701aab8"
    }]
  }]
}
`

var ListResponse = []securitygroups.SecurityGroup{
	{
		ID:          "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
		Name:        "Sys-default",
		Description: "default",
		VpcId:       "default",
		SecurityGroupRules: []securitygroups.SecurityGroupRule{
			{
				Direction:       "egress",
				Ethertype:       "IPv4",
				ID:              "2aafd879-c90f-49e0-87ee-0e4f0721c40c",
				SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			}, {
				Direction:       "egress",
				Ethertype:       "IPv6",
				ID:              "bf04230c-208a-4864-b0ff-876c21b33d0c",
				SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			}, {
				Direction:       "ingress",
				Ethertype:       "IPv6",
				RemoteGroupId:   "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
				ID:              "ee07e0d4-b4ec-4c68-a754-ff01652d4e11",
				SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			}, {
				Direction:       "ingress",
				Ethertype:       "IPv4",
				ID:              "164bcf26-c9c0-417a-abc5-c129ad0c1854",
				Protocol:        "tcp",
				PortRangeMax:    &PortNumber_3389,
				PortRangeMin:    &PortNumber_3389,
				RemoteIpPrefix:  "0.0.0.0/0",
				SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			}, {
				Direction:       "ingress",
				Ethertype:       "IPv4",
				ID:              "ff0eb13a-9ea7-432c-9032-9316a298d43a",
				Protocol:        "tcp",
				PortRangeMax:    &PortNumber_22,
				PortRangeMin:    &PortNumber_22,
				RemoteIpPrefix:  "0.0.0.0/0",
				SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			}, {
				Direction:       "ingress",
				Ethertype:       "IPv4",
				RemoteGroupId:   "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
				ID:              "d534b6e7-d001-4c36-b34d-ae70aa48f988",
				SecurityGroupId: "7844d4b4-d78f-45dc-9465-2b4d1bca83a5",
			},
		},
	},
	{
		ID:          "f7616338-fa30-42b8-bf6b-754c0701aab8",
		Name:        "EricSG",
		Description: "",
		SecurityGroupRules: []securitygroups.SecurityGroupRule{
			{
				Direction:       "egress",
				Ethertype:       "IPv4",
				ID:              "5658ade3-e87c-4016-afc6-b8efd3c2e306",
				SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
			},
			{
				Direction:       "ingress",
				Ethertype:       "IPv4",
				RemoteGroupId:   "f7616338-fa30-42b8-bf6b-754c0701aab8",
				ID:              "4f87f0ea-1927-4b81-a4bb-83943b10585e",
				SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
			},
			{
				Direction:       "ingress",
				Ethertype:       "IPv4",
				ID:              "1438172b-20b8-4454-945c-9b5508a8e4f7",
				Protocol:        "tcp",
				PortRangeMax:    &PortNumber_22,
				PortRangeMin:    &PortNumber_22,
				RemoteIpPrefix:  "0.0.0.0/0",
				SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
			},
			{
				Direction:       "ingress",
				Ethertype:       "IPv4",
				ID:              "43fcaf3d-5406-4fc1-ab42-a18cc2660f01",
				Protocol:        "tcp",
				PortRangeMax:    &PortNumber_3389,
				PortRangeMin:    &PortNumber_3389,
				RemoteIpPrefix:  "0.0.0.0/0",
				SecurityGroupId: "f7616338-fa30-42b8-bf6b-754c0701aab8",
			},
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/security-groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/security-groups/2465d913-1084-4a6a-91e7-2fd6f490ecb3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
