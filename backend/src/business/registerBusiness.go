package business

import (
	"common/returnResult"
	"db/mangodbDriver"
	"encoding/json"
	"fmt"
	loginModel "myModel/dbModel"
	"time"

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

	user := loginModel.UserInfo{
		UserId:    "xxx",
		UserName:  m["userName"].(string),
		Password:  m["pass"].(string),
		Email:     m["email"].(string),
		Token:     "",
		Status:    "Normal",
		HeadImage: "",
		Timepoint: loginModel.TimePoint{StartTime: time.Now().Unix(), EndTime: 0},
	}

	bExist, _ := mangodbDriver.RegisterCollectionFindByEmail(mangodbDriver.Client, &user)
	if bExist {
		c.JSON(200, returnResult.Fail("Email already exists"))
		return
	}

	if err := mangodbDriver.RegisterCollectionAdd(mangodbDriver.Client, &user); err != nil {
		fmt.Println(err)
		c.JSON(200, returnResult.Fail(err.Error()))
		return
	}

	c.JSON(200, returnResult.SuccessObject("Register success!", m))
}
