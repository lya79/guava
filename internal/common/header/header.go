package header

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	headerLang     = "lang"
	headerTimezone = "timezone"

	langEn = "en"
	langTw = "tw"
	langCn = "cn"
)

// InitLangHeader 初始化 lang header
func InitLangHeader(c *gin.Context) {
	lang := c.GetHeader(headerLang)
	if lang == langEn || lang == langTw || lang == langCn {
		return
	}
	c.Set(headerLang, langEn)
}

// GetLangHeader 取得多語系 header
func GetLangHeader(c *gin.Context) string {
	lang := c.GetHeader(headerLang)
	if lang == langEn || lang == langTw || lang == langCn {
		return lang
	}
	return langEn
}

// InitTimeZoneHeader 初始化 timezone header
func InitTimeZoneHeader(c *gin.Context) {
	tz := c.GetHeader(headerTimezone)
	_, err := strconv.ParseInt(tz, 10, 64)
	if err == nil {
		return
	}
	c.Set(headerTimezone, 0)
}

// GetTimeZoneHeader 取得時區 header
func GetTimeZoneHeader(c *gin.Context) int64 {
	tz := c.GetHeader(headerTimezone)
	val, err := strconv.ParseInt(tz, 10, 64)
	if err != nil {
		return 0
	}
	return val
}
