package controllers

import (
	"mybbs/filters"
	"mybbs/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Detail() {
	username := c.Ctx.Input.Param(":username")
	ok, user := models.FindUserByUserName(username)
	if ok {
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
		c.Data["PageTitle"] = "个人主页"
		c.Data["CurrentUserInfo"] = user
		c.Data["Topics"] = models.FindTopicByUser(&user, 7)
		c.Data["Replies"] = models.FindReplyByUser(&user, 7)
	}
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/detail.tpl"
}

func (c *UserController) ToSetting() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	c.Data["PageTitle"] = "用户设置"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/setting.tpl"
}
