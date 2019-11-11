package backups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type CreateBackupsOpts struct {
	InstanceId  string      `json:"instance_id"  required:"true"`
	Name        string      `json:"name" required:"true"`
	Description string      `json:"description,omitempty"`
	Databases   []Databases `json:"databases,omitempty"`
}
type Databases struct {
	Name string `json:"name" required:"true"`
}

type CreateBackupsBuilder interface {
	CreateBackupsMap() (map[string]interface{}, error)
}

func (opts CreateBackupsOpts) CreateBackupsMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateBackupsBuilder) (r CreateBackupsResult) {
	b, err := opts.CreateBackupsMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

type AutoBackupsPolicyOpts struct {
	BackupPolicy *BackupsPolicy `json:"backup_policy"  required:"true"`
}

type BackupsPolicy struct {
	KeepDays  *int   `json:"keep_days"  required:"true"`
	StartTime string `json:"start_time,omitempty"`
	Period    string `json:"period,omitempty"`
}

type AutoBackupsPolicyBuilder interface {
	AutoBackupsPolicyMap() (map[string]interface{}, error)
}

func (opts AutoBackupsPolicyOpts) AutoBackupsPolicyMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func UpdatePolicy(client *gophercloud.ServiceClient, opts AutoBackupsPolicyBuilder, instancesID string) (r ListBackupsPolicyResult) {
	b, err := opts.AutoBackupsPolicyMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(updatepolicyURL(client, instancesID), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}


func GetPolicy(client *gophercloud.ServiceClient, instancesID string) (r BackupsPolicyResult) {

	
	_, r.Err = client.Get(getpolicyURL(client, instancesID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

func Delete(client *gophercloud.ServiceClient, backupid string) (r DeletetBackupsResult) {

	_, r.Err = client.Delete(deleteURL(client, backupid), &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

type ListBackupsOpts struct {
	InstanceId string `q:"instance_id" `
	BackupId   string `q:"backup_id"`
	BackupType string `q:"backup_type"`
	Offset     int    `q:"offset"`
	Limit      int    `q:"limit"`
	BeginTime  int    `q:"begin_time"`
	EndTime    int    `q:"end_time"`
}

type ListBackupsBuilder interface {
	ToBackupsListQuery() (string, error)
}

func (opts ListBackupsOpts) ToBackupsListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}


func List(client *gophercloud.ServiceClient, opts ListBackupsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToBackupsListQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return BackupsPage{pagination.Offset{PageResult: r}}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}

type RestoreNewRdsOpts struct {
	Name                string          `json:"name"  required:"true"`
	Ha                  *Ha             `json:"ha,omitempty"`
	ConfigurationId     string          `json:"configuration_id,omitempty"`
	Port                string          `json:"port,omitempty"`
	Password            string          `json:"password" required:"true"`
	BackupStrategy      *BackupStrategy `json:"backup_strategy,omitempty"`
	EnterpriseProjectId string          `json:"enterprise_project_id,omitempty"`
	DiskEncryptionId    string          `json:"disk_encryption_id,omitempty"`
	FlavorRef           string          `json:"flavor_ref" required:"true"`
	Volume              *Volume         `json:"volume" required:"true"`
	AvailabilityZone    string          `json:"availability_zone" required:"true"`
	VpcId               string          `json:"vpc_id" required:"true"`
	SubnetId            string          `json:"subnet_id" required:"true"`
	SecurityGroupId     string          `json:"security_group_id" required:"true"`
	RestorePoint        *RestorePoint   `json:"restore_point" required:"true"`
	TimeZone            string          `json:"time_zone,omitempty"`
}

type Ha struct {
	Mode            string `json:"mode" required:"true"`
	ReplicationMode string `json:"replication_mode" required:"true"`
}

type BackupStrategy struct {
	StartTime string `json:"start_time" required:"true"`
	KeepDays  int    `json:"keep_days,omitempty"`
}
type Volume struct {
	Type string `json:"type" required:"true"`
	Size int    `json:"size" required:"true"`
}

type RestorePoint struct {
	InstanceId   string            `json:"instance_id" required:"true"`
	Type         string            `json:"type" required:"true"`
	BackupId     string            `json:"backup_id,omitempty"`
	RestoreTime  int               `json:"restore_time,omitempty" `
	DatabaseName map[string]string `json:"database_name,omitempty" `
}

type RestoreNewRdsBuilder interface {
	RestoreNewRdsMap() (map[string]interface{}, error)
}

func (opts RestoreNewRdsOpts) RestoreNewRdsMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Restore(client *gophercloud.ServiceClient, opts RestoreNewRdsBuilder) (r RestoreNewRdsResult) {
	b, err := opts.RestoreNewRdsMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(restoreURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return r
}

type ListBackupFilesOpts struct {
	BackupId string `q:"backup_id"`
}

type BackupFilesBuilder interface {
	ToBackupFilesListQuery() (string, error)
}

func (opts ListBackupFilesOpts) ToBackupFilesListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func ListFiles(client *gophercloud.ServiceClient, opts BackupFilesBuilder) pagination.Pager {
	url := listfilesURL(client)
	if opts != nil {
		query, err := opts.ToBackupFilesListQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return BackupsFilesPage{pagination.Offset{PageResult: r}}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}

type ListRestoreTimeOpts struct {
	Date string `q:"date"`
}

type ListRestoreTimeBuilder interface {
	ToRestoreTimeListQuery() (string, error)
}

func (opts ListRestoreTimeOpts) ToRestoreTimeListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func ListRestoreTime(client *gophercloud.ServiceClient, opts ListRestoreTimeBuilder, instanceId string) pagination.Pager {
	url := getrestoretimeURL(client, instanceId)
	if opts != nil {
		query, err := opts.ToRestoreTimeListQuery()

		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return RestoreTimePage{pagination.Offset{PageResult: r}}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}

type RecoveryOpts struct {
	Source Source `json:"source"  required:"true"`
	Target Target `json:"target" required:"true"`
}
type Source struct {
	InstanceId   string            `json:"instance_id"`
	Type         string            `json:"type,omitempty"`
	BackupId     string            `json:"backup_id,omitempty"`
	RestoreTime  int               `json:"restore_time,omitempty" `
	DatabaseName map[string]string `json:"database_name,omitempty"`
}
type Target struct {
	InstanceId string `json:"instance_id" required:"true"`
}

type RecoveryBuilder interface {
	ToRecoveryMap() (map[string]interface{}, error)
}

func (opts RecoveryOpts) ToRecoveryMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Recovery(client *gophercloud.ServiceClient, opts RecoveryBuilder) (r RecoveryResult) {
	b, err := opts.ToRecoveryMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(recoveryURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
