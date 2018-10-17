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
