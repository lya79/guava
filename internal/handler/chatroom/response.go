package chatroom

import (
	"module/database/internal/common/header"
	"module/database/internal/common/i18n"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response API回傳內容
type Response struct {
	Code    string
	Message string
	Result  interface{}
}

// Send 回傳 API錯誤訊息
func Send(c *gin.Context, code string) {
	lang := header.GetLangHeader(c)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: i18n.GetErrorMsg(lang, code),
	})
}
