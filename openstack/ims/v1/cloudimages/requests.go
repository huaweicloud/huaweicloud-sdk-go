package cloudimages

import (
	"github.com/gophercloud/gophercloud"
)

// ListImageTagsOpts is the options for list image tags operation
type ListImageTagsOpts struct {
	// Limit is Integer value for the limit of values to return
	Limit int `q:"limit"`

	// Page represents the page number listing
	Page int `q:"page"`
}

// ToImageTagsListQuery builds the options to url query parameters
func (opts ListImageTagsOpts) ToImageTagsListQuery() (string, error) {

	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

/*
// ListImageTags list tags of an image
func ListImageTags(client *gophercloud.ServiceClient, opts ListImageTagsOpts) (r ImageTagsResult) {

	b, err := opts.ToImageTagsListQuery()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Get(getImageTagsURL(client)+b, &r.Body, nil)
	return
}
 */
// SetImageTagOpts represents the request body of SetImageTag
type SetImageTagOpts struct {
	// ImageID is the image id
	ImageID string `json:"image_id"`

	// Tag is the tag you wanna set
	Tag string `json:"tag"`
	// set map tag
	ImageTag map[string]interface{} `json:"image_tag"`
}

// ToImageSetMap assembles a request body
func (opts SetImageTagOpts) ToImageTagSetMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/*
// SetImageTag set a tag to an image
func SetImageTag(client *gophercloud.ServiceClient, opts SetImageTagOpts) (r SetImageTagResult) {

	b, err := opts.ToImageTagSetMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(putImageTagsURL(client), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204}})
	return
}

 //ImportImage import an image from OBS bucket to the huawei cloud
func ImportImage(client *gophercloud.ServiceClient, imageID, imageURL string) (r ImportResult) {
	_, r.Err = client.Put(importImageURL(client, imageID), map[string]string{
		"image_url": imageURL,
	}, nil, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return
}

// ExportImage export an image from huawei cloud to OBS bucket
func ExportImage(client *gophercloud.ServiceClient, imageID, bucketURL, fileFormat string) (r ExportResult) {
	_, r.Err = client.Post(exportImageURL(client, imageID), map[string]string{
		"bucket_url":  bucketURL,
		"file_format": fileFormat,
	}, nil, &gophercloud.RequestOpts{OkCodes: []int{200}})
	return
}

*/
// CopyImageOpts represents the request body struct of copy image operation
type CopyImageOpts struct {
	// Name is the image name
	Name string `json:"name"`
	// Description is the description of the image
	Description string `json:"description,omitempty"`
	// CmkID is the key of encryption, default is empty
	CmkID string `json:"cmk_id,omitempty"`
}

// ToImageCopyMap assembles a request body
func (opts CopyImageOpts) ToImageCopyMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/*
// CopyImage copy an existed image to another one
func CopyImage(client *gophercloud.ServiceClient, imageID string, opts CopyImageOpts) (r ImageCopyResult) {

	b, err := opts.ToImageCopyMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(
		copyImageURL(client, imageID),
		b,
		&r.Body,
		&gophercloud.RequestOpts{OkCodes: []int{200}},
	)
	return
}
*/
// AddImageMembersOpts is the request body of image members adding operation
type AddImageMembersOpts struct {
	// Images is the image list to operating
	Images []string `json:"images"`
	// Projects is the member list
	Projects []string `json:"projects"`
}

// ToImageMemberAddMap assembles a request body
func (opts AddImageMembersOpts) ToImageMemberAddMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/*
// AddImageMembers add members to images
func AddImageMembers(client *gophercloud.ServiceClient, opts AddImageMembersOpts) (r AddImageMembersResult) {

	b, err := opts.ToImageMemberAddMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(
		imageMemberOpURL(client),
		b,
		&r.Body,
		&gophercloud.RequestOpts{OkCodes: []int{200}},
	)
	return
}
*/

// UpdateImageMemberOpts is the request body for image member updating
type UpdateImageMemberOpts struct {
	// Images is the image list to update
	Images []string `json:"images"`
	// ProjectiD is the project id used to set
	ProjectID string `json:"project_id"`
	// Status is the status of members, accepted means accept shared image,
	// otherwise rejected
	Status string `json:"status"`
}

// ToImageMemberUpdateMap assembles a request body
func (opts UpdateImageMemberOpts) ToImageMemberUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// UpdateImageMember update member to images
/*
func UpdateImageMember(client *gophercloud.ServiceClient, opts UpdateImageMemberOpts) (r UpdateImageMemberResult) {

	b, err := opts.ToImageMemberUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(
		imageMemberOpURL(client),
		b,
		&r.Body,
		&gophercloud.RequestOpts{OkCodes: []int{200}},
	)
	return
}
*/

// DeleteImageMembersOpts is the request body of image members adding
// operation
type DeleteImageMembersOpts struct {
	// Images is the image list to operating
	Images []string `json:"images"`
	// Projects is the member list
	Projects []string `json:"projects"`
}

// ToImageMemberUpdateMap assembles a request body
func (opts DeleteImageMembersOpts) ToImageMembersDeleteMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/*
// DeleteImageMembers delete members from images
func DeleteImageMembers(client *gophercloud.ServiceClient, opts DeleteImageMembersOpts) (r DeleteImageMembersResult) {

	b, err := opts.ToImageMembersDeleteMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Delete(
		imageMemberOpURL(client),
		&gophercloud.RequestOpts{
			JSONBody:     b,
			JSONResponse: &r.Body,
			OkCodes:      []int{200}},
	)

	return
}

 //GetQuota get the quota of cloud image
func GetQuota(client *gophercloud.ServiceClient) (r QuotaResult) {
	_, r.Err = client.Get(getCloudImagesQuota(client), &r.Body, nil)
	return r
}

 */
