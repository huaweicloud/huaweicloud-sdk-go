package whitelist

import "github.com/gophercloud/gophercloud"

const (
	ROOTPATH     = "lbaas"
	RESOURCEPATH = "whitelists"
)
//GET list and post url
func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(ROOTPATH, RESOURCEPATH)
}
//GET details put delete url
func resourceURL(c *gophercloud.ServiceClient, whitelistId string) string {
	return c.ServiceURL(ROOTPATH, RESOURCEPATH, whitelistId)
}
