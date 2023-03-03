package handler

import (
	"encoding/json"
	"fmt"
	"httpmicroservice/database"
	"httpmicroservice/helper"
	"httpmicroservice/models"
	"httpmicroservice/srverrors"
	"net/http"

	"github.com/gorilla/mux"
)

var pagelimit int = 50

type CustRepsonse struct {
	Customers []models.Customer `json: customers`
	Page      int               `json: page`
}

func CustomerHandler(w http.ResponseWriter, r *http.Request) {

	var cust models.Customer
	var custresp CustRepsonse

	switch r.Method {
	case "GET":
		values := r.URL.Query()
		wh, limit, page, err := helper.BuildWhereClause(values)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		if limit > 0 {
			pagelimit = limit
		}
		offset := (page - 1) * pagelimit 
		customers, err := database.GetAllCustomers(wh, limit, offset)
		if err != nil {
			srverrors.HandleError(err, w)
		}
		custresp.Customers = customers
		custresp.Page = page
		response, errMarshall := json.Marshal(custresp)
		if errMarshall != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintln(w, string(response))
		}

	case "POST":
		err := json.NewDecoder(r.Body).Decode(&cust)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		retcust, err := database.CreateCustomer(cust)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(err.Error()))
		}

		response, errMarshall := json.Marshal(retcust)
		if errMarshall != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintln(w, string(response))
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func CustomerHandlerWID(w http.ResponseWriter, r *http.Request) {

	var cust, retcust models.Customer
	var err error
	//Get the Id path parameter
	params := mux.Vars(r)

	values := r.URL.Query()
	helper.BuildWhereClause(values)
	// Switch on the Crud Operation.
	switch r.Method {

	case "GET":
		retcust, err = database.GetCustomerById(params["id"])
		if err != nil {
			srverrors.HandleError(err, w)
			return
		}
		response, errMarshall := json.Marshal(retcust)
		if errMarshall != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintln(w, string(response))
		}

	case "POST":
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	case "PUT":

		err = json.NewDecoder(r.Body).Decode(&cust)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		cust, err = database.UpdateCustomerByID(params["id"], cust)
		if err != nil {
			srverrors.HandleError(err, w)
		}
		response, errMarshall := json.Marshal(cust)
		if errMarshall != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintln(w, string(response))
		}

	case "DELETE":

		err = database.DeleteCustomerByID(params["id"])
		if err != nil {
			srverrors.HandleError(err, w)
		}

	}

}
