/*

package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

// List Output is a sample response to a List call.
const ListOutput = `
{
    "links": {
      "self": "http://example.com:9001/v2/zones"
    },
    "metadata": {
      "total_count": 2
    },
    "zones": [
        {
            "id": "a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
            "pool_id": "572ba08c-d929-4c70-8e42-03824bb24ca2",
            "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
            "name": "example.org.",
            "email": "joe@example.org",
            "ttl": 7200,
            "serial": 1404757531,
            "status": "ACTIVE",
            "action": "CREATE",
            "description": "This is an example zone.",
            "masters": [],
            "type": "PRIMARY",
            "transferred_at": null,
            "version": 1,
            "created_at": "2014-07-07T18:25:31.275934",
            "updated_at": null,
            "links": {
              "self": "https://127.0.0.1:9001/v2/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3"
            }
        },
        {
            "id": "34c4561c-9205-4386-9df5-167436f5a222",
            "pool_id": "572ba08c-d929-4c70-8e42-03824bb24ca2",
            "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
            "name": "foo.example.com.",
            "email": "joe@foo.example.com",
            "ttl": 7200,
            "serial": 1488053571,
            "status": "ACTIVE",
            "action": "CREATE",
            "description": "This is another example zone.",
            "masters": ["example.com."],
            "type": "PRIMARY",
            "transferred_at": null,
            "version": 1,
            "created_at": "2014-07-07T18:25:31.275934",
            "updated_at": "2015-02-25T20:23:01.234567",
            "links": {
              "self": "https://127.0.0.1:9001/v2/zones/34c4561c-9205-4386-9df5-167436f5a222"
            }
        }
    ]
}
`

// GetOutput is a sample response to a Get call.
const GetOutput = `
{
    "id": "a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
    "pool_id": "572ba08c-d929-4c70-8e42-03824bb24ca2",
    "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
    "name": "example.org.",
    "email": "joe@example.org",
    "ttl": 7200,
    "serial": 1404757531,
    "status": "ACTIVE",
    "action": "CREATE",
    "description": "This is an example zone.",
    "masters": [],
    "type": "PRIMARY",
    "transferred_at": null,
    "version": 1,
    "created_at": "2014-07-07T18:25:31.275934",
    "updated_at": null,
    "links": {
      "self": "https://127.0.0.1:9001/v2/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3"
    }
}
`

// FirstZone is the first result in ListOutput
var FirstZoneCreatedAt, _ = time.Parse(gophercloud.RFC3339MilliNoZ, "2014-07-07T18:25:31.275934")
var FirstZone = zones.Zone{
	ID:          "a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
	PoolID:      "572ba08c-d929-4c70-8e42-03824bb24ca2",
	ProjectID:   "4335d1f0-f793-11e2-b778-0800200c9a66",
	Name:        "example.org.",
	Email:       "joe@example.org",
	TTL:         7200,
	Serial:      1404757531,
	Status:      "ACTIVE",
	Action:      "CREATE",
	Description: "This is an example zone.",
	Masters:     []string{},
	Type:        "PRIMARY",
	Version:     1,
	CreatedAt:   FirstZoneCreatedAt,
	Links: map[string]interface{}{
		"self": "https://127.0.0.1:9001/v2/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
	},
}

var SecondZoneCreatedAt, _ = time.Parse(gophercloud.RFC3339MilliNoZ, "2014-07-07T18:25:31.275934")
var SecondZoneUpdatedAt, _ = time.Parse(gophercloud.RFC3339MilliNoZ, "2015-02-25T20:23:01.234567")
var SecondZone = zones.Zone{
	ID:          "34c4561c-9205-4386-9df5-167436f5a222",
	PoolID:      "572ba08c-d929-4c70-8e42-03824bb24ca2",
	ProjectID:   "4335d1f0-f793-11e2-b778-0800200c9a66",
	Name:        "foo.example.com.",
	Email:       "joe@foo.example.com",
	TTL:         7200,
	Serial:      1488053571,
	Status:      "ACTIVE",
	Action:      "CREATE",
	Description: "This is another example zone.",
	Masters:     []string{"example.com."},
	Type:        "PRIMARY",
	Version:     1,
	CreatedAt:   SecondZoneCreatedAt,
	UpdatedAt:   SecondZoneUpdatedAt,
	Links: map[string]interface{}{
		"self": "https://127.0.0.1:9001/v2/zones/34c4561c-9205-4386-9df5-167436f5a222",
	},
}

// ExpectedZonesSlice is the slice of results that should be parsed
// from ListOutput, in the expected order.
var ExpectedZonesSlice = []zones.Zone{FirstZone, SecondZone}

// HandleListSuccessfully configures the test server to respond to a List request.
func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

// HandleGetSuccessfully configures the test server to respond to a List request.
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

// CreateZoneRequest is a sample request to create a zone.
const CreateZoneRequest = `
{
    "name": "example.org.",
    "email": "joe@example.org",
    "type": "PRIMARY",
    "ttl": 7200,
    "description": "This is an example zone."
}
`

// CreateZoneResponse is a sample response to a create request.
const CreateZoneResponse = `
{
    "id": "a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
    "pool_id": "572ba08c-d929-4c70-8e42-03824bb24ca2",
    "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
    "name": "example.org.",
    "email": "joe@example.org",
    "ttl": 7200,
    "serial": 1404757531,
    "status": "ACTIVE",
    "action": "CREATE",
    "description": "This is an example zone.",
    "masters": [],
    "type": "PRIMARY",
    "transferred_at": null,
    "version": 1,
    "created_at": "2014-07-07T18:25:31.275934",
    "updated_at": null,
    "links": {
      "self": "https://127.0.0.1:9001/v2/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3"
    }
}
`

// CreatedZone is the expected created zone
var CreatedZone = FirstZone

// HandleZoneCreationSuccessfully configures the test server to respond to a Create request.
func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateZoneRequest)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateZoneResponse)
	})
}

// UpdateZoneRequest is a sample request to update a zone.
const UpdateZoneRequest = `
{
    "ttl": 600,
    "description": "Updated Description"
}
`

// UpdateZoneResponse is a sample response to update a zone.
const UpdateZoneResponse = `
{
    "id": "a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
    "pool_id": "572ba08c-d929-4c70-8e42-03824bb24ca2",
    "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
    "name": "example.org.",
    "email": "joe@example.org",
    "ttl": 600,
    "serial": 1404757531,
    "status": "PENDING",
    "action": "UPDATE",
    "description": "Updated Description",
    "masters": [],
    "type": "PRIMARY",
    "transferred_at": null,
    "version": 1,
    "created_at": "2014-07-07T18:25:31.275934",
    "updated_at": null,
    "links": {
      "self": "https://127.0.0.1:9001/v2/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3"
    }
}
`

// HandleZoneUpdateSuccessfully configures the test server to respond to an Update request.
func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, UpdateZoneRequest)

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, UpdateZoneResponse)
		})
}

// DeleteZoneResponse is a sample response to update a zone.
const DeleteZoneResponse = `
{
    "id": "a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
    "pool_id": "572ba08c-d929-4c70-8e42-03824bb24ca2",
    "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
    "name": "example.org.",
    "email": "joe@example.org",
    "ttl": 600,
    "serial": 1404757531,
    "status": "PENDING",
    "action": "DELETE",
    "description": "Updated Description",
    "masters": [],
    "type": "PRIMARY",
    "transferred_at": null,
    "version": 1,
    "created_at": "2014-07-07T18:25:31.275934",
    "updated_at": null,
    "links": {
      "self": "https://127.0.0.1:9001/v2/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3"
    }
}
`

// HandleZoneDeleteSuccessfully configures the test server to respond to an Delete request.
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.WriteHeader(http.StatusAccepted)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, DeleteZoneResponse)
		})
}

*/
package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"net/http"
	"testing"
)

// List Output is a sample response to a List call.
const ListOutput = `
{
	"links": {
		"self": "http://example.com:9001/v2/zones"
	},
	"metadata": {
		"total_count": 2
	},
	"zones": [{
		"id": "ff8080825b8fc86c015b94bc6f8712c3",
		"name": "example.com.",
		"description": "This is an example zone.",
		"email": "xx@example.com",
		"ttl": 300,
		"serial": 0,
		"masters": [],
		"status": "ACTIVE",
		"links": {
			"self": "https://Endpoint/v2/zones/ff8080825b8fc86c015b94bc6f8712c3"
		},
		"pool_id": "ff8080825ab738f4015ab7513298010e",
		"project_id": "e55c6f3dc4e34c9f86353b664ae0e70c",
		"zone_type": "private",
		"created_at": "2017-04-22T08:17:08.997",
		"updated_at": "2017-04-22T08:17:09.997",
		"record_num": 2,
		"routers": [{
			"status": "ACTIVE",
			"router_id": "19664294-0bf6-4271-ad3a-94b8c79c6558",
			"router_region": "xx"
		},
		{
			"status": "ACTIVE",
			"router_id": "f0791650-db8c-4a20-8a44-a06c6e24b15b",
			"router_region": "xx"
		}]
	},
	{
		"id": "ff8080825b8fc86c015b94bc6f871223",
		"name": "example.com.",
		"description": "This is an example zone.",
		"email": "xx@example.com",
		"ttl": 300,
		"serial": 0,
		"masters": [],
		"status": "ACTIVE",
		"links": {
			"self": "https://Endpoint/v2/zones/ff8080825b8fc86c015b94bc6f8712c3"
		},
		"pool_id": "ff8080825ab738f4015ab7513298010e",
		"project_id": "e55c6f3dc4e34c9f86353b664ae0e70c",
		"zone_type": "private",
		"created_at": "2017-04-22T08:17:08.997",
		"updated_at": "2017-04-22T08:17:09.997",
		"record_num": 2,
		"routers": [{
			"status": "ACTIVE",
			"router_id": "19664294-0bf6-4271-ad3a-94b8c79c6558",
			"router_region": "xx"
		},
		{
			"status": "ACTIVE",
			"router_id": "f0791650-db8c-4a20-8a44-a06c6e24b15b",
			"router_region": "xx"
		}]
	}]
}
`

// GetOutput is a sample response to a Get call.
const GetOutput = `
{
    "id": "ff8080825b8fc86c015b94bc6f8712c3",
    "name": "example.com.",
    "description": "This is an example zone.",
    "email": "xx@example.com",
    "ttl": 300,
    "serial": 0,
    "masters": [],
    "status": "ACTIVE",
    "links": {
        "self": "https://Endpoint/v2/zones/ff8080825b8fc86c015b94bc6f8712c3"
    },
    "pool_id": "ff8080825ab738f4015ab7513298010e",
    "project_id": "e55c6f3dc4e34c9f86353b664ae0e70c",
    "zone_type": "private",
    "created_at": "2017-04-22T08:17:08.997",
    "updated_at": "2017-04-22T08:17:09.997",
    "record_num": 2,
    "routers": [
        {
            "status": "ACTIVE",
            "router_id": "19664294-0bf6-4271-ad3a-94b8c79c6558",
            "router_region": "xx"
        },
        {
            "status": "ACTIVE",
            "router_id": "f0791650-db8c-4a20-8a44-a06c6e24b15b",
            "router_region": "xx"
        }
    ]
}
`

var FirstZone = zones.Zone{
	ID:          "ff8080825b8fc86c015b94bc6f8712c3",
	Name:        "example.com.",
	Email:       "xx@example.com",
	Serial:      0,
	Status:      "ACTIVE",
	Description: "This is an example zone.",
	Masters:     []string{},
	CreatedAt:   "2017-04-22T08:17:08.997",
	UpdatedAt:   "2017-04-22T08:17:09.997",
	TTL:         300,
	Links: zones.Link{
		Self: "https://Endpoint/v2/zones/ff8080825b8fc86c015b94bc6f8712c3",
	},
	PoolId:    "ff8080825ab738f4015ab7513298010e",
	ProjectId: "e55c6f3dc4e34c9f86353b664ae0e70c",
	ZoneType:  "private",
	RecordNum: 2,
	Routers: []zones.AssociateRouterResponse{
		{
			RouterId:     "19664294-0bf6-4271-ad3a-94b8c79c6558",
			Status:       "ACTIVE",
			RouterRegion: "xx",
		},
		{
			RouterId:     "f0791650-db8c-4a20-8a44-a06c6e24b15b",
			Status:       "ACTIVE",
			RouterRegion: "xx",
		},
	},
}

var SecondZone = zones.Zone{
	ID:          "ff8080825b8fc86c015b94bc6f871223",
	Name:        "example.com.",
	Email:       "xx@example.com",
	Serial:      0,
	Status:      "ACTIVE",
	Description: "This is an example zone.",
	Masters:     []string{},
	CreatedAt:   "2017-04-22T08:17:08.997",
	UpdatedAt:   "2017-04-22T08:17:09.997",
	TTL:         300,
	Links: zones.Link{
		Self: "https://Endpoint/v2/zones/ff8080825b8fc86c015b94bc6f8712c3",
	},
	PoolId:    "ff8080825ab738f4015ab7513298010e",
	ProjectId: "e55c6f3dc4e34c9f86353b664ae0e70c",
	ZoneType:  "private",
	RecordNum: 2,
	Routers: []zones.AssociateRouterResponse{
		{
			RouterId:     "19664294-0bf6-4271-ad3a-94b8c79c6558",
			Status:       "ACTIVE",
			RouterRegion: "xx",
		},
		{
			RouterId:     "f0791650-db8c-4a20-8a44-a06c6e24b15b",
			Status:       "ACTIVE",
			RouterRegion: "xx",
		},
	},
}

// ExpectedZonesSlice is the slice of results that should be parsed
// from ListOutput, in the expected order.
var ExpectedZonesSlice = []zones.Zone{FirstZone, SecondZone}

// HandleListSuccessfully configures the test server to respond to a List request.
func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

// HandleGetSuccessfully configures the test server to respond to a List request.
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones/ff8080825b8fc86c015b94bc6f8712c3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}

// CreateZoneRequest is a sample request to create a zone.
const CreateZoneRequest = `
{
  "name": "example.org.",
  "email": "joe@example.org",
  "zone_type": "PRIMARY",
  "description": "This is an example zone.",
  "router": {
    "router_id": "19664294-0bf6-4271-ad3a-94b8c79c6558",
    "router_region": "xx"
  }
}
`

// CreateZoneResponse is a sample response to a create request.
const CreateZoneResponse = `
{
    "id": "ff8080825b8fc86c015b94bc6f8712c3",
    "name": "example.com.",
    "description": "This is an example zone.",
    "email": "xx@example.com",
    "ttl": 300,
    "serial": 0,
    "masters": [],
    "status": "ACTIVE",
    "links": {
        "self": "https://Endpoint/v2/zones/ff8080825b8fc86c015b94bc6f8712c3"
    },
    "pool_id": "ff8080825ab738f4015ab7513298010e",
    "project_id": "e55c6f3dc4e34c9f86353b664ae0e70c",
    "zone_type": "private",
    "created_at": "2017-04-22T08:17:08.997",
    "updated_at": "2017-04-22T08:17:09.997",
    "record_num": 2,
    "router": 
        {
            "status": "ACTIVE",
            "router_id": "19664294-0bf6-4271-ad3a-94b8c79c6558",
            "router_region": "xx"
        }
}
`

// CreatedZone is the expected created zone
var CreatedZone = zones.ZoneCreateResponse{
	ID:          "ff8080825b8fc86c015b94bc6f8712c3",
	Name:        "example.com.",
	Email:       "xx@example.com",
	Serial:      0,
	Status:      "ACTIVE",
	Description: "This is an example zone.",
	Masters:     []string{},
	CreatedAt:   "2017-04-22T08:17:08.997",
	UpdatedAt:   "2017-04-22T08:17:09.997",
	TTL:         300,
	Links: zones.Link{
		Self: "https://Endpoint/v2/zones/ff8080825b8fc86c015b94bc6f8712c3",
	},
	PoolId:    "ff8080825ab738f4015ab7513298010e",
	ProjectId: "e55c6f3dc4e34c9f86353b664ae0e70c",
	ZoneType:  "private",
	RecordNum: 2,
	Router: zones.AssociateRouterResponse{
		RouterId:     "19664294-0bf6-4271-ad3a-94b8c79c6558",
		Status:       "ACTIVE",
		RouterRegion: "xx",
	},
}

// HandleZoneCreationSuccessfully configures the test server to respond to a Create request.
func HandleCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateZoneRequest)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateZoneResponse)
	})
}

// UpdateZoneRequest is a sample request to update a zone.
const UpdateZoneRequest = `
{
    "ttl": 600,
    "description": "Updated Description"
}
`

// UpdateZoneResponse is a sample response to update a zone.
const UpdateZoneResponse = `
{
    "id": "a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
    "pool_id": "572ba08c-d929-4c70-8e42-03824bb24ca2",
    "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
    "name": "example.org.",
    "email": "joe@example.org",
    "ttl": 600,
    "serial": 1404757531,
    "status": "PENDING",
    "action": "UPDATE",
    "description": "Updated Description",
    "masters": [],
    "type": "PRIMARY",
    "transferred_at": null,
    "version": 1,
    "created_at": "2014-07-07T18:25:31.275934",
    "updated_at": null,
    "links": {
      "self": "https://127.0.0.1:9001/v2/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3"
    }
}
`

// HandleZoneUpdateSuccessfully configures the test server to respond to an Update request.
func HandleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			th.TestJSONRequest(t, r, UpdateZoneRequest)

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, UpdateZoneResponse)
		})
}

// DeleteZoneResponse is a sample response to update a zone.
const DeleteZoneResponse = `
{
    "id": "ff8080825b8fc86c015b94bc6f8712c3",
    "name": "example.com.",
    "description": "Updated Description",
    "email": "xx@example.com",
    "ttl": 300,
    "serial": 0,
    "masters": [],
    "status": "PENDING",
    "links": {
        "self": "https://Endpoint/v2/zones/ff8080825b8fc86c015b94bc6f8712c3"
    },
    "pool_id": "ff8080825ab738f4015ab7513298010e",
    "project_id": "e55c6f3dc4e34c9f86353b664ae0e70c",
    "zone_type": "private",
    "created_at": "2017-04-22T08:17:08.997",
    "updated_at": "2017-04-22T08:17:09.997",
    "record_num": 2,
    "routers": [
        {
            "status": "ACTIVE",
            "router_id": "19664294-0bf6-4271-ad3a-94b8c79c6558",
            "router_region": "xx"
        },
        {
            "status": "ACTIVE",
            "router_id": "f0791650-db8c-4a20-8a44-a06c6e24b15b",
            "router_region": "xx"
        }
    ]
}
`

// HandleZoneDeleteSuccessfully configures the test server to respond to an Delete request.
func HandleDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/zones/ff8080825b8fc86c015b94bc6f8712c3",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			w.WriteHeader(http.StatusAccepted)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, DeleteZoneResponse)
		})
}
