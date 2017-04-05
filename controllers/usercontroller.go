package controllers

import (
	"encoding/json"
	"time"
	s "weixinapp/structs"

	"log"

	"weixinapp/common"

	"github.com/astaxie/beego"
)

//UserController : controller of login
type UserController struct {
	beego.Controller
}

//LoginVerify : func of verify login info
func (u *UserController) LoginVerify() (token string) {
	var userInfo s.Userinfo
	var userPassword string
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		log.Fatal(err)
	}
	db := common.Connectdatabase()
	sql := "SELECT UserPassword FROM usertable WHERE UserName = ?"
	rows, err := db.Query(sql, userInfo.GetUserName())
	if err != nil {
		log.Fatal(err)
	}
	rows.Scan(userPassword)
	if userPassword == userInfo.GetUserPassword() {
		timestamp := time.Unix(time.Now().Unix(), 0)
		time := timestamp.Format("2006-01-02 03:04:05 PM")
		token = common.GenerateTokenInfo(userInfo.GetUserName(),
			userInfo.GetUserPassword(), time)
	}
	return token
}
