package redis

import (
	"time"

	"github.com/go-redis/redis"
)

// 要有測試案例, 因此要寫出可測試的程式碼

var conn *redis.Client

func GetConnection() *redis.Client { // TODO 需要同步鎖, 為了能夠重啟, 用 once只能跑一次
	initConnection()
	return conn
}

func Close() error { // TODO 需要確認 redis關閉的寫法
	return conn.Close()
}

func initConnection() {
	conn = redis.NewClient(&redis.Options{ // TODO 初始化寫入到設定
		Addr:        "redis:6379", //conf.Host + conf.Port,
		Password:    "",           //conf.Password,
		DB:          0,            //conf.DB, // use default DB
		PoolSize:    10,           // conf.MaxConn,
		IdleTimeout: time.Minute * 5,
	})
}
