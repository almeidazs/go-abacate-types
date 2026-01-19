package v2

import "time"

// https://docs.abacatepay.com/pages/products/reference#estrutura
type APIProduct struct {
	// The ID of your product.
	ID string

	// Unique product identifier in your system.
	ExternalID string

	// Product name.
	Name string

	// Product price in cents.
	Price uint64

	// Product currency.
	Currency string

	// Product status.
	Status ProductStatus

	// Indicates whether the product was created in a testing environment.
	DevMode bool

	// Product creation date.
	CreatedAt time.Time

	// Product update date.
	UpdatedAt time.Time

	// Product description.
	Description *string
}

type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "ACTIVE"
	ProductStatusInactive ProductStatus = "INACTIVE"
)
