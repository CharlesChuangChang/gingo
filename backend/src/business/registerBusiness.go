package business

import (
	"encoding/json"
	"fmt"
	"gingo/src/db/mangodb"
	"time"

	returnResult "gingo/src/common/result"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ErrorMsg struct {
	ErrorCode   string `json:"error_code"`
	ErrorMsg    string `json:"error_msg"`
	Description string `json:"description"`
}

func GenerateUserID() string {
	return uuid.New().String()
}
func RegisterNewUser(c *gin.Context) {
	fmt.Println("RegisterNewUser")
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(b, &m)

	user := mangodb.UserInfo{
		UserId:    "xxx",
		UserName:  m["userName"].(string),
		Password:  m["pass"].(string),
		Email:     m["email"].(string),
		Token:     "",
		Status:    "Normal",
		HeadImage: "",
		Timepoint: mangodb.TimePoint{StartTime: time.Now().Unix(), EndTime: 0},
	}

	if err := mangodb.RegisterCollectionFindByEmail(mangodb.Client, &user); err != nil {
		fmt.Println(err)
		c.JSON(200, returnResult.Fail("Error while registering find by email"))
	}

	if err := mangodb.RegisterCollectionAdd(mangodb.Client, &user); err != nil {
		fmt.Println(err)
	}

	c.JSON(200, m)
}
