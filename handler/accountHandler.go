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

type AccRepsonse struct {
	Accounts []models.Account `json: accounts`
	Page     int              `json: page`
}

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	var acc models.Account
	var accresp AccRepsonse
	switch r.Method {
	case "GET":
		values := r.URL.Query()
		wh, limit, page, err := helper.BuildWhereClause(values)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		if pagelimit == 0 {
			pagelimit = limit
		}
		offset := (page - 1) * pagelimit

		accounts, err := database.GetAccounts(wh, limit, offset)
		if err != nil {
			srverrors.HandleError(err, w)
		}
		accresp.Accounts = accounts
		accresp.Page = page
		response, errMarshall := json.Marshal(accresp)
		if errMarshall != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintln(w, string(response))
		}

	case "POST":

		err := json.NewDecoder(r.Body).Decode(&acc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		retacc, err := database.CreateAccount(acc)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(err.Error()))
		}

		response, errMarshall := json.Marshal(retacc)
		if errMarshall != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintln(w, string(response))
		}
		break

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func AccountHandlerWID(w http.ResponseWriter, r *http.Request) {
	var acc, retacc models.Account
	var err error

	params := mux.Vars(r)

	switch r.Method {
	case "GET":
		retacc, err = database.GetAccountById(params["id"])
		if err != nil {
			srverrors.HandleError(err, w)
			return
		}
		response, errMarshall := json.Marshal(retacc)
		if errMarshall != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintln(w, string(response))
		}

	case "POST":
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	case "PUT":
		err = json.NewDecoder(r.Body).Decode(&acc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		acc, err = database.UpdateAccountByID(params["id"], acc)
		if err != nil {
			srverrors.HandleError(err, w)
		}
		response, errMarshall := json.Marshal(acc)
		if errMarshall != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintln(w, string(response))
		}

	case "DELETE":
		err = database.DeleteAccountByID(params["id"])
		if err != nil {
			srverrors.HandleError(err, w)
		}
	}
}
