package volumetransfer

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVolumeCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	// The volume name
	Name     string `json:"name,omitempty"`
	VolumeID string `json:"volume_id" required:"true"`
}

// ToVolumeCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToVolumeCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "transfer")
}

// Create will create a new Volume based on the values in CreateOpts. To extract
// the Volume object from the response, call the Extract method on the
// CreateResult.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVolumeCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

type AcceptOptsBuilder interface {
	ToVolumeAcceptMap() (map[string]interface{}, error)
}

type AcceptOpts struct {
	AuthKey string `json:"auth_key" required:"true"`
}

func (opts AcceptOpts) ToVolumeAcceptMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "accept")
}

func Accept(client *gophercloud.ServiceClient, transferID string, opts AcceptOptsBuilder) (r AcceptResult) {
	b, err := opts.ToVolumeAcceptMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(acceptURL(client, transferID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, transferID string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, transferID), nil)
	return
}

func Get(client *gophercloud.ServiceClient, transferID string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, transferID), &r.Body, nil)
	return
}

type ListDetailOptsBuilder interface {
	ToVolumeListDetailQuery() (string, error)
}

type ListDetailOpts struct {
	Limit  int `q:"limit"`
	Offset int `q:"offset"`
}

// ToVolumeListQuery formats a ListOpts into a query string.
func (opts ListDetailOpts) ToVolumeListDetailQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func ListDetail(client *gophercloud.ServiceClient, opts ListDetailOptsBuilder) pagination.Pager {
	url := detailURL(client)
	if opts != nil {
		query, err := opts.ToVolumeListDetailQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return TransferDetailsPage{pagination.SinglePageBase(r)}
	})
}

func List(client *gophercloud.ServiceClient, opts ListDetailOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToVolumeListDetailQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return TransferListPage{pagination.SinglePageBase(r)}
	})
}
