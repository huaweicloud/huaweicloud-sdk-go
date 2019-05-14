package cloudservers

import (
	"github.com/gophercloud/gophercloud"
)

// ActionResult represents the result of server action operations, like reboot.
// Call its ExtractErr method to determine if the action succeeded or failed.
type ActionResult struct {
	gophercloud.ErrResult
}

type ChangeResult struct {
	gophercloud.Result
}

type Job struct {
	ID string `json:"job_id"`
}

func (r ChangeResult) ExtractJob() (*Job, error) {
	var j *Job
	err := r.ExtractInto(&j)
	return j, err
}

type FlavorResult struct {
	gophercloud.Result
}

type Flavor struct {
	Name                   string     `json:"name"`
	Links                  []Link     `json:"links"`
	RAM                    int        `json:"ram"`
	OSFLVDISABLEDDisabled  bool       `json:"OS-FLV-DISABLED:disabled"`
	Vcpus                  int        `json:"vcpus"`
	ExtraSpecs             ExtraSpecs `json:"extra_specs"`
	Swap                   string     `json:"swap"`
	OsFlavorAccessIsPublic bool       `json:"os-flavor-access:is_public"`
	RxtxFactor             float64    `json:"rxtx_factor"`
	OSFLVEXTDATAEphemeral  int        `json:"OS-FLV-EXT-DATA:ephemeral"`
	Disk                   int        `json:"disk"`
	ID                     string     `json:"id"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type ExtraSpecs struct {
	EcsPerformancetype string `json:"ecs:performancetype"`
	ResourceType       string `json:"resource_type"`
}

func (r FlavorResult) Extract() (*[]Flavor, error) {
	var f *[]Flavor
	err := r.ExtractInto(&f)
	return f, err
}
