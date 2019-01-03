package testing

import (
	"testing"
	"fmt"
	"net/http"
	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/gophercloud/gophercloud/openstack/ims/v2/cloudimages"
)

var GetJobReuslt = `
{
    "status": "SUCCESS",
    "entities": {
        "image_id": "e9e91bff-14b6-4a0b-8377-4ed0813e3360"
    },
    "job_id": "ff8080814dbd65d7014dbe0d84db0013",
    "job_type": "createImageByInstance",
    "begin_time": "04-Jun-2015 18:11:06:586",
    "end_time": "",
    "error_code": null,
    "fail_reason": null
}
`

var CreateResult = `
{
    "job_id": "ff8080814dbd65d7014dbe0d84db0013"
}
`

var listResult = `
{
  "images": [
    {
      "schema": "/v2/schemas/image",
      "min_disk": 100,
      "created_at": "2018-09-06T14:03:27Z",
      "__image_source_type": "uds",
      "container_format": "bare",
      "file": "/v2/images/bc6bed6e-ba3a-4447-afcc-449174a3eb52/file",
      "updated_at": "2018-09-06T15:17:33Z",
      "protected": true,
      "checksum": "d41d8cd98f00b204e9800998ecf8427e",
      "__support_kvm_fpga_type": "VU9P",
      "id": "bc6bed6e-ba3a-4447-afcc-449174a3eb52",
      "__isregistered": "true",
      "min_ram": 2048,
      "__lazyloading": "true",
      "owner": "1bed856811654c1cb661a6ca845ebc77",
      "__os_type": "Linux",
      "__imagetype": "gold",
      "visibility": "public",
      "virtual_env_type": "FusionCompute",
      "tags": [],
      "__platform": "CentOS",
      "size": 0,
      "__os_bit": "64",
      "__os_version": "CentOS 7.3 64bit",
      "name": "CentOS 7.3 64bit vivado",
      "self": "/v2/images/bc6bed6e-ba3a-4447-afcc-449174a3eb52",
      "disk_format": "zvhd2",
      "virtual_size": null,
      "status": "active"
    },
    {
      "schema": "/v2/schemas/image",
      "min_disk": 100,
      "created_at": "2018-09-06T14:03:05Z",
      "__image_source_type": "uds",
      "container_format": "bare",
      "file": "/v2/images/0328c25e-c840-4496-81ac-c4e01b214b1f/file",
      "updated_at": "2018-09-25T14:27:40Z",
      "protected": true,
      "checksum": "d41d8cd98f00b204e9800998ecf8427e",
      "__support_kvm_fpga_type": "VU9P_COMMON",
      "id": "0328c25e-c840-4496-81ac-c4e01b214b1f",
      "__isregistered": "true",
      "min_ram": 2048,
      "__lazyloading": "true",
      "owner": "1bed856811654c1cb661a6ca845ebc77",
      "__os_type": "Linux",
      "__imagetype": "gold",
      "visibility": "public",
      "virtual_env_type": "FusionCompute",
      "tags": [],
      "__platform": "CentOS",
      "size": 0,
      "__os_bit": "64",
      "__os_version": "CentOS 7.3 64bit",
      "name": "CentOS 7.3 64bit with sdx",
      "self": "/v2/images/0328c25e-c840-4496-81ac-c4e01b214b1f",
      "disk_format": "zvhd2",
      "virtual_size": null,
      "status": "active"
    }
  ]
}
`


var createResult = cloudimages.Job{
	Id: "ff8080814dbd65d7014dbe0d84db0013",
}

var jobResult = cloudimages.JobResult{

	Id:         "ff8080814dbd65d7014dbe0d84db0013",
	Type:       "createImageByInstance",
	Status:     "SUCCESS",
	BeginTime:  "04-Jun-2015 18:11:06:586",
	EndTime:    "",
	ErrorCode:  "",
	FailReason: "",
	Entities: cloudimages.Entity{
		ImageId: "e9e91bff-14b6-4a0b-8377-4ed0813e3360",
	},
}

func HandleImageListSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/cloudimages", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResult)
	})
}

func HandleImageCreateByFileSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/cloudimages/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		w.Header().Add("Content-Type", "application/json")
		//w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, CreateResult)
	})
}
func HandleImageCreateByServerSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/cloudimages/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateResult)
	})

}
func HandleGetJobSuccessfully(t *testing.T) {

	th.Mux.HandleFunc("/jobs/ff8080814dbd65d7014dbe0d84db0013", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetJobReuslt)
	})

}
