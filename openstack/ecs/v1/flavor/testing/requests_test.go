package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/gophercloud/gophercloud/testhelper/client"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/flavor"
)

func TestFlavorList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/cloudservers/flavors", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()
		zone := r.Form.Get("availability_zone")
		switch zone {
		case "":
			fmt.Fprintf(w, FlavorListResponse)
		case "az1.dc1(obt)":
			fmt.Fprintf(w, FlavorListWithFormResponse)
		default:
			t.Fatalf("Unexpected availability_zone: [%s]", zone)
		}
	})

	allPages, err := flavor.List(fake.ServiceClient(), flavor.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	flavors ,err := flavor.ExtractFlavors(allPages)
	th.AssertEquals(t, len(flavors), 2)
	th.AssertEquals(t, flavors[1].ID, "105")
	th.AssertEquals(t, flavors[1].Name, "m2.large")

	allPages, err = flavor.List(fake.ServiceClient(), flavor.ListOpts{AvailabilityZone:"az1.dc1(obt)"}).AllPages()
	th.AssertNoErr(t, err)
	flavors ,err = flavor.ExtractFlavors(allPages)
	th.AssertEquals(t, len(flavors), 1)
	th.AssertEquals(t, flavors[0].ID, "104")
	th.AssertEquals(t, flavors[0].Name, "m1.large")
}

func TestResize(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	serverId := "9a56640e-5503-4b8d-8231-963fc59ff91c"

	th.Mux.HandleFunc(fmt.Sprintf("/cloudservers/%s/resize",serverId), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, ResizeRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ResizeResponse)
	})

	options := flavor.ResizeOpts{
		FlavorRef:"c3.15xlarge.2",
		DedicatedHostId:"459a2b9d-804a-4745-ab19-a113bb1b4ddc",
	}

	jobId, err := flavor.Resize(fake.ServiceClient(), serverId, options)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, jobId, "70a599e0-31e7-49b7-b260-868f441e862b")
}