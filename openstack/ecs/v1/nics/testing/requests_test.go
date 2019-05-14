package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/gophercloud/gophercloud/testhelper/client"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/nics"
)

func TestAddNics(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	serverId := "9a56640e-5503-4b8d-8231-963fc59ff91c"

	th.Mux.HandleFunc(fmt.Sprintf("/cloudservers/%s/nics",serverId), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, AddNicsRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, AddNicsResponse)
	})

	options := nics.AddOpts{
		Nics: []nics.Nic{
			{
				SubnetId:  "d32019d3-bc6e-4319-9c1d-6722fc136a23",
				IpAddress: "",
				SecurityGroups: []nics.SecurityGroup{
					{
						ID: "f0ac4394-7e4a-4409-9701-ba8be283dbc3",
					},
				},
			},
		},
	}
	jobId, err := nics.AddNics(fake.ServiceClient(), serverId, options)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, jobId, "70a599e0-31e7-49b7-b260-868f441e862b")
}

func TestDelNics(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	serverId := "9a56640e-5503-4b8d-8231-963fc59ff91c"

	th.Mux.HandleFunc(fmt.Sprintf("/cloudservers/%s/nics/delete",serverId), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, DelNicsRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, DelNicsResponse)
	})

	options := nics.DelOpts{
		Nics: []nics.Nics{
			{
				ID: "d32019d3-bc6e-4319-9c1d-6722fc136a23",
			},
		},
	}

	jobId, err := nics.DeleteNics(fake.ServiceClient(), serverId, options)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, jobId, "70a599e0-31e7-49b7-b260-868f441e862b")
}

func TestBindNic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	nicId := "9a56640e-5503-4b8d-8231-963fc59ff91c"

	th.Mux.HandleFunc(fmt.Sprintf("/cloudservers/nics/%s",nicId), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, BindNicRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, BindNicResponse)
	})

	reversebinding := true
	options := nics.BindOpts{
		SubnetId:       "d32019d3-bc6e-4319-9c1d-6722fc136a23",
		IpAddress:      "192.168.0.7",
		ReverseBinding: &reversebinding,
	}

	portId, err := nics.BindNic(fake.ServiceClient(), nicId, options).ExtractPortId()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, portId, "d32019d3-bc6e-4319-9c1d-6722fc136a23")
}

func TestUnBindNic(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	nicId := "9a56640e-5503-4b8d-8231-963fc59ff91c"

	th.Mux.HandleFunc(fmt.Sprintf("/cloudservers/nics/%s",nicId), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UnBindNicRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UnBindNicResponse)
	})

	reversebinding := false
	options := nics.UnBindOpts{
		SubnetId:       "",
		IpAddress:      "",
		ReverseBinding: &reversebinding,
	}

	portId, err := nics.UnBindNic(fake.ServiceClient(), nicId, options).ExtractPortId()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, portId, "d32019d3-bc6e-4319-9c1d-6722fc136a23")
}