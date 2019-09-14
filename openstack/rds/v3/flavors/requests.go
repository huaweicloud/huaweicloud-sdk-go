package flavors

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type DbFlavorsOpts struct {
	Versionname string `q:"version_name"`
}

type DbFlavorsBuilder interface {
	ToDbFlavorsListQuery() (string, error)
}

func (opts DbFlavorsOpts) ToDbFlavorsListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts DbFlavorsBuilder, databasename string) pagination.Pager {
	url := listURL(client, databasename)
	if opts != nil {
		query, err := opts.ToDbFlavorsListQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return DbFlavorsPage{pagination.Offset{PageResult: r}}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}
