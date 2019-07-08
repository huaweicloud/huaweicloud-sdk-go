package cloudservers

import (
	"strings"

	"github.com/gophercloud/gophercloud"
	"fmt"
)

func ResetPassword(client *gophercloud.ServiceClient, serverID, newPassword string) (r ActionResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return
	}

	if newPassword == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "newPassword")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return
	}
	body := map[string]interface{}{
		"reset-password": map[string]string{
			"new_password": newPassword,
		},
	}

	//v2替换成v2.1
	newUrl := strings.Replace(resetPwdURL(client, serverID), "/v2/", "/v2.1/", 1)

	_, r.Err = client.Put(newUrl, body, nil, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ChangeOptsBuilder interface {
	// Returns value that can be passed to json.Marshal
	ToServerChangeMap() (map[string]interface{}, error)
}

type ChangeOpts struct {
	AdminPass string    `json:"adminpass,omitempty"`
	KeyName   string    `json:"keyname,omitempty"`
	UserID    string    `json:"userid,omitempty"`
	ImageID   string    `json:"imageid" required:"true"`
	MetaData  *MetaData `json:"metadata,omitempty"`
}

type MetaData struct {
	UserData string `json:"user_data,omitempty"`
}

func (opts ChangeOpts) ToServerChangeMap() (map[string]interface{}, error) {
	body, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"os-change": body}, nil
}

func ChangeOS(client *gophercloud.ServiceClient, serverID string, opts ChangeOptsBuilder) (r ChangeResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return
	}
	body, err := opts.ToServerChangeMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(changeURL(client, serverID), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type ReinstallOptsBuilder interface {
	ToServerReinstallMap() (map[string]interface{}, error)
}

type ReinstallOpts struct {
	AdminPass string    `json:"adminpass,omitempty"`
	KeyName   string    `json:"keyname,omitempty"`
	UserID    string    `json:"userid,omitempty"`
	MetaData  *MetaData `json:"metadata,omitempty"`
}


func (opts ReinstallOpts) ToServerReinstallMap() (map[string]interface{}, error) {

	b, err := gophercloud.BuildRequestBody(opts, "os-reinstall")

	return b, err
}

func ReinstallOS(client *gophercloud.ServiceClient, serverID string, opts ReinstallOptsBuilder) (r ChangeResult) {
	if serverID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "serverID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		r.Err = err
		return
	}
	body, err := opts.ToServerReinstallMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(reinstallOSURL(client, serverID), body, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}


type ResizeOptsBuilder interface {
	ToResizeQuery() (string, error)
}

type ResizeOpts struct {
	// Limit limits the number of Images to return.
	Limit int `q:"limit"`

	// Mark is an Image UUID at which to set a marker.
	Marker string `q:"marker"`

	InstanceUUID string `q:"instance_uuid"`

	SourceFlavorID string `q:"source_flavor_id"`

	SourceFlavorName string `q:"source_flavor_name"`

	SortKey string `q:"sort_key"`

	SortDir string `q:"sort_dir"`
}

func (opts ResizeOpts) ToResizeQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

//func ResizeFlavor(client *gophercloud.ServiceClient, opts ResizeOptsBuilder) (r FlavorResult) {
//
//	url, err := opts.ToResizeQuery()
//	if err != nil {
//		r.Err = err
//		return
//	}
//	url = resizeFlavorURL(client) + url
//
//	//v2替换成v2.1
//	newUrl := strings.Replace(url, "/v2/", "/v2.1/", 1)
//
//	_, r.Err = client.Get(newUrl, &r.Body, &gophercloud.RequestOpts{
//		OkCodes: []int{200},
//	})
//	return
//}
