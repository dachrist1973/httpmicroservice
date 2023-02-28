package handler

import (
	"encoding/json"
	"fmt"
	"interviewMSrvHTTP/database"
	"interviewMSrvHTTP/helper"
	"interviewMSrvHTTP/models"
	"interviewMSrvHTTP/srverrors"
	"net/http"

	"github.com/gorilla/mux"
)

type CardRepsonse struct {
	Cards []models.Card `json: cards`
	Page  int           `json: page`
}

func CardHandler(w http.ResponseWriter, r *http.Request) {
	var card models.Card
	var cardresp CardRepsonse
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

		cards, err := database.GetCards(wh, limit, offset)
		if err != nil {
			srverrors.HandleError(err, w)
		}
		cardresp.Cards = cards
		cardresp.Page = page
		response, errMarshall := json.Marshal(cardresp)
		if errMarshall != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintln(w, string(response))
		}

	case "POST":

		err := json.NewDecoder(r.Body).Decode(&card)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		retcard, err := database.CreateCard(card)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(err.Error()))
		}

		response, errMarshall := json.Marshal(retcard)
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

func CardHandlerWID(w http.ResponseWriter, r *http.Request) {
	var card, retcard models.Card
	var err error
	//Get the Id path parameter
	params := mux.Vars(r)

	switch r.Method {
	case "GET":
		retcard, err = database.GetCardById(params["id"])
		if err != nil {
			srverrors.HandleError(err, w)
			return
		}
		response, errMarshall := json.Marshal(retcard)
		if errMarshall != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintln(w, string(response))
		}

	case "POST":
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	case "PUT":
		err = json.NewDecoder(r.Body).Decode(&card)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		card, err = database.UpdateCardByID(params["id"], card)
		if err != nil {
			srverrors.HandleError(err, w)
		}
		response, errMarshall := json.Marshal(card)
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
