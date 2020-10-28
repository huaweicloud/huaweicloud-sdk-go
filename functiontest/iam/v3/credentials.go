package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/iam/v3/credentials"
)

func main() {

	fmt.Println("main start...")

	provider, err_auth := common.AuthAKSK()
	//provider, err_auth := common.AuthToken()

	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}

	sc, err := openstack.NewIAMV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get IAM v3.0 failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	CreatePermanentAccessKey(sc)
	CreateTemporaryAccessKeyByAgency(sc)
	CreateTemporaryAccessKeyByToken(sc)
	DeletePermanentAk(sc)
	GetAllPermanentAks(sc)
	GetPermanentAk(sc)
	UpdatePermanentAk(sc)

	fmt.Println("main end...")
}

// 创建永久访问密钥
// Create a permanent accesskey
// POST /v3.0/OS-CREDENTIAL/credentials
func CreatePermanentAccessKey(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestCreatePermanentAccessKey")
	opts := credentials.CreateOpts{
		Description: "",
		UserID:      "",
	}
	resp := credentials.Create(sc, opts)
	result, err := resp.Extract()
	if err != nil {
		fmt.Println("CreatePermanentAccessKey failed")
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestCreatePermanentAccessKey")
}

// 通过委托来获取临时访问密钥(临时AK/SK)和securitytoken
// Create a temporary accesskey and securitytoken by the agency
// POST /v3.0/OS-CREDENTIAL/securitytokens
func CreateTemporaryAccessKeyByAgency(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestCreateTemporaryAccessKeyByAegency")
	sessionUser := credentials.SessionUser{Name: "SessionUserName"}
	durationSeconds := 1000
	assumeRole := credentials.AssumeRole{
		AgencyName:      "",
		DomainID:        "",
		DomainName:      "",
		DurationSeconds: &durationSeconds,
		SessionUser:     &sessionUser,
	}
	var methods []string
	methods = append(methods, "assume_role")
	var action []string
	action = append(action, "obs:object:*")
	var condition map[string]map[string][]string
	condition = make(map[string]map[string][]string)
	var prefix []string
	prefix = append(prefix, "public")
	var prefixMap map[string][]string
	prefixMap = make(map[string][]string)
	prefixMap["obs:prefix"] = prefix
	condition["StringEquals"] = prefixMap
	var resource []string
	resource = append(resource, "obs:*:*:object:*")
	statement := credentials.Statement{
		Action:    action,
		Condition: condition,
		Effect:    "allow",
		Resource:  resource,
	}
	var statements []*credentials.Statement
	statements = append(statements, &statement)
	policy := credentials.Policy{
		Version:   "1.1",
		Statement: statements,
	}
	identity := credentials.Identity{
		AssumeRole: &assumeRole,
		Methods:    methods,
		Policy:     &policy,
	}
	auth := credentials.Auth{Identity: &identity}
	opts := credentials.CreateTemporaryOpts{Auth: &auth}
	resp := credentials.CreateTemporaryAk(sc, opts)
	result, err := resp.ExtractTemporary()
	if err != nil {
		fmt.Println("CreatePermanentAccessKey failed")
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestCreateTemporaryAccessKeyByAegency")
}

// 通过token来获取临时访问密钥(临时AK/SK)和securitytoken
// Create a temporary accesskey and securitytoken by the token
// POST /v3.0/OS-CREDENTIAL/securitytokens
func CreateTemporaryAccessKeyByToken(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestCreateTemporaryAccessKeyByToken")
	var methods []string
	methods = append(methods, "token")
	authToken := ""
	durationSeconds := 3600
	token := credentials.Token{
		DurationSeconds: &durationSeconds,
		Id:              authToken,
	}
	var action []string
	action = append(action, "obs:object:*")
	var condition map[string]map[string][]string
	condition = make(map[string]map[string][]string)
	var prefix []string
	prefix = append(prefix, "public")
	var prefixMap map[string][]string
	prefixMap = make(map[string][]string)
	prefixMap["obs:prefix"] = prefix
	condition["StringEquals"] = prefixMap
	var resource []string
	resource = append(resource, "obs:*:*:object:*")
	statement := credentials.Statement{
		Action:    action,
		Condition: condition,
		Effect:    "allow",
		Resource:  resource,
	}
	var statements []*credentials.Statement
	statements = append(statements, &statement)
	policy := credentials.Policy{
		Version:   "1.1",
		Statement: statements,
	}
	identity := credentials.Identity{
		Methods: methods,
		Token:   &token,
		Policy:  &policy,
	}
	auth := credentials.Auth{Identity: &identity}
	opts := credentials.CreateTemporaryOpts{Auth: &auth}

	resp := credentials.CreateTemporaryAk(sc, opts)
	result, err := resp.ExtractTemporary()
	if err != nil {
		fmt.Println("CreatePermanentAccessKey failed")
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestCreateTemporaryAccessKeyByToken")
}

// 删除IAM用户的指定永久访问密钥
// Delete the permanent accesskey
// DELETE /v3.0/OS-CREDENTIAL/credentials/{access_key}
func DeletePermanentAk(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestDeletePermanentAk")
	ak := ""
	resp := credentials.DeletePermanentAk(sc, ak)
	if resp.Err != nil {
		fmt.Println("DeletePermanentAccessKey failed")
	} else {
		fmt.Println("DeletePermanentAccessKey successfully")
	}
	fmt.Println("finish TestDeletePermanentAk")
}

// 查询IAM用户的所有永久访问密钥
// Query all permanent accesskeys
// GET /v3.0/OS-CREDENTIAL/credentials
func GetAllPermanentAks(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestGetAllPermanentAks")
	userID := ""
	resp := credentials.GetAllPermanentAks(sc, userID)
	result, err := resp.ExtractCredentials()
	if err != nil {
		fmt.Println("GetAllPermanentAks failed")
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestGetAllPermanentAks")
}

// 查询指定永久访问密钥
// Query the permanent accesskeys
// GET /v3.0/OS-CREDENTIAL/credentials/{access_key}
func GetPermanentAk(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestGetPermanentAk")
	ak := ""
	resp := credentials.GetPermanentAk(sc, ak)
	result, err := resp.ExtractPermanent()
	if err != nil {
		fmt.Println("CreatePermanentAccessKey failed")
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestGetPermanentAk")
}

// 修改指定永久访问密钥
// Update the permanent accesskeys
// PUT /v3.0/OS-CREDENTIAL/credentials/{access_key}
func UpdatePermanentAk(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestUpdatePermanentAk")
	ak := ""
	status := ""
	description := ""
	opts := credentials.UpdateOpts{
		Description: &description,
		Status:      &status,
	}
	resp := credentials.UpdatePermanentAk(sc, opts, ak)
	result, err := resp.ExtractUpdatePermanent()
	if err != nil {
		fmt.Println("UpdatePermanentAccessKey failed")
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestUpdatePermanentAk")
}
