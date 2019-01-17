package routers

import (
	"mybbs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.IndexController{},"GET:Index")
}
