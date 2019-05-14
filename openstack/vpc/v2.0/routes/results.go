package routes

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

type Route struct {
	Type        string `json:"type"`
	Nexthop     string `json:"nexthop"`
	Destination string `json:"destination"`
	VpcID       string `json:"vpc_id"`
	TenantID    string `json:"tenant_id"`
	ID          string `json:"id"`
}

type CreateResult struct {
	commonResult
}

func (r CreateResult) Extract() (*Route, error) {
	var entity Route
	err := r.ExtractIntoStructPtr(&entity, "route")
	return &entity, err
}

type DeleteResult struct {
	gophercloud.ErrResult
}

type GetResult struct {
	commonResult
}

func (r GetResult) Extract() (*Route, error) {
	var entity Route
	err := r.ExtractIntoStructPtr(&entity, "route")
	return &entity, err
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*[]Route, error) {
	var list []Route
	err := r.ExtractIntoSlicePtr(&list, "routes")
	return &list, err
}
func (r RoutePage) IsEmpty() (bool, error) {
	list, err := ExtractRoutes(r)
	return len(list) == 0, err
}

type RoutePage struct {
	pagination.LinkedPageBase
}

func ExtractRoutes(r pagination.Page) ([]Route, error) {
	var s struct {
		Routes []Route `json:"routes"`
	}
	err := r.(RoutePage).ExtractInto(&s)
	return s.Routes, err
}

func (r RoutePage) NextPageURL() (string, error) {
	s, err := ExtractRoutes(r)
	if err != nil {
		return "", err
	}
	return r.WrapNextPageURL(s[len(s)-1].ID)
}
