package main

import (
	"github.com/kantahasegawa/hair-salon-reservation-system/src/router"
)

func main() {
	r := router.NewRouter()
	r.Run()
}
