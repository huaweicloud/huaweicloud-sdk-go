package enquiry

import "github.com/gophercloud/gophercloud"

type QueryRatingResp struct {
	//Status code
	ErrorCode string `json:"error_code"`

	//Error description.
	ErrorMsg string `json:"error_msg"`

	//Inquiry result
	RatingResult RatingResult `json:"ratingResult"`
}

type RatingResult struct {
	//Total order amount (final order amount after discount)
	Amount *float64 `json:"amount"`

	//Discounted amount in an order.
	DiscountAmount *float64 `json:"discountAmount"`

	//Original order amount (order amount before discount).
	OriginalAmount *float64 `json:"originalAmount"`

	//Measurement unit ID.
	MeasureId *int `json:"measureId"`

	//Currency unit code (complying with the ISO 4217 standard)
	Currency string `json:"currency"`

	//Product price inquiry result.
	ProductRatingResult []ProductRatingResult `json:"productRatingResult"`

	//Extended parameter.
	ExtendParams string `json:"extendParams"`
}

type ProductRatingResult struct {
	//ID, which comes from the ID in the request.
	Id string `json:"id"`

	//Product ID.
	ProductId string `json:"productId"`

	//Total amount (final amount after discount)
	Amount *float64 `json:"amount,omitempty"`

	//Original total amount.
	OriginalAmount *float64 `json:"originalAmount,omitempty"`

	//Discounted amount.
	DiscountAmount *float64 `json:"discountAmount,omitempty"`

	//Measurement unit ID.
	MeasureId *int `json:"measureId,omitempty"`

	//Extended parameter.
	ExtendParams string `json:"extendParams"`
}

type QueryRatingResult struct { 
	gophercloud.Result
}

func (r QueryRatingResult) Extract() (*QueryRatingResp, error) {
	var res *QueryRatingResp
	err := r.ExtractInto(&res)
	return res, err
}
