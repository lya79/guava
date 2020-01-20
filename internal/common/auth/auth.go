package auth

import (
	"golang.org/x/crypto/scrypt"
)

const (
	userRoleOfAdmin = iota + 1 // 管理員
	userRoleOfUser             // 會員
	userRoleOfGuest            // 還未登入的訪客
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
