package debug

import (
	cDebug "rank/controller/api/debug"
	"github.com/gin-gonic/gin"
)

// InitDebugRouter 初始化API路由
func InitDebugRouter(Router *gin.RouterGroup) (router gin.IRoutes) {
	debugRouter := Router.Group("debug")
	{
		debugRouter.GET("init/database", cDebug.InitDbTable) //初始化数据库
	}

	return debugRouter
}
