package main

import (
	"log"
	"module/database/internal"
	"module/database/internal/common/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	log.SetPrefix("[" + config.GetAppSite() + "]")
	// log.Printf("GOMAXPROCS: %v", runtime.GOMAXPROCS(runtime.NumCPU())) // TODO 待確認是否需要多核心
}

func main() {
	engine := internal.NewEngine()
	internal.SetRouter(engine)
	internal.Listener(engine)
}
