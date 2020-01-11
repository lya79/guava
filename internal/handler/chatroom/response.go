package chatroom

import (
	"module/database/internal/common/header"
	"module/database/internal/common/i18n"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response API回傳內容
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// Send 回傳 API錯誤訊息
func Send(c *gin.Context, code string, result interface{}) {
	resp := Response{}

	if code != "" {
		resp.Code = code
		resp.Message = i18n.GetErrorMsg(header.GetLangHeader(c), code)
	}

	if result != nil {
		resp.Result = result
	}

	c.JSON(http.StatusOK, resp)
}
