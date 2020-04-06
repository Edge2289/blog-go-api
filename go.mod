module blog-go-api

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.54.0 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/sirupsen/logrus v1.2.0
	github.com/skip2/go-qrcode v0.0.0-20191027152451-9434209cb086
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
)

replace proxy.golang.org/github.com/jinzhu/gorm v1.9.11 => github.com/jinzhu/gorm v1.9.11-20180820150726-614d502a4dac
