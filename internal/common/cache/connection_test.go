package cache

import (
	"math/rand"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func Test_GetConnection(t *testing.T) {
	for i := 0; i < 3; i++ {
		setting := Setting{
			host:        "localhost",
			port:        "6379",
			password:    "",
			db:          0,
			poolSize:    10,
			idleTimeout: time.Duration(time.Minute * 5),
		}

		{ // 開啟連線
			err := Open(setting)
			if err != nil {
				t.Error("Open:", err)
				return
			}
		}

		sum := 10
		for i := 0; i < sum; i++ { // 測試取出和歸還連線
			var db *redis.Client
			{ // connect
				var err error
				db, err = GetConnection() // 取出連線
				if err != nil {
					t.Error("GetConnection:", err)
					return
				}
			}
			go func(db *redis.Client) {
				defer PutConnection() // 歸還連線

				{ // sleep
					rand.Seed(time.Now().Unix())
					time.Sleep(time.Duration(rand.Intn(100)+1) * time.Microsecond)
				}

				{ // ping
					pong, err := db.Ping().Result()
					if err != nil {
						t.Error("Ping:", err)
						return
					}
					if pong != "PONG" {
						t.Error("pong != PONG, pong:", pong)
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
