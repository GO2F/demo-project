package main

import (
	"rank/config"
	log "rank/init/logger"
	model "rank/model"
	go2fe "github.com/GO2F/Go2Fe"
	"rank/router"
)

type ComponentModel struct {
    ID          string `json:"ID" unique_key:"" show:"" title:"id"`
    DisplayName string `json:"DisplayName" show:"" title:"组件名"`
    PackageName string `json:"PackageName" show:"" title:"包名"`
    DevListJSON string `json:"DevListJSON" show:"" title:"开发者"`
    Description string `json:"Description" show:"" title:"描述"`
    SiteURL     string `json:"SiteURL" show:"" title:"网站主页"`
    Remark      string `json:"Remark" title:"备注"`
}



func main() {
    customerModel := go2fe.ModelDefine{
        DataModel: ComponentModel{},
        Page: go2fe.Page{
            Create: true,
            Update: true,
            Detail: true,
        },
        Name:        "测试模型",
        BasePath:    "/component",
        BaseAPIPath: "/api/component",
    }
    go2fe.RegModel(customerModel)

    if go2fe.Bootstrap() {
        return
	}
	
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
	