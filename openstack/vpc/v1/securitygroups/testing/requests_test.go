package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/securitygroups"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := securitygroups.Create(client.ServiceClient(), securitygroups.CreateOpts{
		Name: "EricSG",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := securitygroups.Get(client.ServiceClient(), "f7616338-fa30-42b8-bf6b-754c0701aab8").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := securitygroups.List(client.ServiceClient(), securitygroups.ListOpts{
		Limit: 2,
	}).AllPages()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	securitygroups.Delete(client.ServiceClient(), "2465d913-1084-4a6a-91e7-2fd6f490ecb3")
}
