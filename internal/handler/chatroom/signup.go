package chatroom

import (
	"log"
	"github.com/lya79/guava/internal/common/auth"
	"github.com/lya79/guava/internal/common/db"
	"github.com/lya79/guava/internal/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

// SignupInput 註冊會員、管理員
type SignupInput struct {
	// 名稱
	Username string `json:"username" example:"yuan"`
	// 密碼
	Password string `json:"password" example:"qwe123"`
	// 暱稱
	Alias string `json:"alias" example:"育安"`
	// 階層
	UserRole int `json:"user_role" example:"2"`
}

// Signup 註冊帳號
func Signup(c *gin.Context) {
	var input SignupInput
	{ // 解析傳遞的參數
		err := c.ShouldBindJSON(&input)
		if err != nil {
			Send(c, "000100010002", nil)
			return
		}
	}

	bussiness := func(input SignupInput) (string, error) {
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
				return "000100010004", nil
			}
		}

		{ // 檢查參數是否符合規則
			if auth.IsVaildUsernameFormat(input.Username) {
				return "000100010005", nil
			}
			if auth.IsVaildPasswordFormat(input.Password) {
				return "000100010006", nil
			}
			if auth.IsVaildAliasFormat(input.Alias) {
				return "000100010007", nil
			}
			if auth.IsVaildUserRole(input.UserRole) {
				return "000100010008", nil
			}
		}

		var permisson string
		{ // 初始化新帳號的權限
			m := auth.GetDefalutPermisson(input.UserRole)
			str, err := auth.GetPermissonMapToJSON(m)
			if err != nil {
				return "000100010011", err
			}
			permisson = str
		}

		var encryptionPwd string
		{ // 密碼加密
			pwd, err := auth.Encryption(input.Password)
			if err != nil {
				return "000100010012", err
			}
			encryptionPwd = pwd
		}

		db, err := db.GetConnection()
		if err != nil {
			return "000100010014", err
		}

		err = repository.CreateAccount(
			db,
			input.UserRole,
			input.Username,
			input.Alias,
			encryptionPwd,
			permisson,
		)
		if err != nil {
			return "000100020003", err
		}

		return "", nil
	}

	{ // 註冊新帳號
		respCode, err := bussiness(input)
		if err != nil {
			log.Println("註冊新帳號失敗, err:", err)
		}
		Send(c, respCode, nil)
		return
	}
}
