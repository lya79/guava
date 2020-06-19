package auth

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/scrypt"
)

const (
	userRoleOfAdmin = iota + 1 // 管理員
	userRoleOfUser             // 會員
	userRoleOfGuest            // 還未登入的訪客
)

const (
	sessionName = "session" // cookie內的 session的 tag名稱
)

const (
	salt = `bV82w-SfW\d"b]E;U>$q**;[43T#DGt!` // 加密鹽, 密碼加密用途
)

// IsVaildUserRole 檢查是否為有效的 user role
func IsVaildUserRole(userRole int) bool {
	if userRole >= userRoleOfAdmin && userRole <= userRoleOfGuest {
		return true
	}
	return false
}

// IsAdminUserRole 檢查是否為管理員
func IsAdminUserRole(userRole int) bool {
	return userRole == userRoleOfAdmin
}

// IsMemberUserRole 檢查是否為會員
func IsMemberUserRole(userRole int) bool {
	return userRole == userRoleOfUser
}

// IsGuestUserRole 檢查是否為會員
func IsGuestUserRole(userRole int) bool {
	return userRole == userRoleOfGuest
}

// IsVaildUsernameFormat 檢查 Username規則是否正確  // TODO 實作 Username規則
func IsVaildUsernameFormat(str string) bool {
	return true
}

// IsVaildPasswordFormat 檢查 Password規則是否正確  // TODO 實作 Password規則
func IsVaildPasswordFormat(str string) bool {
	return true
}

// IsVaildAliasFormat 檢查 Alias規則是否正確  // TODO 實作 Alias規則
func IsVaildAliasFormat(str string) bool {
	return true
}

// Encryption 密碼加密
func Encryption(pwd string) (string, error) {
	hash, err := scrypt.Key([]byte(pwd), []byte(salt), 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// GenerateRangeNum 隨機產生大於等於 min, 且小於等於 max的數字
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// GenerateRandomSessionID 隨機產生指定長度的 session
func GenerateRandomSessionID() (string, error) {
	len := GenerateRangeNum(25, 32)
	session := make([]byte, len)
	if _, err := rand.Read(session); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(session), nil
}

// GetSession 取得 cookie內的 session
func GetSession(c *gin.Context) (string, error) {
	return c.Cookie(sessionName)
}

// SetSession 設定 cookie內的 session
func SetSession(c *gin.Context, session string, maxAge int) error {
	referer, err := url.Parse(c.Request.Referer()) // TODO 確認為什麼 domin要這樣寫
	if err != nil {
		return err
	}
	domain := referer.Host
	c.SetCookie(sessionName, session, maxAge, "/", domain, false, true)
	return nil
}

// DelSession 刪除 cookie內的 session
func DelSession(c *gin.Context) {
	SetSession(c, "", -1)
}

// getUserRedisKey 取使用者的redis中的key值
func getUserRedisKey(id, role int64) string {
	return fmt.Sprintf("Pepper:Auth:User:Role_%d:ID_%d", role, id)
}

// getSessionRedisKey 取session的redis中的key值
func getSessionRedisKey(session string) string {
	return fmt.Sprintf("Pepper:Auth:Session:%s", session)
}
