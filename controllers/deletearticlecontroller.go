package controllers

import (
	s "weixinapp/structs"

	"log"

	"weixinapp/common"

	"github.com/astaxie/beego"
)

//DeleteArtcileController : controller of delete article
type DeleteArtcileController struct {
	beego.Controller
}

//DeleteArticle : func of delete article
func (d *DeleteArtcileController) DeleteArticle() (result s.ReturnStructs) {
	articleID, err := d.GetInt("ArticleID")
	if err != nil {
		log.Fatal(err)
	}
	db := common.Connectdatabase()
	sql := "DELETE FROM articledetailinfotable where ArticleID = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		/*result.SetErrorCode(common.ExecSQLFailed)
		result.SetErrorMsg(err.Error())
		result.SetDataStruct(false)*/
		log.Fatal(err)
	}
	res, err := stmt.Exec(articleID)
	if err != nil {
		result.SetErrorCode(common.ExecSQLFailed)
		result.SetErrorMsg(err.Error())
		result.SetDataStruct(false)
		log.Fatal(err)
	}
	effectrow, err := res.RowsAffected()
	if err != nil {
		result.SetErrorCode(common.NoDataInDB)
		result.SetErrorMsg(err.Error())
		result.SetDataStruct(true)
		log.Fatal(err)
	} else if effectrow == int64(oneRecord) {
		result.SetErrorCode(common.ResultAsNormal)
		result.SetErrorMsg("")
		result.SetDataStruct(true)
	}
	return result
}
