package testing

import (
	"testing"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/bandwidths"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"fmt"
)

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/" + "128a7bf965154373a7b73c89eb6b65aa/"
	return sc
}

func TestUpdateBandwidthSize(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWithBSSInfoSuccessfully(t)
	sc := ServiceClient()

	actual, err := bandwidths.Update(sc, "3fa5b383-5a73-4dcb-a314-c6128546d855", requestOpts2)
	th.AssertNoErr(t, err)
	if data, ok := actual.(bandwidths.PrePaid); ok {
		th.CheckDeepEquals(t, data.OrderID, "dd0bdea6efe0")
	}
}


func TestUpdateBandwidthName(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWithNameSuccessfully(t)
	sc := ServiceClient()

	actual, err := bandwidths.Update(sc, "3fa5b383-5a73-4dcb-a314-c6128546d855", requestOpts1)
	th.AssertNoErr(t, err)
	fmt.Println(actual)
	if data, ok := actual.(bandwidths.PostPaid); ok {
		th.CheckDeepEquals(t, data.Name, "test")
	}
}