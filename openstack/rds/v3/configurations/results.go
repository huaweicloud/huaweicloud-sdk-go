package configurations

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type GetConfigurationsResponse struct {
	ConfigurationsList []Configurations `json:"configurations"`
	Bucket             string           `json:"bucket"`
}

type Configurations struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	DatastoreVersionName string `json:"datastore_version_name"`
	DatastoreName        string `json:"datastore_name"`
	Created              string `json:"created"`
	Updated              string `json:"updated"`
	UserDefined          bool   `json:"user_defined"`
}
type ConfigurationsPage struct {
	pagination.Offset
}


func (r ConfigurationsPage) IsEmpty() (bool, error) {
	data, err := ExtractGetConfigurations(r)
	if err != nil {
		return false, err
	}
	return len(data.ConfigurationsList) == 0, err
}


func ExtractGetConfigurations(r pagination.Page) (GetConfigurationsResponse, error) {
	var s GetConfigurationsResponse
	err := (r.(ConfigurationsPage)).ExtractInto(&s)
	return s, err
}

type CreateConfigurationsRes struct {
	Id                   string `json:"id"`
	Name                 string `json:"name" `
	Description          string `json:"description"`
	DatastoreVersionName string `json:"datastore_version_name" `
	DatastoreName        string `json:"datastore_name"`
	Created              string `json:"created"`
	Updated              string `json:"updated" `
}

type CreateConfigurations struct {
	CreateConfigurationsRes `json:"configuration" required:"true"`
}

type commonResult struct {
	gophercloud.Result
}
type CreateConfigurationsResult struct {
	commonResult
}

func (r CreateConfigurationsResult)Extract() (*CreateConfigurations, error) {
	var response CreateConfigurations
	err := r.ExtractInto(&response)
	return &response, err
}
