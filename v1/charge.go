package v1

// https://docs.abacatepay.com/pages/payment/reference#estrutura
type APICharge struct {
	// Unique billing ID at AbacatePay.
	ID string `json:"id"`

	// Billing frequency. It can be "ONE_TIME" or "MULTIPLE_PAYMENTS".
	Frequency PaymentFrequency `json:"frequency"`

	// URL for your customer to make payment for the charge.
	URL string `json:"url"`

	// Billing status. Can be "PENDING", "EXPIRED", "CANCELLED", "PAID", "REFUNDED".
	Status PaymentStatus `json:"status"`

	// Indicates whether the charge was created in a development (true) or production (false) environment.
	DevMode bool

	// Supported payment types: "PIX" and "CARD" (beta).
	Methods []PaymentMethod `json:"methods"`

	// List of products included in the charge.
	Products []APIProduct `json:"products"`

	// Customer you are billing. Optional. See structure reference [here](https://docs.abacatepay.com/pages/payment/client/reference.mdx).
	Customer APICustomer `json:"customer"`

	// Object with metadata about the charge.
	Metadata APIChargeMetadata `json:"metadata"`

	// Date and time of next charge, or null for one-time charges.
	NextBilling *string `json:"nextBilling,omitempty"`

	// Whether or not to allow coupons for billing.
	AllowCoupons *bool `json:"allowCoupons,omitempty"`

	// Coupons allowed in billing. Coupons are only considered if "AllowCoupons" is true.
	Coupons []APICoupon `json:"coupons"`

	// Charge creation date and time.
	CreatedAt string `json:"createdAt"`

	// Charge last updated date and time.
	UpdatedAt string `json:"updatedAt"`
}

// https://docs.abacatepay.com/pages/payment/reference#metadata
type APIChargeMetadata struct {
	// Fee applied by AbacatePay.
	Fee int `json:"fee"`

	// URL that the customer will be redirected to when clicking the “back” button.
	ReturnURL string `json:"returnUrl"`

	// URL that the customer will be redirected to when making payment.
	CompletionURL string `json:"completionUrl"`
}

// https://docs.abacatepay.com/pages/payment/reference#methods
type PaymentMethod string

// https://docs.abacatepay.com/pages/payment/reference#methods
const (
	PaymentMethodPix  PaymentMethod = "PIX"
	PaymentMethodCard PaymentMethod = "CARD"
)

// https://docs.abacatepay.com/pages/payment/reference#frequency
type PaymentFrequency string

// https://docs.abacatepay.com/pages/payment/reference#frequency
const (
	PaymentFrequencyOneTime          PaymentFrequency = "ONE_TIME"
	PaymentFrequencyMultiplePayments PaymentFrequency = "MULTIPLE_PAYMENTS"
)

// https://docs.abacatepay.com/pages/payment/reference#products
type APIProduct struct {
	// The product id on your system. We use this id to create your product on AbacatePay automatically, so make sure your id is unique.
	ExternalID string `json:"externalId"`

	// Product name.
	Name string `json:"name"`

	// Quantity of product being purchased.
	Quantity int `json:"quantity"`

	// Price per unit of product in cents. The minimum is 100 (1 BRL).
	Price int `json:"price"`

	Description *string `json:"description,omitempty"`
}
