package chatroom

import (
	"module/database/internal/common/header"
	"module/database/internal/common/i18n"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response API回傳內容 // TODO 改放在通用程式碼目錄
type Response struct {
	Code    string
	Message string
	Result  interface{}
}

// Signup 註冊帳號
func Signup(c *gin.Context) {
	type Input struct {
		// 名稱
		Username string `json:"username" example:"yuan"`
		// 密碼
		Password string `json:"password" example:"qwe123"`
		// 暱稱
		Alias string `json:"alias" example:"育安"`
		// 階層
		UserRole int64 `json:"user_role" example:"2"`
	}

	var input Input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		Send(c, "000100010002")
		return
	}

	// TODO 檢查傳入的參數

	// TODO 執行註冊帳號
	/*
		如果要註冊管理者帳號, 就需要先驗證操作者是否管理者
		if Input.UserRole = 1 {
			operator := GetOperator(c) // 取得 登入者資料從 session
			if !operator.IsAdmin(){
				Send(c, lang, "000100010003")
				return
			}
			...
		}

		如果註冊會員帳號, 不需要驗證操作者是否管理者
		if Input.UserRole = 2 {
			...
		}
	*/

	Send(c, "000100020001")
}

// Send 回傳 API錯誤訊息
func Send(c *gin.Context, code string) {
	lang := header.GetLangHeader(c)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: i18n.GetErrorMsg(lang, code),
	})
}
