package router

import (
	"github.com/gin-gonic/gin"
	"rank/config"
	logger "rank/init/logger"
	"rank/router/api"
	"rank/router/debug"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	if config.Current == config.Prod {
		gin.SetMode(gin.ReleaseMode)
	}
	Router := gin.Default()
	// 添加日志中间件
	Router.Use(logger.LogerMiddleware())
	RouterGroup := Router.Group("")
	api.InitAPIRouter(RouterGroup)
	debug.InitDebugRouter(RouterGroup)

	// 添加静态资源支持
	Router.Static("/static", "./static/static")
	Router.Static("/polyfill", "./static/polyfill")
	staticFileList := []string{"asset-manifest.json", "favicon.ico", "favicon.jpg", "favicon.png", "index.html", "manifest.json", "precache-manifest.38c9e9291e422a3f49dcee478a04bffd.js", "service-worker.js"}
	for _, fileName := range staticFileList {
		Router.StaticFile("/"+fileName, "./static/"+fileName)
	}
	// 手工指定静态文件
	// 设置默认404
	Router.NoRoute(func(c *gin.Context) {
		// 默认返回index.html
		c.File("./static/index.html")
	})
	return Router
}
