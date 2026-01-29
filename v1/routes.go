package v1

import "net/url"

const APIBaseURL = "https://api.abacatepay.com/"

const (
	RouteCreateCustomer = "/customer/create"
	RouteListCustomers  = "/customer/list"
)

const (
	RouteCreateCharge = "/billing/create"
	RouteListCharges  = "/billing/list"
)

const (
	RouteCreatePIXQRCode = "/pixQrCode/create"
	RouteSimulatePayment = "/pixQrCode/simulate-payment"
	RouteCheckQRCodePIX  = "/pixQrCode/check"
)

func BuildSimulatePaymentURL(params RESTPostSimulatePaymentQueryParams) string {
	query := url.Values{}

	query.Set("id", params.ID)

	return RouteSimulatePayment + "?" + query.Encode()
}

func BuildCheckQRCodePIXStatusURL(params RESTGetCheckQRCodePixStatusQueryParams) string {
	query := url.Values{}

	query.Set("id", params.ID)

	return RouteCheckQRCodePIX + "?" + query.Encode()
}

const (
	RouteCreateCoupon = "/coupon/create"
	RouteListCoupons  = "/coupon/list"
)

const (
	RouteCreateWithdraw = "/withdraw/create"
	RouteGetWithdraw    = "/withdraw/get"
	RouteListWithdraws  = "/withdraw/list"
)

func BuildGetWithdrawURL(params RESTGetSearchWithdrawQueryParams) string {
	query := url.Values{}

	query.Set("externalId", params.ExternalID)

	return RouteGetWithdraw + "?" + query.Encode()
}

const RouteStoreDetails = "/store/get"

const (
	RouteGetMRR      = "/public-mrr/mrr"
	RouteGetMerchant = "/public-mrr/merchant-info"
	RouteGetRevenue  = "/public-mrr/revenue"
)

func BuildGetRevenueByPeriod(start, end string) string {
	query := url.Values{}

	query.Set("startDate", start)
	query.Set("endDate", end)

	return RouteGetRevenue + "?" + query.Encode()
}
