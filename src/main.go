package main

import (
	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/router"
)

func main() {
	db := database.NewDatabaseHandler()
	defer db.Close()
	r := router.NewRouter(db)
	r.Run()
}
