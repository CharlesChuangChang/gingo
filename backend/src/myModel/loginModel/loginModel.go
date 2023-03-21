package mymodel

import (
	dbModel "myModel/dbModel"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	UserInfo dbModel.UserInfo
	jwt.StandardClaims
}
