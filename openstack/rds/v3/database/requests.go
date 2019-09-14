package database

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)


type CreateOpts struct {
	Dbname       string `json:"name"  required:"true"`
	Characterset string `json:"character_set"  required:"true"`
}

type CreateOptsBuilder interface {
	ToDatabaseCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToDatabaseCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder, instanceId string) (r CreateDatabaseResult) {
	b, err := opts.ToDatabaseCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client, instanceId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return
}

type ListOpts struct {
	Page  int `q:"page" `
	Limit int `q:"limit" `
}

type ListOptsBuilder interface {
	ToDataBaseListQuery() (string, error)
}

func (opts ListOpts) ToDataBaseListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}


func List(client *gophercloud.ServiceClient, opts ListOptsBuilder, instance string) pagination.Pager {
	url := listURL(client, instance)
	if opts != nil {
		query, err := opts.ToDataBaseListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return DataBasePage{pagination.Offset{PageResult: r}}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}

func Delete(client *gophercloud.ServiceClient, instanceid string, dbName string) (r commonResult) {


	_, r.Err = client.Delete(deleteURL(client, instanceid,dbName), &gophercloud.RequestOpts{JSONResponse:&r.Body,OkCodes: []int{202},MoreHeaders: map[string]string{"Content-Type": "application/json"}})
	return
}
