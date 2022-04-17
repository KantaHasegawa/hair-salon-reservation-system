package controller

import (
	"errors"

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

func (controller *BeauticianController) ShowHandler(c *gin.Context) {
	result, err := controller.interactor.GetBeautician(c.Param("id"))
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"result": result})
}

func (controller *BeauticianController) NewHandler(c *gin.Context) {
	type Request struct {
		Name  string `json:"name"`
		Sex   string `json:"sex"`
		Price int    `json:"price"`
	}
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		errorHandler.ControllerError(c, errors.New("bad request"))
		return
	}
	result, err := controller.interactor.AddBeautician(body.Name, body.Sex, body.Price)
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"id": result})
}

func (controller *BeauticianController) UpdateHandler(c *gin.Context) {
	type Request struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Sex   string `json:"sex"`
		Price int    `json:"price"`
	}
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		errorHandler.ControllerError(c, errors.New("bad request"))
		return
	}
	err := controller.interactor.UpdateBeautician(body.Id, body.Name, body.Sex, body.Price)
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (controller *BeauticianController) DeleteHandler(c *gin.Context) {
	err := controller.interactor.DeleteBeautician(c.Param("id"))
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
