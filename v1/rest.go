package v1

type APIResponse[D any] struct {
	/// The data of the response.
	Data *D `json:"data,omitempty"`

	// Error message returned from the API.
	Error *string `json:"error,omitempty"`
}

// Checks whether the response of the API is an error or not
func (response *APIResponse[D]) IsError() bool {
	return response.Error != nil
}

// https://api.abacatepay.com/v1/customer/create
type RESTPostCreateCustomerBody = APICustomerMetadata

// https://api.abacatepay.com/v1/customer/create
type RESTPostCreateCustomerData = APIResponse[APICustomer]

// https://api.abacatepay.com/v1/billing/create
type RESTPostCreateNewChargeBody struct {
	// Payment methods that will be used. Currently, only `PIX` is supported, `CARD` is in beta.
	Methods []PaymentMethod `json:"methods"`

	// Defines the type of billing frequency. For one-time charges, use "ONE_TIME". For charges that can be paid more than once, use "MULTIPLE_PAYMENTS".
	Frequency PaymentFrequency `json:"frequency"`

	// List of products your customer is paying for.
	Products []RESTChargeProduct `json:"products"`

	// URL to redirect the customer if they click on the "Back" option.
	ReturnURL string `json:"returnUrl"`

	// URL to redirect the customer when payment is completed.
	CompletionURL string `json:"completionUrl"`

	// The ID of a customer already registered in your store.
	CustomerID *string `json:"customerId,omitempty"`

	// Your customer's data to create it.
	// The customer object is not mandatory, but when entering any `customer` information, all `name`, `cellphone`, `email` and `taxId` fields are mandatory.
	Customer *APICustomerMetadata `json:"customer,omitempty"`

	// If true coupons can be used in billing.
	// Default false
	AllowCoupons *bool `json:"allowCoupons,omitempty"`

	// List of coupons available for resem used with billing (0-50 max.).
	Coupons []string `json:"coupons,omitempty"`

	// If you have a unique identifier for your billing application, completely optional.
	ExternalID *string `json:"externalId,omitempty"`

	// Optional billing metadata.
	Metadata map[string]any `json:"metadata,omitempty"`
}

type RESTChargeProduct struct {
	// The product id on your system. We use this id to create your product on AbacatePay automatically, so make sure your id is unique.
	ExternalID string `json:"externalId"`

	// Product name.
	Name string `json:"name"`

	// Quantity of product being purchased.
	Quantity int `json:"quantity"`

	// Price per unit of product in cents. The minimum is 100 (1 BRL).
	Price int `json:"price"`

	// Detailed product description.
	Description *string `json:"description,omitempty"`
}

// https://api.abacatepay.com/v1/billing/create
type RESTPostCreateNewChargeData = APIResponse[APICharge]

// https://api.abacatepay.com/v1/pixQrCode/create
type RESTPostCreateQRCodePixBody struct {
	// Charge amount in cents.
	Amount int `json:"amount"`

	// Billing expiration time in seconds.
	ExpiresIn *int `json:"expiresIn,omitempty"`

	// Message that will appear when paying the PIX.
	Description *string `json:"description,omitempty"`

	// Your customer's data to create it.
	// The customer object is not mandatory, but when entering any `customer` information, all `name`, `cellphone`, `email` and `taxId` fields are mandatory.
	Customer *APICustomerMetadata `json:"customer,omitempty"`

	// Optional billing metadata.
	Metadata map[string]any `json:"metadata,omitempty"`
}

// https://api.abacatepay.com/v1/pixQrCode/create
type RESTPostCreateQRCodePixData = APIResponse[APIQRCodePIX]

// https://api.abacatepay.com/v1/pixQrCode/simulate-payment
type RESTPostSimulatePaymentQueryParams struct {
	// QRCode Pix ID.
	ID string `json:"id"`
}

// https://api.abacatepay.com/v1/pixQrCode/simulate-payment
type RESTPostSimulatePaymentData = APIResponse[APIQRCodePIX]

// https://api.abacatepay.com/v1/pixQrCode/check
type RESTGetCheckQRCodePixStatusQueryParams struct {
	// QRCode Pix ID.
	ID string `json:"id"`
}

// https://api.abacatepay.com/v1/pixQrCode/check
type RESTGetCheckQRCodePixStatusData = APIResponse[RESTQRCodePixStatus]

type RESTQRCodePixStatus struct {
	// QRCode Pix expiration date.
	ExpiresAt string `json:"expiresAt"`

	// Information about the progress of QRCode Pix.
	Status PaymentStatus `json:"status"`
}

// https://api.abacatepay.com/v1/billing/list
type RESTGetListBillingsData = APIResponse[[]APICharge]

// https://api.abacatepay.com/v1/customer/list
type RESTGetListCustomersData = APIResponse[[]APICustomer]

// https://api.abacatepay.com/v1/coupon/create
type RESTPostCreateCouponBody struct {
	// Coupon data.
	Data RESTCreateCouponData `json:"data"`
}

type RESTCreateCouponData struct {
	// Unique coupon identifier.
	Code string `json:"code"`

	// Discount amount to be applied.
	Discount int `json:"discount"`

	// Type of discount applied, percentage or fixed.
	DiscountKind CouponDiscountKind `json:"discountKind"`

	// Coupon description.
	Notes *string `json:"notes,omitempty"`

	// Number of times the coupon can be redeemed. -1 means this coupon can be redeemed without limits.
	// Default -1
	MaxRedeems *int `json:"maxRedeems,omitempty"`

	// Key value object for coupon metadata.
	Metadata map[string]any `json:"metadata,omitempty"`
}

// https://api.abacatepay.com/v1/coupon/create
type RESTPostCreateCouponData = APIResponse[APICoupon]

// https://api.abacatepay.com/v1/coupon/list
type RESTGetListCouponsData = APIResponse[[]APICoupon]

type PixKeyType string

const (
	PixKeyCPF    PixKeyType = "CPF"
	PixKeyCNPJ   PixKeyType = "CNPJ"
	PixKeyPhone  PixKeyType = "PHONE"
	PixKeyEmail  PixKeyType = "EMAIL"
	PixKeyRandom PixKeyType = "RANDOM"
	PixKeyBRCode PixKeyType = "BR_CODE"
)

// https://api.abacatepay.com/v1/withdraw/create
type RESTPostCreateNewWithdrawBody struct {
	// Unique identifier of the withdrawal in your system.
	ExternalID string `json:"externalId"`

	// Always "PIX"
	Method string `json:"method"`

	// Withdrawal value in cents (Min 350).
	Amount int `json:"amount"`

	// PIX key data to receive the withdrawal.
	Pix struct {
		// PIX key type.
		Key string `json:"key"`

		// PIX key value.
		Type PixKeyType `json:"type"`
	} `json:"pix"`

	// Optional description of the withdrawal.
	Description *string `json:"description,omitempty"`
}

// https://api.abacatepay.com/v1/withdraw/create
type RESTPostCreateNewWithdrawData = APIResponse[APIWithdraw]

// https://api.abacatepay.com/v1/withdraw/get
type RESTGetSearchWithdrawQueryParams struct {
	// Unique identifier of the withdrawal in your system.
	ExternalID string `json:"externalId"`
}

type RESTGetSearchWithdrawData = APIResponse[APIWithdraw]

// https://api.abacatepay.com/v1/withdraw/list
type RESTGetListWithdrawsData = APIResponse[[]APIWithdraw]

// https://api.abacatepay.com/v1/public-mrr/revenue
type RESTGetRevenueByPeriodQueryParams struct {
	// Period start date (YYYY-MM-DD format).
	StartDate string `json:"startDate"`

	// Period end date (YYYY-MM-DD format).
	EndDate string `json:"endDate"`
}

// https://api.abacatepay.com/v1/public-mrr/merchant-info
type RESTGetMerchantData = APIResponse[RESTMerchantInfo]

type RESTMerchantInfo struct {
	// Store name.
	Name string `json:"name"`

	// Store website.
	Website string `json:"website"`

	// Store creation date.
	CreatedAt string `json:"createdAt"`
}

// https://api.abacatepay.com/v1/public-mrr/mrr
type RESTGetMRRData = APIResponse[RESTMRRInfo]

type RESTMRRInfo struct {
	// Monthly recurring revenue in cents. Value 0 indicates that there is no recurring revenue at the moment.
	MRR int `json:"mrr"`

	// Total active subscriptions. Value 0 indicates that there are no currently active subscriptions.
	TotalActiveSubscriptions int `json:"totalActiveSubscriptions"`
}

// https://api.abacatepay.com/v1/store/get
type RESTGetStoreDetailsData = APIResponse[APIStore]
