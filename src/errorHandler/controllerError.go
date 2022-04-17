package errorHandler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ControllerError(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(500, gin.H{"message": "sorry, internet server error..."})
}
