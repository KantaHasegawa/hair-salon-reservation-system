package main

import (
	"log"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/beautician"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/menu"
)

func main() {
	log.Println("insert seed data start")

	db := database.NewTestDatabaseHandler()

	var err error

	err = beautician.Factory(db)
	err = menu.Factory(db)

	if err != nil {
		panic(err.Error())
	}

	log.Println("insert seed data done")
}
