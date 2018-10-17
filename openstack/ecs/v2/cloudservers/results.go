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
