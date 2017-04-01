package controllers

import (
	"weixinapp/common"
	"weixinapp/structs"

	"log"

	"github.com/astaxie/beego"
)

//BasicConfigController : controll of basic config
type BasicConfigController struct {
	beego.Controller
}

//GetBasicConfig : get basic cofigure from database
func (b *BasicConfigController) GetBasicConfig() (config structs.BasicConfig) {
	db := common.Connectdatabase()
	UserID := b.GetString("UserID")
	sql := "SELECT MAinPageTitle, CategoryName, Mode, ThemeColor from " +
		"basicconfiguretable where UserID = ?"
	rows, err := db.Query(sql, UserID)
	if err != nil {
		log.Fatal(err)
	}
	/*if rows.Next() {
		rows.Scan(config.GetMainPageTitle, config.GetCategoryName,
			config.GetMode, config.GetThemeColor)
	}*/
	rows.Scan(config.GetMainPageTitle, config.GetCategoryName,
		config.GetMode, config.GetThemeColor)
	return config
}

//UpdateBasicConfig : modfiy basic configure
func (b *BasicConfigController) UpdateBasicConfig() (result bool) {
	result = false
	db := common.Connectdatabase()
	UserID := b.GetString("UserID")
	sql := ""
	rows, err := db.Query(sql, UserID)
	result = true
	return result
}
