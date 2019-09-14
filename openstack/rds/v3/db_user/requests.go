package db_user

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type CreateDbUserOpts struct {
	Username string `json:"name"  required:"true"`
	Password string `json:"password"  required:"true"`
}

type CreateDbUserBuilder interface {
	ToDbUserCreateMap() (map[string]interface{}, error)
}

func (opts CreateDbUserOpts) ToDbUserCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateDbUserBuilder, instanceId string) (r DbUserResult) {
	b, err := opts.ToDbUserCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client, instanceId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return
}

type ListDbUsersOpts struct {
	Page  int `q:"page" `
	Limit int `q:"limit" `
}

type DbUsersBuilder interface {
	ToDbUsersListQuery() (string, error)
}

func (opts ListDbUsersOpts) ToDbUsersListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts DbUsersBuilder, instance string) pagination.Pager {
	url := listURL(client, instance)
	if opts != nil {
		query, err := opts.ToDbUsersListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return DbUsersPage{pagination.Offset{PageResult: r}}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}

func Delete(client *gophercloud.ServiceClient, instanceid string, dbUser string) (r commonResult) {

	_, r.Err = client.Delete(deleteURL(client, instanceid, dbUser), &gophercloud.RequestOpts{JSONResponse:&r.Body,OkCodes: []int{202},MoreHeaders: map[string]string{"Content-Type": "application/json"}})

	return
}
