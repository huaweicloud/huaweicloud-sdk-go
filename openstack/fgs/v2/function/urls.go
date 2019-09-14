package function

import "github.com/gophercloud/gophercloud"

const (
    FGS      = "fgs"
    FUNCTION = "functions"
    CODE     = "code"
    CONFIG   = "config"
    VERSION  = "versions"
    ALIAS    = "aliases"
    INVOKE   = "invocations"
    ASINVOKE = "invocations-async"
)

func createURL(c *gophercloud.ServiceClient) string {
    return listURL(c)
}

func listURL(c *gophercloud.ServiceClient) string {
    return c.ServiceURL(FGS, FUNCTION)
}

func deleteURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return c.ServiceURL(FGS, FUNCTION, functionUrn)
}

//function code
func getCodeURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return c.ServiceURL(FGS, FUNCTION, functionUrn, CODE)
}

func updateCodeURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return getCodeURL(c, functionUrn)
}

//function metadata
func getMetadataURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return c.ServiceURL(FGS, FUNCTION, functionUrn, CONFIG)
}

func updateMetadataURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return getMetadataURL(c, functionUrn)
}

//function invoke
func invokeURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return c.ServiceURL(FGS, FUNCTION, functionUrn, INVOKE)
}

func asyncInvokeURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return c.ServiceURL(FGS, FUNCTION, functionUrn, ASINVOKE)
}

//function version
func createVersionURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return c.ServiceURL(FGS, FUNCTION, functionUrn, VERSION)
}

func listVersionURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return createVersionURL(c, functionUrn)
}

//function alias
func createAliasURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return c.ServiceURL(FGS, FUNCTION, functionUrn, ALIAS)
}

func updateAliasURL(c *gophercloud.ServiceClient, functionUrn, aliasName string) string {
    return c.ServiceURL(FGS, FUNCTION, functionUrn, ALIAS, aliasName)
}

func deleteAliasURL(c *gophercloud.ServiceClient, functionUrn, aliasName string) string {
    return updateAliasURL(c, functionUrn, aliasName)
}

func getAliasURL(c *gophercloud.ServiceClient, functionUrn, aliasName string) string {
    return updateAliasURL(c, functionUrn, aliasName)
}

func listAliasURL(c *gophercloud.ServiceClient, functionUrn string) string {
    return createAliasURL(c, functionUrn)
}
