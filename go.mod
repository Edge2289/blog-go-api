module blog-go-api

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.54.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.12
	github.com/qiniu/api.v7 v7.2.5+incompatible // indirect
	github.com/qiniu/api.v7/v7 v7.4.1 // indirect
	github.com/sirupsen/logrus v1.2.0
	github.com/skip2/go-qrcode v0.0.0-20191027152451-9434209cb086
	github.com/smartystreets/goconvey v1.6.4 // indirect
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85
	golang.org/x/net => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a
	proxy.golang.org/github.com/jinzhu/gorm v1.9.11 => github.com/jinzhu/gorm v1.9.11-20180820150726-614d502a4dac
)
