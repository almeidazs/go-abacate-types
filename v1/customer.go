package v1

// https://docs.abacatepay.com/pages/client/reference#estrutura
type APICustomer struct {
	// Unique customer identifier.
	ID string `json:"id"`

	// Customer data.
	Metadata APICustomerMetadata `json:"metadata"`
}

// https://docs.abacatepay.com/pages/client/reference#metadados
type APICustomerMetadata struct {
	// Customer's full name.
	Name string `json:"name"`

	// Customer's email.
	Email string `json:"email"`

	// Customer's CPF or CNPJ.
	TaxID string `json:"taxId"`

	// Customer's cell phone.
	Cellphone string `json:"cellphone"`
}
