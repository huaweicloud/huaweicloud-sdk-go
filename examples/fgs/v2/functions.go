package main

import (
    "fmt"
    //"github.com/gophercloud/gophercloud/auth/token"
    "github.com/gophercloud/gophercloud/auth/aksk"
    "github.com/gophercloud/gophercloud/openstack"
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/openstack/fgs/v2/function"
    "github.com/gophercloud/gophercloud/functiontest/fgs/Common"
    "strings"
)

var (
    VERSION = "1.0.1"
    NAME    = "testName"
)

func main() {
    gophercloud.EnableDebug = true
    //Set authentication parameters, use AKSK
    tokenOpts := aksk.AKSKOptions{
        IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
        ProjectID:        "{ProjectID}",
        AccessKey:        "your AK string",
        SecretKey:        "your SK string",
        Cloud:            "yyy.com",
        Region:           "xxx",
        DomainID:         "{domainID}",
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

    //Querying a Function List
    TestFunctionList(sc)
    //Creating a Function
    res := TestFunctionCreate(sc)
    if res.FuncName != "" {
        //Executing a Function Synchronously
        TestInvoke(sc, res.FuncName)
        //Modifying the Metadata Information of a Function
        TestFunctionMetadataUpdate(sc, res.FuncName)
        //Querying the Metadata Information of a Function
        TestFunctionMetadataGet(sc, res.FuncName)
        //Modifying the Code of a Function
        TestFunctionCodeUpdate(sc, res.FuncName)
        //Querying the Code of a Function
        TestFunctionCodeGet(sc, res.FuncName)
        //Publishing a Function Version
        TestVersionCreate(sc, res.FuncName)
        //Querying the Aliases of a Function's All Versions
        TestVersionList(sc, res.FuncName)
        //Creating an Alias for a Function Version
        alias := TestAliasCreate(sc, res.FuncName)
        //Querying the Alias Information of a Function Version
        TestAliasGet(sc, res.FuncName, alias.Name)
        //Modifying the Alias Information of a Function Version
        TestAliasUpdate(sc, res.FuncName, alias.Name)
        //Querying the Aliases of a Function's All Versions
        TestAliasList(sc, res.FuncName)
        //Executing a Function Asynchronously
        TestAsyncInvoke(sc, res.FuncName)
        //Deleting an Alias of a Function Version
        TestAliasDelete(sc, res.FuncName, alias.Name)
        //Querying the Alias Information of a Function Version
        TestAliasGet(sc, res.FuncName, alias.Name)
        //Deleting a Function or Function Version
        TestFunctionDelete(sc, res.FuncName)
    }
}

func TestVersionList(sc *gophercloud.ServiceClient, funcUrn string) {
    listOpt := function.ListOpts{
        Marker:   "0",
        MaxItems: "40",
    }
    allPages, err := function.ListVersions(sc, listOpt, funcUrn).AllPages()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    //Parse the return value
    allFunctions, err := function.ExtractVersionlist(allPages)
    if Common.CheckErr(err) {
        return
    }

    fmt.Println("Get version list success!")
    Common.Show(allFunctions)
}

func TestVersionCreate(sc *gophercloud.ServiceClient, funcUrn string) {
    opts := function.CreateVersionOpts{
        Description: "test1.0.1↵函数服务发布版本1.0.1",
        Version:     VERSION,
    }
    resp := function.CreateVersion(sc, opts, funcUrn)
    if Common.CheckErr(resp.Err) {
        return
    }
    fmt.Println("Get function metadata success!")
    Common.Show(resp)
}

func TestFunctionList(sc *gophercloud.ServiceClient) {
    listOpt := function.ListOpts{
        Marker:   "0",
        MaxItems: "40",
    }
    allPages, err := function.List(sc, listOpt).AllPages()
    if Common.CheckErr(err) {
        return
    }
    //Parse the return value
    allFunctions, err := function.ExtractList(allPages)
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Get function list success!")
    Common.Show(allFunctions)
}

func TestFunctionMetadataGet(sc *gophercloud.ServiceClient, funcUrn string) {
    resp, err := function.GetMetadata(sc, funcUrn).Extract()
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Get function metadata success!")
    Common.Show(resp)
}

func TestFunctionCodeGet(sc *gophercloud.ServiceClient, funcUrn string) {
    resp, err := function.GetCode(sc, funcUrn).Extract()
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Get function code success!")
    Common.Show(resp)
}

func TestFunctionCreate(sc *gophercloud.ServiceClient) (*function.Function) {
    //Build a function to be created
    codeOpt := function.FunctionCodeOpts{
        File: "aW1wb3J0IGpzb24KZGVmIGhhbmRsZXIgKGV2ZW50LCBjb250ZXh0KToKICAgIG91dHB1dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganNvbi5kdW1wcyhldmVudCkKICAgIHJldHVybiBvdXRwdXQ=",
        Link: "test",
    }
    opts := function.CreateOpts{
        FuncName:   "TestCreateFunctionInGoSdk",
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

func TestFunctionDelete(sc *gophercloud.ServiceClient, funcUrn string) {
    //First remove the version number
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

func TestFunctionCodeUpdate(sc *gophercloud.ServiceClient, funcUrn string) {
    //Build a function to be created
    codeOpt := function.FunctionCodeOpts{
        File: "UEsDBAoAAAAAAO5Qak5BtFZEugAAALoAAAAIAAAAaW5kZXgucHlpbXBvcnQganNvbgpkZWYgaGFuZGxlciAoZXZlbnQsIGNvbnRleHQpOgogICAgb3V0cHV0ID0gJ0hlbGxvIG1lc3NhZ2U6ICcgKyBqc29uLmR1bXBzKGV2ZW50KQogICAgcHJpbnQoIjIzMjM0MjQzNDM0MzQzNDM0MzQzNTQ1MzQzPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09IikKICAgIHJldHVybiBvdXRwdXRQSwECHgMKAAAAAADuUGpOQbRWRLoAAAC6AAAACAAAAAAAAAAAAAAAtIEAAAAAaW5kZXgucHlQSwUGAAAAAAEAAQA2AAAA4AAAAAAA=",
        Link: "",
    }
    opts := function.UpdateCodeOpts{
        CodeType: "inline",
        FuncCode: codeOpt,
    }
    functionCode, err := function.UpdateCode(sc, funcUrn, opts).Extract()
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Update function code success!")
    Common.Show(functionCode)
}

func TestFunctionMetadataUpdate(sc *gophercloud.ServiceClient, funcUrn string) {
    opts := function.UpdateMetadataOpts{
        Runtime:     "Python2.7",
        CodeType:    "inline",
        MemorySize:  128,
        Handler:     "index.py",
        Timeout:     30,
        Description: "test go sdk function",
    }

    functionMetadata, err := function.UpdateMetadata(sc, funcUrn, opts).Extract()
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Update function metadata success!")
    Common.Show(functionMetadata)
}

func TestInvoke(sc *gophercloud.ServiceClient, funcUrn string) {
    m := make(map[string]interface{})
    m["message"] = "hello world!"
    r, err := function.Invoke(sc, m, funcUrn).ExtractInvoke()
    if Common.CheckErr(err) {
        return
    }
    Common.Show(r)
}

func TestAsyncInvoke(sc *gophercloud.ServiceClient, funcUrn string) {
    m := make(map[string]interface{})
    m["message"] = "hello world!"
    r, err := function.AsyncInvoke(sc, m, funcUrn).ExtractInvoke()
    if Common.CheckErr(err) {
        return
    }
    Common.Show(r)
}

func TestAliasCreate(sc *gophercloud.ServiceClient, funcUrn string) *function.AliasResult {
    opts := function.CreateAliasOpts{
        Name:    NAME,
        Version: VERSION,
    }
    versionAlias, err := function.CreateAlias(sc, opts, funcUrn).ExtractAlias()
    if Common.CheckErr(err) {
        return versionAlias
    }
    fmt.Printf("Create version alias:%s success! \n", opts.Name)
    Common.Show(versionAlias)
    return versionAlias
}

func TestAliasDelete(sc *gophercloud.ServiceClient, funcUrn string, aliasName string) {
    err := function.DeleteAlias(sc, funcUrn, aliasName).ExtractErr()
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Test delete alias success")
    return
}

func TestAliasGet(sc *gophercloud.ServiceClient, funcUrn string, aliasName string) {
    versionAlias, err := function.GetAlias(sc, funcUrn, aliasName).ExtractAlias()
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Get version alias success!")
    Common.Show(versionAlias)
}

func TestAliasUpdate(sc *gophercloud.ServiceClient, funcUrn string, aliasName string) {
    opts := function.UpdateAliasOpts{
        Version:     VERSION,
        Description: "this is my version alias",
    }
    versionAlias, err := function.UpdateAlias(sc, funcUrn, aliasName, opts).ExtractAlias()
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Update version alias success!")
    Common.Show(versionAlias)
}

func TestAliasList(sc *gophercloud.ServiceClient, funcUrn string) {
    allVersionAlias, err := function.ListAlias(sc, funcUrn).AllPages()
    if Common.CheckErr(err) {
        return
    }
    alias, err := function.ExtractAliasList(allVersionAlias)
    if Common.CheckErr(err) {
        return
    }
    fmt.Println("Get version alias list success!")
    Common.Show(alias)
}
