package nics

import "github.com/gophercloud/gophercloud"

type Job struct {
	//jobId
	Id string `json:"job_id"`
}

// AddResult represents the result of a add operation.
// Call its ExtractJob method to get the JobId.
type AddResult struct {
	nicsResult
}

type nicsResult struct {
	gophercloud.Result
}

// ExtractJob is a function that accepts a result get the JobId
func (r nicsResult) ExtractJob() (Job, error) {
	var j Job
	err := r.ExtractInto(&j)
	return j, err
}

// DelResult represents the result of a delete operation.
// Call its ExtractJob method to get the JobId.
type DelResult struct {
	nicsResult
}

// BindResult represents the result of bind operation and unbind operation.
// Call its ExtractPortId method to get the portId.
type BindResult struct {
	nicsResult
}

type Port struct {
	// Specifies the ECS NIC ID.
	PortId string `json:"port_id"`
}

// ExtractPortId is a function that accepts a result get the PortId
func (r BindResult) ExtractPortId() (string, error) {
	var j Port
	err := r.ExtractInto(&j)
	return j.PortId, err
}
