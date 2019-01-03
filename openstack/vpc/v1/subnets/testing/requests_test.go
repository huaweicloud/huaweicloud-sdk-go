package testing

import (
	"testing"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/subnets"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/gophercloud/gophercloud/pagination"
	"fmt"
)

func TestListAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := subnets.List(client.ServiceClient(), subnets.ListOpts{}).AllPages()
	resp, err := subnets.ExtractSubnets(actual)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, &resp)
}

func TestListEachPage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)
	count := 0
	err := subnets.List(client.ServiceClient(), subnets.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		data, paErr := subnets.ExtractSubnets(page)

		if paErr != nil {
			t.Errorf("Failed to extract subnets: %v", paErr)
			return false, paErr
		}
		fmt.Println(len(data))
		th.CheckDeepEquals(t, &ListResponse, &data)

		return true, nil

	})
	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
	th.AssertNoErr(t, err)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := subnets.Create(client.ServiceClient(), subnets.CreateOpts{
		Name:         "subnet",
		Cidr:         "192.168.20.0/24",
		GatewayIP:    "192.168.20.1",
		DhcpEnable:   &enable,
		PrimaryDNS:   "114.114.114.114",
		SecondaryDNS: "114.114.115.115",
		DNSList: []string{
			"114.114.114.114",
			"114.114.115.115",
		},
		AvailabilityZone: "cn-north-1a",
		VpcID:            "ea3b0efe-0d6a-4288-8b16-753504a994b9",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreateResponse, actual)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateSuccessfully(t)

	actual, err := subnets.Update(client.ServiceClient(), "ea3b0efe-0d6a-4288-8b16-753504a994b9", "c9aba52d-ec14-40cb-930f-c8153e93c2db", subnets.UpdateOpts{
		Name: "ABC-back",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdateResponse, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := subnets.Get(client.ServiceClient(), "c9aba52d-ec14-40cb-930f-c8153e93c2db").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponse, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	subnets.Delete(client.ServiceClient(), "ea3b0efe-0d6a-4288-8b16-753504a994b9", "c9aba52d-ec14-40cb-930f-c8153e93c2db")
}
