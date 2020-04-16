package jwt
//
//import (
//	"blog-go-api/app/config"
//	"errors"
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"time"
//)
//
//type JWT struct {
//	SigningKey []byte
//}
//
//var (
//	TokenExpired     error  = errors.New("Token is expired")
//	TokenNotValidYet error  = errors.New("Token not active yet")
//	TokenMalformed   error  = errors.New("That's not even a token")
//	TokenInvalid     error  = errors.New("Couldn't handle this token:")
//	SignKey          string = "hlms_yeeuu"
//)
//
//// 自定义结构体参数
//type CustomClaims struct {
//	User      string `form:"phone"`
//	Level     string `form:"type"`
//	HotelId   string `form:"hotel_id"`
//	HotelName string `form:"hotelName"`
//	UserName  string `form:"username"`
//	jwt.StandardClaims
//}
//
//// JWT验证
//func JWTAuth() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		token, err := c.Cookie("claims")
//		if err != nil {
//			c.Redirect(302, "/login")
//			return
//		}
//		j := NewJWT()
//		// token := t.(string)
//		// claims, err := j.ParseToken(token)
//		claims, err := j.ParseToken(token)
//		if err != nil {
//			if err == TokenExpired {
//				if token, err = j.RefreshToken(token); err == nil {
//					c.JSON(http.StatusOK, gin.H{"error": 0, "message": "refresh token", "token": token})
//					return
//				}
//			}
//			c.Redirect(302, "/login")
//			return
//		}
//		c.Set("claims", claims)
//		c.Next()
//		return
//	}
//}
//
//func NewJWT() *JWT {
//	return &JWT{
//		[]byte(GetSignKey()),
//	}
//}
//func GetSignKey() string {
//	return SignKey
//}
//func SetSignKey(key string) string {
//	SignKey = key
//	return SignKey
//}
//
//// parse
//func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return j.SigningKey, nil
//	})
//	if err != nil {
//		if ve, ok := err.(*jwt.ValidationError); ok {
//			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
//				return nil, TokenMalformed
//			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
//				// Token is expired
//				return nil, TokenExpired
//			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
//				return nil, TokenNotValidYet
//			} else {
//				return nil, TokenInvalid
//			}
//		}
//	}
//	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
//		return claims, nil
//	}
//	return nil, TokenInvalid
//}
//func (j *JWT) RefreshToken(tokenString string) (string, error) {
//	jwt.TimeFunc = func() time.Time {
//		return time.Unix(0, 0)
//	}
//	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return j.SigningKey, nil
//	})
//	if err != nil {
//		return "", err
//	}
//	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
//		jwt.TimeFunc = time.Now
//		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
//		return j.CreateToken(*claims)
//	}
//	return "", TokenInvalid
//}
//
//// create
//func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
//
//	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	//return token.SignedString(j.SigningKey)
//
//	/**
//
//	Audience  string `json:"aud,omitempty"`
//	ExpiresAt int64  `json:"exp,omitempty"`
//	Id        string `json:"jti,omitempty"`
//	IssuedAt  int64  `json:"iat,omitempty"`
//	Issuer    string `json:"iss,omitempty"`
//	NotBefore int64  `json:"nbf,omitempty"`
//	Subject   string `json:"sub,omitempty"`
//	 */
//
//	expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
//	claims = jwt.StandardClaims{
//		Audience:  "Username",        // 受众
//		ExpiresAt: expiresTime,       // 失效时间
//		Id:        "123123",          // 编号
//		IssuedAt:  time.Now().Unix(), // 签发时间
//		Issuer:    "gin hello",       // 签发人
//		NotBefore: time.Now().Unix(), // 生效时间
//		Subject:   "login",           // 主题
//	}
//
//	var jwtSecret = []byte(config.Secret)
//	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err := tokenClaims.SignedString(jwtSecret);
//	if  err != nil {
//		// 生成token失败
//	}
//	fmt.Print(token)
//	return "", nil
//}
//
