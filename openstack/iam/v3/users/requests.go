package users

import "github.com/gophercloud/gophercloud"

func Get(client *gophercloud.ServiceClient, userId string) (r GetResult) {
	url := queryUserDetailUrl(client, userId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type CreateUserOpts struct {
	Areacode    string `json:"areacode,omitempty"`
	Description string `json:"description,omitempty"`
	DomainId    string `json:"domain_id,omitempty"`
	Email       string `json:"email,omitempty"`
	Enabled     *bool  `json:"enabled,omitempty"`
	Name        string `json:"name" required:"true"`
	Password    string `json:"password,omitempty"`
	Phone       string `json:"phone,omitempty"`
	PwdStatus   *bool  `json:"pwd_status,omitempty"`
	XuserId     string `json:"xuser_id,omitempty"`
	XuserType   string `json:"xuser_type,omitempty"`
}

type CreateUserOptsBuilder interface {
	ToUserCreateMap() (map[string]interface{}, error)
}

func (opts CreateUserOpts) ToUserCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "user")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opt CreateUserOptsBuilder) (r CreatResult) {
	b, err := opt.ToUserCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	url := createUserUrl(client)
	_, r.Err = client.Post(url, &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201}})
	return
}

type UpdateUserInfoOpts struct {
	Email  string `json:"email,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}

type UpdateUserInfoOptsBuilder interface {
	ToUserInfoUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateUserInfoOpts) ToUserInfoUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "user")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateUserInfo(client *gophercloud.ServiceClient, userId string, opt UpdateUserInfoOptsBuilder) (r UpdateResult) {
	b, err := opt.ToUserInfoUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	url := updateUserInfoUrl(client, userId)
	_, r.Err = client.Put(url, &b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204}})
	return
}

type UpdateUserOpts struct {
	Areacode    string `json:"areacode,omitempty"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email,omitempty"`
	Enabled     *bool  `json:"enabled,omitempty"`
	Name        string `json:"name,omitempty"`
	Password    string `json:"password,omitempty"`
	Phone       string `json:"phone,omitempty"`
	PwdStatus   *bool  `json:"pwd_status,omitempty"`
	XuserId     string `json:"xuser_id,omitempty"`
	XuserType   string `json:"xuser_type,omitempty"`
}

type UpdateUserOptsBuilder interface {
	ToUserUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateUserOpts) ToUserUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "user")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateUserInfoByAdmin(client *gophercloud.ServiceClient, userId string, opt UpdateUserOptsBuilder) (r UpdateResult) {
	b, err := opt.ToUserUpdateMap()
	if err != nil {
		r.Err = err
	}

	url := updateUserUrl(client, userId)
	_, r.Err = client.Put(url, &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200}})
	return
}
