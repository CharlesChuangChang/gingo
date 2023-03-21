package AuthCheck

import (
	logger "common/myLogger"
	"common/returnResult"
	"errors"
	"fmt"
	"io/ioutil"
	dbModel "myModel/dbModel"
	loginModel "myModel/loginModel"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	TokenExpireTime        = time.Hour * 2
	AutoFailStr     string = "您没有权限访问"
	AutoFailStr2    string = "您没有权限操作"
	AutoFailStr3    string = "用户会话信息不存在"
)

var tokenKeyPriFilePath = "./src/key/token.rsa"
var tokenKeyPubFilePath = "./src/key/token.rsa.pub"

func GetToken(userInfo *dbModel.UserInfo) (string, error) {
	claims := loginModel.MyClaims{
		UserInfo: *userInfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireTime).Unix(),
			Issuer:    "gingo",
		},
	}
	privateKey, err := ioutil.ReadFile(tokenKeyPriFilePath)
	if err != nil {
		pwd, _ := os.Getwd()
		logger.Info(pwd)
		logger.Error("error reading private key file")
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", fmt.Errorf("error parsing RSA private key: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(key)
}

func ParseToken(tokenString string) (*loginModel.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &loginModel.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		publicKey, err := ioutil.ReadFile(tokenKeyPubFilePath)
		if err != nil {
			return nil, fmt.Errorf("error reading public key file: %v", err)
		}

		key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, fmt.Errorf("error parsing RSA public key: %v", err)
		}

		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*loginModel.MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, returnResult.FailWarn(AutoFailStr))
			ctx.Abort()
		} else {
			myClaims, err := ParseToken(token)
			if (err == nil && myClaims != nil) {
				sessionID, _ := strconv.Atoi(myClaims.UserInfo.UserId)
				logger.Info(sessionID)
			} else {
				ctx.JSON(http.StatusUnauthorized, returnResult.FailWarn(AutoFailStr))
				ctx.Abort()
			}
		}
	}
}
