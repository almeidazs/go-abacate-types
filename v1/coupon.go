package v1

// https://docs.abacatepay.com/pages/coupon/reference#estrutura
type APICoupon struct {
	// Unique coupon code that your customers will use to apply the discount.
	ID string `json:"id"`

	// Type of discount applied by the coupon.
	DiscountKind CouponDiscountKind `json:"discountKind"`

	// Discount amount.
	// For `PERCENTAGE` use numbers from 1-100 (ex: 10 = 10%).
	// For `FIXED` use the value in cents (ex: 500 = R$5.00).
	Discount int `json:"discount"`

	// Limit on the number of times the coupon can be used.
	// Use `-1` for unlimited coupons.
	MaxRedeems int `json:"maxRedeems"`

	// Counter of how many times the coupon has been used by customers.
	RedeemsCount int `json:"redeemsCount"`

	// Coupon status.
	Status CouponStatus `json:"status"`

	// Indicates whether the coupon was created in a development (true)
	// or production (false) environment.
	DevMode bool `json:"devMode"`

	// Internal description of the coupon for your organization and control.
	Notes *string `json:"notes,omitempty"`

	// Coupon creation date and time (ISO 8601).
	CreatedAt string `json:"createdAt"`

	// Coupon last updated date and time (ISO 8601).
	UpdatedAt string `json:"updatedAt"`

	// Additional custom metadata.
	Metadata map[string]any `json:"metadata,omitempty"`
}

// https://docs.abacatepay.com/pages/coupon/reference#discountkind
type CouponDiscountKind string

// https://docs.abacatepay.com/pages/coupon/reference#discountkind
const (
	CouponDiscountKindFixed      CouponDiscountKind = "FIXED"
	CouponDiscountKindPercentage CouponDiscountKind = "PERCENTAGE"
)

// https://docs.abacatepay.com/pages/coupon/reference#status
type CouponStatus string

// https://docs.abacatepay.com/pages/coupon/reference#status
const (
	CouponStatusActive   CouponStatus = "ACTIVE"
	CouponStatusDeleted  CouponStatus = "DELETED"
	CouponStatusDisabled CouponStatus = "DISABLED"
)
