package cloudimages

import "github.com/gophercloud/gophercloud"

// getImageTagsURL generate a url for getting tags of an image
func getImageTagsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("cloudimages", "tags")
}

// putImageTagsURL generate a url for adding tag to an image
func putImageTagsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("cloudimages", "tags")
}

// importImageURL generate a url for registering an image
func importImageURL(c *gophercloud.ServiceClient, imageID string) string {
	return c.ServiceURL("cloudimages", imageID, "upload")
}

// exportImageURL generate a url for download an image to OBS
func exportImageURL(c *gophercloud.ServiceClient, imageID string) string {
	return c.ServiceURL("cloudimages", imageID, "file")
}

// copyImageURL generate a url for copy image between regions
func copyImageURL(c *gophercloud.ServiceClient, imageID string) string {
	return c.ServiceURL("cloudimages", imageID, "copy")
}

// imageMemberOpURL generate a URL for adding, updating or removing members to
// or from an image
func imageMemberOpURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("cloudimages", "members")
}

// getCloudImagesQuota generate a URL for query the quota of cloud images
func getCloudImagesQuota(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("cloudimages", "quota")
}
