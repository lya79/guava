package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
)

// TODO 定義資料表
/*
1.帳戶資料表(描述管理者、會員帳號密碼和權限之類的狀態)
2.行為紀錄資料表(描述每一個帳戶執行任何API的紀錄, 包含紀錄傳遞的參數還有回傳的參數)
*/

const account = "account" // 資料表名稱

// Account 帳戶資料表
type Account struct {
	gorm.Model
	UserRole  int    `gorm:"column:user_role;type:int(2);NOT NULL;DEFAULT:1"` // 權限階層, 1:admin, 2:member
	Permisson string `gorm:"coulumn:permission;varchar(255);NOT NULL"`        // api權限, 用json存標籤和值, 先做出可以踢人的 api就好
	Username  string `gorm:"column:username;type:varchar(50);NOT NULL;"`      // 帳號
	Password  string `gorm:"column:password;type:varchar(255);NOT NULL;"`     // 密碼
	Alias     string `gorm:"column:alias;type:varchar(50);NOT NULL;"`         // 暱稱
	Enable    int    `gorm:"column:enable;type:int(2);NOT NULL;DEFAULT:0"`    // 狀態, 1:啟用, 2:停用
}

// CreateAccount 建立新的帳戶
func CreateAccount(
	db *gorm.DB,
	userRole int,
	username string,
	alias string,
	encryptionPwd string,
	permisson string,
) error {
	record := Account{
		UserRole:  userRole,
		Username:  username,
		Alias:     alias,
		Enable:    1,
		Permisson: permisson,
		Password:  encryptionPwd,
	}

	if err := db.Table(account).Create(&record).Error; err != nil {
		return errors.New("帳戶建立失敗, err:" + err.Error())
	}

	return nil
}

// CheckAccountPwd 查詢指定帳戶是否存在
func CheckAccountPwd(
	db *gorm.DB,
	userRole int,
	username string,
	encryptionPwd string,
) (bool, error) {
	var accountArr []Account
	if err := db.Table(account).
		Where("`user_role` = ?", userRole).
		Where("`username` = ?", username).
		Where("`password` = ?", encryptionPwd).
		Find(&accountArr).Error; err != nil {
		return false, errors.New("帳戶查詢失敗, err:" + err.Error())
	}
	if len(accountArr) != 1 {
		return false, errors.New("帳戶查詢失敗, err: len(accountArr)!=1")
	}
	return true, nil
}
