package chatroom

import (
	"github.com/lya79/guava/internal/handler/chatroom"
	"github.com/lya79/guava/internal/middleware/common"

	"github.com/gin-gonic/gin"
)

// Provider 設定 router
func Provider(engine *gin.Engine) { // TODO 實作 chatroom router
	setAPIRouter(engine)
}

// SetAPIRouter 設定 API與 handler
func setAPIRouter(engine *gin.Engine) { // TODO 待確認需要哪些中間件
	api := engine.Group(
		"/api",
		common.GetMiddlewareGroup()...,
	)

	/*
		語系與時區 API
	*/
	// api.POST("/personalization/lang", chatroom.Signup)   // 語系
	// api.POST("/personalization/timezone", chatroom.Signup)   // 時區

	/*
		身份驗證 API
	*/

	/*
		帳號管理 API
	*/
	api.POST("/account/signup", chatroom.Signup) // 註冊
	// api.POST("/account/update/pwd", chatroom.Echo) // 更新密碼
	// api.POST("/account/login", chatroom.Echo)      // 登入
	// api.POST("/account/logout", chatroom.Echo)     // 登出
	// api.POST("/account/online-account-list", chatroom.Echo) // 在線用戶清單(管理者、會員)
	// api.POST("/account/force-logout", chatroom.Echo)        // 強制登出指定用戶(管理者、會員)

	/*
		發文管理 API
	*/
	// api.POST("/chat/text", chatroom.Echo) // 發送文字訊息
	// api.POST("/chat/img", chatroom.Echo)  // 發送圖片訊息
}
