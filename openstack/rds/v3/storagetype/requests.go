package storagetype

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)
type ListOpts struct {
	VersionName string `q:"version_name"`
}

type StorageTypeBuilder interface {
	ToStorageTypeListQuery() (string, error)
}

func (opts ListOpts) ToStorageTypeListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *gophercloud.ServiceClient,opts StorageTypeBuilder, databasesname string) pagination.Pager {
	url := listURL(client, databasesname)
	if opts != nil {
		query, err := opts.ToStorageTypeListQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return StorageTypePage{pagination.Offset{PageResult: r}}
	})

	pageRdsList.Headers = map[string]string{"Content-Type": "application/json"}
	return pageRdsList
}
