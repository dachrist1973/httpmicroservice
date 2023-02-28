package main

import (
	"interviewMSrvHTTP/common"
	"interviewMSrvHTTP/database"
	"interviewMSrvHTTP/handler"
	"interviewMSrvHTTP/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	err error
}

func (suite *TestSuite) SetupTest() {
	common.DB, suite.err = database.ConnectToDB()
}
func Test_Cusmtomers(t *testing.T) {

	tests := []struct {
		Path   string
		Method string
		Cust   models.Customer
	}{
		{
			Path:   "/customer",
			Method: "POST",
			Cust: models.Customer{
				CustomerNumber: 4,
				CustomerName:   "Foo",
			},
		},
	}

	for _, test := range tests {

		req, err := http.NewRequest(test.Method, test.Path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler.CustomerHandler)

		handler.ServeHTTP(rr, req)
	}

}
