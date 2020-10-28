package enterpriseprojects

import (
	"github.com/gophercloud/gophercloud"
)

type CreateOpts struct {
	Name        string `json:"name"  required:"true"`
	Description string `json:"description"`
}

type CreateBuilder interface {
	CreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) CreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateBuilder) (r CommonResult) {
	b, err := opts.CreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createEPURL(client), b, &r.Body, nil)
	return
}

type UpdateOpts struct {
	Name        string `json:"name"  required:"true"`
	Description string `json:"description"`
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

func Update(client *gophercloud.ServiceClient, opts UpdateBuilder, epID string) (r CommonResult) {
	b, err := opts.UpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, epID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Get(client *gophercloud.ServiceClient, epID string) (r CommonResult) {
	_, r.Err = client.Get(getURL(client, epID), &r.Body, nil)
	return
}

func GetQuotas(client *gophercloud.ServiceClient) (r CommonResult) {
	_, r.Err = client.Get(getQuotasURL(client), &r.Body, nil)
	return
}

type ListOpts struct {
	Id         string `q:"id" `
	Name       string `q:"name" `
	AuthAction string `q:"auth_action"`
	Offset     int    `q:"offset"`
	Limit      int    `q:"limit"`
	SortKey    string `q:"sort_key"`
	SortDiy    string `q:"sort_dir"`
}

type ListBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts ListBuilder) (r ListResult) {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	_, r.Err = client.Get(url, &r.Body, nil)
	return
}

type ActionOpts struct {
	Action string `json:"action"  required:"true"`
}

type ActionBuilder interface {
	ActionMap() (map[string]interface{}, error)
}

func (opts ActionOpts) ActionMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func EnableOrDisableEP(client *gophercloud.ServiceClient, opts ActionBuilder, epID string) (r ErrorResult) {
	b, err := opts.ActionMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(actionURL(client, epID), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

type Match struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FilterResourcesOpts struct {
	Projects      []string `json:"projects"`
	ResourceTypes []string `json:"resource_types"  required:"true"`
	Offset        int      `json:"offset"`
	Limit         int      `json:"limit"`
	Matches       []Match  `json:"matches"`
}

type FilterResourcesBuilder interface {
	FilterResourcesMap() (map[string]interface{}, error)
}

func (opts FilterResourcesOpts) FilterResourcesMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func FilterResources(client *gophercloud.ServiceClient, opts FilterResourcesBuilder, epID string) (r CommonResult) {
	b, err := opts.FilterResourcesMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(filterResourcesURL(client, epID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type MigrateResourceOpts struct {
	Action       string `json:"action" required:"true"`
	ProjectId    string `json:"project_id"`
	Associated   bool   `json:"associated"`
	ResourceType string `json:"resource_type" required:"true"`
	ResourceId   string `json:"resource_id" required:"true"`
	RegionId     string `json:"region_id"`
}

type MigrateResourcesBuilder interface {
	MigrateResourcesMap() (map[string]interface{}, error)
}

func (opts MigrateResourceOpts) MigrateResourcesMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func MigrateResources(client *gophercloud.ServiceClient, opts MigrateResourcesBuilder, epID string) (r ErrorResult) {
	b, err := opts.MigrateResourcesMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(migrateResourcesURL(client, epID), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
