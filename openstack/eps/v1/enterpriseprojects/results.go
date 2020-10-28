package enterpriseprojects

import (
	"github.com/gophercloud/gophercloud"
)

type CommonResult struct {
	gophercloud.Result
}

type ErrorResult struct {
	gophercloud.ErrResult
}

type EnterpriseProject struct {
	Id          string `q:"id"`
	Name        string `q:"name"`
	Description string `q:"description"`
	Status      int    `q:"status"`
	CreatedAt   string `q:"created_at"`
	UpdatedAt   string `q:"updated_at"`
}

type CreateEP struct {
	Enterprise_Project EnterpriseProject `q:"enterprise_project"`
}

type ListEPResp struct {
	Enterprise_projects []EnterpriseProject `q:"enterprise_projects"`
	Total_count         int                 `q:"total_count"`
}

type ListResult struct {
	gophercloud.Result
}

func (r ListResult) Extract() (*ListEPResp, error) {
	var response ListEPResp
	err := r.ExtractInto(&response)
	return &response, err
}

type VersionResp struct {
	Id          string            `q:"id"`
	Links       []VersionLinkResp `json:"links,omitempty"`
	Status      string            `q:"status"`
	Updated     string            `q:"updated"`
	Min_version string            `q:"min_version"`
}

type VersionLinkResp struct {
	Rel  string `q:"rel"`
	Href string `q:"href"`
}

type QuotasResource struct {
	Type  string `q:"type"`
	Used  int    `q:"used"`
	Quota int    `q:"quota"`
}

type QuotasResp struct {
	Resources [] QuotasResource `q:"resources"`
}

type Resource struct {
	Project_id            string `q:"project_id"`
	Resource_type         string `q:"resource_type"`
	Project_name          string `q:"project_name"`
	Resource_id           string `q:"resource_id"`
	Resource_name         string `q:"resource_name"`
	Resource_detail       string `q:"resource_detail"`
	Enterprise_project_id string `q:"enterprise_project_id"`
}

type Error struct {
	Project_id    string `q:"project_id"`
	Resource_type string `q:"resource_type"`
	Error_code    string `q:"error_code"`
	Error_msg     string `q:"error_msg"`
}

type FilterResourcesResp struct {
	Resources   [] Resource `q:"resources"`
	Errors      [] Error    `q:"errors"`
	Total_count int         `q:"total_count"`
}

func (r ErrorResult) Extract() (*EnterpriseProject, error) {
	var response EnterpriseProject
	err := r.ExtractInto(&response)
	return &response, err
}

func (r CommonResult) ExtractEP() (EnterpriseProject, error) {
	var s struct {
		Enterprise_project EnterpriseProject `json:"enterprise_project"`
	}
	err := r.ExtractInto(&s)
	return s.Enterprise_project, err
}

func (r CommonResult) ExtractQuotas() (QuotasResp, error) {
	var s struct {
		Quotas QuotasResp `json:"quotas"`
	}
	err := r.ExtractInto(&s)
	return s.Quotas, err
}

func (r CommonResult) ExtractFilterResources() (FilterResourcesResp, error) {
	var s FilterResourcesResp
	err := r.ExtractInto(&s)
	return s, err
}
