package registerHandler

import (
	"gingo/src/business"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	business.RegisterNewUser(c)
}
