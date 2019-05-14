package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/extensions/volumetypes"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)
	pages := 0
	err := volumetypes.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := volumetypes.ExtractVolumeTypes(page)
		if err != nil {
			return false, err
		}
		expected := []volumetypes.VolumeType{
			{
				ID:          "6c81c680-df58-4512-81e7-ecf66d160638",
				Name:        "SAS",
				ExtraSpecs:  map[string]string{"volume_backend_name": "SAS", "availability-zone": "az1.dc1"},
				Description: "",
				QosSpecID:   "",
				IsPublic:    true,
			}, {
				ID:          "ea6e3c13-aac5-46e0-b280-745ed272e662",
				Name:        "SATA",
				ExtraSpecs:  map[string]string{"volume_backend_name": "SATA", "availability-zone": "az1.dc1"},
				IsPublic:    true,
				Description: "",
				QosSpecID:   "585f29d6-7147-42e7-bfb8-ca214f640f6f",
			},
		}
		th.CheckDeepEquals(t, expected, actual)
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, pages, 1)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	v, err := volumetypes.Get(client.ServiceClient(), "ea6e3c13-aac5-46e0-b280-745ed272e662").Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v.Name, "SATA")
	th.AssertEquals(t, v.ID, "ea6e3c13-aac5-46e0-b280-745ed272e662")
	th.AssertEquals(t, v.ExtraSpecs["volume_backend_name"], "SATA")
	th.AssertEquals(t, v.IsPublic, true)
}
