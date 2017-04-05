package controllers

import (
	s "weixinapp/structs"

	"encoding/json"

	"weixinapp/common"

	"log"

	"github.com/astaxie/beego"
)

//ArticleDetailInfo : controller of article detail info
type ArticleDetailInfo struct {
	beego.Controller
}

//InsertArticleDetailInfo : func of insert article detail info into db
func (a *ArticleDetailInfo) InsertArticleDetailInfo() (result s.ReturnStructs) {
	var articleDeatilInfo s.ArticleDetailInfo
	json.Unmarshal(a.Ctx.Input.RequestBody, &articleDeatilInfo)
	db := common.Connectdatabase()
	sql := "INSERT INTO articledetailinfotable(UserID, ArticleTitle, PictureUrl," +
		" Content, ReadTargetNum, IncreaseStartTime, IncreaseContinuousTime," +
		" RecommandOrDateSet) VALUES(?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {

		log.Fatal(err)
	}
	res, err := stmt.Exec(articleDeatilInfo.GetUserID(), articleDeatilInfo.GetUserID(),
		articleDeatilInfo.GetUserID(), articleDeatilInfo.GetUserID(), articleDeatilInfo.GetUserID(),
		articleDeatilInfo.GetUserID(), articleDeatilInfo.GetUserID(), articleDeatilInfo.GetUserID())
	if err != nil {

		log.Fatal(err)
	}
	effectrow, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	} else if effectrow == int64(oneRecord) {
		result.SetErrorCode(common.ResultAsNormal)
	}
	return result
}

//GetArticleDetailInfo : func of get article info from db
func (a *ArticleDetailInfo) GetArticleDetailInfo() (result s.ReturnStructs) {
	var articleDatailInfo s.ArticleDetailInfo
	articleID, err := a.GetInt("ArticleID")
	if err != nil {
		log.Fatal(err)
	}
	db := common.Connectdatabase()
	sql := "SELECT ArticleTitle, PictureUrl, Content, ReadTargetNum," +
		" IncreaseStartTime, IncreaseContinuousTime, RecommandOrDateSet" +
		" from articledetailinfotable where ArticleID = ?"
	err = db.QueryRow(sql, articleID).Scan(articleDatailInfo.GetMainPageTitle(),
		articleDatailInfo.GetPictureURL(),
		articleDatailInfo.GetArticleContent(),
		articleDatailInfo.GetReadTargetNum(),
		articleDatailInfo.GetIncreaseStartTime(),
		articleDatailInfo.GetIncreaseContinuousTime(),
		articleDatailInfo.GetRecommandOrDataset())
	if err != nil {
		log.Fatal(err)
	}
	result.SetDataStruct(articleDatailInfo)
	result.SetErrorCode(common.ResultAsNormal)
	result.SetErrorMsg("")
	return result
}
