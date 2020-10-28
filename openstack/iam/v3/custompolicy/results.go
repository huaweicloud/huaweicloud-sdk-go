package custompolicy

import (
	"github.com/gophercloud/gophercloud"
)

type CustomPolicyResult struct {
	gophercloud.Result
}

type ListResult struct {
	CustomPolicyResult
}

type ListResponse struct {
	Roles []struct {
		DomainID      string `json:"domain_id"`
		UpdatedTime   string `json:"updated_time"`
		CreatedTime   string `json:"created_time"`
		DescriptionCn string `json:"description_cn"`
		Catalog       string `json:"catalog"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		References    int    `json:"references"`
		Links         struct {
			Self string `json:"self"`
		} `json:"links"`
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		Policy      struct {
			Version   string `json:"Version"`
			Statement []struct {
				Condition interface{} `json:"Condition"`
				Action    []string    `json:"Action"`
				Resource  interface{} `json:"Resource"`
				Effect    string      `json:"Effect"`
			} `json:"Statement"`
		} `json:"policy"`
	} `json:"roles"`
	Links struct {
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
		Self     string      `json:"self"`
	} `json:"links"`
}

func (r ListResult) ExtractList() (*ListResponse, error) {
	var s ListResponse
	err := r.ExtractInto(&s)
	return &s, err
}

type QueryResult struct {
	CustomPolicyResult
}

type CreateCustomPolicyResponse struct {
	Role struct {
		DomainID      string `json:"domain_id"`
		UpdatedTime   string `json:"updated_time"`
		CreatedTime   string `json:"created_time"`
		DescriptionCn string `json:"description_cn"`
		Catalog       string `json:"catalog"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		References    int    `json:"references"`
		Links         struct {
			Self string `json:"self"`
		} `json:"links"`
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		Policy      struct {
			Version   string `json:"Version"`
			Statement []struct {
				Action   []string    `json:"Action"`
				Resource interface{} `json:"Resource"`
				Effect   string      `json:"Effect"`
			} `json:"Statement"`
		} `json:"policy"`
	} `json:"role"`
}

type CustomPolicyResponse struct {
	Role struct {
		DomainID      string `json:"domain_id"`
		UpdatedTime   string `json:"updated_time"`
		CreatedTime   string `json:"created_time"`
		DescriptionCn string `json:"description_cn"`
		Catalog       string `json:"catalog"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		References    int    `json:"references"`
		Links         struct {
			Self string `json:"self"`
		} `json:"links"`
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		Policy      struct {
			Version   string `json:"Version"`
			Statement []struct {
				Condition interface{} `json:"Condition"`
				Action    []string    `json:"Action"`
				Resource  interface{} `json:"Resource"`
				Effect    string      `json:"Effect"`
			} `json:"Statement"`
		} `json:"policy"`
	} `json:"role"`
}

func (r QueryResult) ExtractQuery() (*CustomPolicyResponse, error) {
	var s CustomPolicyResponse
	err := r.ExtractInto(&s)
	return &s, err
}

func (r QueryResult) ExtractCreate() (*CustomPolicyResponse, error) {
	var s CustomPolicyResponse
	err := r.ExtractInto(&s)
	return &s, err
}

func (r QueryResult) ExtractAgencyCustomCreate() (*CreateCustomPolicyResponse, error) {
	var s CreateCustomPolicyResponse
	err := r.ExtractInto(&s)
	return &s, err
}

func (r QueryResult) ExtractPatch() (*CustomPolicyResponse, error) {
	var s CustomPolicyResponse
	err := r.ExtractInto(&s)
	return &s, err
}

type DeleteResult struct {
	gophercloud.ErrResult
}
