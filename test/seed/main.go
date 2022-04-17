package main

import (
	"log"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/beautician"
)

func main() {
	log.Println("insert seed data start")

	db := database.NewTestDatabaseHandler()

	err := beautician.Factory(db)

	if err != nil {
		panic(err.Error())
	}

	log.Println("insert seed data done")
}
