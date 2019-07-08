package volumes

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVolumeCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating a Volume. This object is passed to
// the volumes.Create function. For more information about these parameters,
// see the Volume object.
type CreateOpts struct {
	// The size of the volume, in GB
	Size int `json:"size" required:"true"`
	// The availability zone
	AvailabilityZone string `json:"availability_zone" required:"true"`
	// ConsistencyGroupID is the ID of a consistency group
	ConsistencyGroupID string `json:"consistencygroup_id,omitempty"`
	// The volume description
	Description string `json:"description,omitempty"`
	// One or more metadata key and value pairs to associate with the volume
	Metadata map[string]string `json:"metadata,omitempty"`
	// The volume name
	Name string `json:"name,omitempty"`
	// The ID of the existing volume snapshot
	SnapshotID string `json:"snapshot_id,omitempty"`
	// SourceReplica is a UUID of an existing volume to replicate with
	SourceReplica string `json:"source_replica,omitempty"`
	// The ID of the existing volume
	SourceVolID string `json:"source_volid,omitempty"`
	// The ID of the image from which you want to create the volume.
	// Required to create a bootable volume.
	ImageID string `json:"imageRef,omitempty"`
	// The associated volume type
	VolumeType string `json:"volume_type,omitempty"`

	//The scheduling parameter currently supports the dedicated_storage_id field, indicating that the cloud disk is created in the DSS storage pool.
	SchedulerHints map[string]string `json:"OS-SCH-HNT:scheduler_hints,omitempty"`

	//Share the cloud drive flag. The default is false.
	Multiattach *bool `json:"multiattach,omitempty"`
}

// ToVolumeCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToVolumeCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volume")
}

// Create will create a new Volume based on the values in CreateOpts. To extract
// the Volume object from the response, call the Extract method on the
// CreateResult.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVolumeCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// DeleteOptsBuilder allows extensions to add additional parameters to the Delete
// request.
type DeleteOptsBuilder interface {
	ToVolumeDeleteQuery() (string, error)
}

// DeleteOpts holds options for delete Volumes. It is passed to the volumes.Delete
// function.
type DeleteOpts struct {
	// Delete all snapshots associated with the cloud drive. The default value is false.
	Cascade *bool `q:"cascade,omitempty"`
}

// ToVolumeDeleteQuery formats a DeleteOpts into a query string.
func (opts DeleteOpts) ToVolumeDeleteQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// Delete will delete the existing Volume with the provided ID.
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}

// Delete will delete the existing Volume with the provided ID,Delete all snapshots associated with the Volume.
func DeleteCascade(client *gophercloud.ServiceClient, id string, opts DeleteOptsBuilder) (r DeleteResult) {
	url := deleteURL(client, id)
	if opts != nil {
		query, err := opts.ToVolumeDeleteQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	_, r.Err = client.Delete(url, nil)
	return
}

// Get retrieves the Volume with the provided ID. To extract the Volume object
// from the response, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToVolumeListQuery() (string, error)
}

// ListOpts holds options for listing Volumes. It is passed to the volumes.List
// function.
type ListOpts struct {
	// AllTenants will retrieve volumes of all tenants/projects.
	AllTenants bool `q:"all_tenants"`

	// Metadata will filter results based on specified metadata.
	Metadata map[string]string `q:"metadata"`

	// Name will filter by the specified volume name.
	Name string `q:"name"`

	Maker string `q:"maker"`

	SortDir string `q:"sort_dir"`

	SortKey string `q:"sort_key"`

	AvailabilityZone string `q:"availability_zone"`

	// Status will filter by the specified status.
	Status string `q:"status"`

	// TenantID will filter by a specific tenant/project ID.
	// Setting AllTenants is required for this.
	TenantID string `q:"project_id"`
}

// ToVolumeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns Volumes optionally limited by the conditions provided in ListOpts.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToVolumeListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VolumePage{pagination.SinglePageBase(r)}
	})
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVolumeUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contain options for updating an existing Volume. This object is passed
// to the volumes.Update function. For more information about the parameters, see
// the Volume object.
type UpdateOpts struct {
	// The volume name
	Name string `json:"name,omitempty"`
	// The volume description
	Description string `json:"description,omitempty"`
	// Metadata will filter results based on specified metadata.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The volume name show for users
	DisplayName string `json:"display_name,omitempty"`
	// The volume description show for users
	DisplayDescription string `json:"display_description,omitempty"`
}

// ToVolumeUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToVolumeUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volume")
}

// Update will update the Volume with provided information. To extract the updated
// Volume from the response, call the Extract method on the UpdateResult.
func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVolumeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// ListBriefOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListBriefOptsBuilder interface {
	ToVolumeListBriefQuery() (string, error)
}

// ListBriefOpts holds options for listing Volumes. It is passed to the volumes.List
// function.
type ListBriefOpts struct {
	// AllTenants will retrieve volumes of all tenants/projects.
	//AllTenants bool `q:"all_tenants"`

	// Metadata will filter results based on specified metadata.
	Metadata map[string]string `q:"metadata"`

	// Name will filter by the specified volume name.
	Name string `q:"name"`

	Marker string `q:"marker"`

	SortDir string `q:"sort_dir"`

	SortKey string `q:"sort_key"`

	AvailabilityZone string `q:"availability_zone"`

	// Status will filter by the specified status.
	Status string `q:"status"`
	// Number of resources returned on each page.Value range:
	// 0-500 Commonly used values are 10, 20, and 50.
	Limit int `q:"limit"`

	//Used when paginating snapshots, used in conjunction with limit.
	Offset int `q:"offset"`

	// TenantID will filter by a specific tenant/project ID.
	// Setting AllTenants is required for this.
	//TenantID string `q:"project_id"`
}

// ToVolumeListBriefQuery formats a ListBriefOpts into a query string.
func (opts ListBriefOpts) ToVolumeListBriefQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns Volumes optionally limited by the conditions provided in ListOpts.
func ListBrief(client *gophercloud.ServiceClient, opts ListBriefOptsBuilder) pagination.Pager {
	url := listBriefURL(client)
	if opts != nil {
		query, err := opts.ToVolumeListBriefQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VolumeListPage{pagination.SinglePageBase(r)}
	})
}

// IDFromName is a convienience function that returns a server's ID given its name.
func IDFromName(client *gophercloud.ServiceClient, name string) (string, error) {
	count := 0
	id := ""
	pages, err := List(client, nil).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractVolumes(pages)
	if err != nil {
		return "", err
	}

	for _, s := range all {
		if s.Name == name {
			count++
			id = s.ID
		}
	}

	switch count {
	case 0:
		//return "", gophercloud.ErrResourceNotFound{Name: name, ResourceType: "volume"}

		message := fmt.Sprintf(gophercloud.CE_ResourceNotFoundMessage, "volume", name)
		err := gophercloud.NewSystemCommonError(gophercloud.CE_ResourceNotFoundCode, message)
		return "", err
	case 1:
		return id, nil
	default:
		//return "", gophercloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "volume"}

		message := fmt.Sprintf(gophercloud.CE_MultipleResourcesFoundMessage, count, "volume", name)
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MultipleResourcesFoundCode, message)
		return "", err
	}
}

// GetQuotaSet allows extensions to query project quota data via projectId
func GetQuotaSet(client *gophercloud.ServiceClient, projectId string) (*QuotaSetInfo, error) {
	var r GetResult
	_, r.Err = client.Get(getQuotaSetURL(client, projectId), &r.Body, nil)
	return r.ExtractQuotaSet()
}

// MetadataOptsBuilder allows extensions to add additional parameters to
// the meatadata requests.
type MetadataOptsBuilder interface {
	ToVolumeMetadataMap() (map[string]interface{}, error)
}

// MetadataOpts contain options for creating or updating an existing Volume. This
// object is passed to the volumes create and update function. For more information
// about the parameters, see the Volume object.
type MetadataOpts struct {
	Metadata map[string]string `json:"metadata,omitempty"`
}

// ToSnapshotMetadataMap assembles a request body based on the contents of
// an MetadataOpts.
func (opts MetadataOpts) ToVolumeMetadataMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type MetaOptsBuilder interface {
	ToVolumeMetaMap() (map[string]interface{}, error)
}

type MetaOpts struct {
	Meta map[string]string `json:"meta,omitempty"`
}

func (opts MetaOpts) ToVolumeMetaMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// CreateMetadata create metadata for Volume.
func CreateMetadata(client *gophercloud.ServiceClient, id string, opts MetadataOptsBuilder) (r MetadataResult) {
	b, err := opts.ToVolumeMetadataMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(metadataURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// GetMetadata returns exist metadata of Volume.
func GetMetadata(client *gophercloud.ServiceClient, id string) (r MetadataResult) {
	_, r.Err = client.Get(metadataURL(client, id), &r.Body, nil)
	return
}

// UpdateMetadata will update metadata according to request map.
func UpdateMetadata(client *gophercloud.ServiceClient, id string, opts MetadataOptsBuilder) (r MetadataResult) {
	b, err := opts.ToVolumeMetadataMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(metadataURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// GetMetadataKey return specific key value in metadata.
func GetMetadataKey(client *gophercloud.ServiceClient, id, key string) (r MetadataResult) {
	_, r.Err = client.Get(metadataKeyURL(client, id, key), &r.Body, nil)
	return
}

// UpdateMetadataKey update specific key to the given map key value.
func UpdateMetadataKey(client *gophercloud.ServiceClient, id, key string, opts MetaOptsBuilder) (r MetadataResult) {
	b, err := opts.ToVolumeMetaMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(metadataKeyURL(client, id, key), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// DeleteMetadataKey delete specific key in metadata
func DeleteMetadataKey(client *gophercloud.ServiceClient, id, key string) (r DeleteMetadataKeyResult) {
	_, r.Err = client.Delete(metadataKeyURL(client, id, key), &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// ExtendSizeOptsBuilder allows extensions to add additional parameters to the
// ExtendSize request.
type ExtendSizeOptsBuilder interface {
	ToVolumeExtendSizeMap() (map[string]interface{}, error)
}

// ExtendSizeOpts contains options for extending the size of an existing Volume.
// This object is passed to the volumes.ExtendSize function.
type ExtendSizeOpts struct {
	// NewSize is the new size of the volume, in GB.
	NewSize int `json:"new_size" required:"true"`
}

// ToVolumeExtendSizeMap assembles a request body based on the contents of an
// ExtendSizeOpts.
func (opts ExtendSizeOpts) ToVolumeExtendSizeMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "os-extend")
}

// ExtendSize will extend the size of the volume based on the provided information.
// This operation does not return a response body.
func ExtendSize(client *gophercloud.ServiceClient, id string, opts ExtendSizeOptsBuilder) (r ExtendSizeResult) {
	b, err := opts.ToVolumeExtendSizeMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, id), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// SetBootableOptsBuilder allows extensions to add additional parameters to the
// SetBootable request.
type SetBootableOptsBuilder interface {
	ToVolumeSetBootableMap() (map[string]interface{}, error)
}

// SetBootableOpts contains options for setting bootable flag of an existing Volume.
// This object is passed to the volumes.SetBootable function.
type SetBootableOpts struct {
	// Bootable is bool of true or false
	Bootable bool `json:"bootable" required:"true"`
}

// ToVolumeSetBootableMap assembles a request body based on the contents of an
// SetBootableOpts.
func (opts SetBootableOpts) ToVolumeSetBootableMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "os-set_bootable")
}

// SetBootable will set bootable flag of the volume based on the provided information.
// This operation does not return a response body.
func SetBootable(client *gophercloud.ServiceClient, id string, opts SetBootableOptsBuilder) (r SetBootableResult) {
	b, err := opts.ToVolumeSetBootableMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, id), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// SetReadOnlyOptsBuilder allows extensions to add additional parameters to the
// SetReadOnly request.
type SetReadOnlyOptsBuilder interface {
	ToVolumeSetReadOnlyMap() (map[string]interface{}, error)
}

// SetReadOnlyOpts contains options for setting readonly flag of an existing Volume.
// This object is passed to the volumes.SetReadOnly function.
type SetReadOnlyOpts struct {
	// ReadOnly is bool of true or false
	ReadOnly bool `json:"readonly" required:"true"`
}

// ToVolumeSetReadOnlyMap assembles a request body based on the contents of an
// SetReadOnlyOpts.
func (opts SetReadOnlyOpts) ToVolumeSetReadOnlyMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "os-update_readonly_flag")
}

// SetReadOnly will set readonly flag of the volume based on the provided information.
// This operation does not return a response body.
func SetReadOnly(client *gophercloud.ServiceClient, id string, opts SetReadOnlyOptsBuilder) (r SetReadOnlyResult) {
	b, err := opts.ToVolumeSetReadOnlyMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, id), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

type ExportVolumesOptsBuilder interface {
	ToExportVolumesMap() (map[string]interface{}, error)
}

type ExportVolumesOpts struct {
	ImageName       string `json:"image_name" required:"true"`
	Force           bool   `json:"force,omitempty"`
	ContainerFormat string `json:"container_format,omitempty"`
	DiskFormat      string `json:"disk_format,omitempty"`
	OsType          string `json:"__os_type,omitempty"`
}

func (opts ExportVolumesOpts) ToExportVolumesMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "os-volume_upload_image")
}

// ExportVolumes will Export volume as image
func ExportVolumes(client *gophercloud.ServiceClient, id string, opts ExportVolumesOptsBuilder) (r ExportVolumesResult) {
	b, err := opts.ToExportVolumesMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
