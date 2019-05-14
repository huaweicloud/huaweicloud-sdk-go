package configures

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type Disk struct {
	// Specifies the disk size. The unit is GB.The system disk size
	// ranges from 40 to 32768 and must be greater than or equal to the minimum size
	// (min_disk value) of the system disk specified in the image.The data disk size ranges
	// from 10 to 32768.
	Size int `json:"size" required:"true"`

	// Specifies the ECS system disk type, which must be the same as
	// the disk type available in the system.Enumerated values of the disk type:SATA: common
	// I/O disk type,SAS: high I/O disk type,SSD: ultra-high I/O disk type,co-pl: high I/O
	// (performance-optimized I) disk type,uh-l1: ultra-high I/O (latency-optimized) disk
	// type.NOTE:For HANA and HL1 ECSs, use co-p1 or uh-l1 disks. For other ECSs, do not use
	// co-p1 or uh-l1 disks.
	VolumeType string `json:"volume_type" required:"true"`

	// Specifies whether the disk is a system disk or a data disk.
	// DATA indicates a data disk. SYS indicates a system disk.
	DiskType string `json:"disk_type" required:"true"`

	// Specifies a DSS device ID for creating an ECS disk.NOTE:Specify
	// DSS devices for all disks in an AS configuration or not. If DSS devices are
	// specified, all the data stores must belong to the same AZ, and the disk types
	// supported by a DSS device for a disk must be the same as the volume_type value.
	DedicatedStorageId string `json:"dedicated_storage_id,omitempty"`

	// Specifies a data disk image for exporting data.
	DataDiskImageId string `json:"data_disk_image_id,omitempty"`

	// Specifies the disk backup snapshot ID for restoring the system
	// disk and data disk at the ECS level when the ECS-level image is used.NOTE:You can
	// obtain the disk backup snapshot ID by using the mirrored ECS-level backup ID to query
	// backup details in CSBS. Regarding snapshot_id, each disk in an AS configuration must
	// correspond to a disk backup at the ECS-level.
	SnapshotId string `json:"snapshot_id,omitempty"`
}

type Personality struct {
	// Specifies the path of the injected file.For Linux OSs, specify
	// the path, for example, /etc/foo.txt, for storing the injected file.For Windows OSs,
	// the injected file is automatically stored in the root directory of disk C. You only
	// need to specify the file name, for example foo. The file name can contain only
	// letters and digits.
	Path string `json:"path" required:"true"`

	// Specifies the content of the injected file, which must be
	// encoded with base64.
	Content string `json:"content" required:"true"`
}

type PublicIP struct {
	// Specifies the configuration parameter for creating an EIP that
	// will be automatically assigned to the ECS.
	EIP EIP `json:"eip" required:"true"`
}

type EIP struct {
	// Specifies the IP address type.Enumerated values of the IP
	// address type:5_bgp: indicates the dynamic BGP.5_lxbgp: BGP,5_telcom: indicates China
	// Telecom.5_union: indicates China Unicom.
	IpType string `json:"ip_type" required:"true"`

	// Specifies the bandwidth.
	Bandwidth BandwidthInfo `json:"bandwidth" required:"true"`
}

type BandwidthInfo struct {
	ID string `json:"id,omitempty"`
	// Specifies the bandwidth (Mbit/s). The value range is 1 to
	// 100.
	Size int `json:"size,omitempty"`

	// Specifies the bandwidth sharing type.Enumerated value: PER
	// (indicates exclusive bandwidth).Only exclusive bandwidth is available.
	ShareType string `json:"share_type" required:"true"`

	// Specifies the bandwidth charging mode.If the field value is
	// bandwidth, the ECS service is charged by bandwidth.If the field value is traffic, the
	// ECS service is charged by traffic.If the field value is others, the ECS creation will
	// fail.
	ChargingMode string `json:"charging_mode,omitempty"`
}

type CreateInstanceConfig struct {
	// This field is reserved.
	//InstanceName string `json:"instance_name,omitempty"`

	// Specifies the ECS ID. When using the existing ECS
	// specifications as the template to create AS configurations, specify this parameter.
	// In this case, flavorRef, imageRef, and disk fields do not take effect.If the
	// instance_id field is not specified, flavorRef, imageRef, and disk fields are
	// mandatory.
	InstanceId string `json:"instance_id,omitempty"`

	// Specifies the ECS specifications ID, which defines the
	// specifications of the CPU and memory for the ECS. You can obtain its value from the
	// API used to query specifications and expansion details about ECSs. For details, see
	// section Querying Specifications and Expansion Details About ECSs in the Elastic Cloud
	// Server API Reference.
	FlavorRef string `json:"flavorRef,omitempty"`

	// Specifies the image ID. It is the same as image_id. You can
	// obtain its value from the API used to query IMS images. For details, see section
	// Querying Images in the Image Management Service API Reference.
	ImageRef string `json:"imageRef,omitempty"`

	// Specifies the disk group information. System disks are
	// mandatory and data disks are optional.
	Disk []Disk `json:"disk,omitempty"`

	// This field is reserved.
	//AdminPass string `json:"adminPass,omitempty"`

	// Specifies the name of the SSH key pair used to log in to the
	// ECS.
	KeyName string `json:"key_name,omitempty"`

	// Specifies information about the injected file. Only text files
	// can be injected. A maximum of five files can be injected at a time and the maximum
	// size of each file is 1 KB.
	Personality []Personality `json:"personality,omitempty"`

	// Specifies the cloud-init user data.Text, text files, and gzip
	// files can be injected. The file content must be encoded with Base64, and the maximum
	// allowed file size is 32 KB.
	UserData string `json:"user_data,omitempty"`

	// Specifies the EIP of the ECS. The EIP can be configured in the
	// following two ways:Not configured (delete this field),Assigned automatically
	PublicIP       *PublicIP       `json:"public_ip,omitempty"`
	SecurityGroups []SecurityGroup `json:"security_groups,omitempty"`

	ServerGroupID string `json:"server_group_id,omitempty"`

	Tenancy string `json:"tenancy,omitempty"`

	DedicatedHostID string `json:"dedicated_host_id,omitempty"`

	// Specifies the metadata of ECSs to be created.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type SecurityGroup struct {
	ID string `json:"id" required:"true"`
}

type CreateOpts struct {
	// Specifies the AS configuration name. The name can contain
	// letters, digits, underscores (_), and hyphens (-), and must be between 1 and 64
	// characters in length.

	ScalingConfigurationName string `json:"scaling_configuration_name" required:"true"`

	// Specifies the information about instance configurations.
	InstanceConfig CreateInstanceConfig `json:"instance_config" required:"true"`
}

type CreateOptsBuilder interface {
	ToConfiguresCreateMap() (map[string]interface{}, error)
}

func (opts CreateOpts) ToConfiguresCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToConfiguresCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(CreateURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func Delete(client *gophercloud.ServiceClient, scalingConfigurationId string) (r DeleteResult) {
	url := DeleteURL(client, scalingConfigurationId)
	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		JSONResponse: nil,
		OkCodes:      []int{204},
	})
	return
}

type DeleteWithBatchOpts struct {
	// Specifies the AS configuration ID.
	ScalingConfigurationId []string `json:"scaling_configuration_id" required:"true"`
}

type DeleteWithBatchOptsBuilder interface {
	ToConfiguresDeleteWithBatchMap() (map[string]interface{}, error)
}

func (opts DeleteWithBatchOpts) ToConfiguresDeleteWithBatchMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func DeleteWithBatch(client *gophercloud.ServiceClient, opts DeleteWithBatchOptsBuilder) (r DeleteWithBatchResult) {
	b, err := opts.ToConfiguresDeleteWithBatchMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(DeleteWithBatchURL(client), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

func Get(client *gophercloud.ServiceClient, scalingConfigurationId string) (r GetResult) {
	url := GetURL(client, scalingConfigurationId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ListOpts struct {
	// Specifies the AS configuration name.
	ScalingConfigurationName string `q:"scaling_configuration_name"`

	// Specifies the image ID. It is same as imageRef.
	ImageId string `q:"image_id"`

	// Specifies the start line number. The default value is 0.
	StartNumber int `q:"start_number"`

	// Specifies the number of query records. The default value is
	// 20.
	Limit int `q:"limit"`
}

type ListOptsBuilder interface {
	ToListQuery() (string, error)
}

func (opts ListOpts) ToListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := ListURL(client)
	if opts != nil {
		query, err := opts.ToListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		p := ConfigPage{pagination.NumberPageBase{PageResult: r}}
		p.NumberPageBase.Owner = p
		return p
	})
}
