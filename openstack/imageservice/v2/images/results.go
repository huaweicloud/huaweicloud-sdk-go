package images

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/internal"
	"github.com/gophercloud/gophercloud/pagination"
)

// Image represents an image found in the OpenStack Image service.
type Image struct {
	// ID is the image UUID.
	ID string `json:"id"`

	// Name is the human-readable display name for the image.
	Name string `json:"name"`

	// Status is the image status. It can be "queued" or "active"
	// See imageservice/v2/images/type.go
	Status ImageStatus `json:"status"`

	// Tags is a list of image tags. Tags are arbitrarily defined strings
	// attached to an image.
	Tags []string `json:"tags"`

	// ContainerFormat is the format of the container.
	// Valid values are ami, ari, aki, bare, and ovf.
	ContainerFormat string `json:"container_format"`

	// DiskFormat is the format of the disk.
	// If set, valid values are ami, ari, aki, vhd, vmdk, raw, qcow2, vdi,
	// and iso.
	DiskFormat string `json:"disk_format"`

	// MinDiskGigabytes is the amount of disk space in GB that is required to
	// boot the image.
	MinDiskGigabytes int `json:"min_disk"`

	// MinRAMMegabytes [optional] is the amount of RAM in MB that is required to
	// boot the image.
	MinRAMMegabytes int `json:"min_ram"`

	// Owner is the tenant ID the image belongs to.
	Owner string `json:"owner"`

	// Protected is whether the image is deletable or not.
	Protected bool `json:"protected"`

	// Visibility defines who can see/use the image.
	Visibility ImageVisibility `json:"visibility"`

	// Checksum is the checksum of the data that's associated with the image.
	Checksum string `json:"checksum"`

	// SizeBytes is the size of the data that's associated with the image.
	SizeBytes int64 `json:"size"`

	// Metadata is a set of metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags.
	// See http://docs.openstack.org/developer/glance/metadefs-concepts.html.
	Metadata map[string]string `json:"metadata"`

	// Properties is a set of key-value pairs, if any, that are associated with
	// the image.
	Properties map[string]interface{} `json:"-"`

	// CreatedAt is the date when the image has been created.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt is the date when the last change has been made to the image or
	// it's properties.
	UpdatedAt time.Time `json:"updated_at"`

	// File is the trailing path after the glance endpoint that represent the
	// location of the image or the path to retrieve it.
	File string `json:"file"`

	// Schema is the path to the JSON-schema that represent the image or image
	// entity.
	Schema string `json:"schema"`

	// VirtualSize is the virtual size of the image
	VirtualSize int64 `json:"virtual_size"`

	// Specifies the image URL.
	Self string `json:"self"`

	// Specifies whether the image has been deleted.
	// The value can be true&nbsp;or&nbsp;false.
	Deleted bool `json:"deleted"`

	// Specifies the environment where the image is used.
	// The value can be FusionCompute,Ironic, or DataImage.
	// For an ECS image, the value is FusionCompute.
	// For a data disk image, the value is DataImage.
	// For a BMS image, the value is&nbsp;Ironic.
	VirtualEnvType ImageVirtualEnvType `json:"virtual_env_type"`

	// Specifies the time when the image was deleted.
	// The value is in UTC format.
	DeletedAt string `json:"deleted_at"`
}

func (r *Image) UnmarshalJSON(b []byte) error {
	type tmp Image
	var s struct {
		tmp
		SizeBytes interface{} `json:"size"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Image(s.tmp)

	switch t := s.SizeBytes.(type) {
	case nil:
		return nil
	case float32:
		r.SizeBytes = int64(t)
	case float64:
		r.SizeBytes = int64(t)
	default:
		return fmt.Errorf("Unknown type for SizeBytes: %v (value: %v)", reflect.TypeOf(t), t)
	}

	// Bundle all other fields into Properties
	var result interface{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return err
	}
	if resultMap, ok := result.(map[string]interface{}); ok {
		delete(resultMap, "self")
		r.Properties = internal.RemainingKeys(Image{}, resultMap)
	}

	return err
}

type commonResult struct {
	gophercloud.Result
}

// Extract interprets any commonResult as an Image.
func (r commonResult) Extract() (*Image, error) {
	var s *Image
	err := r.ExtractInto(&s)
	return s, err
}

// CreateResult represents the result of a Create operation. Call its Extract
// method to interpret it as an Image.
type CreateResult struct {
	commonResult
}

// UpdateResult represents the result of an Update operation. Call its Extract
// method to interpret it as an Image.
type UpdateResult struct {
	commonResult
}

// GetResult represents the result of a Get operation. Call its Extract
// method to interpret it as an Image.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation. Call its
// ExtractErr method to interpret it as an Image.
type DeleteResult struct {
	gophercloud.ErrResult
}

// PutTagResult represents the result of a put tag operation.
type PutTagResult struct {
	gophercloud.ErrResult
}

// DeleteTagResult represents the result of a delete tag operation.
type DeleteTagResult struct {
	gophercloud.ErrResult
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
// ImageSchemas presents the result of getting image schemas request
type ImageSchemas struct {
	// AdditionalProperties presents the additional properties
	AdditionalProperties map[string]string `json:"additionalProperties"`
	// Name is the name of schemas
	Name string `json:"name"`
	// Links is the links of schemas
	Links []map[string]string `json:"links"`
	// Properties is the explaination of schemas properties
	Properties *json.RawMessage `json:"properties"`
}

// ImageSchemasResult represents the result of Image schemas request
type ImageSchemasResult struct {
	gophercloud.Result
}

// Extract interprets the result as an ImageSchemas
func (r ImageSchemasResult) Extract() (*ImageSchemas, error) {
	var s *ImageSchemas
	err := r.ExtractInto(&s)
	return s, err
}

// ImagesSchemas presents the result of getting images schemas request
type ImagesSchemas struct {
	// Name is the name of schemas
	Name string `json:"name"`
	// Links is the links of schemas
	Links []map[string]string `json:"links"`
	// Properties is the explaination of schemas properties
	Properties *json.RawMessage `json:"properties"`
}

// ImagesSchemasResult represents the result of Images schemas request
type ImagesSchemasResult struct {
	gophercloud.Result
}

// Extract interprets the result as an ImagesSchemas
func (r ImagesSchemasResult) Extract() (*ImagesSchemas, error) {
	var s *ImagesSchemas
	err := r.ExtractInto(&s)
	return s, err
}