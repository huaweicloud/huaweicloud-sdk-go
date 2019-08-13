package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/privateips"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	result, err := privateips.Create(client.ServiceClient(), privateips.CreateOpts{
		Privateips: []privateips.PrivateIpCreate{
			{
				SubnetId:  "e103753f-44b1-4741-984f-c99f03bc86c8",
				IpAddress: "192.168.0.231",
			},
		},
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, result)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := privateips.Get(client.ServiceClient(), "ea274524-f1cc-4078-8e67-c002be25c413").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	subnetID := "5ae24488-454f-499c-86c4-c0355704005d"
	pageData, err := privateips.List(client.ServiceClient(), subnetID, privateips.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := privateips.ExtractPrivateIps(pageData)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	privateips.Delete(client.ServiceClient(), "ea274524-f1cc-4078-8e67-c002be25c413")
}
