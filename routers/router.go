package routers

import (
	"weixinapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/login/VerifyLoginInfo", &controllers.UserController{},"post:LoginVerify")
}
