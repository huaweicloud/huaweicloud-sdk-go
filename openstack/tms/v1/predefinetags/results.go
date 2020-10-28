package predefinetags

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)


type commonResult struct {
	gophercloud.Result
}

type ErrorResult struct {
	gophercloud.ErrResult
}

type TagResp struct {
	Key         string     `q:"key"`
	Value       string  `q:"value"`
	Update_Time  string  `q:"update_time"`
}

func ExtractTags(r pagination.Page) ([]TagResp, error) {
	var s struct {
		Tags []TagResp `json:"tags"`
	}
	err := r.(ListPage).ExtractInto(&s)
	return s.Tags, err
}

func (r ErrorResult) Extract() (*Tag, error) {
	var response Tag
	err := r.ExtractInto(&response)
	return &response, err
}