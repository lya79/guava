package chatroom

import (
	"module/database/internal/common/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

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
	{ // 解析傳遞的參數
		err := c.ShouldBindJSON(&input)
		if err != nil {
			Send(c, "000100010002")
			return
		}
	}

	{ // 檢查缺少傳入的參數
		missingParams := []string{}
		if strings.TrimSpace(input.Username) == "" {
			missingParams = append(missingParams, "username")
		}
		if strings.TrimSpace(input.Password) == "" {
			missingParams = append(missingParams, "password")
		}
		if strings.TrimSpace(input.Alias) == "" {
			missingParams = append(missingParams, "alias")
		}
		if len(missingParams) > 0 {
			Send(c, "000100010004")
			return
		}
	}

	{ // 檢查參數是否符合規則
		if auth.IsVaildUsernameFormat(input.Username) {
			Send(c, "000100010005")
			return
		}
		if auth.IsVaildPasswordFormat(input.Password) {
			Send(c, "000100010006")
			return
		}
		if auth.IsVaildAliasFormat(input.Alias) {
			Send(c, "000100010007")
			return
		}
		if auth.IsVaildUserRole(input.UserRole) {
			Send(c, "000100010008")
			return
		}
	}

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
