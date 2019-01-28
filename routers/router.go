package routers

import (
	"mybbs/controllers"
	"mybbs/filters"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.IndexController{}, "GET:Index")
	beego.Router("/login", &controllers.IndexController{}, "GET:LoginPage")
	beego.Router("/login", &controllers.IndexController{}, "POST:Login")
	beego.Router("/register", &controllers.IndexController{}, "GET:RegisterPage")
	beego.Router("/register", &controllers.IndexController{}, "POST:Register")
	beego.Router("/logout", &controllers.IndexController{}, "GET:Logout")
	beego.Router("/about", &controllers.IndexController{}, "GET:About")

	beego.Router("/user/:username", &controllers.UserController{}, "GET:Detail")
	beego.Router("/user/setting", &controllers.UserController{}, "GET:ToSetting")
	beego.InsertFilter("/user/setting", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/user/setting", &controllers.UserController{}, "POST:Setting")
	beego.Router("/user/updatepwd", &controllers.UserController{}, "POST:UpdatePwd")
	beego.Router("/user/updateavatar", &controllers.UserController{}, "POST:UpdateAvatar")
	beego.Router("/user/list", &controllers.UserController{}, "GET:List")
	beego.Router("/user/delete/:id([0-9]+)", &controllers.UserController{}, "GET:Delete")
	beego.Router("/user/edit/:id([0-9]+)", &controllers.UserController{}, "GET:Edit")
	beego.Router("/user/edit/:id([0-9]+)", &controllers.UserController{}, "POST:Update")

	beego.Router("/role/list", &controllers.RoleController{}, "GET:List")
	beego.Router("/role/add", &controllers.RoleController{}, "GET:Add")
	beego.Router("/role/add", &controllers.RoleController{}, "Post:Save")
	beego.Router("/role/edit/:id([0-9]+)", &controllers.RoleController{}, "GET:Edit")
	beego.Router("/role/edit/:id([0-9]+)", &controllers.RoleController{}, "Post:Update")
	beego.Router("/role/delete/:id([0-9]+)", &controllers.RoleController{}, "GET:Delete")
}
