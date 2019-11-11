package cloudservers

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Get requests details on a single server, by ID.
func Get(client *gophercloud.ServiceClient, serverID string) (r GetResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}
	_, r.Err = client.Get(getURL(client, serverID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return
}

func GetServerRecoveryStatus(client *gophercloud.ServiceClient, serverID string) (r RecoveryResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}
	_, r.Err = client.Get(autorecoveryURL(client, serverID), &r.Body, nil)
	return
}

func ConfigServerRecovery(client *gophercloud.ServiceClient, serverID string, opts string) (r ErrResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}

	b := map[string]string{
		"support_auto_recovery": opts,
	}

	_, r.Err = client.Put(autorecoveryURL(client, serverID), &b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

//func AddServerOnMonitorList(client *gophercloud.ServiceClient, serverID string) (r ErrResult) {
//	if serverID == "" {
//		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
//		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
//		r.Err = err
//		return r
//	}
//
//	b := map[string]interface{}{
//		"monitorMetrics": nil,
//	}
//
//	_, r.Err = client.Post(actionURL(client, serverID), &b, nil, &gophercloud.RequestOpts{
//		OkCodes: []int{200},})
//	return
//}

//BatchChangeOptsBuilder allows extensions to add additional parameters to the BatchChangeOpts request.
type BatchChangeOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerBatchChangeMap() (map[string]interface{}, error)
}

//BatchChangeOpts defining the configuration to batch change OS of servers
type BatchChangeOpts struct {
	AdminPass string    `json:"adminpass,omitempty"`
	KeyName   string    `json:"keyname,omitempty"`
	UserID    string    `json:"userid,omitempty"`
	ImageID   string    `json:"imageid" required:"true"`
	Servers   []Server  `json:"servers" required:"true"`
	MetaData  *MetaData `json:"metadata,omitempty"`
}

//Server defining the server configuration in BatchChangeOpts struct.
type Server struct {
	ID string `json:"id" required:"true"`
}

//MetaData defining the metadata configuration in BatchChangeOpts struct.
type MetaData struct {
	UserData string `json:"user_data,omitempty"`
}

//ToServerBatchChangeMap builds a request body from BatchChangeOpts.
func (opts BatchChangeOpts) ToServerBatchChangeMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"os-change": body}, nil
}

//BatchChangeOS batch change OS of servers based on the configuration defined in the BatchChangeOpts struct.
func BatchChangeOS(client *gophercloud.ServiceClient, opts BatchChangeOptsBuilder) (r BatchChangeResult) {
	body, err := opts.ToServerBatchChangeMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(batchChangeURL(client), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToServerListDetailQuery() (string, error)
}

// ListOpts allows the filtering and sorting of collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned.
type ListOpts struct {
	// Offset is the current page number.
	Offset int `q:"offset"`

	// Flavor is the ID of the flavor.
	Flavor string `q:"flavor"`

	// Name is the name of the server.
	Name string `q:"name"`

	// Status is the value of the status of the server so that you can filter on
	// "ACTIVE" for example.
	Status string `q:"status"`

	// Limit is an integer value for the limit of values to return.
	Limit int `q:"limit"`

	// Tags is used to filter out the servers with the specified tags
	Tags string `q:"tags"`

	// NotTags queries the cloud server that does not contain this value in the tag field.
	NotTags string `q:"not-tags"`

	// When you create an elastic cloud server in batches, you can specify the returned ID to query the elastic cloud server created in batches.
	ReservationID string `q:"reservation_id"`

	// EnterpriseProjectID specifies the server that is bound to an enterprise project.
	EnterpriseProjectID string `q:"enterprise_project_id"`
}

// ToServerListDetailQuery formats a ListOpts into a query string.
func (opts ListOpts) ToServerListDetailQuery() (string, error) {
	if opts.Offset == 0 {
		opts.Offset = 1
	}
	if opts.Limit == 0 {
		opts.Limit = 25
	}
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// ListDetail makes a request against the API to list servers accessible to you.
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listDetailURL(client)
	if opts != nil {
		query, err := opts.ToServerListDetailQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return CloudServerPage{pagination.OffsetPage{PageResult: r}}
	})
}

//BatchStartOptsBuilder allows extensions to add additional parameters to the BatchStartOpts request.
type BatchStartOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerBatchStartMap() (map[string]interface{}, error)
}

//BatchStartOpts defining the configuration to batch start servers
type BatchStartOpts struct {
	Servers []Server `json:"servers" required:"true"`
}

//ToServerBatchStartMap builds a request body from BatchStartOpts.
func (opts BatchStartOpts) ToServerBatchStartMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"os-start": body}, nil
}

//BatchStart makes a request against the API to batch start servers.
func BatchStart(client *gophercloud.ServiceClient, opts BatchStartOpts) (r JobResult) {
	body, err := opts.ToServerBatchStartMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchActionURL(client), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

//BatchRebootOptsBuilder allows extensions to add additional parameters to the BatchRebootOpts request.
type BatchRebootOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerBatchRebootMap() (map[string]interface{}, error)
}

//BatchRebootOpts defining the configuration to batch reboot servers
type BatchRebootOpts struct {
	// Type is the type of reboot to perform on the server.
	Type    Type     `json:"type" required:"true"`
	Servers []Server `json:"servers" required:"true"`
}

// Type describes the mechanisms by which a server reboot or stop can be requested.
type Type string

// These constants determine how a server should be rebooted or stopped.
const (
	Soft Type = "SOFT"
	Hard Type = "HARD"
)

//ToServerBatchRebootMap builds a request body from BatchRebootOpts.
func (opts BatchRebootOpts) ToServerBatchRebootMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"reboot": body}, nil
}

//BatchReboot makes a request against the API to batch reboot servers.
func BatchReboot(client *gophercloud.ServiceClient, opts BatchRebootOpts) (r JobResult) {
	body, err := opts.ToServerBatchRebootMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchActionURL(client), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

//BatchStopOptsBuilder allows extensions to add additional parameters to the BatchStopOpts request.
type BatchStopOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerBatchStopMap() (map[string]interface{}, error)
}

//BatchStopOpts defining the configuration to batch stop servers
type BatchStopOpts struct {
	// Type is the type of stop to perform on the server.
	Type    Type     `json:"type,omitempty"`
	Servers []Server `json:"servers" required:"true"`
}

//ToServerBatchStopMap builds a request body from BatchStopOpts.
func (opts BatchStopOpts) ToServerBatchStopMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"os-stop": body}, nil
}

//BatchStop makes a request against the API to batch stop servers.
func BatchStop(client *gophercloud.ServiceClient, opts BatchStopOpts) (r JobResult) {
	body, err := opts.ToServerBatchStopMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchActionURL(client), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

//BatchUpdateOptsBuilder allows extensions to add additional parameters to the BatchUpdateOpts request.
type BatchUpdateOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerBatchUpdateMap() (map[string]interface{}, error)
}

//BatchUpdateOpts defining the configuration to batch update servers
type BatchUpdateOpts struct {
	Name    string   `json:"name" required:"true"`
	DryRun  bool     `json:"dry_run,omitempty"`
	Servers []Server `json:"servers" required:"true"`
}

//ToServerBatchUpdateMap builds a request body from BatchUpdateOpts.
func (opts BatchUpdateOpts) ToServerBatchUpdateMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return body, nil
}

//BatchUpdate makes a request against the API to batch update servers.
func BatchUpdate(client *gophercloud.ServiceClient, opts BatchUpdateOpts) (r BatchUpdateResult) {
	body, err := opts.ToServerBatchUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(batchUpdateURL(client), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
		HandleError: func(httpStatus int, responseContent string) error {
			var batchUpdateError BatchOperateError
			message := responseContent
			err := json.Unmarshal([]byte(responseContent), &batchUpdateError)
			if err == nil {
				return &batchUpdateError
			}
			return &gophercloud.UnifiedError{
				ErrCode:    gophercloud.MatchErrorCode(httpStatus, message),
				ErrMessage: message,
			}
		},
	})

	return
}

//TagCreate defining the key and value of a tag for creating
type TagCreate struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value" required:"true"`
}

//TagDelete defining the key and value of a tag for deleting
type TagDelete struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value,omitempty"`
}

//BatchTagCreateOpts defining the configuration for batch server tags action
type BatchTagCreateOpts struct {
	Tags []TagCreate `json:"tags" required:"true"`
}

//BatchTagDeleteOpts defining the configuration for batch server tags action
type BatchTagDeleteOpts struct {
	Tags []TagDelete `json:"tags" required:"true"`
}

//BatchTagCreateOptsBuilder allows extensions to add additional parameters to the BatchTagActionOpts request.
type BatchTagCreateOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToTagBatchCreateMap() (map[string]interface{}, error)
}

//BatchTagDeleteOptsBuilder allows extensions to add additional parameters to the BatchTagActionOpts request.
type BatchTagDeleteOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToTagBatchDeleteMap() (map[string]interface{}, error)
}

//ToTagBatchCreateMap builds a request body from BatchTagActionOpts.
func (opts BatchTagCreateOpts) ToTagBatchCreateMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return body, nil
}

//ToTagBatchDeleteMap builds a request body from BatchTagActionOpts.
func (opts BatchTagDeleteOpts) ToTagBatchDeleteMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	return body, nil
}

//BatchCreateServerTags requests to batch create server tags
func BatchCreateServerTags(client *gophercloud.ServiceClient, serverID string, opts BatchTagCreateOptsBuilder) (r ErrResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}

	b, err := opts.ToTagBatchCreateMap()
	b["action"] = "create"
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchTagActionURL(client, serverID), &b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

//BatchDeleteServerTags requests to batch delete server tags
func BatchDeleteServerTags(client *gophercloud.ServiceClient, serverID string, opts BatchTagDeleteOptsBuilder) (r ErrResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}

	b, err := opts.ToTagBatchDeleteMap()
	b["action"] = "delete"
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(batchTagActionURL(client, serverID), &b, nil, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

// ListProjectTags makes a request against the API to list project tags accessible to you.
func ListProjectTags(client *gophercloud.ServiceClient) (r ProjectTagsResult) {
	_, r.Err = client.Get(listProjectTagsURL(client), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// ListServerTags makes a request against the API to list server tags accessible to you.
func ListServerTags(client *gophercloud.ServiceClient, serverID string) (r ServerTagsResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return r
	}
	_, r.Err = client.Get(listServerTagsURL(client, serverID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
