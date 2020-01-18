package internal

import (
	"context"
	"log"
	"module/database/internal/common/config"
	"module/database/internal/common/logger"
	"module/database/internal/router/chatroom"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// NewEngine 產生服務 Engine
func NewEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	mws := []gin.HandlerFunc{gin.Recovery()}
	if config.IsLocalEnv() {
		mws = append(mws, gin.Logger())
	}
	engine.Use(mws...)
	// mws = append(mws,
	// 	Graceful(appConf.Site), // TODO 待確認 Graceful用途
	// 	OperationRecord,
	// )
	return engine
}

// Listener 啟動服務
func Listener(engine *gin.Engine) {
	addr := config.GetConfig().Chatroom.Host
	addr += ":" + strconv.FormatInt(config.GetConfig().Chatroom.Port, 10)

	// 建立 Server
	srv := &http.Server{ // TODO 待確認 http.Server設定
		Addr:    addr,
		Handler: engine,
		// ReadTimeout:  30 * time.Second,
		// WriteTimeout: 30 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		err := srv.ListenAndServe()
		msg := "服務被關閉"
		if err != nil {
			msg += " error:" + err.Error()
		}
		log.Println(logger.WARN, msg)
	}()

	// graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	<-sigs

	// stop gin engine
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	msg := "服務強制關閉"
	if err != nil {
		msg += " error:" + err.Error()
	}
	log.Println(logger.WARN, msg)
}

// SetRouter 設定 Router
func SetRouter(engine *gin.Engine) {
	name := config.GetAppSite()
	switch name {
	case "chatroom":
		chatroom.Provider(engine)
	default:
		log.Fatalf("無效的服務名稱 %v", name)
	}
	log.Printf("啟動 %v服務 %v:%v", name, config.GetConfig().Chatroom.Host, config.GetConfig().Chatroom.Port)
}
