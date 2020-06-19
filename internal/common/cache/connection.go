package cache

import (
	"errors"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// Setting redis連線初始化時用的設定
type Setting struct {
	host, port, password string
	db                   int
	poolSize             int
	idleTimeout          time.Duration
}

var (
	dbMutex *sync.Cond
	conn    *redis.Client
	count   int  // 目前 db使用中的連線數
	running bool // 目前是否運行中
)

func init() {
	dbMutex = sync.NewCond(&sync.Mutex{})
}

// GetConnection 取得 redis連線
func GetConnection() (*redis.Client, error) {
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()
	defer dbMutex.Broadcast()

	if conn == nil || !running {
		return nil, errors.New("connection nil")
	}

	count++

	return conn, nil
}

// PutConnection 歸還 db連線
func PutConnection() error {
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()
	defer dbMutex.Broadcast()

	if conn == nil {
		return errors.New("connection nil")
	}

	if count <= 0 {
		return nil
	}
	count--

	return nil
}

// Close 關閉 redis連線
func Close() error { // TODO 服務器關閉時需要呼叫
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()

	if conn == nil {
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
	}

	err := conn.Close()

	conn = nil

	return err
}

// Open 開啟 redis連線
func Open(setting Setting) error {
	dbMutex.L.Lock()
	defer dbMutex.L.Unlock()

	if conn != nil {
		return nil
	}

	running = true

	addr := setting.host + ":" + setting.port

	conn = redis.NewClient(&redis.Options{ // TODO 初始化寫入到設定
		Addr:        addr,             //conf.Host + conf.Port,
		Password:    setting.password, //conf.Password,
		DB:          setting.db,       //conf.DB, // use default DB
		PoolSize:    setting.poolSize, // conf.MaxConn,
		IdleTimeout: setting.idleTimeout,
	})

	{ // ping
		_, err := conn.Ping().Result()
		if err != nil {
			return err
		}
	}

	return nil
}
