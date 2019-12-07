package main

import (
	"log"
	"module/database/internal"
	"module/database/internal/common/config"
	"module/database/internal/common/logger"
	"runtime"
)

func init() {
	logger.SetPrefix("[" + config.GetAppSite() + "]")
	if config.GetAppDebug() == config.DEBUG {
		logger.EnableDebug()
	}

	log.Printf("GOMAXPROCS: %v", runtime.GOMAXPROCS(runtime.NumCPU()))
}

func main() {
	engine := internal.NewEngine()
	internal.SetRouter(engine)
	internal.Listener(engine)
}
