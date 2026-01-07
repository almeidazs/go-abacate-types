package v1

// https://docs.abacatepay.com/pages/withdraw/reference#estrutura
type APIWithdraw struct {
	// Unique ID of the AbacatePay withdrawal transaction.
	ID string `json:"id"`

	// Current status of the withdrawal transaction.
	Status WithdrawStatus `json:"status"`

	// Indicates whether the loot was created in a development environment (sandbox) or production. AbacatePay currently only supports withdrawals in production.
	DevMode bool `json:"devMode"`

	// Withdrawal transaction receipt URL.
	ReceiptURL string `json:"receiptUrl"`

	// Withdrawal value in cents.
	Amount int `json:"amount"`

	// Platform fee charged for withdrawal in cents.
	PlatformFee int `json:"platformFee"`

	// Unique identifier of the withdrawal in your system. Optional.
	ExternalID *string `json:"externalId,omitempty"`

	// Date and time of withdrawal creation.
	CreatedAt string `json:"createdAt"`

	// Date and time of last withdrawal update.
	UpdatedAt string `json:"updatedAt"`
}

// https://docs.abacatepay.com/pages/withdraw/reference#atributos
type WithdrawStatus string

// https://docs.abacatepay.com/pages/withdraw/reference#atributos
const (
	WithdrawStatusPending   WithdrawStatus = "PENDING"
	WithdrawStatusExpired   WithdrawStatus = "EXPIRED"
	WithdrawStatusCancelled WithdrawStatus = "CANCELLED"
	WithdrawStatusComplete  WithdrawStatus = "COMPLETE"
	WithdrawStatusRefunded  WithdrawStatus = "REFUNDED"
)
