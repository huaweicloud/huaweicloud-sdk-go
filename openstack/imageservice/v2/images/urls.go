package images

import (
	"net/url"

	"github.com/gophercloud/gophercloud"
)

// `listURL` is a pure function. `listURL(c)` is a URL for which a GET
// request will respond with a list of images in the service `c`.
func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("images")
}

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("images")
}

// `imageURL(c,i)` is the URL for the image identified by ID `i` in
// the service `c`.
func imageURL(c *gophercloud.ServiceClient, imageID string) string {
	return c.ServiceURL("images", imageID)
}

// `getURL(c,i)` is a URL for which a GET request will respond with
// information about the image identified by ID `i` in the service
// `c`.
func getURL(c *gophercloud.ServiceClient, imageID string) string {
	return imageURL(c, imageID)
}

func updateURL(c *gophercloud.ServiceClient, imageID string) string {
	return imageURL(c, imageID)
}

func deleteURL(c *gophercloud.ServiceClient, imageID string) string {
	return imageURL(c, imageID)
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

// putTagURL add a tag to the image
func putTagURL(c *gophercloud.ServiceClient, imageID, tag string) string {
	return c.ServiceURL("images", imageID, "tags", tag)
}

// deleteTagURL remove the tag from a image
func deleteTagURL(c *gophercloud.ServiceClient, imageID, tag string) string {
	return c.ServiceURL("images", imageID, "tags", tag)
}

// getImageSchemas generate url to get image schemas
func getImageSchemas(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("schemas", "image")
}

// getImagesSchemas generate url to get images schemas
func getImagesSchemas(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("schemas", "images")
}
