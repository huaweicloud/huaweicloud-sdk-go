package flavor

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type Flavor struct {
	// Specifies the ID of ECS specifications.
	ID string `json:"id"`

	// Specifies the name of the ECS specifications.
	Name string `json:"name"`

	// Specifies the number of CPU cores in the ECS specifications.
	Vcpus string `json:"vcpus"`

	// Specifies the memory size (MB) in the ECS specifications.
	Ram int64 `json:"ram"`

	// Specifies the system disk size in the ECS specifications.
	// The value 0 indicates that the disk size is not limited.
	Disk string `json:"disk"`

	// Specifies shortcut links for ECS flavors.
	Links []Link `json:"links"`

	// Specifies extended ECS specifications.
	OsExtraSpecs OsExtraSpecs `json:"os_extra_specs"`

	// Reserved
	Swap string `json:"swap"`

	// Reserved
	FlvEphemeral int64 `json:"OS-FLV-EXT-DATA:ephemeral"`

	// Reserved
	FlvDisabled bool `json:"OS-FLV-DISABLED:disabled"`

	// Reserved
	RxtxFactor int64 `json:"rxtx_factor"`

	// Reserved
	RxtxQuota string `json:"rxtx_quota"`

	// Reserved
	RxtxCap string `json:"rxtx_cap"`

	// Reserved
	AccessIsPublic bool `json:"os-flavor-access:is_public"`

	// quota:attachableQuantity
	AttachableQuantity AttachableQuantity `json:"attachableQuantity"`
}

type Link struct {
	// Specifies the shortcut link marker name.
	Rel string `json:"rel"`

	// Provides the corresponding shortcut link.
	Href string `json:"href"`

	// Specifies the shortcut link type.
	Type string `json:"type"`
}

type OsExtraSpecs struct {
	// Specifies the ECS specifications types
	PerformanceType string `json:"ecs:performancetype"`

	// Specifies the resource type.
	ResourceType string `json:"resource_type"`

	// Specifies the vnic type.
	InstanceVnicType string `json:"instance_vnic:type"`

	// Specifies the vnic instance bandwidth.
	InstanceVnicBandwidth int64 `json:"instance_vnic:instance_bandwidth"`

	// Specifies the vnic maxCount.
	InstanceVnicMaxCount int `json:"instance_vnic:max_count"`

	// Specifies the quota local disk.
	QuotaLocalDisk string `json:"quota:local_disk"`

	// Specifies the quota nvme ssd.
	QuotaLocalNvmeSsd string `json:"quota:nvme_ssd"`

	// Specifies the io persistent grant.
	IoPersistentGrant bool `json:"extra_spec:io:persistent_grant"`

	// Specifies the generation of an ECS type
	Generation string `json:"ecs:generation"`

	// Specifies a virtualization type
	VirtualizationEnvTypes string `json:"ecs:virtualization_env_types"`

	// Indicates whether the GPU is passthrough.
	PciPassthroughEnableGpu string `json:"pci_passthrough:enable_gpu"`

	// Indicates the technology used on the G1 and G2 ECSs,
	// including GPU virtualization and GPU passthrough.
	PciPassthroughGpuSpecs string `json:"pci_passthrough:gpu_specs"`

	// Indicates the model and quantity of passthrough-enabled GPUs on P1 ECSs.
	PciPassthroughAlias string `json:"pci_passthrough:alias"`

	// cond:operation:status
	CondOperationStatus string `json:"cond:operation:status"`

	// cond:operation:az
	CondOperationAz string `json:"cond:operation:az"`

	// quota:max_rate
	QuotaMaxRate string `json:"quota:max_rate"`

	// quota:min_rate
	QuotaMinRate string `json:"quota:min_rate"`

	// quota:max_pps
	QuotaMaxPps string `json:"quota:max_pps"`

}

type AttachableQuantity struct {
	FreeScsi int `json:"free_scsi"`
	FreeBlk  int `json:"free_blk"`
	FreeDisk int `json:"free_disk"`
	FreeNik  int `json:"free_nic"`
}

// FlavorsPage is the page returned by a pager when traversing over a
// collection of flavor.
type FlavorsPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a FlavorsPage struct is empty.
func (r FlavorsPage) IsEmpty() (bool, error) {
	is, err := ExtractFlavors(r)
	return len(is.Fs) == 0, err
}

// ExtractFlavors accepts a Page struct, specifically a FlavorsPage struct,
// and extracts the elements into a slice of flavor structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractFlavors(r pagination.Page) (Flavors, error) {
	var s Flavors
	err := (r.(FlavorsPage)).ExtractInto(&s)
	return s, err
}

type Flavors struct {
	Fs []Flavor `json:"flavors"`
}

type Job struct {
	Id string `json:"job_id"`
}

// ResizeResult represents the result of a create operation. Call its ExtractJob
// method to interpret it as a Job.
type ResizeResult struct {
	flavorResult
}

type flavorResult struct {
	gophercloud.Result
}

// ExtractJob is a function that accepts a result and extracts a Job.
func (r ResizeResult) ExtractJob() (Job, error) {
	var j Job
	err := r.ExtractInto(&j)
	return j, err
}
