package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/subnets"
	"github.com/gophercloud/gophercloud/testhelper/client"
	th "github.com/gophercloud/gophercloud/testhelper"
)


var ListOutput = `
{
  "subnets": [{
    "id": "9a78d8b5-4f00-4de9-b0d8-1228afb27726",
    "name": "subnet-3d95",
    "cidr": "192.168.0.0/24",
    "dnsList": ["100.125.1.250", "114.114.114.114"],
    "status": "ACTIVE",
    "vpc_id": "ea3b0efe-0d6a-4288-8b16-753504a994b9",
    "gateway_ip": "192.168.0.1",
    "dhcp_enable": true,
    "primary_dns": "100.125.1.250",
    "secondary_dns": "114.114.114.114",
    "availability_zone": "cn-north-1a",
    "neutron_network_id": "9a78d8b5-4f00-4de9-b0d8-1228afb27726",
    "neutron_subnet_id": "4d1d4f45-6375-4821-822e-44c92d12a58c"
  }, {
    "id": "c9aba52d-ec14-40cb-930f-c8153e93c2db",
    "name": "subnet",
    "cidr": "192.168.20.0/24",
    "dnsList": ["114.114.114.114", "114.114.115.115"],
    "status": "ACTIVE",
    "vpc_id": "ea3b0efe-0d6a-4288-8b16-753504a994b9",
    "gateway_ip": "192.168.20.1",
    "dhcp_enable": true,
    "primary_dns": "114.114.114.114",
    "secondary_dns": "114.114.115.115",
    "availability_zone": "cn-north-1a",
    "neutron_network_id": "c9aba52d-ec14-40cb-930f-c8153e93c2db",
    "neutron_subnet_id": "c557e272-dea4-40ee-931b-36c33fb192b2"
  }]
}
`


var ListResponse = []subnets.Subnet{
		{
			ID:               "9a78d8b5-4f00-4de9-b0d8-1228afb27726",
			Name:             "subnet-3d95",
			Cidr:             "192.168.0.0/24",
			DNSList:          []string{"100.125.1.250", "114.114.114.114"},
			Status:           "ACTIVE",
			VpcID:            "ea3b0efe-0d6a-4288-8b16-753504a994b9",
			GatewayIP:        "192.168.0.1",
			DhcpEnable:       enable,
			PrimaryDNS:       "100.125.1.250",
			SecondaryDNS:     "114.114.114.114",
			AvailabilityZone: "cn-north-1a",
			NeutronNetworkID: "9a78d8b5-4f00-4de9-b0d8-1228afb27726",
			NeutronSubnetID:  "4d1d4f45-6375-4821-822e-44c92d12a58c",
		},
		{
			ID:               "c9aba52d-ec14-40cb-930f-c8153e93c2db",
			Name:             "subnet",
			Cidr:             "192.168.20.0/24",
			DNSList:          []string{"114.114.114.114", "114.114.115.115"},
			Status:           "ACTIVE",
			VpcID:            "ea3b0efe-0d6a-4288-8b16-753504a994b9",
			GatewayIP:        "192.168.20.1",
			DhcpEnable:       enable,
			PrimaryDNS:       "114.114.114.114",
			SecondaryDNS:     "114.114.115.115",
			AvailabilityZone: "cn-north-1a",
			NeutronNetworkID: "c9aba52d-ec14-40cb-930f-c8153e93c2db",
			NeutronSubnetID:  "c557e272-dea4-40ee-931b-36c33fb192b2",
		},
	}



func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, ListOutput)
		case "c9aba52d-ec14-40cb-930f-c8153e93c2db":
			fmt.Fprintf(w, `{"subnets": []}`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/vpcs/ea3b0efe-0d6a-4288-8b16-753504a994b9/subnets/c9aba52d-ec14-40cb-930f-c8153e93c2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
var CreateOutput = `
{
  "subnet": {
    "id": "c9aba52d-ec14-40cb-930f-c8153e93c2db",
    "name": "subnet",
    "cidr": "192.168.20.0/24",
    "dnsList": ["114.114.114.114", "114.114.115.115"],
    "status": "UNKNOWN",
    "vpc_id": "ea3b0efe-0d6a-4288-8b16-753504a994b9",
    "gateway_ip": "192.168.20.1",
    "dhcp_enable": true,
    "primary_dns": "114.114.114.114",
    "secondary_dns": "114.114.115.115",
    "availability_zone": "cn-north-1a",
    "neutron_network_id": "c9aba52d-ec14-40cb-930f-c8153e93c2db",
    "neutron_subnet_id": "c557e272-dea4-40ee-931b-36c33fb192b2"
  }
}
`

var enable=true

var CreateResponse = subnets.Subnet{
	ID:               "c9aba52d-ec14-40cb-930f-c8153e93c2db",
	Name:             "subnet",
	Cidr:             "192.168.20.0/24",
	DNSList:          []string{"114.114.114.114", "114.114.115.115"},
	Status:           "UNKNOWN",
	VpcID:            "ea3b0efe-0d6a-4288-8b16-753504a994b9",
	GatewayIP:        "192.168.20.1",
	DhcpEnable:       enable,
	PrimaryDNS:       "114.114.114.114",
	SecondaryDNS:     "114.114.115.115",
	AvailabilityZone: "cn-north-1a",
	NeutronNetworkID: "c9aba52d-ec14-40cb-930f-c8153e93c2db",
	NeutronSubnetID:  "c557e272-dea4-40ee-931b-36c33fb192b2",
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "subnet": {
    "id": "c9aba52d-ec14-40cb-930f-c8153e93c2db",
    "status": "ACTIVE"
  }
}
`

var UpdateResponse = subnets.UpdateResp{
	ID:     "c9aba52d-ec14-40cb-930f-c8153e93c2db",
	Status: "ACTIVE",
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/vpcs/ea3b0efe-0d6a-4288-8b16-753504a994b9/subnets/c9aba52d-ec14-40cb-930f-c8153e93c2db", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "subnet": {
    "id": "c9aba52d-ec14-40cb-930f-c8153e93c2db",
    "name": "subnet",
    "cidr": "192.168.20.0/24",
    "dnsList": ["114.114.114.114", "114.114.115.115"],
    "status": "ACTIVE",
    "vpc_id": "ea3b0efe-0d6a-4288-8b16-753504a994b9",
    "gateway_ip": "192.168.20.1",
    "dhcp_enable": true,
    "primary_dns": "114.114.114.114",
    "secondary_dns": "114.114.115.115",
    "availability_zone": "cn-north-1a",
    "neutron_network_id": "c9aba52d-ec14-40cb-930f-c8153e93c2db",
    "neutron_subnet_id": "c557e272-dea4-40ee-931b-36c33fb192b2"
  }
}
`

var GetResponse = subnets.Subnet{
	ID:               "c9aba52d-ec14-40cb-930f-c8153e93c2db",
	Name:             "subnet",
	Cidr:             "192.168.20.0/24",
	DNSList:          []string{"114.114.114.114", "114.114.115.115"},
	Status:           "ACTIVE",
	VpcID:            "ea3b0efe-0d6a-4288-8b16-753504a994b9",
	GatewayIP:        "192.168.20.1",
	DhcpEnable:       enable,
	PrimaryDNS:       "114.114.114.114",
	SecondaryDNS:     "114.114.115.115",
	AvailabilityZone: "cn-north-1a",
	NeutronNetworkID: "c9aba52d-ec14-40cb-930f-c8153e93c2db",
	NeutronSubnetID:  "c557e272-dea4-40ee-931b-36c33fb192b2",
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/subnets/c9aba52d-ec14-40cb-930f-c8153e93c2db", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}