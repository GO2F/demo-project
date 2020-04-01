package middleware

import (
	"regexp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rank/config"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	if config.Current == config.Prod {
		// 生产环境需要配置跨域域名，否则403
		corsConfig.AllowOrigins = []string{config.App.Host}
	} else {
		// 测试环境下模糊匹配本地开头的请求
		corsConfig.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	corsConfig.AllowCredentials = true
	return cors.New(corsConfig)
}
