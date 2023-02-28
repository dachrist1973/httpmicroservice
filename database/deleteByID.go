package database

import (
	"fmt"
	"interviewMSrvHTTP/common"
	"interviewMSrvHTTP/models"
	"interviewMSrvHTTP/srverrors"
	"strconv"
)

func DeleteCustomerByID(id string) error {

	var Cust models.Customer

	custid, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf(srverrors.ConversionError)
	}

	result := common.DB.Where("customer_number = ?", custid).Delete(&Cust)

	if result.RowsAffected == 0 {
		return fmt.Errorf(srverrors.RecordNotFound)
	}
	return result.Error
}

func DeleteAccountByID(id string) error {

	var acc models.Account

	accid, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf(srverrors.ConversionError)
	}

	result := common.DB.Where("account_number = ?", accid).Delete(&acc)

	if result.RowsAffected == 0 {
		return fmt.Errorf(srverrors.RecordNotFound)
	}
	return result.Error
}

func DeleteCardByID(id string) error {

	var card models.Card

	cardid, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf(srverrors.ConversionError)
	}

	result := common.DB.Where("payment_card_number = ?", cardid).Delete(&card)

	if result.RowsAffected == 0 {
		return fmt.Errorf(srverrors.RecordNotFound)
	}
	return result.Error
}
