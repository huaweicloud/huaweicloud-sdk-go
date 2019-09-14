package testing

import (
    "fmt"
    "testing"
    "net/http"
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/testhelper"
    "github.com/gophercloud/gophercloud/openstack/fgs/v2/function"
    "github.com/gophercloud/gophercloud/testhelper/client"
)

var url string

const functionUrn = "functionUrn"

func ServiceClient() *gophercloud.ServiceClient {
    sc := client.ServiceClient()
    sc.ResourceBase = sc.Endpoint + "v2/"
    return sc
}

func HandleFgsSuccessfully(t *testing.T, url, method, result string, head int) {
    testhelper.Mux.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
        testhelper.TestMethod(t, request, method)
        testhelper.TestHeader(t, request, "X-Auth-Token", client.TokenID)
        writer.Header().Add("Content-Type", "application/json")
        writer.WriteHeader(head)
        fmt.Fprintf(writer, result)
    })
}

func getMap() map[string]interface{} {
    m := make(map[string]interface{})
    m["message"] = "Hello world"
    return m
}

func TestCreate(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = "/v2/fgs/functions"
    HandleFgsSuccessfully(t, url, "POST", CreateFunction, http.StatusOK)
    funcCode := function.FunctionCodeOpts{
        File: "test",
    }
    opts := function.CreateOpts{
        FuncName:   "TestCreateFunctionInGoSdk",
        Package:    "default",
        CodeType:   "inline",
        Handler:    "index.handler",
        MemorySize: 128,
        Runtime:    FuncOne.Runtime,
        Timeout:    30,
        FuncCode:   funcCode,
    }
    res, err := function.Create(ServiceClient(), opts).Extract()
    testhelper.AssertNoErr(t, err)

    testhelper.CheckDeepEquals(t, res.Runtime, FuncOne.Runtime)
}

func TestList(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = "/v2/fgs/functions"
    HandleFgsSuccessfully(t, url, "GET", ListBody, http.StatusOK)

    listOpts := function.ListOpts{
        Marker:   "0",
        MaxItems: "40",
    }
    allPage, err := function.List(ServiceClient(), listOpts).AllPages()

    testhelper.AssertNoErr(t, err)
    functions, err := function.ExtractList(allPage)

    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, functions.Functions[0].Runtime, FuncOne.Runtime)
    testhelper.CheckDeepEquals(t, functions.Functions[0].Package, FuncOne.Package)
    testhelper.CheckDeepEquals(t, functions.Functions[0].Handler, FuncOne.Handler)
}

func TestGetMetadata(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s/config", functionUrn)
    HandleFgsSuccessfully(t, url, "GET", GetFunctionMetadata, http.StatusOK)

    funcMetadata, err := function.GetMetadata(ServiceClient(), functionUrn).Extract()
    testhelper.AssertNoErr(t, err)

    testhelper.CheckDeepEquals(t, funcMetadata.Handler, FuncOne.Handler)
}

func TestGetCode(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s/code", functionUrn)
    HandleFgsSuccessfully(t, url, "GET", GetFunctionCode, http.StatusOK)

    code, err := function.GetCode(ServiceClient(), functionUrn).Extract()
    testhelper.AssertNoErr(t, err)

    testhelper.CheckDeepEquals(t, code.Runtime, FuncOne.Runtime)
    testhelper.CheckDeepEquals(t, code.FuncUrn, FuncOne.FuncUrn)
}

func TestUpdateCode(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s/code", functionUrn)
    HandleFgsSuccessfully(t, url, "PUT", UpdateFunctionCode, http.StatusOK)
    opts := function.UpdateCodeOpts{
        CodeType: FuncOne.CodeType,
    }
    updateCode, err := function.UpdateCode(ServiceClient(), functionUrn, opts).Extract()
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, updateCode.Runtime, FuncOne.Runtime)
}

func TestUpdateMetadata(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s/config", functionUrn)
    HandleFgsSuccessfully(t, url, "PUT", UpdateFunctionMetadata, http.StatusOK)
    opts := function.UpdateMetadataOpts{
        Runtime:     FuncOne.Runtime,
        CodeType:    FuncOne.CodeType,
        MemorySize:  128,
        Handler:     "index.py",
        Timeout:     30,
        Description: "replace-you-description",
    }
    config, err := function.UpdateMetadata(ServiceClient(), functionUrn, opts).Extract()
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, config.Runtime, FuncOne.Runtime)
}

func TestCreateVersion(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s/versions", functionUrn)
    HandleFgsSuccessfully(t, url, "POST", CreateVersion, http.StatusOK)

    opts := function.CreateVersionOpts{
        Description: Version.Version,
        Version:     "",
    }
    createVersion, err := function.CreateVersion(ServiceClient(), opts, functionUrn).Extract()
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, createVersion.Version, Version.Version)
}

func TestVersionList(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s/versions", functionUrn)
    HandleFgsSuccessfully(t, url, "GET", VersionList, http.StatusOK)
    opts := function.ListOpts{
        Marker:   "0",
        MaxItems: "40",
    }
    allVersion, err := function.ListVersions(ServiceClient(), opts, functionUrn).AllPages()
    testhelper.AssertNoErr(t, err)
    version, err := function.ExtractVersionlist(allVersion)
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, version.Versions[0].Version, Version.Version)
}

func TestCreateAlias(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s/aliases", functionUrn)
    HandleFgsSuccessfully(t, url, "POST", VersionAlias, http.StatusOK)
    opts := function.CreateAliasOpts{
        Name:    Alias.Name,
        Version: Alias.Version,
    }
    versionAlias, err := function.CreateAlias(ServiceClient(), opts, functionUrn).ExtractAlias()
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, versionAlias.Version, Alias.Version)
    testhelper.CheckDeepEquals(t, versionAlias.Name, Alias.Name)
}

func TestUpdateAlias(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    aliasName := "v1"
    url = fmt.Sprintf("/v2/fgs/functions/%s/aliases/%s", functionUrn, aliasName)
    HandleFgsSuccessfully(t, url, "PUT", VersionAlias, http.StatusOK)
    opts := function.UpdateAliasOpts{
        Version:     Alias.Version,
        Description: "this is my version alias",
    }
    versionAlias, err := function.UpdateAlias(ServiceClient(), functionUrn, aliasName, opts).ExtractAlias()
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, versionAlias.Version, Alias.Version)
    testhelper.CheckDeepEquals(t, versionAlias.Name, Alias.Name)
}

func TestGetAlias(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    aliasName := "v1"
    url = fmt.Sprintf("/v2/fgs/functions/%s/aliases/%s", functionUrn, aliasName)
    HandleFgsSuccessfully(t, url, "GET", VersionAlias, http.StatusOK)
    versionAlias, err := function.GetAlias(ServiceClient(), functionUrn, aliasName).ExtractAlias()
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, versionAlias.Version, Alias.Version)
}

func TestListAlias(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s/aliases", functionUrn)
    HandleFgsSuccessfully(t, url, "GET", ListVersionAlias, http.StatusOK)
    allVersionAlias, err := function.ListAlias(ServiceClient(), functionUrn).AllPages()
    testhelper.AssertNoErr(t, err)
    alias, err := function.ExtractAliasList(allVersionAlias)
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, alias[0], Alias)
}

func TestDeleteAlias(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    aliasName := "v1"
    url = fmt.Sprintf("/v2/fgs/functions/%s/aliases/%s", functionUrn, aliasName)
    HandleFgsSuccessfully(t, url, "DELETE", "", http.StatusNoContent)

    err := function.DeleteAlias(ServiceClient(), functionUrn, aliasName).ExtractErr()
    testhelper.AssertNoErr(t, err)
}

func TestInvoke(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s/invocations", functionUrn)
    result := "\"Hello {\\r\\n  \\\"Message\\\" : \\\"Hello world\\\"\\r\\n}.\""
    HandleFgsSuccessfully(t, url, "POST", result, http.StatusOK)

    _, err := function.Invoke(ServiceClient(), getMap(), functionUrn).ExtractInvoke()
    testhelper.AssertNoErr(t, err)
}

func TestAsyncInvoke(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s/invocations-async", functionUrn)
    HandleFgsSuccessfully(t, url, "POST", asInvoke, http.StatusAccepted)

    _, err := function.AsyncInvoke(ServiceClient(), getMap(), functionUrn).ExtractInvoke()
    testhelper.AssertNoErr(t, err)
}

func TestDelete(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/functions/%s", functionUrn)
    HandleFgsSuccessfully(t, url, "DELETE", "", http.StatusNoContent)

    result := function.Delete(ServiceClient(), functionUrn)
    testhelper.AssertNoErr(t, result.Err)
}
