package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/bandwidths"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := bandwidths.Update(client.ServiceClient(), "3c43e46e-4af1-45b8-a84d-ee6d04488d2a", bandwidths.UpdateOpts{
		Name: "bandwidth-ABCD",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := bandwidths.Get(client.ServiceClient(), "3c43e46e-4af1-45b8-a84d-ee6d04488d2a").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := bandwidths.List(client.ServiceClient(), bandwidths.ListOpts{
		Limit: 2,
	}).AllPages()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}
