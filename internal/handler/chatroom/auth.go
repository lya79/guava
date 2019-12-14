package chatroom

import (
	"log"
	"module/database/internal/common/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Echo 回傳時間
func Echo(c *gin.Context) {
	// c.JSON(http.StatusOK, Response{
	// 	// Code:    code,
	// 	// Message: i18n.GetErrorMsg(lang, code),
	// 	Result: time.Now(),
	// })

	db, err := gorm.Open(
		"mysql",
		getConnectName(
			"mysql",
			"3306",
			"PEPPER",
			"root",
			"qwe123",
		),
	)
	defer db.Close()
	if err != nil {
		log.Println("gorm.Open: err:", err)
		c.JSON(http.StatusOK, Response{
			// Code:    code,
			// Message: i18n.GetErrorMsg(lang, code),
			Result: "gorm.Open: err:" + err.Error(),
		})
		return
	}
	log.Println("gorm.Open: ok:")
	c.JSON(http.StatusOK, Response{
		// Code:    code,
		// Message: i18n.GetErrorMsg(lang, code),
		Result: "gorm.Open: ok",
	})
}

func getConnectName(host, port, database, username, password string) string {
	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Asia%2FTaipei"
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

	{ // 註冊新帳號
		if auth.IsAdminUserRole(input.UserRole) { // 只有登入中的管理者才可以註冊新的管理者帳號
			//  operator := GetOperator(c) // 取得 登入者資料從 session
			// 	if !operator.IsAdmin(){
			// 		Send(c, lang, "000100010003")
			// 		return
			// 	}
			// 	...
		} else if auth.IsUserUserRole(input.UserRole) {
		}

	}

	Send(c, "000100020001")
}
