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
	mws := []gin.HandlerFunc{ // TODO 待確認 api router需要哪些中間件, 部分 API並不需要檢查 seesion, 例如註冊新會員
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

	/*
		測試用途 API
	*/
	api.GET("/echo", chatroom.Echo)

	/*
		語系與時區 API
	*/
	// api.POST("/personalization/lang", chatroom.Signup)   // 語系
	// api.POST("/personalization/timezone", chatroom.Signup)   // 時區

	/*
		身份驗證 API
	*/
	api.POST("/auth/signup", chatroom.Signup) // 註冊
	// api.POST("/auth/update/pwd", chatroom.Echo) // 更新密碼
	// api.POST("/auth/login", chatroom.Echo)      // 登入
	// api.POST("/auth/logout", chatroom.Echo)     // 登出

	/*
		帳號管理 API
	*/
	// api.POST("/account/online-account-list", chatroom.Echo) // 在線用戶清單(管理者、會員)
	// api.POST("/account/force-logout", chatroom.Echo)        // 強制登出指定用戶(管理者、會員)

	/*
		發文管理 API
	*/
	// api.POST("/chat/text", chatroom.Echo) // 發送文字訊息
	// api.POST("/chat/img", chatroom.Echo)  // 發送圖片訊息
}
