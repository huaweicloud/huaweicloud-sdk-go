package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/extensions/apiversions"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

const ManilaAPIVersionResponse = `
{ 
    "versions": [ 
        { 
            "min_version": "",  
            "media-types": [ 
                { 
                    "type": "application/vnd.openstack.volume+json;version=1",  
                    "base": "application/json" 
                },  
                { 
                    "type": "application/vnd.openstack.volume+xml;version=1",  
                    "base": "application/xml" 
                } 
            ],  
            "links": [ 
                { 
                    "rel": "describedby",  
                    "href": "http://docs.openstack.org/",  
                    "type": "text/html" 
                },  
                { 
                    "rel": "self",  
                    "href": "https://evs.localdomain.com/v2" 
                } 
            ],  
            "id": "v2.0",  
            "updated": "2014-06-28T12:20:21Z",  
            "version": "",  
            "status": "SUPPORTED" 
        } 
    ] 
}
`

const ManilaAllAPIVersionsResponse = `
{ 
    "versions": [ 
        { 
            "min_version": "",  
            "media-types": [ 
                { 
                    "type": "application/vnd.openstack.volume+json;version=1",  
                    "base": "application/json" 
                },  
                { 
                    "type": "application/vnd.openstack.volume+xml;version=1",  
                    "base": "application/xml" 
                } 
            ],  
            "links": [ 
                { 
                    "rel": "describedby",  
                    "href": "http://docs.openstack.org/"
                },  
                { 
                    "rel": "self",  
                    "href": "https://evs.localdomain.com/v1" 
                } 
            ],  
            "id": "v1.0",  
            "updated": "2014-06-28T12:20:21Z",  
            "version": "",  
            "status": "SUPPORTED" 
        },  
        { 
            "min_version": "",  
            "media-types": [ 
                { 
                    "type": "application/vnd.openstack.volume+json;version=1",  
                    "base": "application/json" 
                },  
                { 
                    "type": "application/vnd.openstack.volume+xml;version=1",  
                    "base": "application/xml" 
                } 
            ],  
            "links": [ 
                { 
                    "rel": "describedby",  
                    "href": "http://docs.openstack.org/"
                },  
                { 
                    "rel": "self",  
                    "href": "https://evs.localdomain.com/v2" 
                } 
            ],  
            "id": "v2.0",  
            "updated": "2014-06-28T12:20:21Z",  
            "version": "",  
            "status": "SUPPORTED" 
        }
    ] 
}
`

var ManilaAPIVersion1Result = apiversions.APIVersion{
	ID:         "v1.0",
	Status:     "SUPPORTED",
	Updated:    time.Date(2014, 6, 28, 12, 20, 21, 0, time.UTC),
	MinVersion: "",
	MediaTypes: []map[string]string{
		{
			"type": "application/vnd.openstack.volume+json;version=1",
			"base": "application/json",
		},
		{
			"type": "application/vnd.openstack.volume+xml;version=1",
			"base": "application/xml",
		},
	},
}

var ManilaAPIVersion2Result = apiversions.APIVersion{
	ID:         "v2.0",
	Status:     "SUPPORTED",
	Updated:    time.Date(2014, 6, 28, 12, 20, 21, 0, time.UTC),
	MinVersion: "",
	MediaTypes: []map[string]string{
		{
			"type": "application/vnd.openstack.volume+json;version=1",
			"base": "application/json",
		},
		{
			"type": "application/vnd.openstack.volume+xml;version=1",
			"base": "application/xml",
		},
	},
}

var ManilaAllAPIVersionResults = []apiversions.APIVersion{
	ManilaAPIVersion1Result,
	ManilaAPIVersion2Result,
}

func MockListResponse(t *testing.T) {
	th.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ManilaAllAPIVersionsResponse)
	})
}

func MockGetResponse(t *testing.T) {
	th.Mux.HandleFunc("/v2/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ManilaAPIVersionResponse)
	})
}
