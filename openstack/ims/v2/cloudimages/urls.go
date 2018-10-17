package cloudimages

import (
	"net/url"

	"github.com/gophercloud/gophercloud"
)

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("cloudimages")
}

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("cloudimages/action")
}

func jobURL(c *gophercloud.ServiceClient, jobId string) string {
	return c.ServiceURL("jobs", jobId)
}

// builds next page full url based on current url
func nextPageURL(currentURL string, next string) (string, error) {
	base, err := url.Parse(currentURL)
	if err != nil {
		return "", err
	}
	rel, err := url.Parse(next)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(rel).String(), nil
}
