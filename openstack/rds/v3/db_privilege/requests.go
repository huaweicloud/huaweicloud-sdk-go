package db_privilege

import (
	"github.com/gophercloud/gophercloud"
)


type DbprivilegeOpts struct {
	Dbname string `json:"db_name"  required:"true"`
	Users  []User `json:"users"  required:"true"`
}

type User struct {
	Name     string `json:"name"  required:"true"`
	Readonly bool   `json:"readonly"  required:"true"`
}

type DbprivilegeBuilder interface {
	ToDbprivilegeHaMap() (map[string]interface{}, error)
}

func (opts DbprivilegeOpts) ToDbprivilegeHaMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts DbprivilegeBuilder, instanceId string) (r DbprivilegeResult) {
	b, err := opts.ToDbprivilegeHaMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client, instanceId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}


type DeleteDbprivilegeOpts struct {
	Dbname string        `json:"db_name"  required:"true"`
	Users  []DeleteUsers `json:"users"  required:"true"`
}

type DeleteUsers struct {
	Name string `json:"name"  required:"true"`
}

type DeleteDbprivilegeBuilder interface {
	DeleteDbprivilegeMap() (map[string]interface{}, error)
}

func (opts DeleteDbprivilegeOpts) DeleteDbprivilegeMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Delete(client *gophercloud.ServiceClient, opts DeleteDbprivilegeBuilder, instanceId string) (r DbprivilegeResult) {

	b, err := opts.DeleteDbprivilegeMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Delete(deleteURL(client, instanceId),&gophercloud.RequestOpts{JSONResponse:&r.Body,JSONBody:b,
		OkCodes: []int{200},MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})

	return
}
