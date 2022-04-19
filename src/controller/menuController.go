package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/errorHandler"
)

type MenuController struct {
	interactor entity.MenuInteractorInterface
}

func NewMenuController(interactor entity.MenuInteractorInterface) *MenuController {
	return &MenuController{interactor: interactor}
}

func (controller *MenuController) IndexHandler(c *gin.Context) {
	result, err := controller.interactor.GetMenus()
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"result": result})
}

func (controller *MenuController) ShowHandler(c *gin.Context) {
	result, err := controller.interactor.GetMenu(c.Param("id"))
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"result": result})
}

func (controller *MenuController) NewHandler(c *gin.Context) {
	type Request struct {
		Name  string `json:"name"`
		Sex   string `json:"sex"`
		Price int    `json:"price"`
		Time  int    `json:"time"`
	}
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		errorHandler.ControllerError(c, errors.New("bad request"))
		return
	}
	result, err := controller.interactor.AddMenu(body.Name, body.Sex, body.Price, body.Time)
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"id": result})
}

func (controller *MenuController) UpdateHandler(c *gin.Context) {
	type Request struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Sex   string `json:"sex"`
		Price int    `json:"price"`
		Time  int    `json:"time"`
	}
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		errorHandler.ControllerError(c, errors.New("bad request"))
		return
	}
	err := controller.interactor.UpdateMenu(body.Id, body.Name, body.Sex, body.Price, body.Time)
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (controller *MenuController) DeleteHandler(c *gin.Context) {
	err := controller.interactor.DeleteMenu(c.Param("id"))
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
