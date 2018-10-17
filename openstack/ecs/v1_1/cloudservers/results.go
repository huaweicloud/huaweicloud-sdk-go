package cloudservers

import (
	"github.com/gophercloud/gophercloud"
)

//执行创建image异步接口时返回的jobid结构
type Job struct {
	// job id of create image
	Id string `json:"job_id"`
}

type Order struct {
	Id string `json:"order_id"`
}

//根据jobId查询job的结构
type JobResult struct {
	//任务ID
	Id string `json:"job_id"`

	//任务类型
	Type string `json:"job_type"`

	/*任务状态，目前取值如下：
	  SUCCESS：表示该任务执行已经结束，任务执行成功。
	  FAIL：表示该任务执行已经结束，任务执行失败。
	  RUNNING：表示该任务正在执行。
	  INIT：表示给任务还未执行，正在初始化。*/
	Status string `json:"status"`

	//任务开始时间
	BeginTime string `json:"begin_time"`

	//任务结束时间
	EndTime string `json:"end_time"`

	//错误码
	ErrorCode string `json:"error_code"`

	//失败原因
	FailReason string `json:"fail_reason"`

	//任务自定义属性。任务正常时，会返回镜像的ID。
	Entities JobEntity `json:"entities"`
}

type JobEntity struct {
	//子任务数量。没有子任务时为0。
	SubJobsTotal int `json:"sub_jobs_total"`

	//每个子任务的执行信息。没有子任务时为空列表。
	SubJobs []SubJob `json:"sub_jobs"`
}

type SubJob struct {
	//任务ID
	Id string `json:"job_id"`

	//任务类型
	Type string `json:"job_type"`

	/*任务状态，目前取值如下：
	  SUCCESS：表示该任务执行已经结束，任务执行成功。
	  FAIL：表示该任务执行已经结束，任务执行失败。
	  RUNNING：表示该任务正在执行。
	  INIT：表示给任务还未执行，正在初始化。*/
	Status string `json:"status"`

	//任务开始时间
	BeginTime string `json:"begin_time"`

	//任务结束时间
	EndTime string `json:"end_time"`

	//错误码
	ErrorCode string `json:"error_code"`

	//失败原因
	FailReason string `json:"fail_reason"`

	//任务自定义属性。任务正常时，会返回镜像的ID。
	Entities *SubJobEntity `json:"entities"`
}

type SubJobEntity struct {
	ServerId string `json:"server_id"`
	NicId    string `json:"nic_id"`
}

type commonResult struct {
	gophercloud.Result
}

type CreateResult struct {
	commonResult
}

type JobExecResult struct {
	commonResult
}

func (r commonResult) ExtractJob() (Job, error) {
	var j Job
	err := r.ExtractInto(&j)
	return j, err
}

func (r commonResult) ExtractOrder() (Order, error) {
	var o Order
	err := r.ExtractInto(&o)
	return o, err
}

func (r commonResult) ExtractJobResult() (JobResult, error) {
	var jr JobResult
	err := r.ExtractInto(&jr)
	return jr, err
}
