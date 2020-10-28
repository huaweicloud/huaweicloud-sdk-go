package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
)

func main() {

	fmt.Println("main start...")

	provider, err_auth := common.AuthAKSK()
	//provider, err_auth := common.AuthToken()

	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}

	sc, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get IAM v3 failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	CreateTokenForProjectByPassword(sc)
	CreateTokenForDomainByPassword(sc)
	CreateTokenForProjectByPasswordAndMFA(sc)
	CreateTokenForDomainByPasswordAndMFA(sc)
	ValidateToken(sc)
	CreateAgencyToken(sc)
	fmt.Println("main end...")
}

// 通过用户名/密码的方式进行认证来获取token, 作用范围为项目
// Get the token by name/password, and the function scope is project
// POST /v3/auth/tokens
func CreateTokenForProjectByPassword(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestCreateTokenForProjectByPassword")
	userName := ""
	userPwd := ""
	domainName := ""
	projectName := ""
	nocatalog := ""
	var method []string
	method = append(method, "password")
	user := tokens.UserReq{
		Name:     userName,
		Password: userPwd,
		Domain:   tokens.DomainReq{Name: domainName},
	}
	project := tokens.ProjectReq{
		Name: projectName,
	}
	scope := tokens.ScopeReq{Project: &project}
	password := tokens.PasswordReq{User: user}
	identity := tokens.IdentityReq{Password: &password, Methods: method}
	auth := tokens.AuthReq{Identity: &identity, Scope: &scope}
	opts := tokens.PwdTokenOptions{Auth: &auth}
	resp := tokens.CreateTokenByPassword(sc, opts, nocatalog)
	result, err := resp.ExtractTokenBody()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	bytes, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("finish TestCreateTokenForProjectByPassword")
}

// 通过用户名/密码的方式进行认证来获取token, 作用范围为整个账号
// Get the token by name/password, and the function scope is the account
// POST /v3/auth/tokens
func CreateTokenForDomainByPassword(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestCreateTokenForDomainByPassword")
	userName := ""
	userPwd := ""
	domainName := ""
	domainId := ""
	nocatalog := ""
	var method []string
	method = append(method, "password")
	user := tokens.UserReq{
		Name:     userName,
		Password: userPwd,
		Domain:   tokens.DomainReq{Name: domainName},
	}
	domain := tokens.DomainReq{
		ID:   domainId,
		Name: domainName,
	}
	scope := tokens.ScopeReq{Domain: &domain}
	password := tokens.PasswordReq{User: user}
	identity := tokens.IdentityReq{Password: &password, Methods: method}
	auth := tokens.AuthReq{Identity: &identity, Scope: &scope}
	opts := tokens.PwdTokenOptions{Auth: &auth}
	resp := tokens.CreateTokenByPassword(sc, opts, nocatalog)
	result, err := resp.ExtractTokenBody()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(bytes))

	fmt.Println("finish TestCreateTokenForDomainByPassword")
}

// 获取用户Token（使用密码+虚拟MFA）,作用范围为项目
// Get the token by name/password and virtual MFA, and the function scope is project
// POST /v3/auth/tokens
func CreateTokenForProjectByPasswordAndMFA(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestCreateTokenForProjectByPasswordAndMFA")
	userName := ""
	userId := ""
	userPwd := ""
	domainName := ""
	regionName := ""
	totpCode := ""
	nocatalog := ""
	var method []string
	method = append(append(method, "password"), "totp")
	user := tokens.UserReq{
		Name:     userName,
		Password: userPwd,
		Domain:   tokens.DomainReq{Name: domainName},
	}
	project := tokens.ProjectReq{
		Name: regionName,
	}
	totpUser := tokens.UserReq{
		ID:       userId,
		PassCode: totpCode,
	}
	totp := tokens.TotpReq{User: totpUser}
	scope := tokens.ScopeReq{Project: &project}
	password := tokens.PasswordReq{User: user}
	identity := tokens.IdentityReq{Password: &password, Methods: method, Totp: &totp}
	auth := tokens.AuthReq{Identity: &identity, Scope: &scope}
	opts := tokens.PwdTokenOptions{Auth: &auth}
	resp := tokens.CreateTokenByPassword(sc, opts, nocatalog)
	result, err := resp.ExtractTokenBody()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(bytes))

	fmt.Println("finish TestCreateTokenForProjectByPasswordAndMFA")
}

// 获取用户Token（使用密码+虚拟MFA）, 作用范围为整个账号
// Get the token by name/password and virtual MFA, and the function scope is the account
// POST /v3/auth/tokens
func CreateTokenForDomainByPasswordAndMFA(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestCreateTokenForDomainByPasswordAndMFA")
	userName := ""
	userId := ""
	userPwd := ""
	domainName := ""
	domainId := ""
	totpCode := ""
	nocatalog := ""
	var method []string
	method = append(append(method, "password"), "totp")
	user := tokens.UserReq{
		Name:     userName,
		Password: userPwd,
		Domain:   tokens.DomainReq{Name: domainName},
	}
	domain := tokens.DomainReq{
		ID:   domainId,
		Name: domainName,
	}
	totpUser := tokens.UserReq{
		ID:       userId,
		PassCode: totpCode,
	}
	totp := tokens.TotpReq{User: totpUser}
	scope := tokens.ScopeReq{Domain: &domain}
	password := tokens.PasswordReq{User: user}
	identity := tokens.IdentityReq{Password: &password, Methods: method, Totp: &totp}
	auth := tokens.AuthReq{Identity: &identity, Scope: &scope}
	opts := tokens.PwdTokenOptions{Auth: &auth}
	resp := tokens.CreateTokenByPassword(sc, opts, nocatalog)
	result, err := resp.ExtractTokenBody()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(bytes))

	fmt.Println("finish TestCreateTokenForDomainByPasswordAndMFA")
}

// 校验本账号中用户token的有效性
// Verify the validity of the user token in the account
// GET /v3/auth/tokens
func ValidateToken(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestValidateToken")
	token := ""
	nocatalog := ""
	resp := tokens.ValidateToken(sc, token, nocatalog)
	result, err := resp.ExtractTokenBody()
	if err != nil {
		fmt.Println("validate failed:", err)
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
		fmt.Println("validate successfully")
	}

	fmt.Println("finish TestValidateToken")
}

// 获取委托方的token
// Get the token of the agency
// POST /v3/auth/tokens
func CreateAgencyToken(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestCreateAgencyToken")
	nocatalog := ""
	assumeRole := tokens.AssumeRoleReq{
		DomainID:  "",
		XroleName: "",
	}
	var methods []string
	methods = append(methods, "assume_role")

	identity := tokens.IdentityReq{
		Methods:    methods,
		AssumeRole: &assumeRole,
	}
	domain := tokens.DomainReq{
		ID: "",
	}
	scope := tokens.ScopeReq{
		Domain: &domain,
	}
	auth := tokens.AuthReq{
		Identity: &identity,
		Scope:    &scope,
	}
	opts := tokens.AgencyTokenOptions{Auth: auth}
	resp := tokens.CreateTokenByAgency(sc, opts, nocatalog)
	result, err := resp.ExtractAgencyTokenBody()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(bytes))

	fmt.Println("finish TestCreateAgencyToken")
}
