package cloudservers

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"time"
)

type cloudServerResult struct {
	gophercloud.Result
}

type Flavor struct {
	Disk  string `json:"disk"`
	Vcpus string `json:"vcpus"`
	RAM   string `json:"ram"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

// Image defines a image struct in details of a server.
type Image struct {
	ID string `json:"id"`
}

type SysTags struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OsSchedulerHints struct {
	Group []string `json:"group"`
}

// Metadata is only used for method that requests details on a single server, by ID.
// Because metadata struct must be a map.
type Metadata struct {
	ChargingMode      string `json:"charging_mode"`
	OrderID           string `json:"metering.order_id"`
	ProductID         string `json:"metering.product_id"`
	VpcID             string `json:"vpc_id"`
	EcmResStatus      string `json:"EcmResStatus"`
	ImageID           string `json:"metering.image_id"`
	Imagetype         string `json:"metering.imagetype"`
	Resourcespeccode  string `json:"metering.resourcespeccode"`
	ImageName         string `json:"image_name"`
	OsBit             string `json:"os_bit"`
	LockCheckEndpoint string `json:"lock_check_endpoint"`
	LockSource        string `json:"lock_source"`
	LockSourceID      string `json:"lock_source_id"`
	LockScene         string `json:"lock_scene"`
	VirtualEnvType    string `json:"virtual_env_type"`
}

type Address struct {
	Version string `json:"version"`
	Addr    string `json:"addr"`
	MacAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
	PortID  string `json:"OS-EXT-IPS:port_id"`
	Type    string `json:"OS-EXT-IPS:type"`
}

type VolumeAttached struct {
	ID                  string `json:"id"`
	DeleteOnTermination string `json:"delete_on_termination"`
	BootIndex           string `json:"bootIndex"`
	Device              string `json:"device"`
}

type SecurityGroups struct {
	Name string `json:"name"`
}

// CloudServer is only used for method that requests details on a single server, by ID.
// Because metadata struct must be a map.
type CloudServer struct {
	Status              string               `json:"status"`
	Updated             time.Time            `json:"updated"`
	HostID              string               `json:"hostId"`
	Addresses           map[string][]Address `json:"addresses"`
	ID                  string               `json:"id"`
	Name                string               `json:"name"`
	AccessIPv4          string               `json:"accessIPv4"`
	AccessIPv6          string               `json:"accessIPv6"`
	Created             time.Time            `json:"created"`
	Tags                []string             `json:"tags"`
	Description         string               `json:"description"`
	Locked              *bool                `json:"locked"`
	ConfigDrive         string               `json:"config_drive"`
	TenantID            string               `json:"tenant_id"`
	UserID              string               `json:"user_id"`
	HostStatus          string               `json:"host_status"`
	EnterpriseProjectID string               `json:"enterprise_project_id"`
	SysTags             []SysTags            `json:"sys_tags"`
	Flavor              Flavor               `json:"flavor"`
	Metadata            Metadata             `json:"metadata"`
	SecurityGroups      []SecurityGroups     `json:"security_groups"`
	KeyName             string               `json:"key_name"`
	Image               Image                `json:"image"`
	Progress            *int                 `json:"progress"`
	PowerState          *int                 `json:"OS-EXT-STS:power_state"`
	VMState             string               `json:"OS-EXT-STS:vm_state"`
	TaskState           string               `json:"OS-EXT-STS:task_state"`
	DiskConfig          string               `json:"OS-DCF:diskConfig"`
	AvailabilityZone    string               `json:"OS-EXT-AZ:availability_zone"`
	LaunchedAt          string               `json:"OS-SRV-USG:launched_at"`
	TerminatedAt        string               `json:"OS-SRV-USG:terminated_at"`
	RootDeviceName      string               `json:"OS-EXT-SRV-ATTR:root_device_name"`
	RamdiskID           string               `json:"OS-EXT-SRV-ATTR:ramdisk_id"`
	KernelID            string               `json:"OS-EXT-SRV-ATTR:kernel_id"`
	LaunchIndex         *int                 `json:"OS-EXT-SRV-ATTR:launch_index"`
	ReservationID       string               `json:"OS-EXT-SRV-ATTR:reservation_id"`
	Hostname            string               `json:"OS-EXT-SRV-ATTR:hostname"`
	UserData            string               `json:"OS-EXT-SRV-ATTR:user_data"`
	Host                string               `json:"OS-EXT-SRV-ATTR:host"`
	InstanceName        string               `json:"OS-EXT-SRV-ATTR:instance_name"`
	HypervisorHostname  string               `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
	VolumeAttached      []VolumeAttached     `json:"os-extended-volumes:volumes_attached"`
	OsSchedulerHints    OsSchedulerHints     `json:"os:scheduler_hints"`
}

// NewCloudServer defines the response from details on a single server, by ID.
type NewCloudServer struct {
	CloudServer
	Metadata map[string]string `json:"metadata"`
}

// GetResult is the response from a Get operation. Call its Extract
// method to interpret it as a Server.
type GetResult struct {
	cloudServerResult
}

func (r GetResult) Extract() (*CloudServer, error) {
	var s struct {
		Server *CloudServer `json:"server"`
	}
	err := r.ExtractInto(&s)
	return s.Server, err
}

type RecoveryResult struct {
	gophercloud.Result
}

type Recovery struct {
	SupportAutoRecovery string `json:"support_auto_recovery"`
}

func (r RecoveryResult) Extract() (*Recovery, error) {
	var s Recovery

	err := r.ExtractInto(&s)

	return &s, err
}

//BatchChangeResult defining the result struct of batch change OS function.
type BatchChangeResult struct {
	gophercloud.Result
}

//JobResult defining the result struct of job.
type JobResult struct {
	gophercloud.Result
}

//Job defining the struct of job.
type Job struct {
	ID string `json:"job_id"`
}

//ExtractJob defining the result of batch change OS function by extracting job.
func (r BatchChangeResult) ExtractJob() (*Job, error) {
	var j *Job
	err := r.ExtractInto(&j)
	return j, err
}

//ExtractJob defining the result by extracting job.
func (r JobResult) ExtractJob() (Job, error) {
	var j Job
	err := r.ExtractInto(&j)
	return j, err
}

type ErrResult struct {
	gophercloud.ErrResult
}

// CloudServerDetail defines struct of server detail list result.
type CloudServerDetail struct {
	Servers []NewCloudServer `json:"servers"`
	Count   int              `json:"count"`
}

// CloudServerPage is a pagination.Pager that is returned from a call to the List function.
type CloudServerPage struct {
	pagination.OffsetPage
}

// IsEmpty returns true if a ListResult contains no services.
func (r CloudServerPage) IsEmpty() (bool, error) {
	data, err := ExtractCloudServers(r)
	return len(data.Servers) == 0, err
}

// ExtractCloudServers is a function that takes a ListResult and returns the services' information.
func ExtractCloudServers(r pagination.Page) (CloudServerDetail, error) {
	var s CloudServerDetail
	err := (r.(CloudServerPage)).ExtractInto(&s)
	return s, err
}

//BatchUpdateResult defining the result struct of batch updating.
type BatchUpdateResult struct {
	gophercloud.Result
}

//ServerID defines struct of batch update response element.
type ServerID struct {
	ID string `json:"id"`
}

//BatchUpdateResp defines struct of batch update response.
type BatchUpdateResp struct {
	Response []ServerID `json:"response"`
}

//ExtractBatchUpdate defining the result by extracting response.
func (r BatchUpdateResult) ExtractBatchUpdate() (BatchUpdateResp, error) {
	var j BatchUpdateResp
	err := r.ExtractInto(&j)
	return j, err
}

//ProjectTagsResult defining the result struct of project tags.
type ProjectTagsResult struct {
	gophercloud.Result
}

//Tags defining the struct of tags.
type Tags struct {
	Tags []Tag `json:"tags"`
}

//Tag defining the struct of tag element.
type Tag struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

//Extract defining the result of listing tags by extracting
func (r ProjectTagsResult) Extract() (Tags, error) {
	var s Tags
	err := r.ExtractInto(&s)
	return s, err
}

//ServerTagsResult defining the result struct of server tags.
type ServerTagsResult struct {
	gophercloud.Result
}

//ServerTags defining the struct of tags.
type ServerTags struct {
	Tags []ResourceTag `json:"tags"`
}

//ResourceTag defining the struct of server tag element.
type ResourceTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Extract defining the result of listing tags by extracting
func (r ServerTagsResult) Extract() (ServerTags, error) {
	var s ServerTags
	err := r.ExtractInto(&s)
	return s, err
}
