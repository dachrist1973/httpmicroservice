package database

import (
	"httpmicroservice/common"
	"httpmicroservice/models"
)

func CreateCustomer(newCust models.Customer) (models.Customer, error) {

	result := common.DB.Create(&newCust)
	return newCust, result.Error
}

func CreateAccount(newAcc models.Account) (models.Account, error) {

	result := common.DB.Create(&newAcc)
	return newAcc, result.Error
}
func CreateCard(newCard models.Card) (models.Card, error) {

	result := common.DB.Create(&newCard)
	return newCard, result.Error
}
