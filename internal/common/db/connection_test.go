package db

import (
	"math/rand"
	"module/database/internal/repository"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Test_GetConnection(t *testing.T) {
	for i := 0; i < 3; i++ {
		setting := Setting{
			host:            "mysql",
			port:            "3306",
			database:        "chatroom",
			username:        "root",
			password:        "qwe123",
			logMode:         true,
			maxIdleConns:    10,
			maxOpenConns:    10,
			connMaxLifetime: time.Duration(time.Minute * 5),
		}

		{ // 開啟連線
			err := Open(setting)
			if err != nil {
				t.Error("Open:", err)
				return
			}
		}

		{ // 初始化資料表
			err := InitTable(true, &repository.Account{})
			if err != nil {
				t.Error("InitTable:", err)
				return
			}
		}

		sum := 10
		for i := 0; i < sum; i++ { // 測試取出和歸還連線
			var db *gorm.DB
			{ // connect
				var err error
				db, err = GetConnection() // 取出連線
				if err != nil {
					t.Error("GetConnection:", err)
					return
				}
			}
			go func(db *gorm.DB) {
				defer PutConnection() // 歸還連線

				{ // sleep
					rand.Seed(time.Now().Unix())
					time.Sleep(time.Duration(rand.Intn(100)+1) * time.Microsecond)
				}
				{ // ping
					if err := db.DB().Ping(); err != nil {
						t.Error("Ping:", err)
						return
					}
				}

			}(db)
		}

		{
			err := Close() // 結束連線
			if err != nil {
				t.Error("Close:", err)
				return
			}
		}
	}
}
