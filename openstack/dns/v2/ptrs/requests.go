package ptrs

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

func Get(client *gophercloud.ServiceClient, region string, floatingIPId string) (r GetResult) {
	url := GetURL(client, region, floatingIPId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{})
	return
}

type ListOpts struct {
	// Start resource ID of pagination query.If the parameter is left
	// blank, only resources on the first page are queried.
	Marker string `q:"marker"`

	// Number of resources returned on each page.Value range:
	// 0–500.Commonly used values are 10, 20, and 50.
	Limit int `q:"limit"`

	Offset int `q:"offset"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

//Query PTR records of EIPs.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return PtrPage{pagination.LinkedPageBase{PageResult: r}}
	})

}

type RestoreOpts struct {
}

type RestoreOptsBuilder interface {
	ToPtrsRestoreMap() (map[string]interface{}, error)
}

func (opts RestoreOpts) ToPtrsRestoreMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Restore(client *gophercloud.ServiceClient, region string, floatingIPId string) (r RestoreResult) {
	opts := RestoreOpts{}
	b, err := opts.ToPtrsRestoreMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Patch(RestoreURL(client, region, floatingIPId), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

type UpdateOpts struct {
	// Domain name of the PTR record
	Ptrdname string `json:"ptrdname" required:"true"`

	// PTR record description
	Description string `json:"description,omitempty"`

	// Caching period of a PTR record (in seconds).The default value
	// is 300s.The value range is 300–2147483647.
	TTL int `json:"ttl,omitempty"`
}

func (opts UpdateOpts) ToPtrsSetupMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, region string, floatingipId string, opts SetupOptsBuilder) (r SetupResult) {
	b, err := opts.ToPtrsSetupMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Patch(SetupURL(client, region, floatingipId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

type SetupOpts struct {
	// Domain name of the PTR record
	Ptrdname string `json:"ptrdname" required:"true"`

	// PTR record description
	Description string `json:"description,omitempty"`

	// Caching period of a PTR record (in seconds).The default value
	// is 300s.The value range is 300–2147483647.
	TTL int `json:"ttl,omitempty"`
}

type SetupOptsBuilder interface {
	ToPtrsSetupMap() (map[string]interface{}, error)
}

func (opts SetupOpts) ToPtrsSetupMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Setup(client *gophercloud.ServiceClient, region string, floatingipId string, opts SetupOptsBuilder) (r SetupResult) {
	b, err := opts.ToPtrsSetupMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Patch(SetupURL(client, region, floatingipId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
