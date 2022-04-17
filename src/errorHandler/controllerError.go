package errorHandler

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var ErrBadRequest = errors.New("bad request")

func ControllerError(c *gin.Context, err error) {
	if errors.Is(err, ErrBadRequest) {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"message": "not found"})
		return
	}
	log.Println(err)
	c.JSON(500, gin.H{"message": "sorry, internet server error..."})
}
