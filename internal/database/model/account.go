package model

import "github.com/shopspring/decimal"

type Account struct {
	UserId string `db:"user_id"`
	Balance decimal.Decimal `db:"balance"`	
	CurrencyCode string `db:"currency_code"`
	Status string `db:"status"`
	Type   string  `db:"type"`
	AccountNumber string `db:"account_number"`
}