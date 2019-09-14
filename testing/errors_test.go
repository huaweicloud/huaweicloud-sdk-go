package testing

import (
	"testing"
	"github.com/gophercloud/gophercloud"
	th "github.com/gophercloud/gophercloud/testhelper"
	"encoding/json"
)

func TestOneLevelType1(t *testing.T) {
	expectedCode := "404"
	expectedMsg := "Instance *89973356-f733-418b-95b2-f6fc27244f18 could not be found."
	var resp struct {
		ErrorMsg string `json:"error_msg"`
		ErrCode  string `json:"error_code"`
	}

	resp.ErrCode = expectedCode
	resp.ErrorMsg = expectedMsg

	b, _ := json.Marshal(resp)

	err := gophercloud.NewSystemServerError(400, string(b))

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			th.CheckDeepEquals(t, expectedCode, ue.ErrorCode())
			th.CheckDeepEquals(t, expectedMsg, ue.Message())
		}
	}

}

func TestOneLevelType2(t *testing.T) {
	expectedCode := "VPC.0101"
	expectedMsg := "Instance *89973356-f733-418b-95b2-f6fc27244f18 could not be found."
	var resp struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}

	resp.Message = expectedMsg
	resp.Code = expectedCode

	b, _ := json.Marshal(resp)

	err := gophercloud.NewSystemServerError(400, string(b))

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			th.CheckDeepEquals(t, expectedCode, ue.ErrorCode())
			th.CheckDeepEquals(t, expectedMsg, ue.Message())
		}
	}
}

func TestTwoLevelType1(t *testing.T) {

	expectedCode := "404"
	expectedMsg := "Instance *89973356-f733-418b-95b2-f6fc27244f18 could not be found."

	type itemNotFound struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}
	var resp struct {
		ItemNotFound itemNotFound `json:"itemNotFound"`
	}

	resp.ItemNotFound.Message = expectedMsg
	resp.ItemNotFound.Code = expectedCode

	b, _ := json.Marshal(resp)

	err := gophercloud.NewSystemServerError(400, string(b))

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			th.CheckDeepEquals(t, "Ecs.1544", ue.ErrorCode())
			th.CheckDeepEquals(t, expectedMsg, ue.Message())
		}
	}

}

func TestTwoLevelType2(t *testing.T) {
	expectedCode := "404"
	expectedMsg := "Instance *89973356-f733-418b-95b2-f6fc27244f18 could not be found."

	type Error struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}

	var resp struct {
		Error Error `json:"error"`
	}

	resp.Error.Message = expectedMsg
	resp.Error.Code = expectedCode

	b, _ := json.Marshal(resp)

	err := gophercloud.NewSystemServerError(400, string(b))

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			th.CheckDeepEquals(t, "Ecs.1544", ue.ErrorCode())
			th.CheckDeepEquals(t, expectedMsg, ue.Message())
		}
	}
}
