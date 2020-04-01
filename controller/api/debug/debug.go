package debug

import (
	"rank/controller/base"

	"net/http"
	mComponent "rank/model/component"

	"github.com/gin-gonic/gin"
)

// InitDbTable debug:创建数据库表
func InitDbTable(c *gin.Context) {
	mComponent.DebugCreateTable()
	mComponent.DebugAutoMigrate()
	mComponent.DebugFillRecord()
	c.JSON(http.StatusOK, base.Success("数据表创建完毕"))
	return
}

// InitDbTableInCommand debug:在命令行中创建数据库表
func InitDbTableInCommand() {
	mComponent.DebugCreateTable()
	mComponent.DebugAutoMigrate()
	return
}
