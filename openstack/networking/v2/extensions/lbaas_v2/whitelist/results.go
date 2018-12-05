package whitelist

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type WhiteList struct {
	WhiteList Whitelist `json:"whitelist"`
}
type WhiteLists struct {
	Whitelists []Whitelist `json:"whitelists"`
}

type Whitelist struct {
	//Specifies the whitelist ID.
	ID              string `json:"id"`
	//Specifies the tenant ID.
	TenantId        string `json:"tenant_id"`
	//Specifies the listener ID.
	ListenerId      string `json:"listener_id"`
	//Specifies whether to enable the access control.
	EnableWhitelist bool   `json:"enable_whitelist"`
	//Lists the IP addresses in the whitelist.
	Whitelist       string `json:"whitelist"`
}

// WhiteListPage is the page returned by a pager when traversing over a
// collection of whitelist.
type WhiteListPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a WhiteListPage struct is empty.
func (r WhiteListPage) IsEmpty() (bool, error) {
	is, err := ExtractWhiteLists(r)
	return len(is.Whitelists) == 0, err
}

// NextPageURL will retrieve the next page URL.
func (page WhiteListPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"whitelists_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// ExtractWhiteLists accepts a Page struct, specifically a WhiteListPage struct,
// and extracts the elements into a slice of whitelist structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractWhiteLists(r pagination.Page) (WhiteLists, error) {
	var s WhiteLists
	s.Whitelists = make([]Whitelist,0)
	err := (r.(WhiteListPage)).ExtractInto(&s)
	return s, err
}

type whitelistResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a whitelist.
func (r whitelistResult) Extract() (WhiteList, error) {
	var s WhiteList
	err := r.ExtractInto(&s)
	return s, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a WhiteList.
type CreateResult struct {
	whitelistResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a WhiteList.
type GetResult struct {
	whitelistResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a WhiteList.
type UpdateResult struct {
	whitelistResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
