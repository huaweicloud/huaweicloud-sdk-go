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

type CheckCustomerRegisterInfoOptsBuilder interface {
	ToCheckCustomerRegisterInfoMap() (map[string]interface{}, error)
}

func (opts CheckCustomerRegisterInfoOpts) ToCheckCustomerRegisterInfoMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type CreateCustomerOpts struct {
	//HUAWEI CLOUD username of the customer.
	DomainName string `json:"domainName,omitempty"`

	//Email address.
	Email string `json:"email,omitempty"`

	//Verification code
	VerificationCode string `json:"verificationCode,omitempty"`

	//Two-letter ID representing the country/region of the customer.
	DomainArea string `json:"domainArea,omitempty"`

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

type FrozenCustomerOpts struct {
	//IDs of customers whose accounts are to be frozen。
	CustomerIds []string `json:"customerIds" required:"true"`

	//Account freezing reason.
	Reason string `json:"reason" required:"true"`
}

type FrozenCustomerOptsBuilder interface {
	ToFrozenCustomerMap() (map[string]interface{}, error)
}

func (opts FrozenCustomerOpts) ToFrozenCustomerMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

type UnFrozenCustomerOpts struct {
	//IDs of customers whose accounts are to be unfrozen
	CustomerIds []string `json:"customerIds" required:"true"`

	//Account unfreezing reason.
	Reason string `json:"reason" required:"true"`
}

type UnFrozenCustomerOptsBuilder interface {
	ToUnFrozenCustomerMap() (map[string]interface{}, error)
}

func (opts UnFrozenCustomerOpts) ToUnFrozenCustomerMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/**
* This API is used to check whether the account name, and mobile number or email address entered by the customer can be used for registration.
* This API can be invoked only by the partner AK/SK or token.
 */
func CheckCustomerRegisterInfo(client *gophercloud.ServiceClient, opts CheckCustomerRegisterInfoOptsBuilder) (r CheckCustomerRegisterInfoResult) {
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

/**
 * A partner can unfreeze an account of a customer associated with the partner by reseller model.
 * This API can be invoked only by the partner account AK/SK or token.
 */
func FrozenCustomer(client *gophercloud.ServiceClient, opts FrozenCustomerOptsBuilder) (r FrozenCustomerResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToFrozenCustomerMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getFrozenCustomerURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}

/**
 * A partner can unfreeze an account of a customer associated with the partner by reseller model.
 * This API can be invoked only by the partner account AK/SK or token.
 */
func UnFrozenCustomer(client *gophercloud.ServiceClient, opts UnFrozenCustomerOptsBuilder) (r UnFrozenCustomerResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToUnFrozenCustomerMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getUnFrozenCustomerURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200},
		})
	}

	return
}