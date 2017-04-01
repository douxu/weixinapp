package controllers

import (
	s "weixinapp/structs"

	"log"

	"weixinapp/common"

	"github.com/astaxie/beego"
)

//ArticleListController : controller of articlelist
type ArticleListController struct {
	beego.Controller
}

//GetArticleList : return articlelist from database
func (a *ArticleListController) GetArticleList() (records []s.ArtcileList) {
	AutherID, err := a.GetInt("AuterID")
	if err != nil {
		log.Fatal(err)
	}
	db := common.Connectdatabase()
	sql := "SELECT ArticleTitle,PictureUrl,ArticleTraffic,DataNum " +
		"from articledetailinfotable where UserID = ?"
	rows, err := db.Query(sql, AutherID)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		var record s.ArtcileList
		rows.Scan(record.GetArticleTitle(), record.GetPictureURL(),
			record.GetArticleTraffic(), record.GetDataNum())
		records = append(records, record)
	}
	return records
}
