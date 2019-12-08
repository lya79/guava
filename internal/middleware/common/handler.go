package common

import (
	"module/database/internal/common/header"

	"github.com/gin-gonic/gin"
)

// GetMiddlewareGroup chatroom中間件
func GetMiddlewareGroup() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		header.InitLangHeader,
		header.InitTimeZoneHeader,
	}
}
