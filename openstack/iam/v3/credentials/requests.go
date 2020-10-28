package credentials

import (
	"github.com/gophercloud/gophercloud"
)

type CreateOptsBuilder interface {
	ToCredentialCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

func (opts CreateOpts) ToCredentialCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "credential")
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r GetResult) {
	b, err := opts.ToCredentialCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, nil)
	return
}

type CreateTemporaryOptsBuilder interface {
	ToTemporaryCredentialCreateMap() (map[string]interface{}, error)
}

type CreateTemporaryOpts struct {
	Auth *Auth `json:"auth,omitempty"`
}

type Auth struct {
	Identity *Identity `json:"identity,omitempty"`
}

type Statement struct {
	Action []string `json:"Action",omitempty`

	Condition map[string]map[string][]string `json:"Condition",omitempty`

	Effect string `json:"Effect",omitempty`

	Resource []string `json:"Resource",omitempty`
}

type Policy struct {
	Version string `json:"Version",omitempty`

	Statement []*Statement `json:"Statement",omitempty`
}

type Identity struct {
	AssumeRole *AssumeRole `json:"assume_role",omitempty`
	Methods    []string   `json:"methods",omitempty`
	Token      *Token      `json:"token",omitempty`
	Policy     *Policy     `json:"policy",omitempty`
}

type Token struct {
	DurationSeconds *int    `json:"duration-seconds",omitempty`
	Id              string `json:"id",omitempty`
}

type AssumeRole struct {
	AgencyName      string      `json:"agency_name",omitempty`
	DomainID        string      `json:"domain_id",omitempty`
	DomainName      string      `json:"domain_name",omitempty`
	DurationSeconds *int         `json:"duration-seconds",omitempty`
	SessionUser     *SessionUser `json:"session_user",omitempty`
}

type SessionUser struct {
	Name string `json:"name",omitempty`
}

func (opts CreateTemporaryOpts) ToTemporaryCredentialCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

func CreateTemporaryAk(client *gophercloud.ServiceClient, opts CreateTemporaryOptsBuilder) (r CreateResult) {
	b, err := opts.ToTemporaryCredentialCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createTemporaryAkURL(client), &b, &r.Body, nil)
	return
}

func DeletePermanentAk(client *gophercloud.ServiceClient, ak string) (r DeleteResult) {
	_, r.Err = client.Delete(deletePermanentAkURL(client, ak), nil)
	return
}

func GetAllPermanentAks(client *gophercloud.ServiceClient, userID string) (r GetResult) {
	_, r.Err = client.Get(getAllPermanentAksURL(client, userID), &r.Body, nil)
	return
}

func GetPermanentAk(client *gophercloud.ServiceClient, aK string) (r GetResult) {
	_, r.Err = client.Get(getPermanentAkURL(client, aK), &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToCredentialUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Description *string `json:"description",omitempty`
	Status      *string `json:"status",omitempty`
}

func (opts UpdateOpts) ToCredentialUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "credential")
}

func UpdatePermanentAk(client *gophercloud.ServiceClient, opts UpdateOptsBuilder, ak string) (r UpdateResult) {
	b, err := opts.ToCredentialUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(updatePermanentAkURL(client, ak), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}