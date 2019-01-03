package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/vpcs"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := vpcs.Create(client.ServiceClient(), vpcs.CreateOpts{
		Name: "ABC",
		Cidr: "192.168.0.0/16",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := vpcs.Update(client.ServiceClient(), "7ffddb5f-6731-43d8-9476-1444aaa40bc0", vpcs.UpdateOpts{
		Name: "ABC-back",
		Cidr: "192.168.0.0/24",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := vpcs.Get(client.ServiceClient(), "7ffddb5f-6731-43d8-9476-1444aaa40bc0").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	allpages, err := vpcs.List(client.ServiceClient(), vpcs.ListOpts{
		Limit: 2,
	}).AllPages()
	th.AssertNoErr(t, err)
	vpcs,err :=vpcs.ExtractVpcs(allpages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ListResponse, vpcs)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	vpcs.Delete(client.ServiceClient(), "7ffddb5f-6731-43d8-9476-1444aaa40bc0")
}
