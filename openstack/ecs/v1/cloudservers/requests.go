package cloudservers

import (
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
