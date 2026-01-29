package v1

// https://docs.abacatepay.com/pages/trustMRR/reference#estrutura
type APIMerchantInfo struct {
	Name      string  `json:"name"`
	Website   *string `json:"website"`
	CreatedAt string  `json:"createdAt"`
}

// https://docs.abacatepay.com/pages/trustMRR/reference#mrr-monthly-recurring-revenue
type APIMRRInfo struct {
	MRR                      uint64 `json:"mrr"`
	TotalActiveSubscriptions uint64 `json:"totalActiveSubscriptions"`
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
