package cloudservers

import (
	"github.com/gophercloud/gophercloud"
	"fmt"
)


// Get requests details on a single server, by ID.
func Get(client *gophercloud.ServiceClient, serverID string) (r GetResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err=err
		return r
	}
	_, r.Err = client.Get(getURL(client, serverID), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return
}

//func GetServerRecoveryStatus(client *gophercloud.ServiceClient, serverID string) (r RecoveryResult) {
//	if serverID == "" {
//		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
//		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
//		r.Err = err
//		return r
//	}
//	_, r.Err = client.Get(autorecoveryURL(client, serverID), &r.Body, nil)
//	return
//}

//func ConfigServerRecovery(client *gophercloud.ServiceClient, serverID string, opts string) (r ErrResult) {
//	if serverID == "" {
//		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
//		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
//		r.Err = err
//		return r
//	}
//
//	b := map[string]string{
//		"support_auto_recovery": opts,
//	}
//
//	_, r.Err = client.Put(autorecoveryURL(client, serverID), &b, nil, &gophercloud.RequestOpts{
//		OkCodes:[]int{204},
//	})
//	return
//}

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
	Servers	  []Server  `json:"servers" required:"true"`
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

