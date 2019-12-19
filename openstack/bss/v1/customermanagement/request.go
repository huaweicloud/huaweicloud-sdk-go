package customermanagement

import (
	"encoding/json"
	"github.com/gophercloud/gophercloud"
)

type CheckCustomerRegisterInfoOpts struct {
	//The options are email, mobile, or name
	SearchType string `json:"searchType" required:"true"`

	//Mobile number, email address, or account name.
	SearchKey string `json:"searchKey" required:"true"`
}

type CheckCustomerRegisterInfoBuilder interface {
	ToCheckCustomerRegisterInfoMap() (map[string]interface{}, error)
}

func (opts CheckCustomerRegisterInfoOpts) ToCheckCustomerRegisterInfoMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type CreateCustomerOpts struct {
	//HUAWEI CLOUD username of the customer.
	DomainName string `json:"domainName,omitempty"`

	//Mobile number.
	MobilePhone string `json:"mobilePhone,omitempty"`

	//Mobile number country code
	CountryCode string `json:"countryCode,omitempty"`

	//Verification code
	VerificationCode string `json:"verificationCode,omitempty"`

	//Unique ID of the user on the third-party system, which is assigned by the partner.
	XAccountId string `json:"xAccountId" required:"true"`

	//Platform ID assigned by Huawei to a partner.
	XAccountType string `json:"xAccountType" required:"true"`

	//password
	Password string `json:"password,omitempty"`

	//Indicates whether to disable the marketing message sending function.
	IsCloseMarketMs string `json:"isCloseMarketMs,omitempty"`
}

type CreateCustomerOptsBuilder interface {
	ToCreateCustomerMap() (map[string]interface{}, error)
}

func (opts CreateCustomerOpts) ToCreateCustomerMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type QueryCustomerOpts struct {
	//Account name
	DomainName string `json:"domainName,omitempty"`

	//Real-name authentication name
	Name string `json:"name,omitempty"`

	//Page to be queried
	Offset int `json:"offset,omitempty"`

	//Number of records on each page
	Limit int `json:"limit,omitempty"`

	//Tag
	Label string `json:"label,omitempty"`

	//Association type
	CooperationType string `json:"cooperationType,omitempty"`

	//Start time of the association time range (UTC time)
	CooperationTimeStart string `json:"cooperationTimeStart,omitempty"`

	//End time of the association time range (UTC time)
	CooperationTimeEnd string `json:"cooperationTimeEnd,omitempty"`
}

type QueryCustomerOptsBuilder interface {
	ToQueryCustomerMap() (map[string]interface{}, error)
}

func (opts QueryCustomerOpts) ToQueryCustomerMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/**
* This API is used to check whether the account name, and mobile number or email address entered by the customer can be used for registration.
* This API can be invoked only by the partner AK/SK or token.
 */
func CheckCustomerRegisterInfo(client *gophercloud.ServiceClient, opts CheckCustomerRegisterInfoBuilder) (r CheckCustomerRegisterInfoResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToCheckCustomerRegisterInfoMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getCheckCustomerRegisterInfoURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}

/**
 * This API is used to create a HUAWEI CLOUD account for a customer when the customer creates an account on your sales platform, and bind the customer account on the partner sales platform to the HUAWEI CLOUD account. In addition, the HUAWEI CLOUD account is bound to the partner account.
 * This API can be invoked only by the partner AK/SK or token..
 */
func CreateCustomer(client *gophercloud.ServiceClient, opts CreateCustomerOptsBuilder) (r CreateCustomerResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToCreateCustomerMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getCreateCustomerURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
			HandleError: func(httpStatus int, responseContent string) error {
				var createCustomerResp CreateCustomerResp
				message := responseContent
				err := json.Unmarshal([]byte(responseContent), &createCustomerResp)
				if err == nil {
					return &createCustomerResp
				}
				return &gophercloud.UnifiedError{
					ErrCode:    gophercloud.MatchErrorCode(httpStatus, message),
					ErrMessage: message,
				}
			},
		})
	}

	return
}

/**
 * This API is used to query your customers.
 * This API can be invoked only by the partner account AK/SK or token.
 */
func QueryCustomer(client *gophercloud.ServiceClient, opts QueryCustomerOptsBuilder) (r QueryCustomerResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToQueryCustomerMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getQueryCustomerURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}
