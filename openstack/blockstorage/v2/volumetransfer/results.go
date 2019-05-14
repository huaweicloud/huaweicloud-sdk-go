package volumetransfer

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type TransferDetailsPage struct {
	pagination.SinglePageBase
}

func (r TransferDetailsPage) IsEmpty() (bool, error) {
	volumes, err := ExtractTransferDetails(r)
	return len(volumes) == 0, err
}

func ExtractTransferDetails(r pagination.Page) ([]TransferInfo, error) {
	var s struct {
		TransferInfoList []TransferInfo `json:"transfers"`
	}
	err := (r.(TransferDetailsPage)).ExtractInto(&s)
	return s.TransferInfoList, err
}

type TransferListPage struct {
	pagination.SinglePageBase
}

func (r TransferListPage) IsEmpty() (bool, error) {
	volumes, err := ExtractTransferList(r)
	return len(volumes) == 0, err
}

func ExtractTransferList(r pagination.Page) ([]TransferAccept, error) {
	var s struct {
		TransferList []TransferAccept `json:"transfers"`
	}
	err := (r.(TransferListPage)).ExtractInto(&s)
	return s.TransferList, err
}

type commonResult struct {
	gophercloud.Result
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

type AcceptResult struct {
	commonResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	gophercloud.ErrResult
}

type Links struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type Transfer struct {
	ID        string  `json:"id"`
	CreatedAt string  `json:"created_at"`
	Name      string  `json:"name"`
	VolumeID  string  `json:"volume_id"`
	AuthKey   string  `json:"auth_key"`
	Links     []Links `json:"links"`
}

type TransferAccept struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	VolumeID string  `json:"volume_id"`
	Links    []Links `json:"links"`
}

type TransferInfo struct {
	ID        string  `json:"id"`
	CreatedAt string  `json:"created_at"`
	Name      string  `json:"name"`
	VolumeID  string  `json:"volume_id"`
	Links     []Links `json:"links"`
}

func (r CreateResult) Extract() (Transfer, error) {
	var s struct {
		Transfer Transfer `json:"transfer"`
	}
	err := r.ExtractInto(&s)
	return s.Transfer, err
}

func (r AcceptResult) Extract() (TransferAccept, error) {
	var s struct {
		Transfer TransferAccept `json:"transfer"`
	}
	err := r.ExtractInto(&s)
	return s.Transfer, err
}

func (r GetResult) Extract() (TransferInfo, error) {
	var s struct {
		Transfer TransferInfo `json:"transfer"`
	}
	err := r.ExtractInto(&s)
	return s.Transfer, err
}
