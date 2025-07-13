package payload

type Account struct {
	UserID        string  `json:"user_id"`
	Balance       string `json:"balance"`
	CurrencyCode  string  `json:"currency_code"`
	Status        string  `json:"status"`
	Type          string  `json:"type"`
	AccountNumber string  `json:"account_number"`
}

type CreateAccountRequest Account

type AccountsResponse struct {
	Accounts []*Account `json:"accounts"`
}



