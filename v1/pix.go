package v1

// https://docs.abacatepay.com/pages/pix-qrcode/reference#estrutura
type APIQRCodePIX struct {
	// Unique billing identifier.
	ID string `json:"id"`

	// Charge amount in cents (e.g. 4000 = R$40.00).
	Amount int `json:"amount"`

	// Billing status. Can be "PENDING", "EXPIRED", "CANCELLED", "PAID", "REFUNDED".
	Status PaymentStatus `json:"status"`

	// Indicates whether the charge is in a testing (true) or production (false) environment.
	DevMode bool `json:"devMode"`

	// PIX code (copy-and-paste) for payment.
	BrCode string `json:"brCode"`

	// PIX code in Base64 format (Useful for displaying in images).
	BrCodeBase64 string `json:"brCodeBase64"`

	// Platform fee in cents. Example: 80 means R$0.80.
	PlatformFee int `json:"platformFee"`

	// Payment description.
	Description string `json:"description"`

	// Payment metadata.
	Metadata map[string]any `json:"metadata,omitempty"`

	// QRCode PIX creation date and time.
	CreatedAt string `json:"createdAt"`

	// QRCode PIX last updated date and time.
	UpdatedAt string `json:"updatedAt"`

	// QRCode expiration date and time.
	ExpiresAt string `json:"expiresAt"`
}

// // https://docs.abacatepay.com/pages/payment/reference#atributos
type PaymentStatus string

// https://docs.abacatepay.com/pages/payment/reference#atributos
const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusExpired   PaymentStatus = "EXPIRED"
	PaymentStatusCancalled PaymentStatus = "CANCELLED"
	PaymentStatusPaid      PaymentStatus = "PAID"
	PaymentStatusRefunded  PaymentStatus = "REFUNDED"
)
