package job

import "github.com/gophercloud/gophercloud"

type jobResult struct {
	gophercloud.Result
}

type JobResult struct {
	// Specifies the task ID.
	Id string `json:"job_id"`

	// Specifies the task type.
	Type string `json:"job_type"`

	//Specifies the task status.
	//  SUCCESS: indicates the task is successfully executed.
	//  RUNNING: indicates that the task is in progress.
	//  FAIL: indicates that the task failed.
	//  INIT: indicates that the task is being initialized.
	Status string `json:"status"`

	// Specifies the time when the task started.
	BeginTime string `json:"begin_time"`

	// Specifies the time when the task finished.
	EndTime string `json:"end_time"`

	// Specifies the returned error code when the task execution fails.
	ErrorCode string `json:"error_code"`

	// Specifies the cause of the task execution failure.
	FailReason string `json:"fail_reason"`

	// Specifies the object of the task.
	Entities JobEntity `json:"entities"`
}

type JobEntity struct {
	// Specifies the number of subtasks.
	// When no subtask exists, the value of this parameter is 0.
	SubJobsTotal int `json:"sub_jobs_total"`

	// Specifies the execution information of a subtask.
	// When no subtask exists, the value of this parameter is left blank.
	SubJobs []SubJob `json:"sub_jobs"`
}

type SubJob struct {
	// Specifies the task ID.
	Id string `json:"job_id"`

	// Task type.
	Type string `json:"job_type"`

	//Specifies the task status.
	//  SUCCESS: indicates the task is successfully executed.
	//  RUNNING: indicates that the task is in progress.
	//  FAIL: indicates that the task failed.
	//  INIT: indicates that the task is being initialized.
	Status string `json:"status"`

	// Specifies the time when the task started.
	BeginTime string `json:"begin_time"`

	// Specifies the time when the task finished.
	EndTime string `json:"end_time"`

	// Specifies the returned error code when the task execution fails.
	ErrorCode string `json:"error_code"`

	// Specifies the cause of the task execution failure.
	FailReason string `json:"fail_reason"`

	// Specifies the object of the task.
	Entities *SubJobEntity `json:"entities"`
}

type SubJobEntity struct {
	// If the task is an ECS-related operation, the value is server_id.
	ServerId string `json:"server_id"`

	// If the task is a NIC-related operation, the value is nic_id.
	NicId    string `json:"nic_id"`
}

// JobExecResult represents the result of a get operation. Call its ExtractJobResult
// method to interpret it as a jobresult.
type JobExecResult struct {
	jobResult
}

// ExtractJobResult is a function that accepts a result and extracts a jobresult.
func (r jobResult) ExtractJobResult() (JobResult, error) {
	var jr JobResult
	err := r.ExtractInto(&jr)
	return jr, err
}
