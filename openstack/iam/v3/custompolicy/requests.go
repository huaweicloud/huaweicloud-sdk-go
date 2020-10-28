package custompolicy

import (
	"github.com/gophercloud/gophercloud"
)

func ListCustomPolicies(client *gophercloud.ServiceClient) (r ListResult) {
	url := listCustomPoliciesUrl(client)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		OkCodes:     []int{200},
	})
	return
}

func QueryCustomPolicyDetails(client *gophercloud.ServiceClient, roleId string) (r QueryResult) {
	url := queryAgencyDetailsUrl(client, roleId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		OkCodes:     []int{200},
	})
	return
}

type CloudServiceCustomPolicyStatement struct {
	Effect    string      `json:"Effect" required:"true"`
	Action    []string    `json:"Action" required:"true"`
	Condition interface{} `json:"Condition,omitempty"`
	Resource  []string    `json:"Resource,omitempty"`
}

type CloudServiceCustomPolicy struct {
	Version   string                              `json:"Version" required:"true"`
	Statement []CloudServiceCustomPolicyStatement `json:"Statement" required:"true"`
}

type CreateCloudServiceCustomPolicyOpts struct {
	DisplayName   string                   `json:"display_name" required:"true"`
	Type          string                   `json:"type" required:"true"`
	Description   string                   `json:"description" required:"true"`
	DescriptionCn string                   `json:"description_cn,omitempty"`
	Policy        CloudServiceCustomPolicy `json:"policy" required:"true"`
}

type UpdateCloudServiceCustomPolicyOpts struct {
	DisplayName   string                   `json:"display_name" required:"true"`
	Type          string                   `json:"type" required:"true"`
	Description   string                   `json:"description" required:"true"`
	DescriptionCn string                   `json:"description_cn,omitempty"`
	Policy        CloudServiceCustomPolicy `json:"policy" required:"true"`
}

type CreateCloudServiceCustomPolicyOptsBuilder interface {
	ToCloudServiceCustomPolicyCreateMap() (map[string]interface{}, error)
}

func (opts CreateCloudServiceCustomPolicyOpts) ToCloudServiceCustomPolicyCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "role")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateCloudServiceCustomPolicy(client *gophercloud.ServiceClient, opts CreateCloudServiceCustomPolicyOptsBuilder) (r QueryResult) {
	b, err := opts.ToCloudServiceCustomPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	url := createCustomPolicyURL(client)
	_, r.Err = client.Post(url, b, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		OkCodes:     []int{201},
	})
	return
}

func (opts UpdateCloudServiceCustomPolicyOpts) ToCloudServiceCustomPolicyUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "role")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateCloudServiceCustomPolicy(client *gophercloud.ServiceClient, roleId string, opts UpdateCloudServiceCustomPolicyOpts) (r QueryResult) {
	b, err := opts.ToCloudServiceCustomPolicyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	url := updateCustomPolicyURL(client, roleId)
	_, r.Err = client.Patch(url, b, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		OkCodes:     []int{200},
	})
	return
}

type AgencyCustomPolicyStatementResource struct {
	Uri []string `json:"uri"`
}

type AgencyCustomPolicyStatement struct {
	Effect   string                              `json:"Effect" required:"true"`
	Action   []string                            `json:"Action" required:"true"`
	Resource AgencyCustomPolicyStatementResource `json:"Resource" required:"true"`
}
type AgencyCustomPolicy struct {
	Version   string                        `json:"Version" required:"true"`
	Statement []AgencyCustomPolicyStatement `json:"Statement" required:"true"`
}

type CreateAgencyCustomPolicyOpts struct {
	DisplayName   string             `json:"display_name" required:"true"`
	Type          string             `json:"type" required:"true"`
	Description   string             `json:"description" required:"true"`
	DescriptionCn string             `json:"description_cn,omitempty"`
	Policy        AgencyCustomPolicy `json:"policy" required:"true"`
}

type UpdateAgencyCustomPolicyOpts struct {
	DisplayName   string             `json:"display_name" required:"true"`
	Type          string             `json:"type" required:"true"`
	Description   string             `json:"description" required:"true"`
	DescriptionCn string             `json:"description_cn,omitempty"`
	Policy        AgencyCustomPolicy `json:"policy" required:"true"`
}

type CreateAgencyCustomPolicyOptsBuilder interface {
	ToAgencyCustomPolicyCreateMap() (map[string]interface{}, error)
}

func (opts CreateAgencyCustomPolicyOpts) ToAgencyCustomPolicyCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "role")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateAgencyCustomPolicy(client *gophercloud.ServiceClient, opts CreateAgencyCustomPolicyOptsBuilder) (r QueryResult) {
	b, err := opts.ToAgencyCustomPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	url := createCustomPolicyURL(client)
	_, r.Err = client.Post(url, b, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		OkCodes:     []int{201},
	})
	return
}

type UpdateAgencyCustomPolicyOptsBuilder interface {
	ToAgencyCustomPolicyUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateAgencyCustomPolicyOpts) ToAgencyCustomPolicyUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "role")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateAgencyCustomPolicy(client *gophercloud.ServiceClient, roleId string, opts UpdateAgencyCustomPolicyOptsBuilder) (r QueryResult) {
	b, err := opts.ToAgencyCustomPolicyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	url := updateCustomPolicyURL(client, roleId)
	_, r.Err = client.Patch(url, b, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		OkCodes:     []int{200},
	})
	return
}

func DeleteCustomPolicy(client *gophercloud.ServiceClient, roleId string) (r DeleteResult) {
	url := deleteCustomPolicyURL(client, roleId)
	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		MoreHeaders:  map[string]string{"Content-Type": "application/json;charset=utf8"},
		JSONResponse: nil,
		OkCodes:      []int{200},
	})
	return
}
