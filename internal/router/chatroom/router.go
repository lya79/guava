package chatroom

import (
	"module/database/internal/handler/chatroom"

	"github.com/gin-gonic/gin"
)

// Provider 設定 router
func Provider(engine *gin.Engine) { // TODO 實作 chatroom router
	setAPIRouter(engine)
}

// SetAPIRouter 設定 API與 handler
func setAPIRouter(engine *gin.Engine) {
	mws := []gin.HandlerFunc{ // TODO 待確認 api router需要哪些中間件
		// gzip.Gzip(gzip.BestCompression),
		// CheckClientIP(false),
		// apiMiddleware.CheckLanguage,
		// apiMiddleware.CheckSession,
		// CheckAPIPermission,
		// apiMiddleware.CheckSelfPermission,
	}

	api := engine.Group(
		"/api",
		mws...,
	)

	api.GET("/echo", chatroom.Echo)
}
