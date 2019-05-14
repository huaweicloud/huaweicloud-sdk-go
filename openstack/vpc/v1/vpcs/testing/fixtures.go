package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/vpcs"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var CreateOutput = `
{
  "vpc": {
    "id": "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
    "name": "ABC",
    "cidr": "192.168.0.0/16",
    "status": "CREATING",
    "routes": null
  }
}
`

var CreateResponse = vpcs.VPC{
	ID:     "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
	Name:   "ABC",
	Cidr:   "192.168.0.0/16",
	Status: "CREATING",
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "vpc": {
    "id": "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
    "name": "ABC-back",
    "cidr": "192.168.0.0/24",
    "status": "OK",
    "routes": null
  }
}
`

var UpdateResponse = vpcs.VPC{
	ID:     "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
	Name:   "ABC-back",
	Cidr:   "192.168.0.0/24",
	Status: "OK",
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/vpcs/7ffddb5f-6731-43d8-9476-1444aaa40bc0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "vpc": {
    "id": "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
    "name": "ABC-back",
    "cidr": "192.168.0.0/24",
    "status": "OK",
    "routes": []
  }
}
`

var GetResponse = vpcs.VPC{
	ID:     "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
	Name:   "ABC-back",
	Cidr:   "192.168.0.0/24",
	Status: "OK",
	Routes: []vpcs.Route{},
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/vpcs/7ffddb5f-6731-43d8-9476-1444aaa40bc0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "vpcs": [{
    "id": "773c3c42-d315-417b-9063-87091713148c",
    "name": "vpc-c8cb",
    "cidr": "192.168.0.0/16",
    "status": "OK",
    "routes": []
  }, {
    "id": "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
    "name": "ABC-back",
    "cidr": "192.168.0.0/24",
    "status": "OK",
    "routes": []
  }]
}
`

var ListResponse = []vpcs.VPC{
	{
		ID:     "773c3c42-d315-417b-9063-87091713148c",
		Name:   "vpc-c8cb",
		Cidr:   "192.168.0.0/16",
		Status: "OK",
		Routes: []vpcs.Route{},
	},
	{
		ID:     "7ffddb5f-6731-43d8-9476-1444aaa40bc0",
		Name:   "ABC-back",
		Cidr:   "192.168.0.0/24",
		Status: "OK",
		Routes: []vpcs.Route{},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/vpcs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		err := r.ParseForm()
		if err != nil {
			t.Fatalf("parse form failed: [%s]", r.Form)
		}
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, ListOutput)
		case "7ffddb5f-6731-43d8-9476-1444aaa40bc0":
			fmt.Fprintf(w, `{"vpcs": []}`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/vpcs/7ffddb5f-6731-43d8-9476-1444aaa40bc0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
