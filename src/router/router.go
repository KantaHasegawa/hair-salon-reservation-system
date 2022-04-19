package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kantahasegawa/hair-salon-reservation-system/di"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/controller"
)

func NewRouter() *gin.Engine {
	rootController := controller.NewRootController()
	beauticianController := di.InitializeBeauticianController()
	menuController := di.InitializeMenuController()
	r := gin.Default()
	r.GET("/", rootController.GreetingHandler)
	r.GET("/beauticians", beauticianController.IndexHandler)
	r.GET("/beautician/:id", beauticianController.ShowHandler)
	r.POST("/beautician", beauticianController.NewHandler)
	r.PATCH("/beautician", beauticianController.UpdateHandler)
	r.DELETE("/beautician/:id", beauticianController.DeleteHandler)
	r.GET("/menus", menuController.IndexHandler)
	r.GET("/menu/:id", menuController.ShowHandler)
	r.POST("/menu", menuController.NewHandler)
	r.PATCH("/menu", menuController.UpdateHandler)
	r.DELETE("/menu/:id", menuController.DeleteHandler)
	return r
}
