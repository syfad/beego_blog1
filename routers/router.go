package routers

import (
	"beego_blog1/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//===================前台页面=====================
	//首页
	beego.Router("/", &controllers.MainController{},"*:Index")

	//test
	beego.Router("/link", &controllers.LinkController{},"*:Link")



	//===================后台页面=====================
}