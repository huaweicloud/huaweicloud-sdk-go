package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/securitygrouprules"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := securitygrouprules.Create(client.ServiceClient(), securitygrouprules.CreateOpts{
		Description:     "Test SecurityGroup",
		SecurityGroupId: "7af80d49-0a43-462d-aed8-a1e12ac91af6",
		Direction:       "egress",
		Protocol:        "tcp",
		RemoteIpPrefix:  "10.10.0.0/24",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := securitygrouprules.Get(client.ServiceClient(), "26243298-ae79-46a3-bad9-34395762e033").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t, client.ServiceClient().Endpoint)

	allPage, err := securitygrouprules.List(client.ServiceClient(), securitygrouprules.ListOpts{
		Limit:    2,
	}).AllPages()

	actual, err := securitygrouprules.ExtractSecurityGroupRules(allPage.(securitygrouprules.ListPage))
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	securitygrouprules.Delete(client.ServiceClient(), "26243298-ae79-46a3-bad9-34395762e033")
}
