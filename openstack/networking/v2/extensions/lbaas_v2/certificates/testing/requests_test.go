package testing

import (
	"testing"

	fake "github.com/gophercloud/gophercloud/openstack/networking/v2/common"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/certificates"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestListCertificate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCertificateALLSuccessfully(t)

	pages := 0
	err := certificates.List(fake.ServiceClient(), certificates.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := certificates.ExtractCertificates(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 certificates, got %d", len(actual))
		}
		th.CheckDeepEquals(t, CertificateOne, actual[0])
		th.CheckDeepEquals(t, CertificateTwo, actual[1])
		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllCertificates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCertificateALLSuccessfully(t)

	allPages, err := certificates.List(fake.ServiceClient(), certificates.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := certificates.ExtractCertificates(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, CertificateOne, actual[0])
	th.CheckDeepEquals(t, CertificateTwo, actual[1])
}

//
func TestCreateCertificate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCertificateCreationSuccessfully(t)

	actual, err := certificates.Create(fake.ServiceClient(), certificates.CreateOpts{
		Certificate: "-----BEGIN CERTIFICATE-----\r\nMIIDpTCCAo2gAwIBAgIJAKdmmOBYnFvoMA0GCSqGSIb3DQEBCwUAMGkxCzAJBgNV\r\nBAYTAnh4MQswCQYDVQQIDAJ4eDELMAkGA1UEBwwCeHgxCzAJBgNVBAoMAnh4MQsw\r\nCQYDVQQLDAJ4eDELMAkGA1UEAwwCeHgxGTAXBgkqhkiG9w0BCQEWCnh4QDE2My5j\r\nb20wHhcNMTcxMjA0MDM0MjQ5WhcNMjAxMjAzMDM0MjQ5WjBpMQswCQYDVQQGEwJ4\r\neDELMAkGA1UECAwCeHgxCzAJBgNVBAcMAnh4MQswCQYDVQQKDAJ4eDELMAkGA1UE\r\nCwwCeHgxCzAJBgNVBAMMAnh4MRkwFwYJKoZIhvcNAQkBFgp4eEAxNjMuY29tMIIB\r\nIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwZ5UJULAjWr7p6FVwGRQRjFN\r\n2s8tZ/6LC3X82fajpVsYqF1xqEuUDndDXVD09E4u83MS6HO6a3bIVQDp6/klnYld\r\niE6Vp8HH5BSKaCWKVg8lGWg1UM9wZFnlryi14KgmpIFmcu9nA8yV/6MZAe6RSDmb\r\n3iyNBmiZ8aZhGw2pI1YwR+15MVqFFGB+7ExkziROi7L8CFCyCezK2/oOOvQsH1dz\r\nQ8z1JXWdg8/9Zx7Ktvgwu5PQM3cJtSHX6iBPOkMU8Z8TugLlTqQXKZOEgwajwvQ5\r\nmf2DPkVgM08XAgaLJcLigwD513koAdtJd5v+9irw+5LAuO3JclqwTvwy7u/YwwID\r\nAQABo1AwTjAdBgNVHQ4EFgQUo5A2tIu+bcUfvGTD7wmEkhXKFjcwHwYDVR0jBBgw\r\nFoAUo5A2tIu+bcUfvGTD7wmEkhXKFjcwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0B\r\nAQsFAAOCAQEAWJ2rS6Mvlqk3GfEpboezx2J3X7l1z8Sxoqg6ntwB+rezvK3mc9H0\r\n83qcVeUcoH+0A0lSHyFN4FvRQL6X1hEheHarYwJK4agb231vb5erasuGO463eYEG\r\nr4SfTuOm7SyiV2xxbaBKrXJtpBp4WLL/s+LF+nklKjaOxkmxUX0sM4CTA7uFJypY\r\nc8Tdr8lDDNqoUtMD8BrUCJi+7lmMXRcC3Qi3oZJW76ja+kZA5mKVFPd1ATih8TbA\r\ni34R7EQDtFeiSvBdeKRsPp8c0KT8H1B4lXNkkCQs2WX5p4lm99+ZtLD4glw8x6Ic\r\ni1YhgnQbn5E0hz55OLu5jvOkKQjPCW+8Kg==\r\n-----END CERTIFICATE-----",
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, CertificateOne, *actual)
}

func TestGetCertificate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCertificateGetSuccessfully(t)

	client := fake.ServiceClient()
	actual, err := certificates.Get(client, "eabfefa3fd1740a88a47ad98e132d238").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, CertificateOne, *actual)
}

func TestDeleteCertificate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCertificateDeletionSuccessfully(t)

	res := certificates.Delete(fake.ServiceClient(), "eabfefa3fd1740a88a47ad98e132d238")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateCertificate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCertificateUpdateSuccessfully(t)

	client := fake.ServiceClient()
	actual, err := certificates.Update(client, "eabfefa3fd1740a88a47ad98e132d238", certificates.UpdateOpts{
		Domain: "testdomain",
	}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}

	th.CheckDeepEquals(t, CertificateOne, *actual)
}
