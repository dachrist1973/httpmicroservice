package main

import (
	"fmt"
	"net/http"

	"httpmicroservice/database"
	"httpmicroservice/handler"

	//"httpmicroservice/models"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {

	db, err := database.ConnectToDB()
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	//Generate Dummy Data
	// for i := 1; i <= 50; i++ {
	// 	db.Create(&models.Customer{
	// 		CustomerNumber:      i,
	// 		CustomerName:        fmt.Sprintf("Customer%d", i),
	// 		CustomerType:        "Corporation",
	// 		CustomerPostalCode:  "27601",
	// 		CustomerCountryCode: "011",
	// 		CustomerCity:        "Raleigh",
	// 		CustomerState:       "NC",
	// 		CustomerPhone:       "9195393711",
	// 	})
	// 	db.Create(&models.Account{
	// 		AccountNumber:      i + 10,
	// 		AccountType:        "Checking",
	// 		AccountPostalCode:  "27601",
	// 		AccountCountryCode: "011",
	// 		CustomerNumber:     i,
	// 	})
	// 	db.Create(&models.Card{
	// 		AccountNumber:     i + 10,
	// 		PaymentCardNumber: fmt.Sprintf("%d23456781234567", i),
	// 		CreditLimit:       300 * i,
	// 	})

	// }

	// List all available end points.
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %s!", r.URL.Path[1:])
	})
	r.HandleFunc("/account", handler.AccountHandler)

	r.HandleFunc("/account/{id}", handler.AccountHandlerWID)

	r.HandleFunc("/customer", handler.CustomerHandler)

	r.HandleFunc("/customer/{id}", handler.CustomerHandlerWID)
	r.HandleFunc("/card", handler.CardHandler)

	r.HandleFunc("/card/{id}", handler.CardHandlerWID)
	fmt.Println("Starting Server on port 8080")
	http.ListenAndServe(":8080", r)

}
