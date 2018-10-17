package auth

import (
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
)

var nilOptions = token.TokenOptions{}

/*
AuthOptionsFromEnv fills out an identity.AuthOptions structure with the
settings found on the various OpenStack OS_* environment variables.

The following variables provide sources of truth: OS_AUTH_URL, OS_USERNAME,
OS_PASSWORD, OS_TENANT_ID, and OS_TENANT_NAME.

Of these, OS_USERNAME, OS_PASSWORD, and OS_AUTH_URL must have settings,
or an error will result.  OS_TENANT_ID, OS_TENANT_NAME, OS_PROJECT_ID, and
OS_PROJECT_NAME are optional.

OS_TENANT_ID and OS_TENANT_NAME are mutually exclusive to OS_PROJECT_ID and
OS_PROJECT_NAME. If OS_PROJECT_ID and OS_PROJECT_NAME are set, they will
still be referred as "tenant" in Gophercloud.

To use this function, first set the OS_* environment variables (for example,
by sourcing an `openrc` file), then:

	opts, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.AuthenticatedClient(opts)
*/
func TokenOptionsFromEnv() (token.TokenOptions, error) {
	authURL := os.Getenv("OS_AUTH_URL")
	username := os.Getenv("OS_USERNAME")
	userID := os.Getenv("OS_USERID")
	password := os.Getenv("OS_PASSWORD")
	tenantID := os.Getenv("OS_TENANT_ID")
	tenantName := os.Getenv("OS_TENANT_NAME")
	domainID := os.Getenv("OS_DOMAIN_ID")
	domainName := os.Getenv("OS_DOMAIN_NAME")

	// If OS_PROJECT_ID is set, overwrite tenantID with the value.
	if v := os.Getenv("OS_PROJECT_ID"); v != "" {
		tenantID = v
	}

	// If OS_PROJECT_NAME is set, overwrite tenantName with the value.
	if v := os.Getenv("OS_PROJECT_NAME"); v != "" {
		tenantName = v
	}

	if authURL == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "authURL")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		return nilOptions, err
		//err := gophercloud.ErrMissingInput{Argument: "authURL"}
		//return nilOptions, err
	}

	if username == "" && userID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "username")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		return nilOptions, err
		//err := gophercloud.ErrMissingInput{Argument: "username"}
		//return nilOptions, err
	}

	if password == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "password")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		return nilOptions, err
		//err := gophercloud.ErrMissingInput{Argument: "password"}
		//return nilOptions, err
	}

	to := token.TokenOptions{
		IdentityEndpoint: authURL,
		UserID:           userID,
		Username:         username,
		Password:         password,
		TenantID:         tenantID,
		TenantName:       tenantName,
		DomainID:         domainID,
		DomainName:       domainName,
	}

	return to, nil
}
