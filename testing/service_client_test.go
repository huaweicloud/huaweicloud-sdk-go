package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud"
	th "github.com/gophercloud/gophercloud/testhelper"
	"encoding/json"
	"bytes"
)

func TestServiceURL(t *testing.T) {
	c := &gophercloud.ServiceClient{Endpoint: "http://123.45.67.8/"}
	expected := "http://123.45.67.8/more/parts/here"
	actual := c.ServiceURL("more", "parts", "here")
	th.CheckEquals(t, expected, actual)
}

func TestMoreHeaders(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	c := new(gophercloud.ServiceClient)
	c.MoreHeaders = map[string]string{
		"custom": "header",
	}
	c.ProviderClient = new(gophercloud.ProviderClient)
	resp, err := c.Get(fmt.Sprintf("%s/route", th.Endpoint()), nil, nil)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, resp.Request.Header.Get("custom"), "header")
}

func TestRequestWithTwoBody(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	c := new(gophercloud.ServiceClient)

	c.ProviderClient = new(gophercloud.ProviderClient)

	expected := "please provide only one of JSONBody or RawBody to gophercloud.Request()"
	rawBody := bytes.NewReader([]byte("test1body"))
	jsonBody, err := json.Marshal("test2body")
	th.AssertNoErr(t, err)

	_, err = c.Post(fmt.Sprintf("%s/route", th.Endpoint()), jsonBody, nil, &gophercloud.RequestOpts{
		RawBody: rawBody,
	})
	th.CheckEquals(t, expected, err.Error())
}
