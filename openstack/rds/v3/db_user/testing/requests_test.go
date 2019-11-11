package testing

import (
	"github.com/gophercloud/gophercloud/openstack/rds/v3/db_user"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"testing"
)

func TestCreateDbUserSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateDbUserSuccessfully(t)
	opts := db_user.CreateDbUserOpts{
		Username:"rds",
		Password:"Huawei_test",
	}
	acrtual,err := db_user.Create(client.ServiceClient(), opts, "dsfae23fsfdsae3435in01").Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "successful", acrtual.Resp)
}
func TestListAllSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListDbUserSuccessfully(t)
	opts := db_user.ListDbUsersOpts{
		Page:1,
		Limit:6,
	}
	allPages, err := db_user.List(client.ServiceClient(), opts, "dsfae23fsfdsae3435in01").AllPages()
	th.AssertNoErr(t, err)
	alldbusers, err := db_user.ExtractDbUsers(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 6, len(alldbusers.UsersList))
}

func TestListPgEachSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListDbUserSuccessfully(t)
	opts := db_user.ListDbUsersOpts{
		Page:1,
		Limit:6,
	}
	count := 0
	err := db_user.List(client.ServiceClient(), opts, "dsfae23fsfdsae3435in01").EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := db_user.ExtractDbUsers(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, "mysql.session", actual.UsersList[1].Name)
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestDeleteDbUserSuccess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteDbUserSuccess(t)

	acrtual,err := db_user.Delete(client.ServiceClient(), "dsfae23fsfdsae3435in01", "rds_009").Extract()
	th.AssertNoErr(t,err)
	th.CheckEquals(t, "successful", acrtual.Resp)
}