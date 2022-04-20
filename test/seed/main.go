package main

import (
	"log"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/beautician"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/customer"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/menu"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/reservation"
)

func main() {
	log.Println("insert seed data start")

	db := database.NewTestDatabaseHandler()

	var err error

	err = beautician.Factory(db)
	err = menu.Factory(db)
	err = customer.Factory(db)
	err = reservation.Factory(db)

	if err != nil {
		panic(err.Error())
	}

	log.Println("insert seed data done")
}
