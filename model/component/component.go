package component

import (
	base "rank/model"
	"strconv"

	"github.com/jinzhu/gorm"
)

// Component 模型
type Component struct {
	gorm.Model
	DisplayName string `gorm:"size:100;default:'';not null;comment:'组件库项目名';"`
	PackageName string `gorm:"size:100;default:'';not null;unique_index;comment:'组件库的npm包名';"`
	ApplyUcid   string `gorm:"size:30;default:'';not null;comment:'申请人ucid';"`
	SiteURL     string `gorm:"size:40;default:'';not null;comment:'网站地址';"`
	DevListJSON string `gorm:"size:300;default:'[]';not null;comment:'json形式的项目开发者列表,格式为[{ucid:'',name:''}]';"`
	IsAllow     int    `gorm:"default:0;comment:'项目是否审核通过.0:未通过,1:已通过';"`
	Description string `gorm:"size:300;default:'';comment:'项目描述';"`
	Remark      string `gorm:"size:300;default:'';comment:'备注';"`
}

const pageSize = 10

// GetList 获取列表
func GetList(offset int) (proj []Component) {
	var recordList []Component
	base.DB.Limit(pageSize).Offset(offset * pageSize).Order(`id desc`).Find(&recordList)
	return recordList
}

// GetAllApproveUIList 获取所有已批准的组件库列表
func GetAllApproveUIList() (proj []Component) {
	var recordList []Component
	base.DB.Where("is_allow = ?", 1).Find(&recordList)
	return recordList
}

// GetAllApprovePackageNameList 获取所有已批准的组件库packageName名
func GetAllApprovePackageNameList() (packageNameList []string) {
	var recordList []Component
	base.DB.Where("is_allow = ?", 1).Find(&recordList)
	for _, record := range recordList {
		packageNameList = append(packageNameList, record.PackageName)
	}
	return packageNameList
}

// Get 获取指定元素
func Get(id uint) (proj Component) {
	var record Component
	base.DB.Where("id = ?", id).First(&record)
	return record
}

// GetByPackageName 获取指定元素
func GetByPackageName(packageName string) (proj Component) {
	var record Component
	base.DB.Where("package_name = ?", packageName).First(&record)
	return record
}

// GetUnscopedByPackageName 获取指定元素-包括被软删除的项目
func GetUnscopedByPackageName(packageName string) (proj Component) {
	var record Component
	base.DB.Unscoped().Where("package_name = ?", packageName).First(&record)
	return record
}

// Update 更新
func Update(record Component) {
	// 自动以id作为更新依据
	base.DB.Model(&record).Where("id = ?", record.ID).Update(Component{
		DisplayName: record.DisplayName,
		PackageName: record.PackageName,
		ApplyUcid:   record.ApplyUcid,
		SiteURL:     record.SiteURL,
		DevListJSON: record.DevListJSON,
		IsAllow:     record.IsAllow,
		Description: record.Description,
		Remark:      record.Remark,
	})
	return
}

// Delete 软删除
func Delete(id uint) {
	var record Component
	record.ID = id
	base.DB.Delete(&record)
	return
}

// Create 创建元素
func Create(record Component) (proj Component) {
	base.DB.Create(&record)
	return record
}

// DebugCreateTable 自动创建数据表
// 仅供调试使用
func DebugCreateTable() {
	if base.DB.HasTable(&Component{}) == false {
		base.DB.CreateTable(&Component{})
	}
}

// DebugAutoMigrate 自动迁移数据(补全缺失的列定义)
// 仅供调试使用
func DebugAutoMigrate() {
	base.DB.AutoMigrate(&Component{})
}

// DebugFillRecord 初始化测试数据
// 仅供调试使用
func DebugFillRecord() {

	var existRecordList []Component
	base.DB.Find(&existRecordList)

	startIndex := len(existRecordList)
	if startIndex >= 50 {
		return
	}

	for i := startIndex; i < startIndex+50; i++ {
		indexStr := strconv.Itoa(i)
		record := Component{
			DisplayName: "测试项目:" + indexStr,
			PackageName: "test-name-" + indexStr,
			ApplyUcid:   "10086",
			SiteURL:     "www.ke.com",
			DevListJSON: "贝壳,链家,德佑",
			IsAllow:     1,
			Description: "第" + indexStr + "个测试项目",
			Remark:      "备注:" + indexStr,
		}
		Create(record)
	}
	return
}
