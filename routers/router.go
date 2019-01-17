package routers

import (
	"mybbs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.IndexController{},"GET:Index")
	beego.Router("/", &controllers.IndexController{}, "GET:LoginPage")
	beego.Router("/about", &controllers.IndexController{}, "GET:About")
}
