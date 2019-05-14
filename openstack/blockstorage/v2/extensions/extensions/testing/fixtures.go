package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/extensions/extensions"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/gophercloud/gophercloud"
)

// ListOutput provides a single page of Extension results.
const ListOutput = `
{ 
    "extensions": [ 
        { 
            "updated": "2013-04-18T00:00:00+00:00",  
            "name": "SchedulerHints",  
            "links": [ ],  
            "namespace": "http://docs.openstack.org/block-service/ext/scheduler-hints/api/v2",  
            "alias": "OS-SCH-HNT",  
            "description": "Pass arbitrary key/value pairs to the scheduler." 
        }
    ] 
}

`

// ListedExtension is the Extension that should be parsed from ListOutput.
var ListedExtension = extensions.Extension{
	Updated:     "2013-04-18T00:00:00+00:00",
	Name:        "SchedulerHints",
	Links:       []gophercloud.Link{},
	Namespace:   "http://docs.openstack.org/block-service/ext/scheduler-hints/api/v2",
	Alias:       "OS-SCH-HNT",
	Description: "Pass arbitrary key/value pairs to the scheduler.",
}

// ExpectedExtensions is a slice containing the Extension that should be parsed from ListOutput.
var ExpectedExtensions = []extensions.Extension{ListedExtension}

// HandleListExtensionsSuccessfully creates an HTTP handler at `/extensions` on the test handler
// mux that response with a list containing a single tenant.
func HandleListExtensionsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/extensions", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(w, ListOutput)
	})
}
