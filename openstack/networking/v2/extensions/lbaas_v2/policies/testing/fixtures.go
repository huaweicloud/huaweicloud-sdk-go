package testing

import (
	"testing"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/policies"
	th "github.com/gophercloud/gophercloud/testhelper"
	"net/http"
	"github.com/gophercloud/gophercloud/testhelper/client"
	"fmt"
)

var CreateResp = `
{
    "l7policy": {
        "redirect_pool_id": "431a03eb-81bb-408e-ae37-7ce19023692b", 
        "redirect_listener_id": null, 
        "description": "", 
        "admin_state_up": true, 
        "rules": [
            {
                "id": "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3"
            }, 
            {
                "id": "f02b3bca-69d2-4335-a3fa-a8054e996213"
            }
        ], 
        "tenant_id": "a31d2bdcf7604c0faaddb058e1e08819", 
        "listener_id": "26058b64-6185-4e06-874e-4bd68b7633d0", 
        "redirect_url": null, 
        "action": "REDIRECT_TO_POOL", 
        "position": 2, 
        "provisioning_status": "ACTIVE",
        "id": "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", 
        "name": "test1"
    }
}`

var GetRespList = `
{
    "l7policies": [
        {
            "redirect_pool_id": "431a03eb-81bb-408e-ae37-7ce19023692b", 
            "redirect_listener_id": null,  
            "description": "", 
            "admin_state_up": true, 
            "rules": [
                {
                    "id": "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3"
                }, 
                {
                    "id": "f02b3bca-69d2-4335-a3fa-a8054e996213"
                }
            ], 
            "tenant_id": "a31d2bdcf7604c0faaddb058e1e08819", 
            "listener_id": "26058b64-6185-4e06-874e-4bd68b7633d0", 
            "redirect_url": null, 
            "action": "REDIRECT_TO_POOL", 
            "position": 2,
            "provisioning_status": "ACTIVE", 
            "id": "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", 
            "name": "test1"
        }, 
        {
            "redirect_pool_id": "59eebd7b-c68f-4f8a-aa7f-e062e84c0690", 
            "redirect_listener_id": null,  
            "description": "", 
            "admin_state_up": true, 
            "rules": [
                {
                    "id": "f4499f48-de3d-4efe-926d-926aa4d6aaf5"
                }
            ], 
            "tenant_id": "a31d2bdcf7604c0faaddb058e1e08819", 
            "listener_id": "e1310063-00de-4867-ab55-ccac4d9db364", 
            "redirect_url": null, 
            "action": "REDIRECT_TO_POOL", 
            "position": 1, 
            "provisioning_status": "ACTIVE",
            "id": "6cfd9d89-1d7e-4d84-ae1f-a8c5ff126f72", 
            "name": "test2"
        }
    ]
}
`

var GetResp = `
{
    "l7policy": {
        "redirect_pool_id": "431a03eb-81bb-408e-ae37-7ce19023692b", 
        "description": "", 
        "admin_state_up": true, 
        "rules": [
            {
                "id": "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3"
            }, 
            {
                "id": "f02b3bca-69d2-4335-a3fa-a8054e996213"
            }
         ],
        "tenant_id": "a31d2bdcf7604c0faaddb058e1e08819", 
        "listener_id": "26058b64-6185-4e06-874e-4bd68b7633d0", 
        "redirect_url": null, 
        "action": "REDIRECT_TO_POOL", 
        "position": 2,
        "provisioning_status": "ACTIVE", 
        "id": "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", 
        "name": "test1"
    }
}`
var UpdateResp = `
{
    "l7policy": {
        "redirect_pool_id": "431a03eb-81bb-408e-ae37-7ce19023692b", 
        "description": "", 
        "admin_state_up": true, 
        "rules": [
            {
                "id": "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3"
            }, 
            {
                "id": "f02b3bca-69d2-4335-a3fa-a8054e996213"
            }
        ], 
        "tenant_id": "a31d2bdcf7604c0faaddb058e1e08819", 
        "listener_id": "26058b64-6185-4e06-874e-4bd68b7633d0", 
        "redirect_url": null, 
        "action": "REDIRECT_TO_POOL", 
        "position": 2, 
        "id": "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", 
        "name": "test"
    }
}
`
var CreateRuleResp = `
{
    "rule": {
        "compare_type": "EQUAL_TO", 
        "provisioning_status": "ACTIVE",
        "admin_state_up": true, 
        "tenant_id": "a31d2bdcf7604c0faaddb058e1e08819", 
        "invert": false, 
        "value": "www.test.com", 
        "key": null, 
        "type": "HOST_NAME", 
        "id": "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3"
    }
}
`

var GetRuleRespList = `
{
    "rules": [
        {
            "compare_type": "EQUAL_TO", 
            "provisioning_status": "ACTIVE",
            "admin_state_up": true, 
            "tenant_id": "a31d2bdcf7604c0faaddb058e1e08819", 
            "invert": false, 
            "value": "www.test.com", 
            "key": null, 
            "type": "HOST_NAME", 
            "id": "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3"
        }, 
        {
            "compare_type": "EQUAL_TO",
            "provisioning_status": "ACTIVE", 
            "admin_state_up": true, 
            "tenant_id": "a31d2bdcf7604c0faaddb058e1e08819", 
            "invert": false, 
            "value": "/aaa.html", 
            "key": null, 
            "type": "PATH", 
            "id": "f02b3bca-69d2-4335-a3fa-a8054e996213"
        }
    ]
}
`
var GetRuleResp = `
{
    "rule": {
        "compare_type": "EQUAL_TO", 
        "provisioning_status": "ACTIVE",
        "admin_state_up": true, 
        "tenant_id": "a31d2bdcf7604c0faaddb058e1e08819", 
        "invert": false, 
        "value": "www.test.com", 
        "key": null, 
        "type": "HOST_NAME", 
        "id": "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3"
    }
}
`
var UpdateRuleResp = `
{
    "rule": {
        "compare_type": "EQUAL_TO", 
        "provisioning_status": "ACTIVE",
        "admin_state_up": true, 
        "tenant_id": "a31d2bdcf7604c0faaddb058e1e08819", 
        "invert": false, 
        "value": "www.test.com", 
        "key": null, 
        "type": "HOST_NAME", 
        "id": "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3"
    }
}
`

var (
	PoliciesOne = policies.Policies{
		ProvisioningStatus: "ACTIVE",
		ID:                 "5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586",
		Name:               "test1",
		Position:           2,
		Action:             "REDIRECT_TO_POOL",
		RedirectURL:        "",
		ListenerID:         "26058b64-6185-4e06-874e-4bd68b7633d0",
		TenantID:           "a31d2bdcf7604c0faaddb058e1e08819",
		AdminStateUp:       true,
		Description:        "",
		RedirectPoolID:     "431a03eb-81bb-408e-ae37-7ce19023692b",
		Rules: []policies.RuleId{
			{ID: "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3",},
			{ID: "f02b3bca-69d2-4335-a3fa-a8054e996213",},
		},
	}
	PoliciesTwo = policies.Policies{
		ProvisioningStatus: "ACTIVE",
		ID:                 "6cfd9d89-1d7e-4d84-ae1f-a8c5ff126f72",
		Name:               "test2",
		Position:           1,
		Action:             "REDIRECT_TO_POOL",
		RedirectURL:        "",
		ListenerID:         "e1310063-00de-4867-ab55-ccac4d9db364",
		TenantID:           "a31d2bdcf7604c0faaddb058e1e08819",
		AdminStateUp:       true,
		Description:        "",
		RedirectPoolID:     "59eebd7b-c68f-4f8a-aa7f-e062e84c0690",
		Rules: []policies.RuleId{
			{ID: "f4499f48-de3d-4efe-926d-926aa4d6aaf5",},
		},
	}
	RuleOne = policies.PolicyRule{
		CompareType:  "EQUAL_TO",
		AdminStateUp: true,
		TenantId:     "a31d2bdcf7604c0faaddb058e1e08819",
		Invert:       false,
		Value:        "www.test.com",
		Key:          "",
		Type:         "HOST_NAME",
		ID:           "67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3",
	}
	RuleTwo = policies.PolicyRule{
		CompareType:  "EQUAL_TO",
		AdminStateUp: true,
		TenantId:     "a31d2bdcf7604c0faaddb058e1e08819",
		Invert:       false,
		Value:        "/aaa.html",
		Key:          "",
		Type:         "PATH",
		ID:           "f02b3bca-69d2-4335-a3fa-a8054e996213",
	}
)

func HandlePolicySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, GetRespList)
		case "6cfd9d89-1d7e-4d84-ae1f-a8c5ff126f72":
			fmt.Fprintf(w, `{ "l7policies": [] }`)
		default:
			t.Fatalf("/v2.0/lbaas/l7policies invoked with unexpected marker=[%s]", marker)
		}
	})
}

func HandlePoliciesALLSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetRespList)

	})
}

func HandlePolicyCreationSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, CreateResp)
	})

}

func HandlePolicyGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies/5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", func(w http.ResponseWriter, r *http.Request) {

		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetResp)
	})
}

func HandlePolicyDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies/5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", func(w http.ResponseWriter, r *http.Request) {

		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}

func HandlePolicyUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies/5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586", func(w http.ResponseWriter, r *http.Request) {

		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetResp)
	})
}

func HandleRulesGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies/5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586/rules/67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w,GetRuleResp)
		})
}

func HandleRulesListALLSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies/5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586/rules",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, GetRuleRespList)

		})
}

func HandleRuleCreationSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies/5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586/rules",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, CreateRuleResp)
		})

}

func HandleRuleDeletionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies/5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586/rules/67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3",
		func(w http.ResponseWriter, r *http.Request) {

			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})
}

func HandleRuleUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v2.0/lbaas/l7policies/5ae0e1e7-5f0f-47a1-b39f-5d4c428a1586/rules/67d8a8fa-b0dd-4bd4-a85b-671db19b2ef3",
		func(w http.ResponseWriter, r *http.Request) {

			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, UpdateRuleResp)
		})
}
