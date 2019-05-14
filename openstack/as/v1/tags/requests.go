package tags

import (
	"github.com/gophercloud/gophercloud"
)

type Tag struct {
	// Specifies the resource tag key.
	Key string `json:"key" required:"true"`

	// Specifies the resource tag values.
	Value string `json:"value,omitempty"`
}

func ListResourceTags(client *gophercloud.ServiceClient, resourceType string, resourceId string) (r ListResourceTagsResult) {
	url := ListResourceTagsURL(client, resourceType, resourceId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func ListTenantTags(client *gophercloud.ServiceClient, resourceType string) (r ListTenantTagsResult) {
	url := ListTenantTagsURL(client, resourceType)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type InstanceOpts struct {
	// Operation ID (case sensitive).update: indicates updating a tag.
	// If the same key value exists, it will be overwritten. If no same key value exists, a
	// new tag will be created.delete: indicates deleting a tag.create: indicates creating a
	// tag. If the same key value already exists, it will be overwritten.
	Action     string `json:"action" required:"true"`
	Offset     string `json:"offset,omitempty"`
	Limit      string `json:"limit,omitempty"`
	Matches    []Tag  `json:"matches,omitempty"`
	NotTags    []Tags `json:"not_tags,omitempty"`
	Tags       []Tags `json:"tags,omitempty"`
	NotTagsAny []Tags `json:"not_tags_any,omitempty"`
	TagsAny    []Tags `json:"tags_any,omitempty"`
}

type InstanceOptsBuilder interface {
	ToInstanceMap() (map[string]interface{}, error)
}

func (opts InstanceOpts) ToInstanceMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ListInstanceTags(client *gophercloud.ServiceClient, resourceType string, opts InstanceOptsBuilder) (r ListInstanceTagsResult) {
	b, err := opts.ToInstanceMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(ListInstanceTagsURL(client, resourceType), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type UpdateOpts struct {
	// Specifies the tag list.If action is set to delete, the tag
	// structure cannot be missing, and the key cannot be left blank or an empty string.
	Tags []Tag `json:"tags" required:"true"`

	// Operation ID (case sensitive).update: indicates updating a tag.
	// If the same key value exists, it will be overwritten. If no same key value exists, a
	// new tag will be created.delete: indicates deleting a tag.create: indicates creating a
	// tag. If the same key value already exists, it will be overwritten.
	Action string `json:"action" required:"true"`
}

type UpdateOptsBuilder interface {
	ToTagsUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateOpts) ToTagsUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Update(client *gophercloud.ServiceClient, resourceType string, resourceId string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToTagsUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(UpdateURL(client, resourceType, resourceId), b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
