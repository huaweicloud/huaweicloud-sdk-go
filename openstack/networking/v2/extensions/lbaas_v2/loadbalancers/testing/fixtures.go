package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/loadbalancers"
)

// LoadbalancersListBody contains the canned body of a loadbalancer list response.
const LoadbalancersListBody = `
{
	"loadbalancers":[
	         {
			"id": "c331058c-6a40-4144-948e-b9fb1df9db4b",
			"tenant_id": "54030507-44f7-473c-9342-b4d14a95f692",
			"name": "web_lb",
			"description": "lb config for the web tier",
			"vip_subnet_id": "8a49c438-848f-467b-9655-ea1548708154",
			"vip_address": "10.30.176.47",
			"vip_port_id": "2a22e552-a347-44fd-b530-1f2b1b2a6735",
			"provider": "haproxy",
			"admin_state_up": true,
			"provisioning_status": "ACTIVE",
			"operating_status": "ONLINE"
		},
		{
			"id": "36e08a3e-a78f-4b40-a229-1e7e23eee1ab",
			"tenant_id": "54030507-44f7-473c-9342-b4d14a95f692",
			"name": "db_lb",
			"description": "lb config for the db tier",
			"vip_subnet_id": "9cedb85d-0759-4898-8a4b-fa5a5ea10086",
			"vip_address": "10.30.176.48",
			"vip_port_id": "2bf413c8-41a9-4477-b505-333d5cbe8b55",
			"provider": "haproxy",
			"admin_state_up": true,
			"provisioning_status": "PENDING_CREATE",
			"operating_status": "OFFLINE"
		}
	]
}
`

// SingleLoadbalancerBody is the canned body of a Get request on an existing loadbalancer.
const SingleLoadbalancerBody = `
{
	"loadbalancer": {
		"id": "36e08a3e-a78f-4b40-a229-1e7e23eee1ab",
		"tenant_id": "54030507-44f7-473c-9342-b4d14a95f692",
		"name": "db_lb",
		"description": "lb config for the db tier",
		"vip_subnet_id": "9cedb85d-0759-4898-8a4b-fa5a5ea10086",
		"vip_address": "10.30.176.48",
		"vip_port_id": "2bf413c8-41a9-4477-b505-333d5cbe8b55",
		"provider": "haproxy",
		"admin_state_up": true,
		"provisioning_status": "PENDING_CREATE",
		"operating_status": "OFFLINE"
	}
}
`

// PostUpdateLoadbalancerBody is the canned response body of a Update request on an existing loadbalancer.
const PostUpdateLoadbalancerBody = `
{
	"loadbalancer": {
		"id": "36e08a3e-a78f-4b40-a229-1e7e23eee1ab",
		"tenant_id": "54030507-44f7-473c-9342-b4d14a95f692",
		"name": "NewLoadbalancerName",
		"description": "lb config for the db tier",
		"vip_subnet_id": "9cedb85d-0759-4898-8a4b-fa5a5ea10086",
		"vip_address": "10.30.176.48",
		"vip_port_id": "2bf413c8-41a9-4477-b505-333d5cbe8b55",
		"provider": "haproxy",
		"admin_state_up": true,
		"provisioning_status": "PENDING_CREATE",
		"operating_status": "OFFLINE"
	}
}
`

// SingleLoadbalancerBody is the canned body of a Get request on an existing loadbalancer.
const LoadbalancerStatuesesTree = `
{
	"statuses": {
		"loadbalancer": {
			"name": "lb-jy",
			"provisioning_status": "ACTIVE",
			"listeners": [{
				"name": "listener-jy-1",
				"provisioning_status": "ACTIVE",
				"pools": [{
					"name": "pool-jy-1",
					"provisioning_status": "ACTIVE",
					"healthmonitor": {
						"type": "TCP",
						"id": "7422b51a-0ed2-4702-9429-4f88349276c6",
						"name": "",
						"provisioning_status": "ACTIVE"
					},
					"members": [{
						"protocol_port": 80,
						"address": "192.168.44.11",
						"id": "7bbf7151-0dce-4087-b316-06c7fa17b894",
						"operating_status": "ONLINE",
						"provisioning_status": "ACTIVE"
					}],
					"id": "c54b3286-2349-4c5c-ade1-e6bb0b26ad18",
					"operating_status": "ONLINE"
				}],
				"l7policies": [],
				"id": "eb84c5b4-9bc5-4bee-939d-3900fb05dc7b",
				"operating_status": "ONLINE"
			}],
			"pools": [{
				"name": "pool-jy-1",
				"provisioning_status": "ACTIVE",
				"healthmonitor": {
					"type": "TCP",
					"id": "7422b51a-0ed2-4702-9429-4f88349276c6",
					"name": "",
					"provisioning_status": "ACTIVE"
				},
				"members": [{
					"protocol_port": 80,
					"address": "192.168.44.11",
					"id": "7bbf7151-0dce-4087-b316-06c7fa17b894",
					"operating_status": "ONLINE",
					"provisioning_status": "ACTIVE"
				}],
				"id": "38278031-cfca-44be-81be-a412f618773b",
				"operating_status": "ONLINE"
			}],
			"id": "38278031-cfca-44be-81be-a412f618773b",
			"operating_status": "ONLINE"
		}
	}
}
`

var (
	LoadbalancerWeb = loadbalancers.LoadBalancer{
		ID:                 "c331058c-6a40-4144-948e-b9fb1df9db4b",
		TenantID:           "54030507-44f7-473c-9342-b4d14a95f692",
		Name:               "web_lb",
		Description:        "lb config for the web tier",
		VipSubnetID:        "8a49c438-848f-467b-9655-ea1548708154",
		VipAddress:         "10.30.176.47",
		VipPortID:          "2a22e552-a347-44fd-b530-1f2b1b2a6735",
		Provider:           "haproxy",
		AdminStateUp:       true,
		ProvisioningStatus: "ACTIVE",
		OperatingStatus:    "ONLINE",
	}
	LoadbalancerDb = loadbalancers.LoadBalancer{
		ID:                 "36e08a3e-a78f-4b40-a229-1e7e23eee1ab",
		TenantID:           "54030507-44f7-473c-9342-b4d14a95f692",
		Name:               "db_lb",
		Description:        "lb config for the db tier",
		VipSubnetID:        "9cedb85d-0759-4898-8a4b-fa5a5ea10086",
		VipAddress:         "10.30.176.48",
		VipPortID:          "2bf413c8-41a9-4477-b505-333d5cbe8b55",
		Provider:           "haproxy",
		AdminStateUp:       true,
		ProvisioningStatus: "PENDING_CREATE",
		OperatingStatus:    "OFFLINE",
	}
	LoadbalancerUpdated = loadbalancers.LoadBalancer{
		ID:                 "36e08a3e-a78f-4b40-a229-1e7e23eee1ab",
		TenantID:           "54030507-44f7-473c-9342-b4d14a95f692",
		Name:               "NewLoadbalancerName",
		Description:        "lb config for the db tier",
		VipSubnetID:        "9cedb85d-0759-4898-8a4b-fa5a5ea10086",
		VipAddress:         "10.30.176.48",
		VipPortID:          "2bf413c8-41a9-4477-b505-333d5cbe8b55",
		Provider:           "haproxy",
		AdminStateUp:       true,
		ProvisioningStatus: "PENDING_CREATE",
		OperatingStatus:    "OFFLINE",
	}
	LoadbalancerStatuses = loadbalancers.LoadbalancerStatus{
		ID:                 "38278031-cfca-44be-81be-a412f618773b",
		Name:               "lb-jy",
		OperatingStatus:    "ONLINE",
		ProvisioningStatus: "ACTIVE",
		Listeners: []loadbalancers.Listener{{
			Name:               "listener-jy-1",
			ProvisioningStatus: "ACTIVE",
			Pools: []loadbalancers.Pool{{
				ID:                 "c54b3286-2349-4c5c-ade1-e6bb0b26ad18",
				OperatingStatus:    "ONLINE",
				Name:               "pool-jy-1",
				ProvisioningStatus: "ACTIVE",
				HealthMonitor: loadbalancers.HealthMonitor{
					Type:               "TCP",
					ID:                 "7422b51a-0ed2-4702-9429-4f88349276c6",
					Name:               "",
					ProvisioningStatus: "ACTIVE",
				},
				Members: []loadbalancers.Member{{
					ProtocolPort:       80,
					Address:            "192.168.44.11",
					ID:                 "7bbf7151-0dce-4087-b316-06c7fa17b894",
					OperatingStatus:    "ONLINE",
					ProvisioningStatus: "ACTIVE",
				}},
			}},
			L7Policies:      []interface{}{},
			ID:              "eb84c5b4-9bc5-4bee-939d-3900fb05dc7b",
			OperatingStatus: "ONLINE",
		}},
		Pools: []loadbalancers.Pool{{
			ID:                 "38278031-cfca-44be-81be-a412f618773b",
			OperatingStatus:    "ONLINE",
			Name:               "pool-jy-1",
			ProvisioningStatus: "ACTIVE",
			HealthMonitor: loadbalancers.HealthMonitor{
				Type:               "TCP",
				ID:                 "7422b51a-0ed2-4702-9429-4f88349276c6",
				Name:               "",
				ProvisioningStatus: "ACTIVE",
			},
			Members: []loadbalancers.Member{{
				ProtocolPort:       80,
				Address:            "192.168.44.11",
				ID:                 "7bbf7151-0dce-4087-b316-06c7fa17b894",
				OperatingStatus:    "ONLINE",
				ProvisioningStatus: "ACTIVE",
			}},
		},
		},
	}
)

// HandleLoadbalancerListSuccessfully sets up the test server to respond to a loadbalancer List request.
func HandleLoadbalancerListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/loadbalancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, LoadbalancersListBody)
		case "45e08a3e-a78f-4b40-a229-1e7e23eee1ab":
			fmt.Fprintf(w, `{ "loadbalancers": [] }`)
		default:
			t.Fatalf("/v2.0/lbaas/loadbalancers invoked with unexpected marker=[%s]", marker)
		}
	})
}

// HandleLoadbalancerCreationSuccessfully sets up the test server to respond to a loadbalancer creation request
// with a given response.
func HandleLoadbalancerCreationSuccessfully(t *testing.T, response string) {
	th.Mux.HandleFunc("/v2.0/lbaas/loadbalancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"loadbalancer": {
				"name": "db_lb",
				"vip_subnet_id": "9cedb85d-0759-4898-8a4b-fa5a5ea10086",
				"vip_address": "10.30.176.48",
				"provider": "haproxy",
				"admin_state_up": true
			}
		}`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
}

// HandleLoadbalancerGetSuccessfully sets up the test server to respond to a loadbalancer Get request.
func HandleLoadbalancerGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/loadbalancers/36e08a3e-a78f-4b40-a229-1e7e23eee1ab", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprintf(w, SingleLoadbalancerBody)
	})
}

// HandleLoadbalancerGetStatusesTree sets up the test server to respond to a loadbalancer Get statuses tree request.
func HandleLoadbalancerGetStatusesTree(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/loadbalancers/38278031-cfca-44be-81be-a412f618773b/statuses", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		fmt.Fprintf(w, LoadbalancerStatuesesTree)

	})

}

// HandleLoadbalancerDeletionSuccessfully sets up the test server to respond to a loadbalancer deletion request.
func HandleLoadbalancerDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/loadbalancers/36e08a3e-a78f-4b40-a229-1e7e23eee1ab", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleLoadbalancerUpdateSuccessfully sets up the test server to respond to a loadbalancer Update request.
func HandleLoadbalancerUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/loadbalancers/36e08a3e-a78f-4b40-a229-1e7e23eee1ab", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `{
			"loadbalancer": {
				"name": "NewLoadbalancerName"
			}
		}`)

		fmt.Fprintf(w, PostUpdateLoadbalancerBody)
	})
}
