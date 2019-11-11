package testing

import (
	"github.com/gophercloud/gophercloud/openstack/rds/v3/storagetype"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"testing"
)

func TestListEachSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)
	opts := storagetype.ListOpts{
		VersionName: "5.7",
	}
	count := 0
	err := storagetype.List(client.ServiceClient(), opts, "mysql").EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := storagetype.ExtractStorageType(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, "COMMON", actual.StorageTypeList[0].Name)
		th.CheckDeepEquals(t, "normal", actual.StorageTypeList[0].AzStatus["az1"])
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}
func TestListAllSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)
	opts := storagetype.ListOpts{
		VersionName: "5.7",
	}
	allPages, err := storagetype.List(client.ServiceClient(), opts, "mysql").AllPages()
	th.AssertNoErr(t, err)
	allstores, err := storagetype.ExtractStorageType(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allstores.StorageTypeList))
}
