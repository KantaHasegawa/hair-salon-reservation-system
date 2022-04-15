package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/repository"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/usecase"
)

type BeauticianController struct {
	interactor entity.BeauticianInteractorInterface
}

func NewBeauticianController() *BeauticianController {
	return &BeauticianController{interactor: usecase.NewBeauticianInteractor(repository.NewBeauticianRepository())}
}

func (controller *BeauticianController) IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Beautician Index Handler"})
}
