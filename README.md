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

下載套件
```shell
go mod tidy  
```

```shell
go mod vendor  
```

## 服務執行
```shell
docker-compose -f ./deployment/local/docker-compose.yml up -d
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