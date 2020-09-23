package jwt

import (
	"blog-go-api/app/config"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"time"
)

type User struct {
	uid  int
	role interface{}
}

/**
 生成token
userData
map[string]interface{}
*/
func GetJwtTokenMiddleware(uid int) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	/**
	Audience:  "Username",        // 受众
	ExpiresAt: expiresTime,       // 失效时间
	Id:        "123123",          // 编号
	IssuedAt:  time.Now().Unix(), // 签发时间
	Issuer:    "gin hello",       // 签发人
	NotBefore: time.Now().Unix(), // 生效时间
	Subject:   "login",           // 主题
	*/
	claims := make(jwt.MapClaims)
	claims["aud"] = config.AppName                   // 受众
	claims["exp"] = time.Now().Add(time.Hour).Unix() // 失效时间 设置超时时间 一分钟 * time.Duration(1)
	claims["iat"] = time.Now().Unix()                // 签发时间
	claims["iss"] = config.AppName                   // 签发人
	claims["nbf"] = time.Now().Unix()                // 生效时间
	claims["sub"] = "go-api-sub"                     // 主题
	claims["id"] = "admin id"                        // userData["uid"]

	// 存储所需要的内容
	claims["uid"] = uid
	token.Claims = claims

	if tokenString, err := token.SignedString([]byte(config.JwtSecretKey)); err == nil {
		return tokenString, err
	} else {
		return "", errors.New("create token error !")
	}
}

/**
JwtParseUser 解析payload的内容,得到用户信息
 */
func JwtParseUser(c *gin.Context) (interface{}, error) {
	claims := make(jwt.MapClaims)
	tokenString, errToken := ExtractToken(c)
	if errToken != "" {
		return nil, errors.New(errToken)
	}

	// 解析token
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.JwtSecretKey), nil
	})

	if err != nil {
		return nil, errors.New("-1002")
	}
	//token, err := RefaceJwtToken(claims)
	//if err == nil {
	//	// 设置token header 头
	//	c.Header("Authorization", "Bearer "+token)
	//}
	return claims["uid"], nil
}

/**
  刷新token
*/
func RefaceJwtToken(claims jwt.MapClaims) (string, error) {

	newToken := jwt.New(jwt.SigningMethodHS256)
	claims["exp"] = time.Now().Add(time.Hour).Unix() // 失效时间 设置超时时间 一分钟 * time.Duration(1)
	// 存储所需要的内容
	newToken.Claims = claims
	if tokenString, err := newToken.SignedString([]byte(config.JwtSecretKey)); err == nil {
		return tokenString, err
	} else {
		return "", errors.New("create token error !")
	}
}

/**
  获取token
*/
func ExtractToken(c *gin.Context) (string, string) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		return "", "auth为空"
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	fmt.Println(parts)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return parts[1], "-1001"
	}
	return parts[1], ""
}

func GetIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
