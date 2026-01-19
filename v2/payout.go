package v2

import "time"

// https://docs.abacatepay.com/pages/payouts/reference
type APIPayout struct {
	// Unique transaction identifier.
	ID string

	// Current transaction status.
	Status PayoutStatus

	// Indicates whether the transaction was created in a testing environment.
	DevMode bool

	// Transaction proof URL.
	ReceiptURL *string

	// Payout value in cents.
	Amount uint64

	// Platform fee in cents.
	PlatformFee uint64

	// External transaction identifier.
	ExternalID string

	// Transaction creation date.
	CreatedAt time.Time

	// Transaction update date.
	UpdatedAt time.Time
}

// https://docs.abacatepay.com/pages/payouts/reference#atributos
type PayoutStatus string

// https://docs.abacatepay.com/pages/payouts/reference#atributos
const (
	PayoutStatusPending   PayoutStatus = "PENDING"
	PayoutStatusExpired   PayoutStatus = "EXPIRED"
	PayoutStatusCancelled PayoutStatus = "CANCELLED"
	PayoutStatusComplete  PayoutStatus = "COMPLETE"
	PayoutStatusRefunded  PayoutStatus = "REFUNDED"
)
