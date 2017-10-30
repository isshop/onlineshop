package routers

import (
	"onlineshop/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.MainController{})
	beego.Router("/register", &controllers.MainController{}, "post:UserRegister")
}
