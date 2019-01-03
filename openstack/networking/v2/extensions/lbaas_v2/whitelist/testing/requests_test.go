package testing

import (
	"testing"

	fake "github.com/gophercloud/gophercloud/openstack/networking/v2/common"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/whitelist"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestListWhitelists(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWhitelistSuccessfully(t)

	pages := 0
	err := whitelist.List(fake.ServiceClient(), whitelist.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := whitelist.ExtractWhiteLists(page)
		if err != nil {
			return false, err
		}

		if len(actual.Whitelists) != 2 {
			t.Fatalf("Expected 2 Whitelists, got %d", len(actual.Whitelists))
		}
		th.CheckDeepEquals(t, WhitelistOne, actual.Whitelists[0])
		th.CheckDeepEquals(t, WhitelistTwo, actual.Whitelists[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllWhitelists(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWhitelistALLSuccessfully(t)

	allPages, err := whitelist.List(fake.ServiceClient(), whitelist.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := whitelist.ExtractWhiteLists(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, WhitelistOne, actual.Whitelists[0])
	th.CheckDeepEquals(t, WhitelistTwo, actual.Whitelists[1])
}

//
func TestCreateWhitelist(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWhitelistCreationSuccessfully(t)

	enable := true
	actual, err := whitelist.Create(fake.ServiceClient(), whitelist.CreateOpts{
		ListenerId:      "eabfefa3fd1740a88a47ad98e132d238",
		EnableWhitelist: &enable,
		Whitelist:       "192.168.11.1,192.168.0.1/24,192.168.201.18/8,100.164.0.1/24",
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, WhitelistOne, actual.WhiteList)
}

//
func TestGetWhitelist(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWhitelistGetSuccessfully(t)

	client := fake.ServiceClient()
	actual, err := whitelist.Get(client, "09e64049-2ab0-4763-a8c5-f4207875dc3e").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, WhitelistOne, actual.WhiteList)
}

func TestDeleteWhitelist(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWhitelistDeletionSuccessfully(t)

	res := whitelist.Delete(fake.ServiceClient(), "09e64049-2ab0-4763-a8c5-f4207875dc3e")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateWhitelist(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWhitelistUpdateSuccessfully(t)
	enable := true
	client := fake.ServiceClient()
	actual, err := whitelist.Update(client, "09e64049-2ab0-4763-a8c5-f4207875dc3e", whitelist.UpdateOpts{
		EnableWhitelist: &enable,
	}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}

	th.CheckDeepEquals(t, WhitelistOne, actual.WhiteList)
}

