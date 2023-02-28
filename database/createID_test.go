package database

import (
	"interviewMSrvHTTP/common"
	"interviewMSrvHTTP/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var cust models.Customer

type CreateSuite struct {
	suite.Suite
	err error
}

func RunCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateSuite))
}

func (suite *CreateSuite) SetupTest() {
	common.DB, suite.err = ConnectToDB()
}

func Test_CreateCustomer(t *testing.T) {
	var err error
	common.DB, err = ConnectToDB()
	if err != nil {
		return
	}
	tests := []struct {
		Name   string
		Cust   models.Customer
		Result error
	}{
		{
			Name: "Customer 1",
			Cust: models.Customer{
				CustomerNumber:      1,
				CustomerName:        "SAS",
				CustomerType:        "Corporation",
				CustomerPostalCode:  "27513",
				CustomerCountryCode: "011",
				CustomerState:       "NC",
				CustomerCity:        "Cary",
				CustomerAddress:     "SAS Campus Dr",
				CustomerPhone:       "(919) 677-4000",
			},
			Result: nil,
		},
		{
			Name: "Customer 2",
			Cust: models.Customer{
				CustomerNumber:      2,
				CustomerName:        "Jump",
				CustomerType:        "Corporation",
				CustomerPostalCode:  "27513",
				CustomerCountryCode: "011",
				CustomerState:       "NC",
				CustomerCity:        "Cary",
				CustomerAddress:     "SAS Campus Dr",
				CustomerPhone:       "(919) 677-4000",
			},
			Result: nil,
		},
		{
			Name: "Customer 3",
			Cust: models.Customer{
				CustomerNumber:      3,
				CustomerName:        "Doug Christie",
				CustomerType:        "Individual",
				CustomerPostalCode:  "27601",
				CustomerCountryCode: "011",
				CustomerState:       "NC",
				CustomerCity:        "Raleigh",
				CustomerAddress:     "541 E Lenoir St",
				CustomerPhone:       "(919) 531-3711",
			},
			Result: nil,
		},
		{
			Name: "Duplicate of 1",
			Cust: models.Customer{
				CustomerNumber:      1,
				CustomerName:        "SAS",
				CustomerType:        "Corporation",
				CustomerPostalCode:  "27513",
				CustomerCountryCode: "011",
				CustomerState:       "NC",
				CustomerCity:        "Cary",
				CustomerAddress:     "SAS Campus Dr",
				CustomerPhone:       "(919) 677-4000",
			},
			Result: nil,
		},
	}

	for _, test := range tests {

		res, err := CreateCustomer(test.Cust)
		assert.Equal(t, res, test.Cust)
		assert.Equal(t, err, test.Result)
	}
}
