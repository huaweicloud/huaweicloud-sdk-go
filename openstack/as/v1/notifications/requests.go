package notifications

import (
	"github.com/gophercloud/gophercloud"
)

func Delete(client *gophercloud.ServiceClient, scalingGroupId string, topicUrn string) (r DeleteResult) {
	url := DeleteURL(client, scalingGroupId, topicUrn)
	_, r.Err = client.Delete(url, &gophercloud.RequestOpts{
		JSONResponse: nil,
		OkCodes:      []int{204},
	})
	return
}

type ConfigNotificationOpts struct {
	// Specifies a unified topic in SMN.
	TopicUrn string `json:"topic_urn" required:"true"`

	// Specifies a notification scenario, which can be one of the
	// following:SCALING_UP: indicates that the capacity is expanded.SCALING_UP_FAIL:
	// indicates that the capacity expansion failed.SCALING_DOWN: indicates that the
	// capacity is reduced.SCALING_DOWN_FAIL: indicates that the capacity reduction
	// failed.SCALING_GROUP_ABNORMAL: indicates that an exception has occurred in the AS
	// group.
	TopicScene []string `json:"topic_scene" required:"true"`
}

type ConfigNotificationOptsBuilder interface {
	ToNotificationsOptsMap() (map[string]interface{}, error)
}

func (opts ConfigNotificationOpts) ToNotificationsOptsMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(&opts, "")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ConfigNotification(client *gophercloud.ServiceClient, scalingGroupId string, opts ConfigNotificationOptsBuilder) (r EnableResult) {
	b, err := opts.ToNotificationsOptsMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(EnableURL(client, scalingGroupId), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

func List(client *gophercloud.ServiceClient, scalingGroupId string) (r ListResult) {
	url := ListURL(client, scalingGroupId)
	_, r.Err = client.Get(url, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
