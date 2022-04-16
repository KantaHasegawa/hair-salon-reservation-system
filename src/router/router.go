package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/controller"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	rootController := controller.NewRootController()
	beauticianController := controller.NewBeauticianController(db)
	r := gin.Default()
	r.GET("/", rootController.GreetingHandler)
	r.GET("/beauticians", beauticianController.IndexHandler)
	return r
}
