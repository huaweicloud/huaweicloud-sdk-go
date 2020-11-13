package instances

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}
type CreateResult struct {
	commonResult
}
type RestartRdsInstanceResult struct {
	commonResult
}
type SingleToHaRdsInstanceResult struct {
	commonResult
}

type ResizeFlavorResult struct {
	commonResult
}

type Instance struct {
	Id                  string         `json:"id" `
	Name                string         `json:"name" `
	Status              string         `json:"status" `
	Datastore           Datastore      `json:"datastore"`
	Ha                  Ha             `json:"ha"`
	ConfigurationId     string         `json:"configuration_id"`
	Port                string         `json:"port"`
	BackupStrategy      BackupStrategy `json:"backup_strategy"`
	EnterpriseProjectId string         `json:"enterprise_project_id"`
	DiskEncryptionId    string         `json:"disk_encryption_id"`
	FlavorRef           string         `json:"flavor_ref"`
	Volume              Volume         `json:"volume"`
	Region              string         `json:"region"`
	AvailabilityZone    string         `json:"availability_zone"`
	VpcId               string         `json:"vpc_id"`
	SubnetId            string         `json:"subnet_id"`
	SecurityGroupId     string         `json:"security_group_id"`
	ChargeInfo          ChargeInfo     `json:"charge_info"`
}

type CreateRds struct {
	Instance Instance `json:"instance"`
	JobId    string   `json:"job_id"`
	OrderId  string   `json:"order_id"`
}

func (r CreateResult) Extract() (*CreateRds, error) {
	var response CreateRds
	err := r.ExtractInto(&response)
	return &response, err
}


type DeleteInstanceRdsResult struct {
	commonResult
}

type DeleteInstanceRdsResponse struct {
	JobId string `json:"job_id"`
}

func (r DeleteInstanceRdsResult) Extract() (*DeleteInstanceRdsResponse, error) {
	var response DeleteInstanceRdsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type RestartRdsResponse struct {
	JobId string `json:"job_id"`
}

func (r RestartRdsInstanceResult) Extract() (*RestartRdsResponse, error) {
	var response RestartRdsResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type SingleToHaResponse struct {
	JobId string `json:"job_id"`
}

func (r SingleToHaRdsInstanceResult)Extract() (*SingleToHaResponse, error) {
	var response SingleToHaResponse
	err := r.ExtractInto(&response)
	return &response, err
}

type ResizeFlavor struct {
	JobId string `json:"job_id"`
}

func (r ResizeFlavorResult) Extract() (*ResizeFlavor, error) {
	var response ResizeFlavor
	err := r.ExtractInto(&response)
	return &response, err
}


type EnlargeVolumeResult struct {
	commonResult
}
type EnlargeVolumeResp struct {
	JobId string `json:"job_id"`
}

func (r EnlargeVolumeResult) Extract() (*EnlargeVolumeResp, error) {
	var response EnlargeVolumeResp
	err := r.ExtractInto(&response)
	return &response, err
}

type ListRdsResult struct {
	commonResult
}

type ListRdsResponse struct {
	Instances  []RdsInstanceResponse `json:"instances"`
	TotalCount int                   `json:"total_count"`
}

type Tag struct {
	Key    string     `json:"key"`
	Value  string     `json:"value"`
}

type RdsInstanceResponse struct {
	Id                  string            `json:"id"`
	Name                string            `json:"name"`
	Status              string            `json:"status"`
	PrivateIps          []string          `json:"private_ips"`
	PublicIps           []string          `json:"public_ips"`
	Port                int               `json:"port"`
	Type                string            `json:"type"`
	Ha                  Ha                `json:"ha"`
	Region              string            `json:"region"`
	DataStore           Datastore         `json:"datastore"`
	Created             string            `json:"created"`
	Updated             string            `json:"updated"`
	DbUserName          string            `json:"db_user_name"`
	VpcId               string            `json:"vpc_id"`
	SubnetId            string            `json:"subnet_id"`
	SecurityGroupId     string            `json:"security_group_id"`
	FlavorRef           string            `json:"flavor_ref"`
	Volume              Volume            `json:"volume"`
	SwitchStrategy      string            `json:"switch_strategy"`
	BackupStrategy      BackupStrategy    `json:"backup_strategy"`
	MaintenanceWindow   string            `json:"maintenance_window"`
	Nodes               []Nodes           `json:"nodes"`
	RelatedInstance     []RelatedInstance `json:"related_instance"`
	DiskEncryptionId    string            `json:"disk_encryption_id"`
	EnterpriseProjectId string            `json:"enterprise_project_id"`
	TimeZone            string            `json:"time_zone"`
	ChargeInfo          ChargeInfoMode    `json:"charge_info"`
	Tags                []Tag             `json:"tags"`
}

type ChargeInfoMode struct {
	ChargeMode            string            `json:"charge_mode"`
}
type Nodes struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Role             string `json:"role"`
	Status           string `json:"status"`
	AvailabilityZone string `json:"availability_zone"`
}
type RelatedInstance struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type RdsPage struct {
	pagination.Offset
}

func (r RdsPage) IsEmpty() (bool, error) {
	data, err := ExtractRdsInstances(r)
	if err != nil {
		return false, err
	}
	return len(data.Instances) == 0, err
}

// ExtractCloudServers is a function that takes a ListResult and returns the services' information.
func ExtractRdsInstances(r pagination.Page) (ListRdsResponse, error) {
	var s ListRdsResponse
	err := (r.(RdsPage)).ExtractInto(&s)
	return s, err
}


type ErrorLogResult struct {
	gophercloud.Result
}
type ErrorLogResp struct {
	ErrorLogList []Errorlog `json:"error_log_list" `
	TotalRecord  int            `json:"total_record"`
}
type Errorlog struct {
	Time    string `json:"time" `
	Level   string `json:"level" `
	Content string `json:"content"`
}

type ErrorLogPage struct {
	pagination.Offset
}

// IsEmpty returns true if a ListResult contains no services.
func (r ErrorLogPage) IsEmpty() (bool, error) {
	data, err := ExtractErrorLog(r)
	if err != nil {
		return false, err
	}
	return len(data.ErrorLogList) == 0, err
}

func ExtractErrorLog(r pagination.Page) (ErrorLogResp, error) {
	var s ErrorLogResp
	err := (r.(ErrorLogPage)).ExtractInto(&s)
	return s, err
}

type SlowLogResp struct {
	Slowloglist []Slowloglist `json:"slow_log_list" `
	TotalRecord int           `json:"total_record"`
}
type Slowloglist struct {
	Count        string `json:"count"`
	Time         string `json:"time" `
	Locktime     string `json:"lock_time" `
	Rowssent     string `json:"rows_sent" `
	Rowsexamined string `json:"rows_examined"`
	Database     string `json:"database"`
	Users        string `json:"users" `
	QuerySample  string `json:"query_sample"`
	Type         string `json:"type" `
	StartTime    string `json:"start_time" `
	ClientIp     string `json:"client_ip" `
}

type SlowLogPage struct {
	pagination.Offset
}

// IsEmpty returns true if a ListResult contains no services.
func (r SlowLogPage) IsEmpty() (bool, error) {
	data, err := ExtractSlowLog(r)
	if err != nil {
		return false, err
	}
	return len(data.Slowloglist) == 0, err
}

// ExtractCloudServers is a function that takes a ListResult and returns the services' information.
func ExtractSlowLog(r pagination.Page) (SlowLogResp, error) {
	var s SlowLogResp
	err := (r.(SlowLogPage)).ExtractInto(&s)
	return s, err
}