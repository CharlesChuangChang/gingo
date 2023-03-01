package business

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrorMsg struct {
	ErrorCode   string `json:"error_code"`
	ErrorMsg    string `json:"error_msg"`
	Description string `json:"description"`
}

func RegisterNewUser(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(b, &m)

	email := m["email"]
	password := m["pass"]
	name := m["userName"]
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println(name)
	if email == "" || password == "" || name == "" {
		c.JSON(200, m)
	}
	c.JSON(200, m)
}
