package common

import (
	"github.com/lya79/guava/internal/common/header"

	"github.com/gin-gonic/gin"
)

// GetMiddlewareGroup chatroom中間件
func GetMiddlewareGroup() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		header.InitLangHeader,
		header.InitTimeZoneHeader,
	}
}
