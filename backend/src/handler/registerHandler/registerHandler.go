package registerHandler

import (
	"fmt"
	"gingo/src/business"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	fmt.Println("registerHandler")
	business.RegisterNewUser(c)
}
