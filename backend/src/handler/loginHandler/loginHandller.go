package loginHandler

import (
	"business"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	fmt.Println("LoginHandler")
	business.Login(c)
}
