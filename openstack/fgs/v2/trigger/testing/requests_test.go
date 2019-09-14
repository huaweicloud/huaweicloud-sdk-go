package testing

import (
    "fmt"
    "testing"
    "net/http"
    "github.com/gophercloud/gophercloud"
    "github.com/gophercloud/gophercloud/testhelper"
    "github.com/gophercloud/gophercloud/openstack/fgs/v2/trigger"
    "github.com/gophercloud/gophercloud/testhelper/client"
)

const functionUrn = "functionUrn"

var url = ""

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

func TestCreate(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/triggers/%s", functionUrn)
    HandleFgsSuccessfully(t, url, "POST", Trigger, http.StatusCreated)
    opts := trigger.CreateOpts{
        TriggerTypeCode: TriggerTwo.TriggerTypeCode,
        EventTypeCode:   TriggerTwo.EventTypeCode,
        EventData:       TriggerTwo.EventData,
    }
    createTrigger, err := trigger.Create(ServiceClient(), opts, functionUrn).Extract()
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, createTrigger.EventData, TriggerTwo.EventData)
    testhelper.CheckDeepEquals(t, createTrigger.TriggerTypeCode, TriggerTwo.TriggerTypeCode)
    testhelper.CheckDeepEquals(t, createTrigger.EventTypeCode, TriggerTwo.EventTypeCode)
}

func TestList(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/triggers/%s", functionUrn)
    HandleFgsSuccessfully(t, url, "GET", Triggers, http.StatusOK)
    allTriggers, err := trigger.List(ServiceClient(), functionUrn).AllPages()
    testhelper.AssertNoErr(t, err)
    triggers, err := trigger.ExtractList(allTriggers)
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, triggers[0], TriggerOne)
}

func TestDeleteAll(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    url = fmt.Sprintf("/v2/fgs/triggers/%s", functionUrn)
    HandleFgsSuccessfully(t, url, "DELETE", "", http.StatusNoContent)
    result := trigger.DeleteAll(ServiceClient(), functionUrn)
    testhelper.AssertNoErr(t, result.Err)
}

func TestGet(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    var typeCode, triggerId string
    typeCode = TriggerTwo.TriggerTypeCode
    triggerId = TriggerTwo.TriggerId
    url = fmt.Sprintf("/v2/fgs/triggers/%s/%s/%s", functionUrn, typeCode, triggerId)
    HandleFgsSuccessfully(t, url, "GET", Trigger, http.StatusOK)
    resp, err := trigger.Get(ServiceClient(), functionUrn, typeCode, triggerId).Extract()
    testhelper.AssertNoErr(t, err)
    testhelper.CheckDeepEquals(t, resp.TriggerId, TriggerTwo.TriggerId)
}

func TestDeleteTrigger(t *testing.T) {
    testhelper.SetupHTTP()
    defer testhelper.TeardownHTTP()
    var typeCode, triggerId string
    typeCode = TriggerTwo.TriggerTypeCode
    triggerId = TriggerTwo.TriggerId
    url = fmt.Sprintf("/v2/fgs/triggers/%s/%s/%s", functionUrn, typeCode, triggerId)
    HandleFgsSuccessfully(t, url, "DELETE", "", http.StatusNoContent)
    result := trigger.Delete(ServiceClient(), functionUrn, typeCode, triggerId)
    testhelper.AssertNoErr(t, result.Err)
}
