package predefinetags

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type CreateOrDeleteOpts struct {
	Action  string    `json:"action"  required:"true"`
	Tags   []Tag     `json:"tags,omitempty"`
}

type Tag struct {
	Key    string     `json:"key"  required:"true"`
	Value  string     `json:"value"  required:"true"`
}

type CreateOrDeleteBuilder interface {
	CreateOrDeleteMap() (map[string]interface{}, error)
}

func (opts CreateOrDeleteOpts) CreateOrDeleteMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateOrDelete(client *gophercloud.ServiceClient, opts CreateOrDeleteBuilder) (r ErrorResult) {
	b, err := opts.CreateOrDeleteMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createOrDeleteURL(client), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

type UpdateOpts struct {
    NewTag   Tag     `json:"new_tag,omitempty"`
    OldTag   Tag     `json:"old_tag,omitempty"`
}

type UpdateBuilder interface {
	UpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) UpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}
	
func Update(client *gophercloud.ServiceClient, opts UpdateBuilder) (r ErrorResult) {
	b, err := opts.UpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

type ListOpts struct {
	Key         string  `q:"key" `
	Value       string  `q:"value"`
	Limit       int     `q:"limit"`
	Marker      string  `q:"marker"`
	OrderField  int     `q:"order_field"`
	OrderMethod int     `q:"order_method"`
}

type ListBuilder interface {
	ToListQuery() (string, error)
}

type ListPage struct {
	pagination.SinglePageBase
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts ListBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ListPage{pagination.SinglePageBase(r)}
	})
}
