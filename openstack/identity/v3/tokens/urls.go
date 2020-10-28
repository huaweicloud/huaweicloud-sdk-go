package tokens

import "github.com/gophercloud/gophercloud"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}

func passwordTokenURL(c *gophercloud.ServiceClient, nocatalog string) string {
	if nocatalog == "" {
		return tokenURL(c)
	} else {
		return tokenURL(c) + "?" + "nocatalog=" + nocatalog
	}
}

func validateURL(c *gophercloud.ServiceClient, nocatalog string) string {
	if nocatalog == "" {
		return tokenURL(c)
	} else {
		return tokenURL(c) + "?" + "nocatalog=" + nocatalog
	}
}

func agencyTokenURL(c *gophercloud.ServiceClient, nocatalog string) string {
	if nocatalog == "" {
		return tokenURL(c)
	} else {
		return tokenURL(c) + "?" + "nocatalog=" + nocatalog
	}
}

