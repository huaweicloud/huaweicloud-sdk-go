package certificates

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type ListOptsBuilder interface {
	ToPolicyListMap() (string, error)
}

type ListOpts struct {
	Marker      string `q:"marker"`
	Limit       int    `q:"limit"`
	PageReverse *bool   `q:"page_reverse"`
}

func (opts ListOpts) ToPolicyListMap() (string, error) {
	s, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return s.String(), err
}

func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		queryString, err := opts.ToPolicyListMap()
		if err != nil {

			return pagination.Pager{Err: err}
		}
		url += queryString

	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return CertificatePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type CreateOptsBuilder interface {
	ToCertificateCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Domain      string `json:"domain,omitempty"`
	PrivateKey  string `json:"private_key"`
	Certificate string `json:"certificate" required:"true"`
}

func (opts CreateOpts) ToCertificateCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCertificateCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, nil)
	return
}

func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToCertificateUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Domain      string `json:"domain,omitempty"`
	PrivateKey  string `json:"private_key,omitempty"`
	Certificate string `json:"certificate,omitempty"`
}

func (opts UpdateOpts) ToCertificateUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

func Update(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToCertificateUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id), nil)
	return
}
