package v1

// https://docs.abacatepay.com/pages/trustMRR/reference#estrutura
type APIMerchantInfo struct {
	// Store name.
	Name string `json:"name"`

	// Store website.
	Website string `json:"website"`

	// Store creation date.
	CreatedAt string `json:"createdAt"`
}

// https://docs.abacatepay.com/pages/trustMRR/reference#mrr-monthly-recurring-revenue
type APIMRRInfo struct {
	// Monthly recurring revenue in cents. Value 0 indicates that there is no recurring revenue at the moment.
	MRR int `json:"mrr"`

	// Total active subscriptions. Value 0 indicates that there are no currently active subscriptions.
	TotalActiveSubscriptions int `json:"totalActiveSubscriptions"`
}

// https://docs.abacatepay.com/pages/trustMRR/reference#receita-por-per%C3%ADodo
type APIRevenueByPeriod struct {
	TotalRevenue       uint64 `json:"totalRevenue"`
	TotalTransactions  uint64 `json:"totalTransactions"`
	TransactionsPerDay map[string]struct {
		Amount uint64 `json:"amount"`
		Count  uint64 `json:"count"`
	} `json:"transactionsPerDay"`
}
