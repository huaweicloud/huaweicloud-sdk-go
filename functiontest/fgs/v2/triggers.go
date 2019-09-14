package main

import (
    "fmt"
    "github.com/gophercloud/gophercloud/auth/token"
    "github.com/gophercloud/gophercloud/openstack"
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/openstack/fgs/v2/function"
    "github.com/gophercloud/gophercloud/openstack/fgs/v2/trigger"
    "github.com/gophercloud/gophercloud/functiontest/fgs/Common"
    "strings"
)

func main() {
    //gophercloud.EnableDebug = true
    //Set authentication parameters, use Username、Password、DomainID、ProjectID
    tokenOpts := token.TokenOptions{
        IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
        Username:         "your username",
        Password:         "your password",
        DomainID:         "your domainId",
        ProjectID:        "your projectID",
        AllowReauth:      true,
    }

    //init provider client
    provider, err := openstack.AuthenticatedClient(tokenOpts)
    if Common.CheckErr(err) {
        return
    }

    //init service client
    sc, err := openstack.NewFGSV2(provider, gophercloud.EndpointOpts{})
    if Common.CheckErr(err) {
        return
    }

    //Creating a Function
    res := CreateFunction(sc)
    if res.FuncName != "" {
        //Creating a Trigger
        trg := TestTriggerCreate(sc, res.FuncUrn)
        //Querying the Information About a Trigger
        TestTriggerGet(sc, res.FuncUrn, trg.TriggerTypeCode, trg.TriggerId)
        //Querying All Triggers of a Function
        TestTriggerList(sc, res.FuncUrn)
        //Deleting a Trigger
        TestTriggerDelete(sc, res.FuncUrn, trg.TriggerTypeCode, trg.TriggerId)
        //Creating a Trigger
        TestTriggerCreate(sc, res.FuncUrn)
        //Deleting All Triggers of a Function
        TestAllTriggersDelete(sc, res.FuncUrn)
        //Deleting a Function
        FunctionDelete(sc, res.FuncUrn)
    }
}

func CreateFunction(sc *gophercloud.ServiceClient) (*function.Function) {
    codeOpt := function.FunctionCodeOpts{
        File: "aW1wb3J0IGpzb24KZGVmIGhhbmRsZXIgKGV2ZW50LCBjb250ZXh0KToKICAgIG91dHB1dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganNvbi5kdW1wcyhldmVudCkKICAgIHJldHVybiBvdXRwdXQ=",
        Link: "",
    }
    opts := function.CreateOpts{
        FuncName:   "TestCreateTriggerInGoSdk",
        Package:    "default",
        CodeType:   "inline",
        Handler:    "index.handler",
        MemorySize: 128,
        Runtime:    "Python2.7",
        Timeout:    30,
        FuncCode:   codeOpt,
    }
    resp, err := function.Create(sc, opts).Extract()
    if Common.CheckErr(err) {
        return resp
    }
    fmt.Println("Create function  success!")
    Common.Show(resp)
    return resp
}

func TestTriggerList(sc *gophercloud.ServiceClient, funcUrn string) {
    allTriggers, err := trigger.List(sc, funcUrn).AllPages()
    if Common.CheckErr(err) {
        return
    }
    triggers, err := trigger.ExtractList(allTriggers)
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Get trigger list success!")
    Common.Show(triggers)
}

func TestTriggerCreate(sc *gophercloud.ServiceClient, funcUrn string) (*trigger.Trigger) {
    type apigTriggereSpecFuncInfo struct {
        Timeout int `json:"timeout"          validate:"min=1,max=60000"`
    }
    eventData := make(map[string]interface{})
    eventData["group_id"] = "f1eb9c97a2d64221b874b430cd06647b"
    eventData["auth"] = "IAM"
    eventData["backend_type"] = "FUNCTION"
    eventData["env_id"] = "DEFAULT_ENVIRONMENT_RELEASE_ID"
    eventData["env_name"] = "RELEASE"
    eventData["func_info"] = apigTriggereSpecFuncInfo{Timeout: 5000}
    eventData["match_mode"] = "SWA"
    eventData["name"] = "TestFunctionInGoSdk"
    eventData["path"] = "/TestFunctionInGoSdk"
    eventData["protocol"] = "HTTPS"
    eventData["req_method"] = "ANY"
    eventData["sl_domain"] = "6f53d3cf3b804cc3a1ea312c36bbc8b5.apigw.southchina.huaweicloud.com"
    eventData["type"] = 1

    opts := trigger.CreateOpts{
        TriggerTypeCode: "APIG",
        EventTypeCode:   "API_test",
        EventData:       eventData,
    }
    createTrigger, err := trigger.Create(sc, opts, funcUrn).Extract()
    if Common.CheckErr(err) {
        return createTrigger
    }
    fmt.Println("Create trigger success!")
    Common.Show(createTrigger)
    return createTrigger
}

func TestAllTriggersDelete(sc *gophercloud.ServiceClient, funcUrn string) {
    err := trigger.DeleteAll(sc, funcUrn).ExtractErr()
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Test delete all trigger success")
}

func TestTriggerGet(sc *gophercloud.ServiceClient, funcUrn string, triggerTypeCode string, triggerId string) {
    resp, err := trigger.Get(sc, funcUrn, triggerTypeCode, triggerId).Extract()
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Get trigger success!")
    Common.Show(resp)
}

func TestTriggerDelete(sc *gophercloud.ServiceClient, funcUrn string, triggerTypeCode string, triggerId string) {
    err := trigger.Delete(sc, funcUrn, triggerTypeCode, triggerId).ExtractErr()
    if Common.CheckErr(err) {
        return
    }
    fmt.Printf("Test delete trigger:%s success \n", triggerId)
}

func FunctionDelete(sc *gophercloud.ServiceClient, funcUrn string) {
    funcUrnTmp := funcUrn
    index := strings.LastIndex(funcUrnTmp, ":")
    if index > 0 {
        funcUrnTmp = string([]byte(funcUrnTmp)[:index])
    }

    err := function.Delete(sc, funcUrnTmp).ExtractErr()
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Test delete function success")
}
