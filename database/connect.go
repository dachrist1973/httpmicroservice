package database

import (
	"httpmicroservice/common"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	connStr := "port=5432 user=postgres dbname=postgres password=test host=localhost sslmode=disable"
	var err error
	common.DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return common.DB, err
}
