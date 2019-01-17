package controllers

import (
	"mybbs/filters"
	"mybbs/models"
	"strconv"

	"github.com/astaxie/beego"
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
