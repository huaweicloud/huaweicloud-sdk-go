package credentials

import (
	"github.com/gophercloud/gophercloud"
)

type credentialResult struct {
	gophercloud.Result
}

// GetResult is the result of a Get request. Call its Extract method to
// interpret it as a Project.
type GetResult struct {
	credentialResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as a Project.
type UpdateResult struct {
	credentialResult
}

type CreateResult struct {
	credentialResult
}

type DeleteResult struct {
	credentialResult
}

type Credential struct {
	Access string `json:"access"`
	CreateTime string `json:"create_time"`
	Description string `json:"description"`
	Secret string `json:"secret"`
	Status string `json:"status"`
	UserID string `json:"user_id"`
}

type AllPermanentCredential struct {
	Access string `json:"access"`
	CreateTime string `json:"create_time"`
	Description string `json:"description"`
	Status string `json:"status"`
	UserID string `json:"user_id"`
}

type PermanentCredential struct {
	Access string `json:"access"`
	CreateTime string `json:"create_time"`
	Description string `json:"description"`
	LastUseTime string `json:"last_use_time"`
	Status string `json:"status"`
	UserID string `json:"user_id"`
}

type TemporaryCredential struct {
	Access string `json:"access"`
	Secret string `json:"secret"`
	ExpiresAt string `json:"expires_at"`
	SecurityToken string `json:"securitytoken"`
}


// Extract interprets any projectResults as a Project.
func (r credentialResult) Extract() (*Credential, error) {
	var s struct {
		Credential *Credential `json:"credential"`
	}
	err := r.ExtractInto(&s)
	return s.Credential, err
}


func (r credentialResult) ExtractTemporary() (*TemporaryCredential, error) {
	var s struct {
		Credential *TemporaryCredential `json:"credential"`
	}
	err := r.ExtractInto(&s)
	return s.Credential, err
}

func (r credentialResult) ExtractPermanent() (*PermanentCredential, error) {
	var s struct {
		Credential *PermanentCredential `json:"credential"`
	}
	err := r.ExtractInto(&s)
	return s.Credential, err
}

func (r credentialResult) ExtractUpdatePermanent() (*AllPermanentCredential, error) {
	var s struct {
		Credential *AllPermanentCredential `json:"credential"`
	}
	err := r.ExtractInto(&s)
	return s.Credential, err
}

func (r credentialResult) ExtractCredentials() ([]AllPermanentCredential, error) {
	var s struct {
		Credentials []AllPermanentCredential `json:"credentials"`
	}
	err := r.ExtractInto(&s)
	return s.Credentials, err
}
