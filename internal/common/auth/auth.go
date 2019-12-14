package auth

const (
	userRoleOfAdmin = iota + 1 // 管理員
	userRoleOfUser             // 會員
)

// IsVaildUserRole 檢查是否為有效的 user role
func IsVaildUserRole(userRole int64) bool {
	if userRole < userRoleOfAdmin {
		return false
	} else if userRole > userRoleOfUser {
		return false
	}
	return true
}

// IsAdminUserRole 檢查是否為管理員
func IsAdminUserRole(userRole int64) bool {
	return userRole == userRoleOfAdmin
}

// IsUserUserRole 檢查是否為會員
func IsUserUserRole(userRole int64) bool {
	return userRole == userRoleOfUser
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
