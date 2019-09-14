package db_user

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ActionResult represents the result of server action operations, like reboot.
// Call its ExtractErr method to determine if the action succeeded or failed.


type DbUserResult struct {
	gophercloud.Result
}

type DbUserResp struct {
	Resp string `json:"resp"`
}

func (r DbUserResult) Extract() (*DbUserResp, error) {
	var response DbUserResp
	err := r.ExtractInto(&response)
	return &response, err
}

type ListDbUsersResp struct {
	UsersList  []NameList `json:"users"`
	Totalcount int        `json:"total_count"`
}
type NameList struct {
	Name string `json:"name"`
}

type DbUsersPage struct {
	pagination.Offset
}

func (r DbUsersPage) IsEmpty() (bool, error) {
	data, err := ExtractDbUsers(r)
	if err != nil {
		return false, err
	}
	return len(data.UsersList) == 0, err
}

func ExtractDbUsers(r pagination.Page) (ListDbUsersResp, error) {
	var s ListDbUsersResp
	err := (r.(DbUsersPage)).ExtractInto(&s)
	return s, err
}


type commonResult struct {
	gophercloud.Result
}

type DeleteDbUser struct {
	Resp string `json:"resp"`
}

func (r commonResult) Extract() (*DeleteDbUser, error) {
	var response DeleteDbUser
	err := r.ExtractInto(&response)
	return &response, err
}