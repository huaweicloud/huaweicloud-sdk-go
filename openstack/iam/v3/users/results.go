package users

import "github.com/gophercloud/gophercloud"

type User struct {
	Areacode          string `json:"areacode"`
	DomainId          string `json:"domain_id"`
	Email             string `json:"email"`
	Enabled           bool   `json:"enabled"`
	Id                string `json:"id"`
	Name              string `json:"name"`
	Phone             string `json:"phone"`
	PwdStatus         bool   `json:"pwd_status"`
	XUserId           string `json:"xuser_id"`
	XUserType         string `json:"xuser_type"`
	CreateTime        string `json:"create_time"`
	Description       string `json:"description"`
	IsDomainOwner     string `json:"is_domain_owner"`
	PasswordExpiresAt string `json:"password_expires_at"`
	XdomainId         string `json:"xdomain_id"`
	XdomainType       string `json:"xdomain_type"`
	Status            int    `json:"status"`
}

type UserDetail struct {
	Areacode   string                 `json:"areacode"`
	DomainId   string                 `json:"domain_id"`
	Email      string                 `json:"email"`
	Enabled    bool                   `json:"enabled"`
	Id         string                 `json:"id"`
	Name       string                 `json:"name"`
	Phone      string                 `json:"phone"`
	PwdStatus  bool                   `json:"pwd_status"`
	UpdateTime string                 `json:"update_time"`
	XUserId    string                 `json:"xuser_id"`
	XUserType  string                 `json:"xuser_type"`
	Links      map[string]interface{} `json:"links"`
}

type UserInfo struct {
	Areacode          string                 `json:"areacode"`
	DomainId          string                 `json:"domain_id"`
	Email             string                 `json:"email"`
	Enabled           bool                   `json:"enabled"`
	Id                string                 `json:"id"`
	Name              string                 `json:"name"`
	Phone             string                 `json:"phone"`
	PwdStatus         bool                   `json:"pwd_status"`
	XUserId           string                 `json:"xuser_id"`
	XUserType         string                 `json:"xuser_type"`
	Description       string                 `json:"description"`
	PasswordExpiresAt string                 `json:"password_expires_at"`
	Links             map[string]interface{} `json:"links"`
}

type userResult struct {
	gophercloud.Result
}

type GetResult struct {
	userResult
}

type CreatResult struct {
	userResult
}

type UpdateResult struct {
	userResult
}

func (r GetResult) ExtractGet() (*UserDetail, error) {
	var s struct {
		User UserDetail `json:"user"`
	}
	err := r.ExtractInto(&s)
	return &s.User, err
}

func (r CreatResult) ExtractCreate() (*User, error) {
	var s struct {
		User User `json:"user"`
	}
	err := r.ExtractInto(&s)
	return &s.User, err
}

func (r UpdateResult) ExtractUpdate() (*UserInfo, error) {
	var s struct {
		UserInfo UserInfo `json:"user"`
	}
	err := r.ExtractInto(&s)
	return &s.UserInfo, err
}
