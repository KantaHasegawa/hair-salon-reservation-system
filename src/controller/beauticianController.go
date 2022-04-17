package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/errorHandler"
)

type BeauticianController struct {
	interactor entity.BeauticianInteractorInterface
}

func NewBeauticianController(interactor entity.BeauticianInteractorInterface) *BeauticianController {
	return &BeauticianController{interactor: interactor}
}

func (controller *BeauticianController) IndexHandler(c *gin.Context) {
	result, err := controller.interactor.GetBeauticians()
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"result": result})
}
