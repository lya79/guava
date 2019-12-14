package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"module/database/internal"
	"module/database/internal/common/config"
	"module/database/internal/common/logger"
)

func init() {
	logger.SetPrefix("[" + config.GetAppSite() + "]")
	if config.GetAppDebug() == config.DEBUG {
		logger.EnableDebug()
	}

	// log.Printf("GOMAXPROCS: %v", runtime.GOMAXPROCS(runtime.NumCPU())) // TODO 待確認是否需要多核心
}

func main() {
	engine := internal.NewEngine()
	internal.SetRouter(engine)
	internal.Listener(engine)
}
