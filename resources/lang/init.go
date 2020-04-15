package lang

import (
	"blog-go-api/app/config"
	cnLang "blog-go-api/resources/lang/cn"
	enLang "blog-go-api/resources/lang/en"
)

var i18n string

/**
 获取语言包下面的内容
 */
func GetLang(text string) (langText string) {
	i18n := config.AppLang
	switch i18n {
		case "cn":
			Lang := cnLang.Lang
			langText := Lang[text]
			if langText == nil {
				return text
			}
			langTextTer := langText.(string)
			return langTextTer
			break
		case "en":
			Lang := enLang.Lang
			langText := Lang[text]
			if langText == nil {
				return text
			}
			langTextTer := langText.(string)
			return langTextTer
			break
	}
	return text
}