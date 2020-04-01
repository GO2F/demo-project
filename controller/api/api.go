package api

import (
	"rank/controller/base"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 第一个api接口
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, base.Result(map[string]string{"success": "done"}, "success", 0, "success", ""))
	return
}
