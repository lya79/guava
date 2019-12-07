package chatroom

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Echo echo hello
func Echo(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
