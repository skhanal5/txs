package model

type Account struct {
	UserId string `db:"user_id"`
	Balance float64 `db:"balance"`	
	CurrencyCode string `db:"currency_code"`
	Status string `db:"status"`
	Type   string  `db:"type"`
	AccountNumber string `db:"account_number"`
}