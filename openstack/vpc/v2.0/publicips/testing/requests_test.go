package testing

import (
	"testing"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/publicips"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/" + "128a7bf965154373a7b73c89eb6b65aa/"
	return sc
}
func TestCreateOndemand(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOndemandSuccessfully(t)
	sc := ServiceClient()

	actual, err := publicips.Create(sc, OndmandOpts)
	th.AssertNoErr(t, err)

	if on, ok := actual.(publicips.PostPaid); ok {
		th.CheckDeepEquals(t, on.Type, OndmandOpts.PublicIP.Type)
		th.CheckDeepEquals(t, on.BandwidthSize, OndmandOpts.Bandwidth.Size)
		th.CheckDeepEquals(t, on.IPVersion, OndmandOpts.PublicIP.IPVersion)

	}
}

func TestCreateWithBSSInfo(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWithBSSInfoSuccessfully(t)
	sc := ServiceClient()

	actual, err := publicips.Create(sc, BSSOpts)
	th.AssertNoErr(t, err)
	if order, ok := actual.(publicips.PrePaid); ok {
		th.CheckDeepEquals(t, order, CreateResultForBSS)
		th.CheckDeepEquals(t, order, CreateResultForBSS)
	}
}
