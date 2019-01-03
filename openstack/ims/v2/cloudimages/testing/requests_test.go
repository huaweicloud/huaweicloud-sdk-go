package testing

import (
	"testing"
	"github.com/gophercloud/gophercloud/openstack/ims/v2/cloudimages"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/pagination"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListImage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageListSuccessfully(t)

	t.Logf("Test setup %+v\n", th.Server)

	t.Logf("Id\tName\tOwner\tChecksum\tSizeBytes")

	pager := cloudimages.List(fakeclient.ServiceClient(), cloudimages.ListOpts{})
	t.Logf("Pager state %v", pager)
	count, pages := 0, 0
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		pages++
		t.Logf("Page %v", page)
		images, err := cloudimages.ExtractImages(page)
		if err != nil {
			return false, err
		}

		for _, i := range images {
			t.Logf("%s\t%s\t%s\t%s\t%v\t\n", i.ID, i.Name, i.Owner, i.Checksum, i.Size)
			count++
		}

		return true, nil
	})
	th.AssertNoErr(t, err)

	t.Logf("--------\n%d images listed on %d pages.\n", count, pages)
	th.AssertEquals(t, 1, pages)
	th.AssertEquals(t, 2, count)
}

func TestAllPagesImage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageListSuccessfully(t)

	pages, err := cloudimages.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	images, err := cloudimages.ExtractImages(pages)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 2, len(images))
}

func TestCreateImageByFile(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleImageCreateByFileSuccessfully(t)

	opts := cloudimages.CreateByFileOpts{
		Name: "test",
	}

	resp, err := cloudimages.CreateImageByFile(fakeclient.ServiceClient(), opts).ExtractJob()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, createResult, *resp)

}
func TestCreateImageByServer(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageCreateByServerSuccessfully(t)
	opts := cloudimages.CreateByServerOpts{
		Name:        "tttt",
		InstanceId:  "877a2cda-ba63-4e1e-b95f-e67e48b6129a",
		Description: "testtestest",
		Tags:        []string{"a:b", "c:d"},
	}

	resp, err := cloudimages.CreateImageByServer(fakeclient.ServiceClient(), opts).ExtractJob()

	th.AssertNoErr(t, err)
	th.AssertEquals(t, createResult, *resp)

}
func TestGetJobResult(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetJobSuccessfully(t)
	id := "ff8080814dbd65d7014dbe0d84db0013"
	actual, err := cloudimages.GetJobResult(fakeclient.ServiceClient(), id).ExtractJobResult()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, jobResult, *actual)

}
