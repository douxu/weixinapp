package controllers

import (
	"weixinapp/common"
	s "weixinapp/structs"

	"log"

	"encoding/json"

	"github.com/astaxie/beego"
)

const (
	noRecord int = iota
	oneRecord
)

//BasicConfigController : controll of basic config
type BasicConfigController struct {
	beego.Controller
}

//GetBasicConfig : get basic cofigure from database
func (b *BasicConfigController) GetBasicConfig() (config s.BasicConfig) {
	db := common.Connectdatabase()
	UserID := b.GetString("UserID")
	sql := "SELECT MAinPageTitle, CategoryName, Mode, ThemeColor from " +
		"basicconfiguretable where UserID = ?"
	rows, err := db.Query(sql, UserID)
	if err != nil {
		log.Fatal(err)
	}
	rows.Scan(config.GetMainPageTitle, config.GetCategoryName,
		config.GetMode, config.GetThemeColor)
	return config
}

//InsertOrUpdateBasicConfig : modfiy basic configure
func (b *BasicConfigController) InsertOrUpdateBasicConfig() (result bool) {
	result = false
	var rowCount int
	var basicconfig s.BasicConfig
	UserID := b.GetString("UserID")
	json.Unmarshal(b.Ctx.Input.RequestBody, &basicconfig)
	db := common.Connectdatabase()
	sql := "SELECT count(0) from basicconfiguretable where UserID = ? "
	rows, err := db.Query(sql, UserID)
	if err != nil {
		log.Fatal(err)
	}
	rows.Scan(rowCount)
	if rowCount == oneRecord {
		sql = "UPDATE basicconfiguretable SET MAinPageTitle =? ,CategoryName ＝？，" +
			"Mode ＝?，ThemeColor ＝？ where UserID =?"
	} else if rowCount == noRecord {
		sql = "INSERT INTO basicconfiguretable VALUES(?,?,?,?,?)"
	} else {
		return result
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(basicconfig.GetMainPageTitle(), basicconfig.GetCategoryName(),
		basicconfig.GetMode(), basicconfig.GetThemeColor())
	if err != nil {
		log.Fatal(err)
	}
	effectrow, err := res.RowsAffected()
	if effectrow == int64(oneRecord) {
		result = true
	}
	return result
}
