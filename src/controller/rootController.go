package controller

import "github.com/gin-gonic/gin"

type RootController struct{}

func NewRootController() *RootController {
	return &RootController{}
}

func (controller *RootController) GreetingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "helloWorld"})
}
