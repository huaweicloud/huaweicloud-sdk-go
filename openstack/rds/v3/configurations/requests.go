package configurations

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

func List(client *gophercloud.ServiceClient) pagination.Pager {
	url := listURL(client)

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ConfigurationsPage{pagination.Offset{PageResult: r}}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}

type CreateConfigurationsOpts struct {
	Name        string     `json:"name" required:"true"`
	Description string     `json:"description,omitempty"`
	Values      map[string]string    `json:"values,omitempty"`
	Datastore   *Datastore `json:"datastore" required:"true"`
}

type Datastore struct {
	Type    string `json:"type" required:"true"`
	Version string `json:"version" required:"true"`
}

type CreateConfigurationsBuilder interface {
	ToCreateConfigurationsMap() (map[string]interface{}, error)
}

func (opts CreateConfigurationsOpts) ToCreateConfigurationsMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateConfigurationsBuilder) (r CreateConfigurationsResult) {
	b, err := opts.ToCreateConfigurationsMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
