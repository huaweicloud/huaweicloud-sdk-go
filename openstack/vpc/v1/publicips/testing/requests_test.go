package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/publicips"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	result, err := publicips.Create(client.ServiceClient(), publicips.CreateOpts{
		Publicip: publicips.PublicIPRequest{
			Type: "5_bgp",
		},
		Bandwidth: publicips.BandWidth{
			Name:       "bandwidth-d62f",
			Size:       1,
			ShareType:  "WHOLE",
			ChargeMode: "traffic"},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, result)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := publicips.Update(client.ServiceClient(), "84a71976-a8c2-42e0-8826-7fc27b876e42", publicips.UpdateOpts{
		IPVersion:4,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := publicips.Get(client.ServiceClient(), "84a71976-a8c2-42e0-8826-7fc27b876e42").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	allpages, err := publicips.List(client.ServiceClient(), publicips.ListOpts{
		Limit: 2,
	}).AllPages()
	th.AssertNoErr(t, err)
	publicipList,err := publicips.ExtractPublicIPs(allpages)
	th.CheckDeepEquals(t, ListResponse, publicipList)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	publicips.Delete(client.ServiceClient(), "7ffddb5f-6731-43d8-9476-1444aaa40bc0")
}
