# 服務安裝與執行

## 安裝

Golang 1.11+

clone專案
```shell
git clone https://github.com/lya79/guava.git
```

啟用 go module
```shell
export GO111MODULE=on
```

下載第三方套件
```shell
go mod tidy  
```

## 服務執行
```shell
PROJECT_ENV=local PROJECT_SITE=chatroom PROJECT_DEBUG=on go run .
```
環境變數 | 值 | 用途
--- | --- | ---
PROJECT_ENV | local | 服務執行的環境
PROJECT_SITE | chatroom | 要執行的服務
PROJECT_DEBUG | on | 是否開啟 `debug模式`, 如果沒帶 `on`參數就代表關閉 `debug模式`. `debug模式`開啟時才會輸出 debug訊息.

# 如何維護程式

## 設定檔使用
設定檔使用 `.toml`格式. 設定檔放置於 `configs`目錄底下. 例如本機端測試環境則會使用 `configs/local/default.toml`.

例如取得服務使用的 Port
```golang
import "module/database/internal/common/config"

port := config.GetConfig().Servers.Port
```

## 環參數使用
環境參數例如 `PROJECT_ENV=local`, 也是藉由 `config`取得參數值.

例如取得執行環境(`PROJECT_ENV`)的值.
```golang
import "module/database/internal/common/config"

env := config.GetAppEnv() // 取得 PROJECT_ENV環境參數的值
```

## log使用
```golang
import "module/database/internal/common/logger"

log.Println("hello") // info level

logger.Println( // debug level
	logger.DEBUG,
	"hello",
	func(logger *log.Logger, str string) {
		logger.Println(str)
	},
)

logger.Println( // warn level
	logger.WARN,
	"hello",
	func(logger *log.Logger, str string) {
		logger.Println(str)
	},
)

logger.Println( // fatal level
	logger.FATAL,
	"hello",
	func(logger *log.Logger, str string) {
		logger.Println(str)
	},
)
```