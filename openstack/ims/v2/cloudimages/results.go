package cloudimages

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Image represents an image found in the OpenStack Image service.
type Image struct {
	File                   string   `json:"file"`
	Owner                  string   `json:"owner"`
	ID                     string   `json:"id"`
	Size                   int64    `json:"size"`
	Self                   string   `json:"self"`
	Schema                 string   `json:"schema"`
	Status                 string   `json:"status"`
	Tags                   []string `json:"tags"`
	Visibility             string   `json:"visibility"`
	Name                   string   `json:"name"`
	Checksum               string   `json:"checksum"`
	Deleted                bool     `json:"deleted"`
	Protected              bool     `json:"protected"`
	ContainerFormat        string   `json:"container_format"`
	MinRam                 int      `json:"min_ram"`
	UpdatedAt              string   `json:"updated_at"`
	OsBit                  string   `json:"__os_bit"`
	OsVersion              string   `json:"__os_version"`
	Description            string   `json:"__description"`
	DiskFormat             string   `json:"disk_format"`
	Isregistered           string   `json:"__isregistered"`
	Platform               string   `json:"__platform"`
	OsType                 string   `json:"__os_type"`
	MinDisk                int      `json:"min_disk"`
	VirtualEnvType         string   `json:"virtual_env_type"`
	ImageSourceType        string   `json:"__image_source_type"`
	Imagetype              string   `json:"__imagetype"`
	CreatedAt              string   `json:"created_at"`
	VirtualSize            int      `json:"virtual_size"`
	DeletedAt              string   `json:"deleted_at"`
	Originalimagename      string   `json:"__originalimagename"`
	BackupID               string   `json:"__backup_id"`
	Productcode            string   `json:"__productcode"`
	ImageSize              string   `json:"__image_size"`
	DataOrigin             string   `json:"__data_origin"`
	SupportKvm             string   `json:"__support_kvm"`
	SupportXen             string   `json:"__support_xen"`
	SupportDiskintensive   string   `json:"__support_diskintensive"`
	SupportHighperformance string   `json:"__support_highperformance"`
	SupportXenGpuType      string   `json:"__support_xen_gpu_type"`
	IsConfigInit           string   `json:"__is_config_init"`
	SystemSupportMarket    bool     `json:"__system_support_market"`
}

//执行创建image异步接口时返回的jobid结构
type Job struct {
	// job id of create image
	Id string `json:"job_id"`
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
	Entities Entity `json:"entities"`
}

type Entity struct {
	ImageId string `json:"image_id"`
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

// ImagePage represents the results of a List request.
type ImagePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if an ImagePage contains no Images results.
func (r ImagePage) IsEmpty() (bool, error) {
	images, err := ExtractImages(r)
	return len(images) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to
// the next page of results.
func (r ImagePage) NextPageURL() (string, error) {
	var s struct {
		Next string `json:"next"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}

	if s.Next == "" {
		return "", nil
	}

	return nextPageURL(r.URL.String(), s.Next)
}

// ExtractImages interprets the results of a single page from a List() call,
// producing a slice of Image entities.
func ExtractImages(r pagination.Page) ([]Image, error) {
	var s struct {
		Images []Image `json:"images"`
	}
	err := (r.(ImagePage)).ExtractInto(&s)
	return s.Images, err
}

func (r commonResult) ExtractJob() (*Job, error) {
	var j *Job
	err := r.ExtractInto(&j)
	return j, err
}

func (r commonResult) ExtractJobResult() (*JobResult, error) {
	var jr *JobResult
	err := r.ExtractInto(&jr)
	return jr, err
}
