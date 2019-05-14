package flavor

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToFlavorListMap() (string, error)
}

//ListOpts allows the filtering and sorting of paginated collections through the API.
type ListOpts struct {
	// Specifies the AZ name.
	AvailabilityZone string `q:"availability_zone"`
}

// ToFlavorListMap formats a ListOpts into a query string.
func (opts ListOpts) ToFlavorListMap() (string, error) {
	s, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return s.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// flavors.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := getListUrl(c)
	if opts != nil {
		queryString, err := opts.ToFlavorListMap()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += queryString
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return FlavorsPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// ResizeOpts represents options for modifying the specifications of an ecs.
type ResizeOpts struct {
	// Specifies the specifications ID of the ECS after the modification.
	FlavorRef     	string 	     `json:"flavorRef" required:"true"`

	// Specifies the new DeH ID, which is applicable only to the ECSs on DeHs.
	DedicatedHostId string        `json:"dedicated_host_id,omitempty"`
}

// ToResizeOptsMap builds a request body from ResizeOpts.
func (opts ResizeOpts) ToResizeOptsMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "resize")
}

// Resize is an operation which modifying the specifications of an ecs.
func Resize(client *gophercloud.ServiceClient,serverId string,opts ResizeOpts) (jobId string, err error) {
	var r ResizeResult
	reqBody, err := opts.ToResizeOptsMap()
	if err != nil {
		return
	}
	_, err = client.Post(resizeURL(client,serverId), reqBody, &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return
	}

	job, err := r.ExtractJob()
	if err != nil {
		return
	}
	jobId = job.Id
	return
}

