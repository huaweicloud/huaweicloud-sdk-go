package agency

import (
	"github.com/gophercloud/gophercloud"
)

type AgenciesListOptsBuilder interface {
	ToAgenciesListQuery() (string, error)
}

type AgenciesListOpts struct {
	DomainId string `q:"domain_id" required:"true`

	TrustDomainId string `q:"trust_domain_id"`

	Name string `q:"name"`
}

func (opts AgenciesListOpts) ToAgenciesListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts AgenciesListOptsBuilder) (r ListResult) {
	url := listAgenciesUrl(client)
	if opts != nil {
		query, err := opts.ToAgenciesListQuery()
		if err != nil {
			return
		}
		url += query
	}
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Get(client *gophercloud.ServiceClient, agencyId string) (r GetResult) {
	url := queryAgencyDetailsUrl(client, agencyId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type CreateAgencyOpts struct {
	Name            string `json:"name" required:"true"`
	DomainID        string `json:"domain_id" required:"true"`
	TrustDomainID   string `json:"trust_domain_id,omitempty"`
	TrustDomainName string `json:"trust_domain_name,omitempty"`
	Duration        string `json:"duration,omitempty"`
	Description     string `json:"description,omitempty"`
}

type CreateAgencyOptsBuilder interface {
	ToAgencyCreateMap() (map[string]interface{}, error)
}

func (opts CreateAgencyOpts) ToAgencyCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "agency")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateAgencyOptsBuilder) (r CreateResult) {
	b, err := opts.ToAgencyCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	url := createAgencyURL(client)
	_, r.Err = client.Post(url, &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

type UpdateAgencyOpts struct {
	TrustDomainID   string `json:"trust_domain_id,omitempty"`
	TrustDomainName string `json:"trust_domain_name,omitempty"`
	Duration        string `json:"duration,omitempty"`
	Description     string `json:"description,omitempty"`
}

type UpdateAgencyOptsBuilder interface {
	ToAgencyCreateMap() (map[string]interface{}, error)
}

func (opts UpdateAgencyOpts) ToAgencyUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "agency")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, agencyId string, opts UpdateAgencyOpts) (r UpdateResult) {
	b, err := opts.ToAgencyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	url := updateAgencyURL(client, agencyId)
	_, r.Err = client.Put(url, &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, agencyId string) (r DeleteResult) {
	url := deleteAgencyURL(client, agencyId)
	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		JSONResponse: nil,
		OkCodes:      []int{204},
	})
	return
}

func ListPermissionsForAgencyOnDomain(client *gophercloud.ServiceClient, domainId string, agencyId string) (r ListRolesResult) {
	url := listPermissionsForAgencyOnDomainURL(client, domainId, agencyId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func ListPermissionsForAgencyOnProject(client *gophercloud.ServiceClient, projectId string, agencyId string) (r ListRolesResult) {
	url := listPermissionsForAgencyOnProjectURL(client, projectId, agencyId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func GrantPermissionToAgencyOnDomain(client *gophercloud.ServiceClient, domainId string, agencyId string, roleId string) (r PutResult) {
	url := grantPermissionToAgencyOnDomainURL(client, domainId, agencyId, roleId)
	_, r.Err = client.Put(url, nil, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

func GrantPermissionToAgencyOnProject(client *gophercloud.ServiceClient, projectId string, agencyId string, roleId string) (r PutResult) {
	url := grantPermissionToAgencyOnProjectURL(client, projectId, agencyId, roleId)
	_, r.Err = client.Put(url, nil, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

func CheckPermissionForAgencyOnDomain(client *gophercloud.ServiceClient, domainId string, agencyId string, roleId string) (r HeadResult) {
	url := checkPermissionForAgencyOnDomainURL(client, domainId, agencyId, roleId)
	_, r.Err = client.Head(url, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

func CheckPermissionForAgencyOnProject(client *gophercloud.ServiceClient, projectId string, agencyId string, roleId string) (r HeadResult) {
	url := checkPermissionForAgencyOnProjectURL(client, projectId, agencyId, roleId)
	_, r.Err = client.Head(url, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

func RemovePermissionFromAgencyOnDomain(client *gophercloud.ServiceClient, domainId string, agencyId string, roleId string) (r DeleteResult) {
	url := removePermissionFromAgencyOnDomainURL(client, domainId, agencyId, roleId)
	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

func RemovePermissionFromAgencyOnProject(client *gophercloud.ServiceClient, projectId string, agencyId string, roleId string) (r DeleteResult) {
	url := removePermissionFromAgencyOnProjectURL(client, projectId, agencyId, roleId)
	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
