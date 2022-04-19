package dto

type NewAccountRequest struct {
	CustomerId  string  `json:"customerId"`
	AccountType string  `json:"accountType"`
	Amount      float64 `json:"amount"`
}
