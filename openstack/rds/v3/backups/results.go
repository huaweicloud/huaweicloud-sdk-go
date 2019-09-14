package backups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)


type commonResult struct {
	gophercloud.Result
}
type ListBackupsPolicyResult struct {
	gophercloud.ErrResult
}

type DeletetBackupsResult struct {
	gophercloud.ErrResult
}

type CreateBackupsResult struct {
	commonResult
}

type CreateBackupsResp struct {
	Backup `json:"backup"`
}
type Backup struct {
	Id          string `json:"id"`
	Instanceid  string `json:"instance_id" `
	Name        string `json:"name" `
	Description string `json:"description" `
	Databases   string `json:"databases" `
	Begintime   string `json:"begin_time"`
	Status      string `json:"status" `
	Type        string `json:"type" `
}

func (r CreateBackupsResult) Extract() (*CreateBackupsResp, error) {
	var response CreateBackupsResp
	err := r.ExtractInto(&response)
	return &response, err
}


type ListAutoBackupsPolicyResp struct {
	ListBackupsPolicy `json:"backup_policy"`
}

type ListBackupsPolicy struct {
	KeepDays  int    `json:"keep_days"`
	StartTime string `json:"start_time"`
	Period    string `json:"period"`
}
type BackupsPolicyResult struct {
	gophercloud.Result
}


func (r BackupsPolicyResult) Extract() (ListAutoBackupsPolicyResp, error) {
	var s ListAutoBackupsPolicyResp
	err := r.ExtractInto(&s)
	return s, err
}

type ListBackupsResp struct {
	Backups    []BackupsResp `json:"backups"`
	TotalCount int               `json:"total_count"`
}

type BackupsResp struct {
	Id         string    `json:"id" `
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Size       int64     `json:"size"`
	Status     string    `json:"status"`
	BeginTime  string    `json:"begin_time"`
	EndTime    string    `json:"end_time"`
	Datastore  Datastore `json:"datastore"`
	Databases  Databases `json:"databases"`
	InstanceId string    `json:"instance_id"`
}
type Datastore struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type BackupsPage struct {
	pagination.Offset
}


func (r BackupsPage) IsEmpty() (bool, error) {
	data, err := ExtractBackups(r)
	if err != nil {
		return false, err
	}
	return len(data.Backups) == 0, err
}


func ExtractBackups(r pagination.Page) (ListBackupsResp, error) {
	var s ListBackupsResp
	err := (r.(BackupsPage)).ExtractInto(&s)
	return s, err
}

type RestoreNewRdsInstanceResp struct {

	Instance RestoreInstanceResp `json:"instance" `
	JobId    string             `json:"job_id" `
}
type RestoreNewRdsResult struct {
	commonResult
}
type RestoreInstanceResp struct {
	ID                   string         `json:"id" `
	Name                 string         `json:"name"`
	Status               string         `json:"status"`
	Datastore            Datastore      `json:"datastore"`
	Ha                   Ha             `json:"ha"`
	ConfigurationId      string         `json:"configuration_id"`
	Port                 string         `json:"port"`
	Password             string         `json:"password"`
	BackupStrategy       BackupStrategy `json:"backup_strategy"`
	EnterpriseProjectTag string         `json:"enterprise_project_tag"`
	DiskEncryptionId     string         `json:"disk_encryption_id"`
	FlavorRef            string         `json:"flavor_ref"`
	Volume               Volume         `json:"volume"`
	Region               string         `json:"region"`
	AvailabilityZone     string         `json:"availability_zone"`
	VpcId                string         `json:"vpc_id"`
	SubnetId             string         `json:"subnet_id"`
	SecurityGroupid      string         `json:"security_group_id" `
}

func (r RestoreNewRdsResult) Extract() (*RestoreNewRdsInstanceResp, error) {
	var response RestoreNewRdsInstanceResp
	err := r.ExtractInto(&response)
	return &response, err
}

type BackupsFilesResp struct {
	FilesList []BackupsFile `json:"files"`
	Bucket    string         `json:"bucket"`
}

type BackupsFile struct {
	DownloadLink    string `json:"download_link"`
	Name            string `json:"name"`
	LinkExpiredLink string `json:"link_expired_time"`
	Size            int64  `json:"size"`
}

type BackupsFilesPage struct {
	pagination.Offset
}


func (r BackupsFilesPage) IsEmpty() (bool, error) {
	data, err := ExtractBackupsFiles(r)
	if err != nil {
		return false, err
	}
	return len(data.FilesList) == 0, err
}

func ExtractBackupsFiles(r pagination.Page) (BackupsFilesResp, error) {
	var s BackupsFilesResp
	err := (r.(BackupsFilesPage)).ExtractInto(&s)
	return s, err
}

type RestoreTimeResp struct {
	RestoreTimeList []RestoreTime `json:"restore_time"`
}
type RestoreTime struct {
	StartTime int `json:"start_time"`
	EndTime   int `json:"end_time"`
}
type RestoreTimePage struct {
	pagination.Offset
}

func (r RestoreTimePage) IsEmpty() (bool, error) {
	data, err := ExtractRestoreTime(r)
	if err != nil {
		return false, err
	}
	return len(data.RestoreTimeList) == 0, err
}

func ExtractRestoreTime(r pagination.Page) (RestoreTimeResp, error) {
	var s RestoreTimeResp
	err := (r.(RestoreTimePage)).ExtractInto(&s)
	return s, err
}

type RecoveryResult struct {
	commonResult
}

type RecoveryResp struct {
	JobId string `json:"job_id"`
}

func (r RecoveryResult) Extract() (*RecoveryResp, error) {
	var response RecoveryResp
	err := r.ExtractInto(&response)
	return &response, err
}
