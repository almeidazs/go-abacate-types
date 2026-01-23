package v2

import (
	"net/url"
	"strconv"
)

const APIBaseURL = "https://api.abacatepay.com"

const (
	RouteCreateCustomer = "/customers/create"
	RouteListCustomers  = "/customers/list"
	RouteGetCustomer    = "/customers/get"
	RouteDeleteCustomer = "/customers/delete"
)

func withQuery(route string, q url.Values) string {
	if len(q) == 0 {
		return route
	}

	return route + "?" + q.Encode()
}

func setInt(q url.Values, key string, v *int) {
	if v != nil {
		q.Set(key, strconv.Itoa(*v))
	}
}

func setString(q url.Values, key string, v *string) {
	if v != nil {
		q.Set(key, *v)
	}
}

func BuildListCustomersURL(page, limit *int) string {
	q := url.Values{}
	setInt(q, "page", page)
	setInt(q, "limit", limit)
	return withQuery(RouteListCustomers, q)
}

func BuildGetCustomerURL(id string) string {
	q := url.Values{}
	q.Set("id", id)
	return withQuery(RouteGetCustomer, q)
}

/* =======================
   Checkouts
======================= */

const (
	RouteCreateCheckout = "/checkouts/create"
	RouteListCheckouts  = "/checkouts/list"
	RouteGetCheckout    = "/checkouts/get"
)

func BuildGetCheckoutURL(id string) string {
	q := url.Values{}
	q.Set("id", id)
	return withQuery(RouteGetCheckout, q)
}

/* =======================
   PIX / Transparents
======================= */

const (
	RouteCreatePixQRCode = "/transparents/create"
	RouteSimulatePayment = "/transparents/simulate-payment"
	RouteCheckQRCodePIX  = "/transparents/check"
)

func BuildSimulatePaymentURL(id string) string {
	q := url.Values{}

	q.Set("id", id)

	return withQuery(RouteSimulatePayment, q)
}

func BuildCheckQRCodePIXStatusURL(id string) string {
	q := url.Values{}

	q.Set("id", id)

	return withQuery(RouteCheckQRCodePIX, q)
}

const (
	RouteCreateCoupon       = "/coupons/create"
	RouteListCoupons        = "/coupons/list"
	RouteGetCoupon          = "/coupons/get"
	RouteDeleteCoupon       = "/coupons/delete"
	RouteToggleStatusCoupon = "/coupons/toggle"
)

func BuildGetCouponURL(id string) string {
	q := url.Values{}

	q.Set("id", id)

	return withQuery(RouteGetCoupon, q)
}

func BuildListCouponsURL(page, limit *int) string {
	q := url.Values{}

	setInt(q, "page", page)
	setInt(q, "limit", limit)

	return withQuery(RouteListCoupons, q)
}

const (
	RouteCreatePayout = "/payouts/create"
	RouteGetPayout    = "/payouts/get"
	RouteListPayouts  = "/payouts/list"
)

func BuildGetPayoutURL(externalID string) string {
	q := url.Values{}
	q.Set("externalId", externalID)

	return withQuery(RouteGetPayout, q)
}

func BuildListPayoutsURL(page, limit *int) string {
	q := url.Values{}

	setInt(q, "page", page)
	setInt(q, "limit", limit)

	return withQuery(RouteListPayouts, q)
}

const RouteStoreDetails = "/store/get"

const (
	RouteGetMRR      = "/public-mrr/mrr"
	RouteGetMerchant = "/public-mrr/merchant-info"
	RouteGetRevenue  = "/public-mrr/revenue"
)

func BuildGetRevenueURL(start, end string) string {
	q := url.Values{}

	q.Set("startDate", start)
	q.Set("endDate", end)

	return withQuery(RouteGetRevenue, q)
}

const (
	RouteCreateSubscription = "/subscriptions/create"
	RouteListSubscriptions  = "/subscriptions/list"
)

func BuildListSubscriptionsURL(limit *int, cursor *string) string {
	q := url.Values{}

	setInt(q, "limit", limit)
	setString(q, "cursor", cursor)

	return withQuery(RouteListSubscriptions, q)
}

const (
	RouteCreateProduct = "/products/create"
	RouteListProducts  = "/products/list"
	RouteGetProduct    = "/products/get"
)

func BuildGetProductURL(id, externalID *string) string {
	q := url.Values{}

	setString(q, "id", id)
	setString(q, "externalId", externalID)

	return withQuery(RouteGetProduct, q)
}

func BuildListProductsURL(page, limit *int) string {
	q := url.Values{}

	setInt(q, "page", page)
	setInt(q, "limit", limit)

	return withQuery(RouteListProducts, q)
}
