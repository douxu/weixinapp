package controllers

import (
	"fmt"
	"log"
	"weixinapp/common"
	s "weixinapp/structs"

	"github.com/astaxie/beego"
)

//BrowseRecordController : controller of browserecord
type BrowseRecordController struct {
	beego.Controller
}

//GetBrowseReocrd ：注解路由
//GetBrowseReocrd : get browse record from database
func (b *BrowseRecordController) GetBrowseReocrd() (records []s.BrowseRecord) {
	articleID, err := b.GetInt("articleID")
	if err != nil {
		log.Fatal(err)
	}
	/*articleID, err := b.GetInt(":articleID")
	if err != nil {
		log.Fatal(err)
	}*/
	db := common.Connectdatabase()
	sql := "SELECT BrowseName,browseTime from browserecordtable where id = ?"
	rows, err := db.Query(sql, articleID)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		var record s.BrowseRecord
		err := rows.Scan(record.GetBrowseName(), record.GetBrowseTime())
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, record)
	}
	return records
}

//InsertRecord ：注解路由
//InsertRecord : api.insertrecord/articleID/:articleID/browseName/:browseName/browseTime/:browseTime
//InsertRecord : api.insertrecord/get?artcileID=xxxxx&browseName=xxxxx&browseTime=xxxxx
//InsertRecord : (POST)insert into database and return affected rowCount
func (b *BrowseRecordController) InsertRecord() (result bool) {
	result = false
	articleID, err := b.GetInt("articleID")
	if err != nil {
		log.Fatal(err)
	}
	browseName := b.GetString("browseName")
	browseTime := b.GetString("browseTime")
	/*articleID, err := b.GetInt(":articleID")
	if err != nil {
		log.Fatal(err)
	}
	browseName := b.GetString(":browseName")
	browseTime := b.GetString(":browseTime")*/
	db := common.Connectdatabase()
	stmt, err := db.Prepare("INSERT INTO browserecordtable values(?,?,?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(articleID, browseName, browseTime)
	if err != nil {
		log.Fatal(err)
	}
	affectrow, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		fmt.Printf("AffectRow:%d", affectrow)
	}
	result = true
	return result
}
