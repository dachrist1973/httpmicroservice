package models

type Account struct {
	AccountNumber      int    `json:"account_number"`
	AccountType        string `json:"account_type"`
	AccountPostalCode  string `json:"account_postal_code"`
	AccountCountryCode string `json:"account_country_code"`
	CustomerNumber     int    `json:"customer_number"`
}
