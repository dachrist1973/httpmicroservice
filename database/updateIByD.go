package database

import (
	"fmt"
	"httpmicroservice/common"
	"httpmicroservice/models"
	"httpmicroservice/srverrors"
	"strconv"
)

func UpdateCustomerByID(id string, cust models.Customer) (models.Customer, error) {

	var foundCust models.Customer

	custid, err := strconv.Atoi(id)
	if err != nil {
		return foundCust, fmt.Errorf(srverrors.ConversionError)
	}
	result := common.DB.Where("customer_number = ?", custid).First(&foundCust)
	if result.Error != nil {
		return foundCust, fmt.Errorf(srverrors.RecordNotFound)
	}

	foundCust = updateCustValues(foundCust, cust)

	result = common.DB.Where("customer_number = ?", custid).Save(&foundCust)

	return foundCust, result.Error
}

func updateCustValues(foundCust, cust models.Customer) models.Customer {

	if cust.CustomerAddress != "" {
		foundCust.CustomerAddress = cust.CustomerAddress
	}
	if cust.CustomerCity != "" {
		foundCust.CustomerCity = cust.CustomerCity
	}
	if cust.CustomerCountryCode != "" {
		foundCust.CustomerCountryCode = cust.CustomerCountryCode
	}
	if cust.CustomerName != "" {
		foundCust.CustomerName = cust.CustomerName
	}
	if cust.CustomerPostalCode != "" {
		foundCust.CustomerPostalCode = cust.CustomerPostalCode
	}
	if cust.CustomerState != "" {
		foundCust.CustomerState = cust.CustomerState
	}
	if cust.CustomerType != "" {
		foundCust.CustomerType = cust.CustomerType
	}
	if cust.CustomerPhone != "" {
		foundCust.CustomerPhone = cust.CustomerPhone
	}
	return foundCust

}

func UpdateAccountByID(id string, cust models.Account) (models.Account, error) {

	var foundAcc models.Account

	accnum, err := strconv.Atoi(id)
	if err != nil {
		return foundAcc, fmt.Errorf(srverrors.ConversionError)
	}
	result := common.DB.Where("account_number = ?", accnum).First(&foundAcc)
	if result.Error != nil {
		return foundAcc, fmt.Errorf(srverrors.RecordNotFound)
	}

	foundAcc = updateAccountValues(foundAcc, cust)

	result = common.DB.Where("account_number = ?", accnum).Save(&foundAcc)

	return foundAcc, result.Error
}

func updateAccountValues(foundAcc, acc models.Account) models.Account {

	if acc.AccountCountryCode != "" {
		foundAcc.AccountCountryCode = acc.AccountCountryCode
	}
	if acc.AccountPostalCode != "" {
		foundAcc.AccountPostalCode = acc.AccountPostalCode
	}
	if acc.AccountType != "" {
		foundAcc.AccountType = acc.AccountType
	}
	return foundAcc

}

func UpdateCardByID(id string, card models.Card) (models.Card, error) {

	var foundCard models.Card

	result := common.DB.Where("payment_card_number = ?", id).First(&foundCard)
	if result.Error != nil {
		return foundCard, fmt.Errorf(srverrors.RecordNotFound)
	}

	foundCard = updateCardValues(foundCard, card)

	result = common.DB.Where("payment_card_number = ?", id).Save(&foundCard)

	return foundCard, result.Error
}

func updateCardValues(foundCard, card models.Card) models.Card {

	if card.CreditLimit != foundCard.CreditLimit {
		foundCard.CreditLimit = card.CreditLimit
	}

	return foundCard

}
