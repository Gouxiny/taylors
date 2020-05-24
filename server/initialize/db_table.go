package initialize

import (
	"gin-vue-admin/dao"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

//注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(model.SysUser{},
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.JwtBlacklist{},
		model.SysWorkflow{},
		model.SysWorkflowStepInfo{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},
		model.ExaCustomer{},
		model.Stock{},
	)
	global.GVA_DB.Model(&model.Stock{}).AddIndex("base_index", "symbol", "name", "exchange", "code")
	dao.Init()
	global.GVA_LOG.Debug("register table success")
}
