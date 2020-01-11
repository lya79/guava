package auth

import "encoding/json"

const (
	createAdmin  = iota + 1 // 建立管理員帳號
	createMember            // 建立會員帳號
	notify                  // 發出系統通知訊息
	talk                    // 發言
)

type api struct {
	Method   string
	Path     string
	userRole []int // 設定允許使用此 api的 userRole
}

var permission = map[int]api{
	createAdmin:  {"POST", "/api/accunt/create/admin", []int{userRoleOfAdmin}},
	createMember: {"POST", "/api/accunt/create/member", []int{userRoleOfAdmin, userRoleOfGuest}},
	notify:       {"POST", "/api/notify", []int{userRoleOfAdmin}},
	talk:         {"POST", "/api/talk", []int{userRoleOfAdmin, userRoleOfUser}},
}

/*
在中間件使用,
為了確認用戶有沒有資格使用指定的 api
*/
func CheckAPIByPermission(userRole int) bool { // TODO
	/*
		1.	從 cookie取出 sesseion取出對應的
		2.	用 session向 redis查找出對應的用戶資料(username、user_id、權限)
		2.	用戶資料中取出用戶權限
		4.  用戶權限 和 permission比對, 假如比對成功, 而且用戶權限是開啟, 就代表可使用, 否則都不可使用
	*/
	return false
}

func GetPermissonMapToJSON(m map[int]bool) (string, error) { // TODO
	str, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func GetPermissonJSONToMap(str string) (map[int]bool, error) { // TODO
	m := make(map[int]bool)
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		return nil, err
	}
	return m, err
}

func GetDefalutPermisson(userRole int) map[int]bool { // TODO
	if IsAdminUserRole(userRole) {
		return getAdminAllPermission()
	} else if IsMemberUserRole(userRole) {
		return getMemberAllPermission()
	}
	return getGuestAllPermission()
}

func getAdminAllPermission() map[int]bool {
	m := make(map[int]bool)
	for i := createAdmin; i < (talk + 1); i++ {
		m[i] = true
	}
	return m
}

func getMemberAllPermission() map[int]bool {
	m := make(map[int]bool)
	for i := createAdmin; i < (talk + 1); i++ {
		m[i] = false
	}
	m[talk] = true
	return m
}

func getGuestAllPermission() map[int]bool {
	m := make(map[int]bool)
	for i := createAdmin; i < (talk + 1); i++ {
		m[i] = false
	}
	m[createMember] = true
	return m
}
