package i18n

import (
	"log"
	"module/database/internal/common/logger"
)

// GetErrorMsg 取得錯誤代碼對應的多語系訊息, lang參數傳入[en|tw|cn]
func GetErrorMsg(lang, code string) string {
	var msg string

	val, ok := errorCodeMap[code]
	if !ok {
		msg = errorCodeMap["000100010001"].En
		logger.Println(
			logger.WARN,
			msg+"(code:"+code+")",
			func(logger *log.Logger, str string) {
				logger.Println(str)
			},
		)
	} else {
		switch lang {
		case "en":
			msg = val.En
		case "tw":
			msg = val.Tw
		case "cn":
			msg = val.Cn
		default:
			msg = val.En
		}
	}

	msg += "(code:" + code + ")"

	return msg
}
