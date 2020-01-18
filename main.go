package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lya79/guava/internal"
	"github.com/lya79/guava/internal/common/config"
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
