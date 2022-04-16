package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/repository"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/usecase"
)

type BeauticianController struct {
	interactor entity.BeauticianInteractorInterface
}

func NewBeauticianController(db *gorm.DB) *BeauticianController {
	return &BeauticianController{interactor: usecase.NewBeauticianInteractor(repository.NewBeauticianRepository(db))}
}

func (controller *BeauticianController) IndexHandler(c *gin.Context) {
	controller.interactor.GetBeauticians()
	c.JSON(200, gin.H{"message": "Beautician Index Handler"})
}
