package errorHandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ControllerError(c *gin.Context, err error) {
	fmt.Println(err)
	c.JSON(500, gin.H{"message": "sorry, internet server error..."})
}
