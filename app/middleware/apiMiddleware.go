package middleware
//
//import (
//	jwtAuth "blog-go-api/app/middleware/jwt"
//	"github.com/dgrijalva/jwt-go"
//	"time"
//)
//
//func AuthInit() (*jwtAuth.GinJWTMiddleware, error) {
//	// the jwt middleware
//	return jwtAuth.New(&jwtAuth.GinJWTMiddleware{
//		Realm:           "test zone",
//		Key:             []byte("secret key"),
//		Timeout:         time.Hour,
//		MaxRefresh:      time.Hour,
//		IdentityKey:     "config2.ApplicationConfig.JwtSecret",
//		PayloadFunc:     nil,//PayloadFunc,
//		IdentityHandler: nil,//"handler.IdentityHandler",
//		Authenticator:   nil,//"handler.Authenticator",
//		Authorizator:    nil,//"handler.Authorizator",
//		Unauthorized:    nil,//"handler.Unauthorized",
//		// TokenLookup is a string in the form of "<source>:<name>" that is used
//		// to extract token from the request.
//		// Optional. Default value "header:Authorization".
//		// Possible values:
//		// - "header:<name>"
//		// - "query:<name>"
//		// - "cookie:<name>"
//		// - "param:<name>"
//		TokenLookup: "header: Authorization, query: token, cookie: jwt",
//		// TokenLookup: "query:token",
//		// TokenLookup: "cookie:token",
//
//		// TokenHeadName is a string in the header. Default value is "Bearer"
//		TokenHeadName: "Bearer",
//
//		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
//		TimeFunc: time.Now,
//	}), nil
//
//}
//
//func PayloadFuncc(data interface{}) jwt.MapClaims {
//	//if v, ok := data.(map[string]interface{}); ok {
//	//	//u, _ := v["user"].(SysUser)
//	//	//r, _ := v["role"].(SysRole)
//	//	return jwt.MapClaims{
//	//		//jwt.IdentityKey:  u.Id,
//	//		//jwt.RoleIdKey:    r.Id,
//	//		//jwt.RoleKey:      r.RoleKey,
//	//		//jwt.NiceKey:      u.Username,
//	//		//jwt.DataScopeKey: r.DataScope,
//	//		//jwt.RoleNameKey:  r.Name,
//	//	}
//	//}
//	return jwt.MapClaims{}
//}
