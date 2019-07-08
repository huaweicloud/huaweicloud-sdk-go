package cloudimages

import "github.com/gophercloud/gophercloud"

// ImageTagsResult represents the result of image tags query
type ImageTagsResult struct {
	gophercloud.Result
}

// ImageTags is the golang struct of ImageTagsResult
type ImageTags struct {
	Tags []string `json:"tags"`
}

// Extract extracts ImageTagsResult to struct
func (r ImageTagsResult) Extract() (*ImageTags, error) {
	var t *ImageTags
	err := r.ExtractInto(&t)
	return t, err
}

// SetImageTagResult is the result of image tag setting operation
type SetImageTagResult struct {
	gophercloud.Result
}

// ImportResult represents the result of the image importing
type ImportResult struct {
	commonResult
}

// ExportResult represents the result of the image exporting
type ExportResult struct {
	commonResult
}

// CopyResult represents the result of the image copying
type ImageCopyResult struct {
	commonResult
}

// ImageCopyTask is the image copy task description
//type ImageTask struct {
//	asyncOperationJob
//}

type commonResult struct {
	gophercloud.Result
}

func (r commonResult) ExtractJob() (*OperationJob, error) {
	var t *OperationJob
	err := r.ExtractInto(&t)
	return t, err
}

type OperationJob struct {
	JobID string `json:"job_id"`
}

//// ImageCopyTask is the image copy task description
//type ImageCopyTask struct {
//	asyncOperationJob
//}
//
//// Extract extracts ImageCopyResult to golangstruct
//func (r ImageCopyResult) Extract() (*ImageCopyTask, error) {
//	var t *ImageCopyTask
//	err := r.ExtractInto(&t)
//	return t, err
//}

// AddImageMembersResult represents the result of image members adding operation
type AddImageMembersResult struct {
	commonResult
}

// AddImageMembersTask is the task description of image members adding operation
//type AddImageMembersTask struct {
//	asyncOperationJob
//}
//
//// Extract extracts AddImageMembersResult to golang struct
//func (r AddImageMembersResult) Extract() (*AddImageMembersTask, error) {
//	var t *AddImageMembersTask
//	err := r.ExtractInto(&t)
//	return t, err
//}

// UpdateImageMemberResult represent the result of image members updating
// operation
type UpdateImageMemberResult struct {
	commonResult
}

// UpdateImageMemberTask is the task description of image member update
//type UpdateImageMemberTask struct {
//	asyncOperationJob
//}

// Extract extracts UpdateImageMemberResult to golang struct
//func (r UpdateImageMemberResult) Extract() (*UpdateImageMemberTask, error) {
//	var t *UpdateImageMemberTask
//	err := r.ExtractInto(&t)
//	return t, err
//}

// DeleteImageMembersResult represents the result of image members deleting
// operation
type DeleteImageMembersResult struct {
	commonResult
}

//
//// DeleteImageMembersTask is the task description of image member deleting
//type DeleteImageMembersTask struct {
//	asyncOperationJob
//}

// Extract extracts DeleteImageMembersResult to golang struct
//func (r DeleteImageMembersResult) Extract() (*DeleteImageMembersTask, error) {
//	var t *DeleteImageMembersTask
//	err := r.ExtractInto(&t)
//	return t, err
//}

// ImagesQuota represents the quota of cloud images
type ImagesQuota struct {
	// Type is the resource type
	Type string `json:"type"`
	// Used is the amount of used resource
	Used int `json:"used"`
	// Quota is the amount of all resource
	Quota int `json:"quota"`
	// Min is the minimum of resource
	Min int `json:"min"`
	// Max is the maximum of resource
	Max int `json:"max"`
}

type Resources struct {
	Resources []ImagesQuota `json:"resources"`
}

type Quota struct {
	Quota Resources `json:"quotas"`
}

// CloudImagesQuotaResult represents the result of cloud image quota request
type QuotaResult struct {
	gophercloud.Result
}

// Extract extracts the result to CloudImagesQuota
func (r QuotaResult) Extract() (*Quota, error) {

	var s *Quota
	err := r.ExtractInto(&s)
	return s, err
}
