package registerHandler

import (
	"business"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	fmt.Println("registerHandler")
	business.RegisterNewUser(c)
}
