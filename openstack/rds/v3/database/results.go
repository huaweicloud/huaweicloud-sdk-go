package database

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ActionResult represents the result of server action operations, like reboot.
// Call its ExtractErr method to determine if the action succeeded or failed.

type CreateDatabaseResult struct {
	gophercloud.Result
}

type CreateDatabaseResp struct {
	Resp string `json:"resp"`
}

func (r CreateDatabaseResult) Extract() (*CreateDatabaseResp, error) {
	var response CreateDatabaseResp
	err := r.ExtractInto(&response)
	return &response, err
}


type ListDataBaseResp struct {
	DatabasesList  []Databases `json:"databases"`
	Totalcount int        `json:"total_count"`
}
type Databases struct {
	Name string `json:"name"`
	CharacterSet string `json:"character_set"`
}

type DataBasePage struct {
	pagination.Offset
}

func (r DataBasePage) IsEmpty() (bool, error) {
	data, err := ExtractDataBase(r)
	if err != nil {
		return false, err
	}
	return len(data.DatabasesList) == 0, err
}


func ExtractDataBase(r pagination.Page) (ListDataBaseResp, error) {
	var s ListDataBaseResp
	err := (r.(DataBasePage)).ExtractInto(&s)
	return s, err
}

type commonResult struct {
	gophercloud.Result
}

type DeleteDataBase struct {
	Resp string `json:"resp"`
}

func (r commonResult) Extract() (*DeleteDataBase, error) {
	var response DeleteDataBase
	err := r.ExtractInto(&response)
	return &response, err
}
