package models

type Card struct {
	PaymentCardNumber string `json:"payment_card_number"`
	CreditLimit       int    `json:"credit_limit"`
	AccountNumber     int    `json:"account_number"`
}
