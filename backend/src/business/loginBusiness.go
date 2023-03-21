package business

import (
	"common/returnResult"
	"db/mangodbDriver"
	"encoding/json"
	"fmt"
	loginModel "myModel/dbModel"
	AuthCheck "mylib/auth"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println("Login")
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(b, &m)

	loginInfo := loginModel.LoginInfo{
		Email:    m["Email"].(string),
		Password: m["Password"].(string),
	}

	userData, err := mangodbDriver.LoginCollectionGetUserInfo(mangodbDriver.Client, &loginInfo)
	if err != nil {
		c.JSON(200, returnResult.Fail("Could not find user:"+err.Error()))
		return
	} else {
		if userData != nil {
			//check password
			if loginInfo.Password != userData.Password {
				c.JSON(200, returnResult.Fail("Login fail:Invalid password!"))
			} else {
				//generate or return jenkens
				token, err := AuthCheck.GetToken(userData)
				if err != nil {
					c.JSON(200, returnResult.Fail("Login fail:"+err.Error()))
				}
				rspdata := gin.H{
					"userInfo": userData,
					"token":    token,
				}

				claims, err := AuthCheck.ParseToken(token)
				if err != nil {
					c.JSON(200, returnResult.Fail("Login fail:"+err.Error()))
				}
				fmt.Println(claims)
				c.JSON(200, returnResult.SuccessObject("Login success!", rspdata))
			}
		} else {
			c.JSON(200, returnResult.Fail("Login fail:Get User information fail!"))
		}
	}

}
