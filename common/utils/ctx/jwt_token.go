package ctx

import (
	"cloud/common/pojo"
	"cloud/error_code"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// 自定义一个字符串
var jwtKey = []byte("maicro@20200924")

// ParseToken @description: 解析token  并返回其中的值
// @parameter tokenString
// @return *jwt.Token
// @return pojo.Claims
// @return error
func ParseToken(tokenString string) (*jwt.Token, pojo.Claims, error) {
	claims := &pojo.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, *claims, err
}

// GetTokenInfo @description: 直接gin路由解析token  并返回其中的值
// @parameter ctx
// @return int
func GetTokenInfo(ctx *gin.Context) int {
	tokenString := ctx.GetHeader("Authorization-Token")
	if tokenString == "" {
		ctx.Abort()
		return error_code.ErrTokenNothing
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil {
		log.Error("token解析出错：", err)
		ctx.Abort()
		return error_code.ErrTokenIssue
	}

	if !token.Valid {
		ctx.Abort()
		return error_code.ErrTokenTimeOut
	}

	ctx.Set("claims", claims)

	return error_code.Success
}

// SettingToken @description: 生成token
// @parameter claims
// @return int
// @return string
func SettingToken(claims *pojo.Claims) (int, string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil || "" == tokenString {
		log.Errorln(err)
		return error_code.ErrTokenSet, ""
	}

	return error_code.Success, tokenString
}
