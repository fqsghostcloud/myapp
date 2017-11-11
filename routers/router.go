package routers

import (
	"myapp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/profile ", &controllers.PageInfoController{})
	beego.Router("/user/signup", &controllers.SignUpController{})
	beego.Router("/user/login", &controllers.LoginController{})
}
