package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type BeauticianController struct {
	interactor entity.BeauticianInteractorInterface
}

func NewBeauticianController(interactor entity.BeauticianInteractorInterface) *BeauticianController {
	return &BeauticianController{interactor: interactor}
}

func (controller *BeauticianController) IndexHandler(c *gin.Context) {
	controller.interactor.GetBeauticians()
	c.JSON(200, gin.H{"message": "Beautician Index Handler"})
}
