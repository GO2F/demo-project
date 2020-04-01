package component

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"rank/controller/base"
	mComponent "rank/model/component"
)

// Get 获取单条记录
func Get(c *gin.Context) {
	type RequestParams struct {
		ID uint `form:"id"`
	}
	var params RequestParams
	c.ShouldBindQuery(&params)
	component := mComponent.Get(params.ID)
	if component.ID == 0 {
		c.JSON(http.StatusOK, base.Failed("项目不存在", 2))
		return
	}
	c.JSON(http.StatusOK, base.Success(component))
	return
}

// GetList 获取列表
func GetList(c *gin.Context) {
	type RequestParams struct {
		Page uint `form:"page"`
	}
	page := c.DefaultQuery("page", "0")
	pageNum, _ := strconv.Atoi(page)
	componentList := mComponent.GetList(pageNum)
	c.JSON(http.StatusOK, base.Success(componentList))
	return
}

// Create 创建文件
func Create(c *gin.Context) {
	type RequestParams struct {
		DisplayName string `json:"DisplayName"`
		PackageName string `json:"PackageName"`
		DevListJSON string `json:"DevListJSON"`
		Description string `json:"Description"`
		SiteURL     string `json:"SiteURL"`
		Remark      string `json:"Remark"`
	}
	var params RequestParams
	c.ShouldBindJSON(&params)

	component := mComponent.Component{
		DisplayName: params.DisplayName,
		PackageName: params.PackageName,
		DevListJSON: params.DevListJSON,
		Description: params.Description,
		ApplyUcid:   "10086",
		SiteURL:     params.SiteURL,
		Remark:      params.Remark,
	}

	existProj := mComponent.GetUnscopedByPackageName(params.PackageName)
	if existProj.ID != 0 {
		if existProj.DeletedAt == nil {
			// 已存在
			c.JSON(http.StatusOK, base.Success(existProj))
			return
		} else {
			c.JSON(http.StatusOK, base.Failed("项目 "+existProj.PackageName+" 已被删除,请联系管理员恢复", 1))
			return
		}

	}

	mComponent.Create(component)

	c.JSON(http.StatusOK, base.Success(params))
	return
}

// Update 更新项目数据
func Update(c *gin.Context) {
	type RequestParams struct {
		ID          uint   `json:"ID"`
		DisplayName string `json:"DisplayName"`
		DevListJSON string `json:"DevListJSON"`
		Description string `json:"Description"`
		SiteURL     string `json:"SiteURL"`
		Remark      string `json:"Remark"`
	}
	var params RequestParams
	c.ShouldBindJSON(&params)
	component := mComponent.Get(params.ID)
	if component.ID == 0 {
		c.JSON(http.StatusOK, base.Failed(string(params.ID), 2))
		return
	}
	if params.DisplayName != "" {
		component.DisplayName = params.DisplayName
	}
	if params.DevListJSON != "" {
		component.DevListJSON = params.DevListJSON
	}
	if params.Description != "" {
		component.Description = params.Description
	}
	if params.SiteURL != "" {
		component.SiteURL = params.SiteURL
	}
	if params.Remark != "" {
		component.Remark = params.Remark
	}

	mComponent.Update(component)

	c.JSON(http.StatusOK, base.Success(component))
	return
}

// Remove 删除项目
func Remove(c *gin.Context) {
	type RequestParams struct {
		ID uint `form:"ID"`
	}
	var params RequestParams
	c.ShouldBindQuery(&params)
	component := mComponent.Get(params.ID)
	if component.ID == 0 {
		c.JSON(http.StatusOK, base.Failed("项目不存在", 2))
		return
	}

	mComponent.Delete(component.ID)

	c.JSON(http.StatusOK, base.Success("删除成功"))
	return
}
