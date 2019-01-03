package openstack

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth"
	akskAuth "github.com/gophercloud/gophercloud/auth/aksk"
	tokenAuth "github.com/gophercloud/gophercloud/auth/token"
	tokens2 "github.com/gophercloud/gophercloud/openstack/identity/v2/tokens"
	tokens3 "github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/gophercloud/gophercloud/openstack/utils"

	"github.com/gophercloud/gophercloud/openstack/identity/v3/endpoints"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/services"
	"github.com/gophercloud/gophercloud/pagination"
	"time"

)

const (
	// v2 represents Keystone v2.
	// It should never increase beyond 2.0.
	v2 = "v2.0"

	// v3 represents Keystone v3.
	// The version can be anything from v3 to v3.x.
	v3 = "v3"
)

const (
	//compute v2 microVersion
	computeMicroVersion       = "2.26"
	objectSstoreMicroVersion  = ""
	networkMicroVersion       = ""
	vpcv1MicroVersion         = ""
	vpcv2MicroVersion         = ""
	volumeMicroVersion        = ""
	volumev2MicroVersion      = ""
	volumev3MicroVersion      = ""
	sharev2MicroVersion       = ""
	cdnMicroVersion           = ""
	orchestrationMicroVersion = ""
	databaseMicroVersion      = ""
	dnsMicroVersion           = ""
	imageMicroVersion         = ""
	loadBalancerMicroVersion  = ""
	ecsMicroVersion           = ""
	ecsv1_1MicroVersion       = ""
	ecsv2MicroVersion         = ""
	imageSelfDevMicroVersion  = ""
	bssMicroVersion           = ""
	cesMicroVersion           = ""
)

/* 重写RoundTrip，实现reauth 限制3次 */
type MyRoundTripper struct {
	rt                http.RoundTripper
	numReauthAttempts int
}

func newHTTPClient() http.Client {
	return http.Client{
		Transport: &MyRoundTripper{
			rt: http.DefaultTransport,
		},
	}
}

func (mrt *MyRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := mrt.rt.RoundTrip(request)
	if response == nil {
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		if mrt.numReauthAttempts == 3 {
			return response, fmt.Errorf("Tried to re-authenticate 3 times with no success.")
		}
		mrt.numReauthAttempts++
	}

	return response, err
}

/*
func NewProviderClientWithOptions(options auth.AuthOptionsProvider, conf *gophercloud.Config) (*gophercloud.ProviderClient, error) {
	client, err := NewClient(options.GetIdentityEndpoint(), options.GetProjectId(), conf)
	if err != nil {
		return nil, err
	}

	err = Authenticate(client, options)
	if err != nil {
		return nil, err
	}
	return client, nil
}
*/

/*
AuthenticatedClient logs in to an OpenStack cloud found at the identity endpoint
specified by the options, acquires a token, and returns a Provider Client
instance that's ready to operate.

If the full path to a versioned identity endpoint was specified  (example:
http://example.com:5000/v3), that path will be used as the endpoint to query.

If a versionless endpoint was specified (example: http://example.com:5000/),
the endpoint will be queried to determine which versions of the identity service
are available, then chooses the most recent or most supported version.

Example:

	ao, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.AuthenticatedClient(ao)
	client, err := openstack.NewNetworkV2(client, EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
*/
func AuthenticatedClient(options auth.AuthOptionsProvider) (*gophercloud.ProviderClient, error) {
	client, err := NewClient(options.GetIdentityEndpoint(), options.GetDomainId(), options.GetProjectId(), gophercloud.NewConfig())
	if err != nil {
		return nil, err
	}

	err = Authenticate(client, options)
	if err != nil {
		return nil, err
	}
	return client, nil
}

/*
NewClient prepares an unauthenticated ProviderClient instance.
Most users will probably prefer using the AuthenticatedClient function
instead.

This is useful if you wish to explicitly control the version of the identity
service that's used for authentication explicitly, for example.

A basic example of using this would be:

	ao, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.NewClient(ao.IdentityEndpoint)
	client, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
*/
func NewClient(endpoint, domainID, projectID string, conf *gophercloud.Config) (*gophercloud.ProviderClient, error) {
	if endpoint == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "IdentityEndpoint")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		return nil, err
	}
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	if domainID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "domainID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		return nil, err
	}

	if projectID == "" {
		message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "projectID")
		err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
		return nil, err
	}

	u.RawQuery, u.Fragment = "", ""

	var base string
	versionRe := regexp.MustCompile("v[0-9.]+/?")
	if version := versionRe.FindString(u.Path); version != "" {
		base = strings.Replace(u.String(), version, "", -1)
	} else {
		base = u.String()
	}

	endpoint = gophercloud.NormalizeURL(endpoint)
	base = gophercloud.NormalizeURL(base)

	p := new(gophercloud.ProviderClient)
	p.IdentityBase = base
	p.IdentityEndpoint = endpoint
	p.DomainID = domainID
	p.ProjectID = projectID
	p.Conf = conf
	p.UseTokenLock()
	p.HTTPClient = newHTTPClient() //自定义httpclient，限制reauth为3次

	return p, nil
}

// Authenticate or re-authenticate against the most recent identity service
// supported at the provided endpoint.
func Authenticate(client *gophercloud.ProviderClient, options auth.AuthOptionsProvider) error {
	versions := []*utils.Version{
		{ID: v2, Priority: 20, Suffix: "/v2.0/"},
		{ID: v3, Priority: 30, Suffix: "/v3/"},
	}

	chosen, endpoint, err := utils.ChooseVersion(client, versions)
	if err != nil {
		return err
	}

	authOptions, isTokenAuthOptions := options.(tokenAuth.TokenOptions)

	if isTokenAuthOptions {
		switch chosen.ID {
		case v2:
			return tokenAuthV2(client, endpoint, authOptions, gophercloud.EndpointOpts{})
		case v3:
			return tokenAuthV3(client, endpoint, &authOptions, gophercloud.EndpointOpts{})
		default:
			// The switch statement must be out of date from the versions list.
			return fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
		}
	} else {
		akskOptions, isAKSKOptions := options.(akskAuth.AKSKOptions)

		if isAKSKOptions {
			return akskAuthV3(client, endpoint, akskOptions, gophercloud.EndpointOpts{})
		} else {
			return fmt.Errorf("Unrecognized auth options provider: %s", reflect.TypeOf(options))
		}
	}
}

func getEntryByServiceId(entries []tokens3.CatalogEntry, serviceId string) *tokens3.CatalogEntry {
	if entries == nil {
		return nil
	}

	for idx, _ := range entries {
		if entries[idx].ID == serviceId {
			return &entries[idx]
		}
	}

	return nil
}

func akskAuthV3(client *gophercloud.ProviderClient, endpoint string, options akskAuth.AKSKOptions, eo gophercloud.EndpointOpts) error {
	v3Client, err := NewIdentityV3(client, eo)
	if err != nil {
		return err
	}

	if endpoint != "" {
		v3Client.Endpoint = endpoint
	}

	v3Client.AKSKOptions = options

	var entries = make([]tokens3.CatalogEntry, 0, 1)
	serviceListErr:=services.List(v3Client, services.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		serviceLst, err := services.ExtractServices(page)
		if err != nil {
			return false, err
		}

		for _, svc := range serviceLst {
			entry := tokens3.CatalogEntry{
				Type: svc.Type,
				Name: svc.Name,
				ID:   svc.ID,
			}
			entries = append(entries, entry)
		}

		return true, nil
	})

	if serviceListErr!=nil{
		return serviceListErr
	}

	endpointListErr:=endpoints.List(v3Client, endpoints.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		endpointList, err := endpoints.ExtractEndpoints(page)
		if err != nil {
			return false, err
		}

		for _, endpoint := range endpointList {
			entry := getEntryByServiceId(entries, endpoint.ServiceID)


			if entry != nil {
				entry.Endpoints = append(entry.Endpoints, tokens3.Endpoint{
					URL:       strings.Replace(endpoint.URL, "$(tenant_id)s", options.ProjectID, -1),
					Region:    endpoint.Region,
					Interface: string(endpoint.Availability),
					ID:        endpoint.ID,
				})
			}
		}

		client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
			return GetEndpointURLForAKSKAuth(&tokens3.ServiceCatalog{
				Entries: entries,
			}, opts, options)
		}

		return true, nil
	})

	if endpointListErr!=nil{
		return endpointListErr
	}

	if client.EndpointLocator == nil {
		return gophercloud.NewSystemCommonError(gophercloud.CE_NoEndPointInCatalogCode, gophercloud.CE_NoEndPointInCatalogMessage)
	} else {
		return nil
	}

}

// AuthenticateV2 explicitly authenticates against the identity v2 endpoint.
func AuthenticateV2(client *gophercloud.ProviderClient, options tokenAuth.TokenOptions, eo gophercloud.EndpointOpts) error {
	return tokenAuthV2(client, "", options, eo)
}

func tokenAuthV2(client *gophercloud.ProviderClient, endpoint string, options tokenAuth.TokenOptions, eo gophercloud.EndpointOpts) error {
	v2Client, err := NewIdentityV2(client, eo)
	if err != nil {
		return err
	}

	if endpoint != "" {
		v2Client.Endpoint = endpoint
	}

	v2Opts := tokens2.AuthOptions{
		IdentityEndpoint: options.IdentityEndpoint,
		Username:         options.Username,
		Password:         options.Password,
		TenantID:         options.TenantID,
		TenantName:       options.TenantName,
		AllowReauth:      options.AllowReauth,
		TokenID:          options.TokenID,
	}

	result := tokens2.Create(v2Client, v2Opts)

	token, err := result.ExtractToken()
	if err != nil {
		return err
	}

	catalog, err := result.ExtractServiceCatalog()
	if err != nil {
		return err
	}

	if options.AllowReauth {
		// here we're creating a throw-away client (tac). it's a copy of the user's provider client, but
		// with the token and reauth func zeroed out. combined with setting `AllowReauth` to `false`,
		// this should retry authentication only once
		tac := *client
		tac.ReauthFunc = nil
		tac.TokenID = ""
		tao := options
		tao.AllowReauth = false
		client.ReauthFunc = func() error {
			err := tokenAuthV2(&tac, endpoint, tao, eo)
			if err != nil {
				return err
			}
			client.TokenID = tac.TokenID
			return nil
		}
	}
	client.TokenID = token.ID
	client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
		return V2EndpointURL(catalog, opts)
	}

	return nil
}

// AuthenticateV3 explicitly authenticates against the identity v3 service.
func AuthenticateV3(client *gophercloud.ProviderClient, options tokens3.AuthOptionsBuilder, eo gophercloud.EndpointOpts) error {
	return tokenAuthV3(client, "", options, eo)
}

func tokenAuthV3(client *gophercloud.ProviderClient, endpoint string, opts tokens3.AuthOptionsBuilder, eo gophercloud.EndpointOpts) error {
	// Override the generated service endpoint with the one returned by the version endpoint.
	v3Client, err := NewIdentityV3(client, eo)
	if err != nil {
		return err
	}

	if endpoint != "" {
		v3Client.Endpoint = endpoint
	}

	result := tokens3.Create(v3Client, opts)

	token, err := result.ExtractToken()
	if err != nil {
		return err
	}

	catalog, err := result.ExtractServiceCatalog()
	if err != nil {
		return err
	}

	client.TokenID = token.ID

	if opts.CanReauth() {
		// here we're creating a throw-away client (tac). it's a copy of the user's provider client, but
		// with the token and reauth func zeroed out. combined with setting `AllowReauth` to `false`,
		// this should retry authentication only once
		tac := *client
		tac.ReauthFunc = nil
		tac.TokenID = ""
		var tao tokens3.AuthOptionsBuilder
		switch ot := opts.(type) {
		case *tokenAuth.TokenOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		case *tokens3.TokenOptions:
			o := *ot
			o.AllowReauth = false
			tao = &o
		default:
			tao = opts
		}
		client.ReauthFunc = func() error {
			err := tokenAuthV3(&tac, endpoint, tao, eo)
			if err != nil {
				return err
			}
			client.TokenID = tac.TokenID
			return nil
		}
	}
	client.EndpointLocator = func(opts gophercloud.EndpointOpts) (string, error) {
		return V3EndpointURL(catalog, opts)
	}

	return nil
}

// NewIdentityV2 creates a ServiceClient that may be used to interact with the
// v2 identity service.
func NewIdentityV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	endpoint := client.IdentityBase + "v2.0/"
	clientType := "identity"
	var err error
	if !reflect.DeepEqual(eo, gophercloud.EndpointOpts{}) {
		eo.ApplyDefaults(clientType)
		endpoint, err = client.EndpointLocator(eo)
		if err != nil {
			return nil, err
		}
	}

	return &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       endpoint,
		Type:           clientType,
	}, nil
}

// NewIdentityV3 creates a ServiceClient that may be used to access the v3
// identity service.
func NewIdentityV3(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	endpoint := client.IdentityBase + "v3/"
	clientType := "identity"
	var err error
	if !reflect.DeepEqual(eo, gophercloud.EndpointOpts{}) {
		eo.ApplyDefaults(clientType)
		endpoint, err = client.EndpointLocator(eo)
		if err != nil {
			return nil, err
		}
	}

	// Ensure endpoint still has a suffix of v3.
	// This is because EndpointLocator might have found a versionless
	// endpoint and requests will fail unless targeted at /v3.
	if !strings.HasSuffix(endpoint, "v3/") {
		endpoint = endpoint + "v3/"
	}

	return &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       endpoint,
		Type:           clientType,
	}, nil
}

func getMicoreVersion(client *gophercloud.ProviderClient, url string) (versionData string) {

	type Links struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
		Type string `json:"type,omitempty"`
	}

	type MediaTypes struct {
		Type string `json:"type"`
		Base string `json:"base"`
	}
	type version struct {
		MinVersion string        `json:"min_version"`
		Links      [] Links      `json:"links"`
		ID         string        `json:"id"`
		Updated    time.Time     `json:"updated"`
		Version    string        `json:"version"`
		Status     string        `json:"status"`
		MediaTypes [] MediaTypes `json:"media-types"`
	}

	type Versions struct {
		Versions [] version `json:"versions"`
	}

	var to Versions
	_, err := client.Request("GET", url, &gophercloud.RequestOpts{JSONResponse: &to, OkCodes: []int{200, 201, 300}})

	if err != nil {
		return
	}

	for _, v := range to.Versions {
		if v.Version != "" {
			return v.Version
		}
	}
	return
}

func initClientOpts(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts, clientType string, microversion string) (*gophercloud.ServiceClient, error) {
	sc := new(gophercloud.ServiceClient)
	eo.ApplyDefaults(clientType)
	url, err := client.EndpointLocator(eo)
	if err != nil {
		return sc, err
	}
	sc.ProviderClient = client
	sc.Endpoint = url
	sc.Type = clientType

	//if clientType != "compute" {
	//	return sc, nil
	//}
	//
	//url1 := strings.Split(url, "/")
	//base := url1[0] + "//" + url1[2]
	//versionData := getMicoreVersion(client, base)
	//if versionData != "" {
	//	if microversion > versionData {
	//		sc.Microversion = versionData
	//	} else {
	//		sc.Microversion = microversion
	//	}
	//}

	return sc, nil
}

// NewObjectStorageV1 creates a ServiceClient that may be used with the v1
// object storage package.
func NewObjectStorageV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "object-store",objectSstoreMicroVersion )
}

// NewComputeV2 creates a ServiceClient that may be used with the v2 compute
// package.
func NewComputeV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "compute", computeMicroVersion)
}

// NewNetworkV2 creates a ServiceClient that may be used with the v2 network
// package.
func NewNetworkV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network",networkMicroVersion)
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc, err
}

// NewBlockStorageV1 creates a ServiceClient that may be used to access the v1
// block storage service.
func NewBlockStorageV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "volume",volumeMicroVersion)
}

// NewBlockStorageV2 creates a ServiceClient that may be used to access the v2
// block storage service.
func NewBlockStorageV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "volumev2",volumev2MicroVersion)
}

// NewBlockStorageV3 creates a ServiceClient that may be used to access the v3 block storage service.
func NewBlockStorageV3(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "volumev3",volumev3MicroVersion)
}

// NewSharedFileSystemV2 creates a ServiceClient that may be used to access the v2 shared file system service.
func NewSharedFileSystemV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "sharev2",sharev2MicroVersion)
}

// NewCDNV1 creates a ServiceClient that may be used to access the OpenStack v1
// CDN service.
func NewCDNV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "cdn",cdnMicroVersion)
}

// NewOrchestrationV1 creates a ServiceClient that may be used to access the v1
// orchestration service.
func NewOrchestrationV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "orchestration",orchestrationMicroVersion)
}

// NewDBV1 creates a ServiceClient that may be used to access the v1 DB service.
func NewDBV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "database",databaseMicroVersion)
}

// NewDNSV2 creates a ServiceClient that may be used to access the v2 DNS
// service.
func NewDNSV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "dns",dnsMicroVersion)
	sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

// NewImageServiceV2 creates a ServiceClient that may be used to access the v2
// image service.
func NewImageServiceV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "image",imageMicroVersion)
	sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

// NewLoadBalancerV2 creates a ServiceClient that may be used to access the v2
// load balancer service.
func NewLoadBalancerV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "load-balancer",loadBalancerMicroVersion)
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc, err
}

/* 自研 */

func NewECSV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "ecs",ecsMicroVersion)
}

func NewECSV1_1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "ecsv1.1",ecsv1_1MicroVersion)
}

func NewECSV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	return initClientOpts(client, eo, "ecsv2",ecsv2MicroVersion)
}

func NewIMSV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "image",imageSelfDevMicroVersion)
	sc.ResourceBase = sc.Endpoint + "v2/"
	return sc, err
}

func NewBSSV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "bss",bssMicroVersion)
	sc.ResourceBase = sc.Endpoint + "v1.0/"
	return sc, err
}

// NewNetworkV1 creates a ServiceClient that may be used with the v1 network
// package.
func NewVPCV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "vpc", vpcv1MicroVersion)
	return sc, err
}

func NewCESV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "cesv1", cesMicroVersion)
	return sc, err
}

func NewVPCV2(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "vpcv2.0", vpcv1MicroVersion)
	return sc, err
}
