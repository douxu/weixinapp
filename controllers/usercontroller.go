package controllers

import (
	"encoding/json"
	"time"
	s "weixinapp/structs"

	"log"

	"weixinapp/common"

	"fmt"

	"github.com/astaxie/beego"
)

//UserController : controller of login
type UserController struct {
	beego.Controller
}

//LoginVerify : func of verify login info
func (u *UserController) LoginVerify() (result s.ReturnStructs) {
	var userInfo s.Userinfo
	var userPassword string
	var token string
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		log.Fatal(err)
	}
	db := common.Connectdatabase()
	sql := "SELECT UserPassword FROM usertable WHERE UserName = ?"
	rows, err := db.Query(sql, userInfo.UserName)
	if err != nil {
		log.Fatal(err)
	}
	rows.Scan(userPassword)
	if userPassword == userInfo.UserPassword {
		timestamp := time.Unix(time.Now().Unix(), 0)
		time := timestamp.Format("2006-01-02 03:04:05 PM")
		token = common.GenerateTokenInfo(userInfo.UserName,
			userInfo.UserPassword, time)
	}
	//add token into redis
	rc := common.RedisClient.Get()
	defer rc.Close()
	res, err := rc.Do("SET", userInfo.UserName, token, "EX", "604800")
	fmt.Printf("%x", res)
	if err != nil {
		log.Fatal(err)
	}
	result.Insert(0, "", token)
	return result
}
