package initialize

import (
	"taylors/dao"
	"taylors/global"
	"taylors/model"
	"taylors/service"
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
		model.StockMonitor{},
	)
	dao.Init()
	service.Init()
	global.GVA_LOG.Debug("register table success")
}
