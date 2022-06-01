//go:build wireinject
// +build wireinject

package main

import (
	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/router"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/utils"
)

func main() {
	utils.LoggingSetting("log/error.log")
	db := database.NewDatabaseHandler()
	r := router.NewRouter(db)
	r.Run()
}
