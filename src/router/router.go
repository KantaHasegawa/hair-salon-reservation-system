package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/controller"
)

func NewRouter() *gin.Engine {
	rootController := controller.NewRootController()
	r := gin.Default()
	r.GET("/", rootController.GreetingHandler)
	return r
}
