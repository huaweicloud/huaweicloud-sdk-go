package agency

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type Agency struct {
	CreateTime      string `json:"create_time"`
	Description     string `json:"description"`
	DomainID        string `json:"domain_id"`
	Duration        string `json:"duration"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	TrustDomainID   string `json:"trust_domain_id"`
	TrustDomainName string `json:"trust_domain_name"`
	ExpireTime      string `json:"expire_time"`
}

type agencyResult struct {
	gophercloud.Result
}

type ListResult struct {
	agencyResult
}

func (r ListResult) ExtractList() (*ListResponse, error) {
	var s ListResponse
	err := r.ExtractInto(&s)
	return &s, err
}

type ListResponse struct {
	Agencies []Agency `json:"agencies"`
}

type GetResult struct {
	agencyResult
}

func (r GetResult) ExtractGet() (*Agency, error) {
	var s struct {
		Agency Agency `json:"agency"`
	}
	err := r.ExtractInto(&s)
	return &s.Agency, err
}

type CreateResult struct {
	agencyResult
}

func (r CreateResult) ExtractCreate() (*Agency, error) {
	var s struct {
		Agency Agency `json:"agency"`
	}
	err := r.ExtractInto(&s)
	return &s.Agency, err
}

type UpdateResult struct {
	agencyResult
}

func (r UpdateResult) ExtractUpdate() (*Agency, error) {
	var s struct {
		Agency Agency `json:"agency"`
	}
	err := r.ExtractInto(&s)
	return &s.Agency, err
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type ListRolesResponse struct {
	Roles []Roles `json:"roles"`
}

type Statement struct {
	Action []string `json:"Action"`
	Effect string   `json:"Effect"`
}

type Policy struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Roles struct {
	Flag          string                 `json:"flag"`
	DisplayName   string                 `json:"display_name"`
	Description   string                 `json:"description"`
	Name          string                 `json:"name"`
	Policy        Policy                 `json:"policy"`
	DescriptionCn string                 `json:"description_cn"`
	DomainID      string                 `json:"domain_id"`
	Type          string                 `json:"type"`
	Catalog       string                 `json:"catalog"`
	ID            string                 `json:"id"`
	CreatedTime   string                 `json:"created_time"`
	UpdatedTime   string                 `json:"updated_time"`
	Links         map[string]interface{} `json:"links"`
}

type rolesResult struct {
	gophercloud.Result
}

type ListRolesResult struct {
	rolesResult
}

func (r ListRolesResult) ExtractListRolesDomain() (*ListRolesResponse, error) {
	var s ListRolesResponse
	err := r.ExtractInto(&s)
	return &s, err
}

func (r ListRolesResult) ExtractListRolesProject() (*ListRolesResponse, error) {
	var s ListRolesResponse
	err := r.ExtractInto(&s)
	return &s, err
}

type PutResult struct {
	gophercloud.ErrResult
}

type HeadResult struct {
	gophercloud.ErrResult
}

type AgencyPage struct {
	pagination.LinkedPageBase
}

func (p AgencyPage) IsEmpty() (bool, error) {
	agency, err := ExtractAgency(p)
	return len(agency) == 0, err
}

func (r AgencyPage) NextPageURL() (string, error) {
	var s struct {
		Links struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
		} `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Links.Next, err
}

func ExtractAgency(r pagination.Page) ([]Agency, error) {
	var s struct {
		Agencies []Agency `json:"agencies"`
	}
	err := (r.(AgencyPage)).ExtractInto(&s)
	return s.Agencies, err
}
