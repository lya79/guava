package chatroom

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lya79/guava/internal/common/auth"
)

type loginInput struct {
	// 名稱
	Username string `json:"username" example:"yuan"`
	// 密碼
	Password string `json:"password" example:"qwe123"`
	// 階層
	UserRole int `json:"user_role" example:"2"`
}

// Login 登入聊天室
func Login(c *gin.Context) {
	var input loginInput
	{ // 解析傳遞的參數
		err := c.ShouldBindJSON(input)
		if err != nil {
			Send(c, "000100010002", nil)
			return
		}
	}

	bussiness := func(input loginInput) (string, error) {
		{ // 檢查缺少傳入的參數
			missingParams := []string{}
			if strings.TrimSpace(input.Username) == "" {
				missingParams = append(missingParams, "username")
			}
			if strings.TrimSpace(input.Password) == "" {
				missingParams = append(missingParams, "password")
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
			if auth.IsVaildUserRole(input.UserRole) {
				return "000100010008", nil
			}
		}

		{ // 將加密後的密碼和 db內的密碼比對是否相同
			var encryptionPwd string
			{ // 密碼加密
				pwd, err := auth.Encryption(input.Password)
				if err != nil {
					return "000100010012", err
				}
				encryptionPwd = pwd
			}

			var dbPwd string
			{ // 取得 db內的密碼 // TODO

			}

			if dbPwd != encryptionPwd { // 比對密碼
				return "000100020004", nil
			}
		}

		{ // 產生新 session // TODO

		}

		{ // 將新 session寫入 cookie // TODO

		}

		{ // 刪除該用戶舊的 session // TODO

		}

		{ // 將新 session寫入 redis // TODO

		}
		return "", nil
	}

	{ // 登入
		respCode, err := bussiness(input)
		if err != nil {
			log.Println("註冊新帳號失敗, err:", err)
		}
		Send(c, respCode, nil)
		return
	}
}
