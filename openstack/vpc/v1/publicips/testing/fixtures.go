package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/publicips"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
)

var CreateOutput = `
{
  "publicip": {
    "id": "84a71976-a8c2-42e0-8826-7fc27b876e42",
    "status": "PENDING_CREATE",
    "type": "5_bgp",
    "public_ip_address": "49.4.68.149",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "create_time": "2018-04-20 15:59:36",
    "bandwidth_size": 0
  }
}
`

var CreateResponse = publicips.PublicIPCreateResp{
	ID:              "84a71976-a8c2-42e0-8826-7fc27b876e42",
	Status:          "PENDING_CREATE",
	Type:            "5_bgp",
	PublicIpAddress: "49.4.68.149",
	TenantId:        "57e98940a77f4bb988a21a7d0603a626",
	CreateTime:      "2018-04-20 15:59:36",
	BandwidthSize:   0,
}

func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/publicips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, CreateOutput)
	})
}

var UpdateOutput = `
{
  "publicip": {
    "id": "84a71976-a8c2-42e0-8826-7fc27b876e42",
    "status": "DOWN",
    "type": "5_bgp",
    "port_id": "",
    "public_ip_address": "49.4.68.149",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "create_time": "2018-04-20 15:59:36",
    "bandwidth_size": 1
  }
}
`

var UpdateResponse = publicips.PublicIPUpdateResp{
	ID:              "84a71976-a8c2-42e0-8826-7fc27b876e42",
	Status:          "DOWN",
	Type:            "5_bgp",
	PortId:          "",
	PublicIpAddress: "49.4.68.149",
	TenantId:        "57e98940a77f4bb988a21a7d0603a626",
	CreateTime:      "2018-04-20 15:59:36",
	BandwidthSize:   1,
}

func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/publicips/84a71976-a8c2-42e0-8826-7fc27b876e42", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateOutput)
	})
}

var GetOutput = `
{
  "publicip": {
    "id": "84a71976-a8c2-42e0-8826-7fc27b876e42",
    "status": "DOWN",
    "type": "5_bgp",
    "public_ip_address": "49.4.68.149",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "create_time": "2018-04-20 15:59:36",
    "bandwidth_id": "ffdaac30-cd95-4955-a672-7b798285b68d",
    "bandwidth_name": "bandwidth-d62f",
    "bandwidth_share_type": "WHOLE",
    "bandwidth_size": 1
  }
}
`

var GetResponse = publicips.PublicIP{
	ID:                 "84a71976-a8c2-42e0-8826-7fc27b876e42",
	Status:             "DOWN",
	Type:               "5_bgp",
	PublicIpAddress:    "49.4.68.149",
	TenantId:           "57e98940a77f4bb988a21a7d0603a626",
	CreateTime:         "2018-04-20 15:59:36",
	BandwidthId:        "ffdaac30-cd95-4955-a672-7b798285b68d",
	BandwidthName:      "bandwidth-d62f",
	BandwidthShareType: "WHOLE",
	BandwidthSize:      1,
}

func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/publicips/84a71976-a8c2-42e0-8826-7fc27b876e42", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

var ListOutput = `
{
  "publicips": [{
    "id": "2412e868-f93a-4dfc-b171-5096baa27403",
    "status": "DOWN",
    "profile": {
      "user_id": null,
      "product_id": null,
      "region_id": null
    },
    "profile": {
              "user_id": "57e98940a77f4bb988a21a7d0603a626",
              "product_id": "da27d47d84ff4adba7de3ca3b0c9ce08",
              "region_id": "cn-north-1",
              "order_id": "CS1803212335SQ1UD"
            },
    "type": "5_sbgp",
    "public_ip_address": "114.115.219.38",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "create_time": "2018-04-20 16:32:16",
    "bandwidth_id": "916409b9-8ec8-41ce-b9d0-26002c4662ff",
    "bandwidth_name": "bandwidth-7eaa",
    "bandwidth_share_type": "WHOLE",
    "bandwidth_size": 1
  }, {
    "id": "3faa05bd-d878-44e2-a363-f6672a9761d3",
    "status": "DOWN",
    "profile": {
      "user_id": "57e98940a77f4bb988a21a7d0603a626",
      "product_id": "da27d47d84ff4adba7de3ca3b0c9ce08",
      "region_id": "cn-north-1",
      "order_id": "CS1803212335SQ1UD"
    },
    "profile": {
              "user_id": "57e98940a77f4bb988a21a7d0603a626",
              "product_id": "da27d47d84ff4adba7de3ca3b0c9ce08",
              "region_id": "cn-north-1",
              "order_id": "CS1803212335SQ1UD"
            },
    "type": "5_bgp",
    "public_ip_address": "49.4.22.32",
    "tenant_id": "57e98940a77f4bb988a21a7d0603a626",
    "create_time": "2018-03-21 15:36:34",
    "bandwidth_id": "3c43e46e-4af1-45b8-a84d-ee6d04488d2a",
    "bandwidth_name": "bandwidth-6f78",
    "bandwidth_share_type": "PER",
    "bandwidth_size": 1
  }]
}
`

var ListResponse = []publicips.PublicIP{
	{
		ID:                 "2412e868-f93a-4dfc-b171-5096baa27403",
		Status:             "DOWN",
		Type:               "5_sbgp",
		PublicIpAddress:    "114.115.219.38",
		TenantId:           "57e98940a77f4bb988a21a7d0603a626",
		CreateTime:         "2018-04-20 16:32:16",
		BandwidthId:        "916409b9-8ec8-41ce-b9d0-26002c4662ff",
		BandwidthName:      "bandwidth-7eaa",
		BandwidthShareType: "WHOLE",
		BandwidthSize:      1,
		Profile: publicips.Profile{
			UserID:    "57e98940a77f4bb988a21a7d0603a626",
			ProductID: "da27d47d84ff4adba7de3ca3b0c9ce08",
			RegionID:  "cn-north-1",
			OrderID:   "CS1803212335SQ1UD",
		},
	},
	{
		ID:                 "3faa05bd-d878-44e2-a363-f6672a9761d3",
		Status:             "DOWN",
		Type:               "5_bgp",
		PublicIpAddress:    "49.4.22.32",
		TenantId:           "57e98940a77f4bb988a21a7d0603a626",
		CreateTime:         "2018-03-21 15:36:34",
		BandwidthId:        "3c43e46e-4af1-45b8-a84d-ee6d04488d2a",
		BandwidthName:      "bandwidth-6f78",
		BandwidthShareType: "PER",
		BandwidthSize:      1,
		Profile: publicips.Profile{
			UserID:    "57e98940a77f4bb988a21a7d0603a626",
			ProductID: "da27d47d84ff4adba7de3ca3b0c9ce08",
			RegionID:  "cn-north-1",
			OrderID:   "CS1803212335SQ1UD",
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/publicips", func(w http.ResponseWriter, r *http.Request) {
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
		case "3faa05bd-d878-44e2-a363-f6672a9761d3":
			fmt.Fprintf(w, `{"publicips": []}`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}

func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/publicips/2412e868-f93a-4dfc-b171-5096baa27403 HTTP/1.1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "")
	})
}
