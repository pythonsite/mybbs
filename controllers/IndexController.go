package controllers

import (
	"mybbs/filters"
	"mybbs/models"
	"strconv"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type IndexController struct {
	beego.Controller
}

// 首页
func (c *IndexController) Index() {
	c.Data["PageTitle"] = "首页"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
	if p == 0 {
		p = 1
	}
	size, _ := beego.AppConfig.Int("page.size")
	s, _ := strconv.Atoi(c.Ctx.Input.Query("s"))
	c.Data["S"] = s
	section := models.Section{Id: s}
	c.Data["Page"] = models.PageTopic(p, size, &section)
	c.Layout = "layout/layout.tpl"
	c.TplName = "index.tpl"
}

func (c *IndexController) About() {
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "关于"
	c.Layout = "layout/layout.tpl"
	c.TplName = "about.tpl"
}

// 登录页
func (c *IndexController) LoginPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		u := models.FindPermissionByUser(1)
		beego.Debug(u)
		c.Data["PageTitle"] = "登录"
		c.Layout = "layout/layout.tpl"
		c.TplName = "login.tpl"
	}
}

// 验证登录
func (c *IndexController) Login() {
	flash := beego.NewFlash()
	username, password := c.Input().Get("username"), c.Input().Get("password")
	if flag, user := models.Login(username); flag && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil {
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30*24*60*60, "/",
			beego.AppConfig.String("cookie.domain"), false, true)
		c.Redirect("/", 302)
	} else {
		flash.Error("用户名或密码错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
}
