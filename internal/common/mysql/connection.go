package main

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

/*
特點:
  1.此包要容易測試.
  2.要可以多次重啟連線池.
  3.關閉 db連線池之前, 要確保都已經完成所有的連線.
  4.關閉 db連線池的過程中或已經關閉連線池的情況下, 要限制不可以有新的連線.
*/

// Setting DB連線初始化時用的設定
type Setting struct {
	host, port, database, username, password string
	logMode                                  bool
	maxIdleConns, maxOpenConns               int
	connMaxLifetime                          time.Duration
}

var (
	dbMutex *sync.Cond //sync.RWMutex
	db      *gorm.DB   // db連線

	count int // 目前 db使用中的連線數
)

func init() {
	dbMutex = sync.NewCond(new(sync.Mutex))
}

// GetConnection 取得 db連線
func GetConnection() (*gorm.DB, error) {
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()
	defer dbMutex.Broadcast()

	if db == nil {
		return nil, errors.New("connection nil")
	}

	count++

	return db.New(), nil
}

// PutConnection 歸還 db連線
func PutConnection() error {
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()
	defer dbMutex.Broadcast()

	if db == nil {
		return errors.New("connection nil")
	}

	if count <= 0 {
		return nil
	}

	count--

	return nil
}

// Close 關閉 db連線
func Close() error { // 服務器關閉時需要呼叫
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()

	if db == nil {
		return nil
	}

	{ // 用於確保使用中的連線都能正常使用完畢
		for {
			if count > 0 {
				dbMutex.Wait()
				continue
			}
			break
		}
	}

	return db.Close()
}

// Open 開啟 db連線
func Open(local bool, setting Setting, models ...interface{}) error {
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()

	if db != nil {
		return nil
	}

	var err error
	db, err = gorm.Open(
		"mysql",
		getConnectName(
			setting.host,     // "mysql",
			setting.port,     // "3306",
			setting.database, // "PEPPER",
			setting.username, // "root",
			setting.password, // "qwe123",
		),
	)
	if err != nil {
		return err
	}

	// local := config.IsLocalByProjectEnv()

	db.LogMode(setting.logMode)

	db.DB().SetMaxIdleConns(setting.maxIdleConns)
	db.DB().SetMaxOpenConns(setting.maxOpenConns)
	db.DB().SetConnMaxLifetime(setting.connMaxLifetime) // (time.Minute * 5)

	initTable := func(db *gorm.DB, local bool, model interface{}) error {
		if local {
			return db.AutoMigrate(model).Error
		}

		if exist := db.HasTable(model); !exist {
			return errors.New("account資料表不存在")
		}

		return nil
	}

	// var models []interface{}
	// models = append(models, &Account{})

	for i := range models {
		if err := initTable(db, local, models[i]); err != nil {
			log.Fatalf("初始化 table錯誤： %v", err)
		}
	}

	return nil
}

func getConnectName(host, port, database, username, password string) string {
	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Asia%2FTaipei"
}
