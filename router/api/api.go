package api

import (
	cApi "rank/controller/api"
	cComponent "rank/controller/api/component"

	"github.com/gin-gonic/gin"
)

// InitAPIRouter 初始化API路由
func InitAPIRouter(Router *gin.RouterGroup) (router gin.IRoutes) {
	apiRouter := Router.Group("api")
	{
		apiRouter.GET("component/get", cComponent.Get)
		apiRouter.GET("component/list", cComponent.GetList)
		apiRouter.POST("component/update", cComponent.Update)
		apiRouter.POST("component/create", cComponent.Create)
		apiRouter.GET("component/remove", cComponent.Remove)
		apiRouter.GET("ping", cApi.Ping)
	}

	return apiRouter
}
