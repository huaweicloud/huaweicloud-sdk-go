package quotas

import (
	"github.com/gophercloud/gophercloud"
)

type ListOpts struct {

	// Specifies the resource type.
	Type string `q:"type"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) (r ListResult) {
	url := ListURL(client)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
