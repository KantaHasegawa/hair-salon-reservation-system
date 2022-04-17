package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kantahasegawa/hair-salon-reservation-system/di"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/controller"
)

func NewRouter() *gin.Engine {
	rootController := controller.NewRootController()
	beauticianController := di.InitializeBeauticianController()
	r := gin.Default()
	r.GET("/", rootController.GreetingHandler)
	r.GET("/beauticians", beauticianController.IndexHandler)
	r.GET("/beautician/:id", beauticianController.ShowHandler)
	r.POST("/beautician", beauticianController.NewHandler)
	r.PATCH("/beautician", beauticianController.UpdateHandler)
	r.DELETE("/beautician/:id", beauticianController.DeleteHandler)
	return r
}
