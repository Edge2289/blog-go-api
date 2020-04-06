package jwt

import (
	"blog-go-api/app/config"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"time"
)

//type MapClaims map[string]interface{}
type UserParam struct {
	uid string
}

/**
 生成token
 */
func GetJwtTokenMiddleware(uid string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix() // 设置超时时间
	claims["iat"] = time.Now().Unix()
	// 存储所需要的内容
	claims["uid"] = uid
	token.Claims = claims

	if tokenString, err := token.SignedString([]byte(config.JwtSecretKey)); err == nil {
		return tokenString, err
	} else {
		return "", errors.New("create token error !")
	}
}


//JwtParseUser 解析payload的内容,得到用户信息
func JwtParseUser(token string) (UserParam, error) {
	claims := make(jwt.MapClaims)
	_, _ = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, err error) {
		//fmt.Println(token)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JwtSecretKey), nil
	})
	return GetIdFromClaims("uid", claims), nil
}

/**
 获取参数里面的值
 */
// 示例 ：GetIdFromClaims("username", token.claims) 其中token是已经解密的token
func GetIdFromClaims(key string, claims jwt.Claims) UserParam {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				uid := fmt.Sprintf("%v", value.Interface())
				return UserParam {
					uid: uid,
				}
			}
		}
	}
	return UserParam{}
}