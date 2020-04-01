package main

import (
	"rank/config"
	log "rank/init/logger"
	model "rank/model"
	"rank/router"
)

func main() {
	dbErr := model.InitDb()
	log.Init(config.App.LogPathURI)
	// log.Info("进程以server模式启动")
	if dbErr != nil {
		// log.Error("数据库初始化失败,程序退出")
		return
	}
	r := router.InitRouter()

	// 程序结束前关闭数据库链接
	defer model.DB.Close()
	r.Run(":" + config.App.Port)
}
