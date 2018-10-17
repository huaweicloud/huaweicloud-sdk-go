package certificates

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type Certificate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Domain      string `json:"domain"`
	PrivateKey  string `json:"private_key"`
	Certificate string `json:"certificate"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

type CertificatePage struct {
	pagination.LinkedPageBase
}

func (r CertificatePage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"certificates_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

func (r CertificatePage) IsEmpty() (bool, error) {
	is, err := ExtractCertificates(r)
	return len(is) == 0, err
}

func ExtractCertificates(r pagination.Page) ([]Certificate, error) {
	var s struct {
		Certificates []Certificate `json:"certificates"`
	}
	err := (r.(CertificatePage)).ExtractInto(&s)
	return s.Certificates, err
}

type commonResult struct {
	gophercloud.Result
}

func (r commonResult) Extract() (*Certificate, error) {
	var c Certificate
	err := r.ExtractInto(&c)
	return &c, err
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	gophercloud.ErrResult
}
