package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/security/groups"
)

var ListOutput = `
{
    "security_groups": [
        {
            "id": "16b6e77a-08fa-42c7-aa8b-106c048884e6", 
            "name": "qq", 
            "description": "qq", 
            "vpc_id": "3ec3b33f-ac1c-4630-ad1c-7dba1ed79d85", 
            "enterprise_project_id": "0aad99bc-f5f6-4f78-8404-c598d76b0ed2",
            "security_group_rules": [
                {
                    "direction": "egress", 
                    "ethertype": "IPv4", 
                    "id": "369e6499-b2cb-4126-972a-97e589692c62", 
                    "description": "",
                    "security_group_id": "16b6e77a-08fa-42c7-aa8b-106c048884e6"
                }, 
                {
                    "direction": "ingress", 
                    "ethertype": "IPv4", 
                    "id": "0222556c-6556-40ad-8aac-9fd5d3c06171", 
                    "description": "",
                    "remote_group_id": "16b6e77a-08fa-42c7-aa8b-106c048884e6", 
                    "security_group_id": "16b6e77a-08fa-42c7-aa8b-106c048884e6"
                }
            ]
        }, 
        {
            "id": "9c0f56be-a9ac-438c-8c57-fce62de19419", 
            "name": "default", 
            "description": "qq", 
            "vpc_id": "13551d6b-755d-4757-b956-536f674975c0", 
            "enterprise_project_id": "0aad99bc-f5f6-4f78-8404-c598d76b0ed2",
            "security_group_rules": [
                {
                    "direction": "egress", 
                    "ethertype": "IPv4", 
                    "id": "95479e0a-e312-4844-b53d-a5e4541b783f", 
                    "description": "",
                    "security_group_id": "9c0f56be-a9ac-438c-8c57-fce62de19419"
                }, 
                {
                    "direction": "ingress", 
                    "ethertype": "IPv4", 
                    "id": "0c4a2336-b036-4fa2-bc3c-1a291ed4c431",
                    "description": "", 
                    "remote_group_id": "9c0f56be-a9ac-438c-8c57-fce62de19419", 
                    "security_group_id": "9c0f56be-a9ac-438c-8c57-fce62de19419"
                }
            ]
        }
    ]
}
`

var ListResponse = []groups.SecGroup{
	{
		ID:                  "16b6e77a-08fa-42c7-aa8b-106c048884e6",
		Name:                "qq",
		Description:         "qq",
		VpcID:               "3ec3b33f-ac1c-4630-ad1c-7dba1ed79d85",
		EnterpriseProjectID: "0aad99bc-f5f6-4f78-8404-c598d76b0ed2",
		SecurityGroupRules: []groups.SecurityGroupRule{{
			ID:              "369e6499-b2cb-4126-972a-97e589692c62",
			Description:     "",
			Direction:       "egress",
			Ethertype:       "IPv4",
			SecurityGroupID: "16b6e77a-08fa-42c7-aa8b-106c048884e6",
		},
			{
				ID:              "0222556c-6556-40ad-8aac-9fd5d3c06171",
				Description:     "",
				Direction:       "ingress",
				Ethertype:       "IPv4",
				RemoteGroupID:   "16b6e77a-08fa-42c7-aa8b-106c048884e6",
				SecurityGroupID: "16b6e77a-08fa-42c7-aa8b-106c048884e6",
			},
		},
	},
	{
		ID:                  "9c0f56be-a9ac-438c-8c57-fce62de19419",
		Name:                "default",
		Description:         "qq",
		VpcID:               "13551d6b-755d-4757-b956-536f674975c0",
		EnterpriseProjectID: "0aad99bc-f5f6-4f78-8404-c598d76b0ed2",
		SecurityGroupRules: []groups.SecurityGroupRule{{
			ID:              "95479e0a-e312-4844-b53d-a5e4541b783f",
			Description:     "",
			Direction:       "egress",
			Ethertype:       "IPv4",
			SecurityGroupID: "9c0f56be-a9ac-438c-8c57-fce62de19419",
		},
			{
				ID:              "0c4a2336-b036-4fa2-bc3c-1a291ed4c431",
				Description:     "",
				Direction:       "ingress",
				Ethertype:       "IPv4",
				RemoteGroupID:   "9c0f56be-a9ac-438c-8c57-fce62de19419",
				SecurityGroupID: "9c0f56be-a9ac-438c-8c57-fce62de19419",
			},
		},
	},
}

func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/57e98940a77f4bb988a21a7d0603a626/security-groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")

		r.ParseForm()
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, ListOutput)
		case "9c0f56be-a9ac-438c-8c57-fce62de19419":
			fmt.Fprintf(w, `{"security_groups": []}`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}
