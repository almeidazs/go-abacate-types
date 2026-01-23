package v2

import "time"

// Any response returned by the AbacatePay API.
type APIResponse[Payload any] struct {
	// The data of the response.
	Data *Payload `json:"data,omitempty"`

	// Error message returned from the API.
	Error *string `json:"error,omitempty"`
}

// Any response returned by the AbacatePay API that has a `pagination` field (Not cursor based).
type APIResponseWithPagination[Payload any] struct {
	// The data of the response.
	Data *Payload `json:"data,omitempty"`

	// Error message returned from the API.
	Error *string `json:"error,omitempty"`

	// Pagination info.
	Pagination APIResponseWithPaginationData `json:"pagination"`
}

type APIResponseWithPaginationData struct {
	// Current page.
	Page uint64 `json:"page"`

	// Number of items per page.
	Limit uint64 `json:"limit"`

	// Number of items.
	Items uint64 `json:"items"`

	// Number of pages.
	TotalPages uint64 `json:"totalPages"`
}

// Any response returned by the AbacatePay API that has a `pagination` field and is cursor-based.
type APIResponseWithCursorBasedPagination[Payload any] struct {
	// The data of the response.
	Data *Payload `json:"data,omitempty"`

	// Error message returned from the API.
	Error *string `json:"error,omitempty"`

	// Pagination info.
	Pagination APIResponseWithCursorBasedPaginationData `json:"pagination"`
}

type APIResponseWithCursorBasedPaginationData struct {
	// Number of items per page.
	Limit uint64 `json:"limit"`

	// Indicates whether there is a next page.
	HasNext bool `json:"hasNext"`

	// Indicates whether there is a previous page.
	HasPrevious bool `json:"hasPrevious"`

	// Cursor for the next page.
	NextCursor *string `json:"nextCursor,omitempty"`
}

// https://api.abacatepay.com/v2/customers/create
type RESTPostCreateCustomerBody struct {
	// Customer's full name.
	Name *string `json:"name,omitempty"`

	// Customer's email.
	Email string `json:"email"`

	// Customer's CPF or CNPJ.
	TaxID *string `json:"taxId,omitempty"`

	// Customer's cell phone.
	Cellphone *string `json:"cellphone,omitempty"`

	// Customer zip code.
	ZipCode *string `json:"zipCode,omitempty"`

	// Additional customer metadata.
	Metadata map[string]any `json:"metadata,omitempty"`
}

// https://api.abacatepay.com/v2/customers/create
type RESTPostCreateCustomerData APIResponse[APICustomer]

// https://api.abacatepay.com/v2/checkouts/create
type RESTPostCreateNewCheckoutBody struct {
	// Payment method that will be used.
	Methods *PaymentMethod `json:"methods,omitempty"`

	// URL to redirect the customer if they click on the "Back" option.
	ReturnURL *string `json:"returnUrl,omitempty"`

	// URL to redirect the customer when payment is completed.
	CompletionURL *string `json:"completionUrl,omitempty"`

	// The ID of a customer already registered in your store.
	CustomerID *string `json:"customerId,omitempty"`

	// Your customer's data to create it.
	// The customer object is not mandatory, but when entering any `customer` information, all `name`, `cellphone`, `email` and `taxId` fields are mandatory.
	Customer *APIBaseCustomer `json:"customer,omitempty"`

	// List of coupons available for resem used with billing (0-50 max.).
	Coupons *[]string `json:"coupons,omitempty"`

	// If you have a unique identifier for your billing application, completely optional.
	ExternalID *string `json:"externalId,omitempty"`

	// Optional billing metadata.
	Metadata map[string]any `json:"metadata,omitempty"`

	// List of items included in the charge.
	// This is the only required field — the total value is calculated from these items.
	Items []APIItem `json:"items"`
}

// https://api.abacatepay.com/v2/checkouts/create
type RESTPostCreateNewCheckoutData APIResponse[APICheckout]

// https://api.abacatepay.com/v2/transparents/create
type RESTPostCreateQRCodePixBody struct {
	// Charge amount in cents.
	Amount uint64 `json:"amount"`

	// Billing expiration time in seconds.
	ExpiresIn *uint64 `json:"expiresIn,omitempty"`

	// Message that will appear when paying the PIX.
	Description *string `json:"description,omitempty"`

	// Your customer's data to create it.
	// The customer object is not mandatory, but when entering any `customer` information, all `name`, `cellphone`, `email` and `taxId` fields are mandatory.
	Customer *APIBaseCustomer `json:"customer,omitempty"`

	// Optional billing metadata.
	Metadata map[string]any `json:"metadata,omitempty"`
}

// https://api.abacatepay.com/v2/transparents/create
type RESTPostCreateQRCodePixData APIResponse[APIQRCodePIX]

// https://api.abacatepay.com/v2/transparents/simulate-payment
type RESTPostSimulateQRCodePixPaymentQueryParams struct {
	// QRCode Pix ID.
	ID string `json:"id"`
}

// https://api.abacatepay.com/v2/transparents/simulate-payment
type RESTPostSimulateQRCodePixPaymentBody struct {
	// Optional metadata for the request.
	Metadata map[string]any `json:"metadata"`
}

// https://api.abacatepay.com/v2/transparents/simulate-payment
type RESTPostSimulateQRCodePixPaymentData APIResponse[APIQRCodePIX]

// https://api.abacatepay.com/v2/pixQrCode/check
type RESTGetCheckQRCodePixStatusQueryParams struct {
	// QRCode Pix ID.
	ID string `json:"id"`
}

// https://api.abacatepay.com/v2/transparents/check
type RESTGetCheckQRCodePixStatusData APIResponse[struct {
	//QRCode Pix expiration date.
	ExpiresAt time.Time `json:"expiresAt"`

	// Information about the progress of QRCode Pix.
	Status PaymentStatus `json:"status"`
}]

// https://api.abacatepay.com/v2/checkouts/list
type RESTGetListCheckoutsData APIResponse[[]APICheckout]

// https://api.abacatepay.com/v2/checkouts/get
type RESTGetCheckoutData APIResponse[APICheckout]

// https://api.abacatepay.com/v2/checkouts/get
type RESTGetCheckoutQueryParams struct {
	// Unique billing identifier.
	ID string `json:"id"`
}

// https://api.abacatepay.com/v2/customers/list
type RESTGetListCustomersData APIResponseWithPagination[[]APICustomer]

// https://api.abacatepay.com/v2/customers/list
type RESTGetListCustomersQueryParams struct {
	// Number of the page.
	Page *uint64 `json:"page,omitempty"`

	// Number of items per page.
	Limit *uint64 `json:"limit,omitempty"`
}

// https://api.abacatepay.com/v2/customers/get
type RESTGetCustomerQueryParams struct {
	// The ID of the customer.
	ID string `json:"id"`
}

// https://api.abacatepay.com/v2/customers/get
type RESTGetCustomerData APIResponse[struct {
	// Unique customer identifier.
	ID string `json:"id"`

	// Indicates whether the client was created in a testing environment.
	DevMode bool `json:"devMode"`

	// Customer's full name.
	Name string `json:"name"`

	// Customer's email.
	Email string `json:"email"`

	// Customer's CPF or CNPJ.
	TaxID string `json:"taxId"`

	// Customer's cell phone.
	Cellphone string `json:"cellphone"`

	// Additional customer metadata.
	Metadata map[string]any `json:"metadata,omitempty"`
}]

// https://api.abacatepay.com/v2/customers/delete
type RESTDeleteCustomerBody struct {
	// Unique public identifier of the customer to be deleted.
	ID string `json:"id"`
}

// https://api.abacatepay.com/v2/customers/delete
type RESTDeleteCustomerData APIResponse[struct {
	// Unique customer identifier.
	ID string `json:"id"`

	// Indicates whether the client was created in a testing environment.
	DevMode bool `json:"devMode"`

	// Customer's full name.
	Name string `json:"name"`

	// Customer's email.
	Email string `json:"email"`

	// Customer's CPF or CNPJ.
	TaxID string `json:"taxId"`

	// Customer's cell phone.
	Cellphone string `json:"cellphone"`

	// Additional customer metadata.
	Metadata map[string]any `json:"metadata,omitempty"`
}]

// https://api.abacatepay.com/v2/coupons/create
type RESTPostCreateCouponBody struct {
	// Unique coupon identifier.
	Code string `json:"code"`

	// Discount amount to be applied.
	Discount uint64 `json:"discount"`

	// Type of discount applied, percentage or fixed.
	DiscountKind CouponDiscountKind `json:"discountKind"`

	// Coupon description.
	Notes *string `json:"notes,omitempty"`

	// Number of times the coupon can be redeemed. -1 means this coupon can be redeemed without limits.
	MaxRedeems *uint64 `json:"maxRedeems"`

	// Key value object for coupon metadata.
	Metadata map[string]any `json:"metadata,omitempty"`
}

// https://api.abacatepay.com/v2/coupon/create
type RESTPostCreateCouponData APIResponse[APICoupon]

// https://api.abacatepay.com/v2/payouts/create
type RESTPostCreateNewPayoutBody struct {
	// Unique identifier of the payout in your system.
	ExternalID string `json:"externalId"`

	// Payout value in cents (Min 350).
	Amount uint64 `json:"amount"`

	// Optional payout description.
	Description *string `json:"description,omitempty"`
}

// https://api.abacatepay.com/v2/payouts/create
type RESTPostCreateNewPayoutData APIResponse[APIPayout]

// https://api.abacatepay.com/v2/payouts/get
type RESTGetSearchPayoutQueryParams struct {
	// Unique payout identifier in your system.
	ExternalID string `json:"externalId"`
}

// https://api.abacatepay.com/v2/payouts/get
type RESTGetSearchPayoutData APIResponse[APIPayout]

// https://api.abacatepay.com/v2/payouts/list
type RESTGetListPayoutsQueryParams struct {
	// Number of the page.
	Page *uint64 `json:"page,omitempty"`

	// Number of items per page.
	Limit *uint64 `json:"limit,omitempty"`
}

// https://api.abacatepay.com/v2/public-mrr/revenue
type RESTGetRevenueByPeriodQueryParams struct {
	// Period start date (YYYY-MM-DD format).
	StartDate time.Time `json:"startDate"`

	// Period end date (YYYY-MM-DD format).
	EndDate time.Time `json:"endDate"`
}

// https://api.abacatepay.com/v2/public-mrr/revenue
type RESTGetRevenueByPeriodData APIResponse[struct {
	// Total revenue for the period in cents.
	TotalRevenue uint64 `json:"totalRevenue"`

	// Total transactions in the period.
	TotalTransactions uint64 `json:"totalTransactions"`

	// Object with transactions grouped by day (key is the date in YYYY-MM-DD format).
	TransactionsPerDay map[string]struct {
		// Total value of the day's transactions in cents.
		Amount uint64 `json:"amount"`

		// Number of transactions for the day.
		Count uint64 `json:"count"`
	} `json:"transactionsPerDay"`
}]

// https://api.abacatepay.com/v2/public-mrr/merchant-info
type RESTGetMerchantData APIResponse[struct {
	// Store name.
	Name string `json:"name"`

	// Store website.
	Website string `json:"website"`

	// Store creation date.
	CreatedAt string `json:"createdAt"`
}]

// https://api.abacatepay.com/v2/public-mrr/mrr
type RESTGetMRRData APIResponse[struct {
	// Monthly recurring revenue in cents. Value 0 indicates that there is no recurring revenue at the moment.
	MRR uint64 `json:"mrr"`

	// Total active subscriptions. Value 0 indicates that there are no currently active subscriptions.
	TotalActiveSubscriptions uint64 `json:"totalActiveSubscriptions"`
}]

// https://api.abacatepay.com/v2/store/get
type RESTGetStoreDetailsData APIResponse[APIStore]

// https://api.abacatepay.com/v2/payouts/list
type RESTGetListPayoutsData APIResponseWithPagination[[]APIPayout]

// https://api.abacatepay.com/v2/coupons/list
type RESTGetListCouponsData APIResponseWithPagination[[]APICoupon]

// https://api.abacatepay.com/v2/coupons/list
type RESTGetListCouponsQueryParams struct {
	// Page number.
	Page *uint64 `json:"page"`

	// Number of items per page.
	Limit *uint64 `json:"limit"`
}

// https://api.abacatepay.com/v2/coupons/get
type RESTGetCouponQueryParams struct {
	// The ID of the coupon.
	ID string `json:"id"`
}

// https://api.abacatepay.com/v2/coupons/get
type RESTGetCouponData APIResponse[APICoupon]

// https://api.abacatepay.com/v2/coupons/delete
type RESTDeleteCouponBody struct {
	// The ID of the coupon.
	ID string `json:"id"`
}

// https://api.abacatepay.com/v2/coupons/delete
type RESTDeleteCouponData APIResponse[APICoupon]

// https://api.abacatepay.com/v2/coupons/toggle
type RESTPatchToggleCouponStatusBody struct {
	// The ID of the coupon.
	ID string `json:"id"`
}

// https://api.abacatepay.com/v2/coupons/toggle
type RESTPatchToggleCouponStatusData APIResponse[APICoupon]

// https://api.abacatepay.com/v2/products/create
type RESTPostCreateProductBody struct {
	// Description for the product.
	Description *string `json:"description,omitempty"`

	// Product name.
	Name string `json:"name"`

	// Product price in cents.
	Price uint64 `json:"price"`

	// Product currency.
	Currency string `json:"currency"`

	// Unique product identifier in your system.
	ExternalID string `json:"externalId"`
}

// https://api.abacatepay.com/v2/products/create
type RESTPostCreateProductData APIResponse[APIProduct]

// https://api.abacatepay.com/v2/products/list
type RESTGetListProductsQueryParams struct {
	// Page number.
	Page *uint64 `json:"page,omitempty"`

	// Limit of products to return.
	Limit *uint64 `json:"limit,omitempty"`
}

// https://api.abacatepay.com/v2/products/list
type RESTGetListProductsData APIResponseWithPagination[[]APIProduct]

// https://api.abacatepay.com/v2/products/get
type RESTGetProductQueryParams struct {
	// The product ID.
	ID *string `json:"id,omitempty"`

	// External ID of the product.
	ExternalId *string `json:"externalId,omitempty"`
}

// https://api.abacatepay.com/v2/products/get
type RESTGetProductData APIResponse[APIProduct]

// https://api.abacatepay.com/v2/subscriptions/create
type RESTPostCreateSubscriptionBody struct {
	// Payment method for the subscription.
	Method PaymentMethod `json:"method"`

	// Billing frequency configuration.
	Frequency APISubscriptionFrequency `json:"frequency"`

	// Identifier of the customer who will have the signature.
	CustomerID string `json:"customerId"`

	// Retry policy in case of payment failure.
	RetryPolicy APISubscriptionRetryPolicy `json:"retryPolicy"`

	// Subscription description.
	Description *string `json:"description,omitempty"`

	// The subscription value in cents.
	Amount uint64 `json:"amount"`

	// Subscription currency.
	Currency string `json:"currency"`

	// Subscription name.
	Name string `json:"name"`

	// Unique identifier of the subscription on your system.
	ExternalID string `json:"externalId"`
}

// https://api.abacatepay.com/v2/subscriptions/create
type RESTPostCreateSubscriptionData APIResponse[APISubscription]

// https://api.abacatepay.com/v2/subscriptions/list
type RESTGetListSubscriptionsQueryParams struct {
	// Cursor for the pagination.
	Cursor *string `json:"cursor,omitempty"`

	// Number of items per page.
	Limit *uint64 `json:"limit,omitempty"`
}

// https://api.abacatepay.com/v2/subscriptions/list
type RESTGetListSubscriptionsData APIResponseWithCursorBasedPagination[[]APISubscription]
