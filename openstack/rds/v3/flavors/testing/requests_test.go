package testing

import (
	"github.com/gophercloud/gophercloud/openstack/rds/v3/flavors"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"testing"
)

func TestListEachSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListMysqlSuccessfully(t)
	opts := flavors.DbFlavorsOpts{
		Versionname:"5.7",
	}
	count := 0
	err := flavors.List(client.ServiceClient(), opts, "MySQL").EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := flavors.ExtractDbFlavors(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, "rds.mysql.c2.medium.ha", actual.Flavorslist[0].Speccode)
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}
func TestListAllSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListMysqlSuccessfully(t)
	opts := flavors.DbFlavorsOpts{
		Versionname:"5.7",
	}
	allPages, err := flavors.List(client.ServiceClient(), opts, "MySQL").AllPages()
	th.AssertNoErr(t, err)
	allstores, err := flavors.ExtractDbFlavors(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allstores.Flavorslist))
}

func TestListPgEachSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListPgSuccessfully(t)
	opts := flavors.DbFlavorsOpts{
		Versionname:"9.6",
	}
	count := 0
	err := flavors.List(client.ServiceClient(), opts, "PostgreSQL").EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := flavors.ExtractDbFlavors(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, "rds.pg.i3.12xlarge.8", actual.Flavorslist[1].Speccode)
		th.CheckDeepEquals(t, "normal", actual.Flavorslist[0].Azstatus["az2xahz"])
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}
func TestListPgAllSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListPgSuccessfully(t)
	opts := flavors.DbFlavorsOpts{
		Versionname:"9.6",
	}
	allPages, err := flavors.List(client.ServiceClient(), opts, "PostgreSQL").AllPages()
	th.AssertNoErr(t, err)
	allstores, err := flavors.ExtractDbFlavors(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 4, len(allstores.Flavorslist))
}
