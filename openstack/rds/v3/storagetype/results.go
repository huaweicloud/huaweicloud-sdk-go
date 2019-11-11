package storagetype

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type StorageTypeResult struct {
	gophercloud.Result
}
type StorageTypeResp struct {
	StorageTypeList []StorageType `json:"storage_type" `
}
type StorageType struct {
	Name     string            `json:"name"`
	AzStatus map[string]string `json:"az_status"`
}

type StorageTypePage struct {
	pagination.Offset
}

func (r StorageTypePage) IsEmpty() (bool, error) {
	data, err := ExtractStorageType(r)
	if err != nil {
		return false, err
	}
	return len(data.StorageTypeList) == 0, err
}

func ExtractStorageType(r pagination.Page) (StorageTypeResp, error) {
	var s StorageTypeResp
	err := (r.(StorageTypePage)).ExtractInto(&s)
	return s, err
}
