package mysql

//	特點:
//	*.可多次重啟連線
//  *.關閉 db連線池之前, 會確保都已經完成所有的連線.
//  *.關閉 db連線池的過程中, 或已經關閉連線池的情況下, 會限制不可以有新的連線.

import (
	"errors"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

// Setting DB連線初始化時用的設定
type Setting struct {
	host, port, database, username, password string
	logMode                                  bool
	maxIdleConns, maxOpenConns               int
	connMaxLifetime                          time.Duration
}

var (
	dbMutex *sync.Cond
	db      *gorm.DB // db連線
	count   int      // 目前 db使用中的連線數
	running bool     // 目前是否運行中
)

func init() {
	dbMutex = sync.NewCond(&sync.Mutex{})
}

// GetConnection 取得 db連線
func GetConnection() (*gorm.DB, error) {
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()
	defer dbMutex.Broadcast()

	if db == nil || !running {
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

	running = false

	{ // 用於確保使用中的連線都能正常使用完畢
		for {
			if count > 0 {
				dbMutex.Wait()
				continue
			}
			break
		}

		for {
			count := db.DB().Stats().InUse
			if count <= 0 {
				break
			}
			time.Sleep(time.Millisecond * 100)
		}
	}

	err := db.Close()

	db = nil

	return err
}

// Open 開啟 db連線
func Open(setting Setting) error {
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()

	if db != nil {
		return nil
	}

	running = true

	var err error
	db, err = gorm.Open(
		"mysql",
		getConnectName(
			setting.host,
			setting.port,
			setting.database,
			setting.username,
			setting.password,
		),
	)
	if err != nil {
		return err
	}

	db.LogMode(setting.logMode)

	db.DB().SetMaxIdleConns(setting.maxIdleConns)
	db.DB().SetMaxOpenConns(setting.maxOpenConns)
	db.DB().SetConnMaxLifetime(setting.connMaxLifetime)

	return db.DB().Ping()
}

func getConnectName(host, port, database, username, password string) string {
	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Asia%2FTaipei"
}

// InitTable 初始化資料表
func InitTable(autoMigrate bool, model interface{}) error {
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()

	if db == nil || !running {
		return errors.New("connection nil")
	}

	if autoMigrate {
		return db.AutoMigrate(model).Error
	}

	if exist := db.HasTable(model); !exist {
		return errors.New("account資料表不存在")
	}

	return nil
}
