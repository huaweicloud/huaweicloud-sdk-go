package utilities

import "github.com/gophercloud/gophercloud"

type SendVerificationCodeOpts struct {
	//Settings.
	ReceiverType *int `json:"receiverType" required:"true"`

	//Timeout duration of a verification code.
	Timeout int `json:"timeout,omitempty"`

	//The sender email address.
	Email string `json:"email,omitempty"`

	//Mobile number.
	MobilePhone string `json:"mobilePhone,omitempty"`

	//Country code.
	CountryCode string `json:"countryCode,omitempty"`

	//If no template information is found for the selected language, the template information for the default language is used.
	Lang string `json:"lang,omitempty"`

	//Supported scenarios
	Scene string `json:"scene,omitempty"`

	//Customer ID.
	CustomerId string `json:"customerId,omitempty"`
}

type SendVerificationCodeOptsBuilder interface {
	ToSendVerificationCodeMap() (map[string]interface{}, error)
}

func (opts SendVerificationCodeOpts) ToSendVerificationCodeMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

/**
 * If customers enter email addresses for registration, this API is used to send a registration verification code to the email addresses to verify the registration information.
 * This API can be invoked only by the partner AK/SK or token.
 */
func SendVerificationCode(client *gophercloud.ServiceClient, opts SendVerificationCodeOptsBuilder) (r SendVerificationCodeResult) {
	domainID := client.ProviderClient.DomainID
	if opts != nil {
		body, err := opts.ToSendVerificationCodeMap()
		if err != nil {
			r.Err = err
			return
		}
		_, r.Err = client.Post(getSendVerificationCodeURL(client, domainID), body, &r.Body, &gophercloud.RequestOpts{
			OkCodes: []int{204},
		})
	}

	return
}