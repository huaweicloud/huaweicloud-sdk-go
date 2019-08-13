package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/ports"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := ports.Create(client.ServiceClient(), ports.CreateOpts{
		Name:      "EricTestPort",
		NetworkId: "5ae24488-454f-499c-86c4-c0355704005d",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := ports.Update(client.ServiceClient(), "5e56a480-f337-4985-8ca4-98546cb4fdae", ports.UpdateOpts{
		Name: "ModifiedPort",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := ports.Get(client.ServiceClient(), "5e56a480-f337-4985-8ca4-98546cb4fdae").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	pageData, err := ports.List(client.ServiceClient(), ports.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := ports.ExtractPorts(pageData)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	ports.Delete(client.ServiceClient(), "5e56a480-f337-4985-8ca4-98546cb4fdae")
}
