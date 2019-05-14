package logs

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type ListOpts struct {
	// Specifies the start time for querying scaling action logs. The
	// format of the start time is yyyy-MM-ddThh:mm:ssZ.
	StartTime string `q:"start_time"`

	// Specifies the end time for querying scaling action logs. The
	// format of the end time is yyyy-MM-ddThh:mm:ssZ.
	EndTime string `q:"end_time"`

	// Specifies the start line number. The default value is 0.
	StartNumber int `q:"start_number"`

	// Specifies the total number of query records. The default value
	// is 20 and the maximum value is 100.
	Limit int `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, scalingGroupId string, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client, scalingGroupId)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		p := LogPage{pagination.NumberPageBase{PageResult: r}}
		p.NumberPageBase.Owner = p
		return p
	})
}
