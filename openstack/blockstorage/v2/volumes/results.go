package volumes

import (
	"encoding/json"
	"io"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type Attachment struct {
	AttachedAt   time.Time `json:"-"`
	AttachmentID string    `json:"attachment_id"`
	Device       string    `json:"device"`
	HostName     string    `json:"host_name"`
	ID           string    `json:"id"`
	ServerID     string    `json:"server_id"`
	VolumeID     string    `json:"volume_id"`
}

func (r *Attachment) UnmarshalJSON(b []byte) error {
	type tmp Attachment
	var s struct {
		tmp
		AttachedAt gophercloud.JSONRFC3339MilliNoZ `json:"attached_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Attachment(s.tmp)

	r.AttachedAt = time.Time(s.AttachedAt)

	return err
}

// Volume contains all the information associated with an OpenStack Volume.
type Volume struct {
	// Unique identifier for the volume.
	ID string `json:"id"`
	// Current status of the volume.
	Status string `json:"status"`
	// Size of the volume in GB.
	Size int `json:"size"`
	// AvailabilityZone is which availability zone the volume is in.
	AvailabilityZone string `json:"availability_zone"`
	// The date when this volume was created.
	CreatedAt time.Time `json:"-"`
	// The date when this volume was last updated
	UpdatedAt time.Time `json:"-"`
	// Instances onto which the volume is attached.
	Attachments []Attachment `json:"attachments"`
	// Human-readable display name for the volume.
	Name string `json:"name"`
	// Human-readable description for the volume.
	Description string `json:"description"`
	// The type of volume to create, either SATA or SSD.
	VolumeType string `json:"volume_type"`
	// The ID of the snapshot from which the volume was created
	SnapshotID string `json:"snapshot_id"`
	// The ID of another block storage volume from which the current volume was created
	SourceVolID string `json:"source_volid"`
	// Arbitrary key-value pairs defined by the user.
	Metadata map[string]string `json:"metadata"`
	// UserID is the id of the user who created the volume.
	UserID string `json:"user_id"`
	// Indicates whether this is a bootable volume.
	Bootable string `json:"bootable"`
	// Encrypted denotes if the volume is encrypted.
	Encrypted bool `json:"encrypted"`
	// ReplicationStatus is the status of replication.
	ReplicationStatus string `json:"replication_status"`
	// ConsistencyGroupID is the consistency group ID.
	ConsistencyGroupID string `json:"consistencygroup_id"`
	// Multiattach denotes if the volume is multi-attach capable.
	Multiattach bool `json:"multiattach"`

	//Cloud hard disk uri self-description information.
	Links []map[string]string `json:"links"`

	//Whether it is a shared cloud drive.
	Shareable bool `json:"shareable"`
	//Volume image metadata
	VolumeImageMetadata map[string]string `json:"volume_image_metadata"`

	//The tenant ID to which the cloud drive belongs.
	TenantAttr string `json:"os-vol-tenant-attr:tenant_id"`

	//The host name to which the cloud drive belongs.
	HostAttr string `json:"os-vol-host-attr:host"`
	//Reserved attribute
	RepAttrDriverData string `json:"os-volume-replication:driver_data"`
	//Reserved attribute
	RepAttrExtendedStatus string `json:"os-volume-replication:extended_status"`
	//Reserved attribute
	MigAttrStat string `json:"os-vol-mig-status-attr:migstat"`
	//Reserved attribute
	MigAttrNameID string `json:"os-vol-mig-status-attr:name_id"`
}

func (r *Volume) UnmarshalJSON(b []byte) error {
	type tmp Volume
	var s struct {
		tmp
		CreatedAt gophercloud.JSONRFC3339MilliNoZ `json:"created_at"`
		UpdatedAt gophercloud.JSONRFC3339MilliNoZ `json:"updated_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Volume(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	return err
}

// VolumePage is a pagination.pager that is returned from a call to the List function.
type VolumePage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ListResult contains no Volumes.
func (r VolumePage) IsEmpty() (bool, error) {
	volumes, err := ExtractVolumes(r)
	return len(volumes) == 0, err
}

// ExtractVolumes extracts and returns Volumes. It is used while iterating over a volumes.List call.
func ExtractVolumes(r pagination.Page) ([]Volume, error) {
	var s []Volume
	err := ExtractVolumesInto(r, &s)
	return s, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract will get the Volume object out of the commonResult object.
func (r commonResult) Extract() (*Volume, error) {
	var s Volume
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "volume")
}

func ExtractVolumesInto(r pagination.Page, v interface{}) error {
	return r.(VolumePage).Result.ExtractIntoSlicePtr(v, "volumes")
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// UpdateResult contains the response body and error from an Update request.
type UpdateResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	gophercloud.ErrResult
}

/*************** 自研 ********************/
type QuotaSetInfo struct {
	//查询请求返回的配额信息
	QuoSet QuotaSet `json:"quota_set"`
}

//查询请求返回的配额信息
type QuotaSet struct {
	//租户id
	Id string `json:"id"`

	//云硬盘数量，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）
	Volumes map[string]int `json:"volumes"`

	//快照数量，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）
	Snapshots map[string]int `json:"snapshots"`

	//总大小（快照+云硬盘），单位为GB，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）
	Gigabytes map[string]int `json:"gigabytes"`

	//为某个volume_type预留的云硬盘个数，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）
	VolumesType map[string]int `json:"volumes_TYPE"`

	//为某个volume_type预留快照个数，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）
	SnapshotsType map[string]int `json:"snapshots_TYPE"`

	//为某个volume_type预留的size大小，单位为GB，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）
	GigabytesType map[string]int `json:"gigabytes_TYPE"`

	//备份个数，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）
	Backups map[string]int `json:"backups"`

	//备份大小，单位为GB，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）
	BackupGigabytes map[string]int `json:"backup_gigabytes"`

	//出现错误时，返回的错误码，具体含义参考下面的返回值列表
	Code string `json:"code"`

	//出现错误时，返回的错误消息
	Message string `json:"message"`
}

func (r commonResult) ExtractQuotaSet() (*QuotaSetInfo, error) {
	var qs *QuotaSetInfo
	err := r.ExtractIntoQuotaSet(&qs)
	return qs, err
}

func (r commonResult) ExtractIntoQuotaSet(to interface{}) error {
	if r.Err != nil {
		return r.Err
	}

	if reader, ok := r.Body.(io.Reader); ok {
		if readCloser, ok := reader.(io.Closer); ok {
			defer readCloser.Close()
		}
		return json.NewDecoder(reader).Decode(to)
	}

	b, err := json.Marshal(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, to)

	return err
}
