package controller

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/errorHandler"
)

type ReservationController struct {
	interactor entity.ReservationInteractorInterface
}

func NewReservationController(interactor entity.ReservationInteractorInterface) *ReservationController {
	return &ReservationController{interactor: interactor}
}

func (controller *ReservationController) IndexHandler(c *gin.Context) {
	result, err := controller.interactor.GetReservations(c)
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"result": result})
}

func (controller *ReservationController) ShowHandler(c *gin.Context) {
	result, err := controller.interactor.GetReservation(c, c.Param("id"))
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"result": result})
}

func (controller *ReservationController) NewHandler(c *gin.Context) {
	type Request struct {
		CustomerId   string    `json:"customer_id"`
		BeauticianId string    `json:"beautician_id"`
		MenuId       string    `json:"menu_id"`
		StartTime    time.Time `json:"start_time"`
	}
	var body Request
	if err := c.ShouldBindJSON(&body); err != nil {
		errorHandler.ControllerError(c, errors.New("bad request"))
		return
	}
	result, err := controller.interactor.AddReservation(c, body.CustomerId, body.BeauticianId, body.MenuId, body.StartTime)
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"id": result})
}

func (controller *ReservationController) DeleteHandler(c *gin.Context) {
	err := controller.interactor.DeleteReservation(c, c.Param("id"))
	if err != nil {
		errorHandler.ControllerError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
