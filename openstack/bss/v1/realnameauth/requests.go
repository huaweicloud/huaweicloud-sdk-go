package realnameauth

import (
	"github.com/gophercloud/gophercloud"
)

type IndividualRealNameAuthOpts struct {
	//Customer ID.
	CustomerId string `json:"customerId" required:"true"`

	//Authentication method.
	IdentifyType *int `json:"identifyType" required:"true"`

	//Certificate type.
	VerifiedType int `json:"verifiedType,omitempty"`

	//Attachment URL for individual certificate authentication.
	VerifiedFileURL []string `json:"verifiedFileURL" required:"true"`

	//Name.
	Name string `json:"name" required:"true"`

	//Certificate number.
	VerifiedNumber string `json:"verifiedNumber" required:"true"`

	//Change type.
	ChangeType int `json:"changeType,omitempty"`

	//Platform ID assigned by Huawei to a partner.
	XaccountType string `json:"xaccountType" required:"true"`

	//bank card information
	BankCardInfo BankCardInfo `json:"bankCardInfo,omitempty"`

}

type BankCardInfo struct {
	//Bank account
	BankAccount string `json:"bankAccount" required:"true"`

	//area code
	Areacode string `json:"areacode" required:"true"`

	//mobile phone number
	Mobile string `json:"mobile" required:"true"`

	//Verification code.
	VerificationCode string `json:"verificationCode" required:"true"`
}




type EnterpriseRealNameAuthOpts struct {
	//Customer ID.
	CustomerId string `json:"customerId" required:"true"`

	//Authentication method.
	IdentifyType *int `json:"identifyType" required:"true"`

	//Enterprise certificate type.
	CertificateType int `json:"certificateType,omitempty"`

	//URL of the certificate attachment file used for enterprise certificate authentication.
	VerifiedFileURL []string `json:"verifiedFileURL" required:"true"`

	//Organization name.
	CorpName string `json:"corpName" required:"true"`

	//Enterprise certificate number.
	VerifiedNumber string `json:"verifiedNumber" required:"true"`

	//Registration country entered for real-name authentication.
	RegCountry string `json:"regCountry,omitempty"`

	//Enterprise registration address for real-name authentication.
	RegAddress string `json:"regAddress,omitempty"`

	//Platform ID assigned by Huawei to a partner.。
	XaccountType string `json:"xaccountType" required:"true"`

	//Enterprise person information.
	EnterprisePerson EnterprisePerson `json:"enterprisePerson,omitempty"`
}

type ChangeEnterpriseRealNameAuthOpts struct {
	//Customer ID.
	CustomerId string `json:"customerId" required:"true"`

	//Authentication method.
	IdentifyType *int `json:"identifyType" required:"true"`

	//Enterprise certificate type.
	CertificateType int `json:"certificateType,omitempty"`

	//URL of the certificate attachment file used for enterprise certificate authentication.
	VerifiedFileURL []string `json:"verifiedFileURL" required:"true"`

	//Organization name.
	CorpName string `json:"corpName" required:"true"`

	//Enterprise certificate number.
	VerifiedNumber string `json:"verifiedNumber" required:"true"`

	//Registration country entered for real-name authentication.
	RegCountry string `json:"regCountry,omitempty"`

	//Enterprise registration address for real-name authentication.
	RegAddress string `json:"regAddress,omitempty"`

	//Change type
	ChangeType *int `json:"changeType" required:"true"`

	//Platform ID assigned by Huawei to a partner.
	XaccountType string `json:"xaccountType" required:"true"`

	//Enterprise person information.
	EnterprisePerson EnterprisePerson `json:"enterprisePerson,omitempty"`
}



type EnterprisePerson struct {
	//Legal entity name.
	LegelName string `json:"legelName" required:"true"`

	//Legal entity card ID.
	LegelIdNumber string `json:"legelIdNumber" required:"true"`

	//Legal entity role.
	CertifierRole string `json:"certifierRole,omitempty"`
}

type QueryRealNameAuthOpts struct {
	//Customer ID.
	CustomerId string `q:"customerId" required:"true"`
}

type QueryRealNameAuthOptsBuilder interface {
	ToQueryRealNameAuthMap() (string, error)
}

func (opts QueryRealNameAuthOpts) ToQueryRealNameAuthMap() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

type EnterpriseRealNameAuthOptsBuilder interface {
	ToEnterpriseRealNameAuthMap() (map[string]interface{}, error)
}

func (opts EnterpriseRealNameAuthOpts) ToEnterpriseRealNameAuthMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type IndividualRealNameAuthOptsBuilder interface {
	ToIndividualRealNameAuthMap() (map[string]interface{}, error)
}

func (opts IndividualRealNameAuthOpts) ToIndividualRealNameAuthMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type ChangeEnterpriseRealNameAuthOptsBuilder interface {
	ToChangeEnterpriseRealNameAuthMap() (map[string]interface{}, error)
}

func (opts ChangeEnterpriseRealNameAuthOpts) ToChangeEnterpriseRealNameAuthMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/**
 * An individual customer can apply for real-name authentication on the partner sales platform. Currently, two authentication methods are supported: using the individual certificate and using individual bank card.
 * This API can be invoked only by the partner account AK/SK or token.ken.
 */
func IndividualRealNameAuth(client *gophercloud.ServiceClient, opts IndividualRealNameAuthOptsBuilder) (r IndividualRealNameAuthResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToIndividualRealNameAuthMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getIndividualRealNameAuthURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}

/**
 * Enterprise customers can perform enterprise real-name authentication on the partner sales platform.
 * This API can be invoked only by the partner account AK/SK or token.
 */
func EnterpriseRealNameAuth(client *gophercloud.ServiceClient, opts EnterpriseRealNameAuthOptsBuilder) (r EnterpriseRealNameAuthResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToEnterpriseRealNameAuthMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getEnterpriseRealNameAuthURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}

/**
 * If the response to a real-name authentication application or real-name authentication change application indicates that manual review is required, this API can be used to query the review result.
 * This API can be invoked only by the partner account AK/SK or token.
 */
func QueryRealNameAuth(client *gophercloud.ServiceClient, opts QueryRealNameAuthOptsBuilder)(r QueryRealNameAuthResult)  {
	domainID := client.ProviderClient.DomainID
	url := getQueryRealNameAuthURL(client, domainID)
	if opts != nil {
		query, err := opts.ToQueryRealNameAuthMap()
		if err != nil {
			r.Err = err
			return
		}
		url += query
		_, r.Err = client.Get(url, &r.Body, nil)
	}

	return
}

/**
 * Customers can submit real-name authentication change applications on the partner sales platform. Currently, customers can change real-name authentication using individual information to using enterprise information.
 * This API can be invoked only by the partner account AK/SK or token
 */
func ChangeEnterpriseRealNameAuth(client *gophercloud.ServiceClient, opts ChangeEnterpriseRealNameAuthOptsBuilder) (r ChangeEnterpriseRealNameAuthResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToChangeEnterpriseRealNameAuthMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Put(getChangeEnterpriseRealNameAuthURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}








