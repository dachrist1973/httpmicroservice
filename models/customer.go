package models

type Customer struct {
	CustomerNumber      int    `json:"customer_number"`
	CustomerName        string `json:"customer_name"`
	CustomerType        string `json:"customer_type"`
	CustomerPostalCode  string `json:"customer_postal_code"`
	CustomerCountryCode string `json:"customer_country_code"`
	CustomerState       string `json:"customer_state"`
	CustomerCity        string `json:"customer_city"`
	CustomerAddress     string `json:"customer_address"`
	CustomerPhone       string `json:"customer_phone"`
}
