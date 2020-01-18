package i18n

import (
	"log"
	"github.com/lya79/guava/internal/common/logger"
)

const (
	langEn = "en"
	langTw = "tw"
	langCn = "cn"
)

// GetErrorMsg 取得錯誤代碼對應的多語系訊息, lang參數傳入[en|tw|cn]
func GetErrorMsg(lang, code string) string {
	var msg string

	val, ok := errorCodeMap[code]
	if !ok {
		msg = errorCodeMap["000100010001"].En
		log.Println(logger.WARN, msg+"(code:"+code+")")
	} else {
		switch lang {
		case langEn:
			msg = val.En
		case langTw:
			msg = val.Tw
		case langCn:
			msg = val.Cn
		default:
			msg = val.En
		}
	}

	msg += "(code:" + code + ")"

	return msg
}
