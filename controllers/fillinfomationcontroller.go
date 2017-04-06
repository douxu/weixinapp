package controllers

import (
	"fmt"
	"log"

	"weixinapp/common"
	s "weixinapp/structs"

	"github.com/astaxie/beego"
)

//FillInfomationController : controller of fillinfomation
type FillInfomationController struct {
	beego.Controller
}

//GetRecord : get fillinfo record from database
func (f *FillInfomationController) GetRecord() (records []s.DataInfo) {
	articleID, err := f.GetInt(":articleID")
	if err != nil {
		log.Fatal(err)
	}
	/*articleID, err := f.GetInt(":articleID")
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
		var record s.DataInfo
		err := rows.Scan(record.ArticleID, record.InfomationList)
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, record)
	}
	return records
}

//Insertrecord : insert record into database
func (f *FillInfomationController) Insertrecord() (result bool) {
	result = false
	articleID, err := f.GetInt(":articleID")
	if err != nil {
		log.Fatal(err)
	}
	fillInfomation := f.GetString("InfomationStruct")
	/*articleID, err := f.GetInt(":articleID")
	if err != nil {
		log.Fatal(err)
	}
	InfomationStruct := f.GetString(":InfomationStruct")*/
	db := common.Connectdatabase()
	stmt, err := db.Prepare("INSERT INTO browserecordtable values(?,?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(articleID, fillInfomation)
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
