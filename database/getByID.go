package database

import (
	"interviewMSrvHTTP/common"
	"interviewMSrvHTTP/models"
	"strconv"
)

func GetCustomerById(id string) (models.Customer, error) {
	var cust models.Customer

	custid, err := strconv.Atoi(id)
	if err != nil {
		cust.CustomerNumber = 0
		return cust, err
	}
	result := common.DB.Where("customer_number = ?", custid).First(&cust)
	if result.Error != nil {
		return cust, result.Error
	}
	return cust, nil
}

func GetAccountById(id string) (models.Account, error) {
	var acc models.Account

	accid, err := strconv.Atoi(id)
	if err != nil {
		acc.AccountNumber = 0
		return acc, err
	}
	common.DB.Where("account_number = ?", accid).First(&acc)

	return acc, nil
}
func GetCardById(id string) (models.Card, error) {
	var card models.Card

	common.DB.Where("payment_card_number = ?", id).First(&card)

	return card, nil
}
