package database

import (
	"interviewMSrvHTTP/common"
	"interviewMSrvHTTP/models"

	"gorm.io/gorm"
)

func GetAllCustomers(wh string, limit, offset int) ([]models.Customer, error) {
	var cust []models.Customer
	var result *gorm.DB

	if limit != 0 && offset != 0 {
		result = common.DB.Offset(offset).Limit(limit).Find(&cust)
	} else if limit != 0 {
		result = common.DB.Limit(limit).Find(&cust)
	} else {
		result = common.DB.Where(wh).First(&cust)
	}

	return cust, result.Error
}

func GetAccounts(wh string, limit, offset int) ([]models.Account, error) {
	var acc []models.Account
	var result *gorm.DB
	if limit != 0 && offset != 0 {
		result = common.DB.Offset(offset).Limit(limit).Find(&acc)
	} else if limit != 0 {
		result = common.DB.Limit(limit).Find(&acc)
	} else {
		result = common.DB.Where(wh).First(&acc)
	}

	return acc, result.Error
}
func GetCards(wh string, limit, offset int) ([]models.Card, error) {

	var card []models.Card
	var result *gorm.DB

	if limit != 0 && offset != 0 {
		result = common.DB.Offset(offset).Limit(limit).Find(&card)
	} else if limit != 0 {
		result = common.DB.Limit(limit).Find(&card)
	} else {
		result = common.DB.Where(wh).First(&card)
	}

	return card, result.Error
}
