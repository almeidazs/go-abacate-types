package v1

// https://docs.abacatepay.com/pages/webhooks#eventos-suportados
type WebhookEventType string

// https://docs.abacatepay.com/pages/webhooks#eventos-suportados
const (
	WebhookEventWithdrawFailed WebhookEventType = "withdraw.failed"
	WebhookEventWithdrawDone   WebhookEventType = "withdraw.done"
	WebhookEventBillingPaid    WebhookEventType = "billing.paid"
)

type BaseWebhookEvent struct {
	// Unique identifier for the webhook.
	ID string `json:"id"`

	// Event type.
	Event WebhookEventType `json:"event"`

	// Indicates whether the event occurred in dev mode.
	DevMode bool `json:"devMode"`
}

// https://docs.abacatepay.com/pages/webhooks#withdraw-failed
type WebhookWithdrawFailedEvent struct {
	BaseWebhookEvent

	Data struct {
		// Status is always CANCELLED.
		Transaction APIWithdraw `json:"transaction"`
	} `json:"data"`
}

// https://docs.abacatepay.com/pages/webhooks#withdraw-done
type WebhookWithdrawDoneEvent struct {
	BaseWebhookEvent

	Data struct {
		// Status is always COMPLETE.
		Transaction APIWithdraw `json:"transaction"`
	} `json:"data"`
}

type WebhookBillingPaidPayment struct {
	Amount int           `json:"amount"`
	Fee    int           `json:"fee"`
	Method PaymentMethod `json:"method"`
}

type WebhookBillingPaidPix struct {
	Amount int           `json:"amount"`
	ID     string        `json:"id"`
	Kind   PaymentMethod `json:"kind"`   // Always PIX
	Status PaymentStatus `json:"status"` // Always PAID
}

type WebhookBillingPaidBilling struct {
	Amount      int              `json:"amount"`
	CouponsUsed []string         `json:"couponsUsed"`
	Customer    APICustomer      `json:"customer"`
	Frequency   PaymentFrequency `json:"frequency"`
	ID          string           `json:"id"`

	// @unstable
	Kind []PaymentMethod `json:"kind"`

	// @unstable
	PaidAmount int `json:"paidAmount"`

	Products []APIProduct `json:"products"`

	Status PaymentStatus `json:"status"` // Always PAID
}

// https://docs.abacatepay.com/pages/webhooks#billing-paid
type WebhookBillingPaidEvent struct {
	BaseWebhookEvent

	Data struct {
		Payment WebhookBillingPaidPayment `json:"payment"`

		// Exactly one of these will be non-nil

		PixQrCode *WebhookBillingPaidPix     `json:"pixQrCode,omitempty"`
		Billing   *WebhookBillingPaidBilling `json:"billing,omitempty"`
	} `json:"data"`
}

type WebhookEnvelope struct {
	Event WebhookEventType `json:"event"`
}

func (event WebhookEnvelope) IsWithdrawFailed() bool {
	return event.Event == WebhookEventWithdrawFailed
}

func (event WebhookEnvelope) IsWithdrawDone() bool {
	return event.Event == WebhookEventWithdrawDone
}

func (event WebhookEnvelope) IsBillingPaid() bool {
	return event.Event == WebhookEventBillingPaid
}

func (event WebhookBillingPaidEvent) IsPix() bool {
	return event.Data.PixQrCode != nil
}

func (event WebhookBillingPaidEvent) IsBilling() bool {
	return event.Data.Billing != nil
}
