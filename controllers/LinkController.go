package controllers

import (
	"github.com/astaxie/beego"
	"mybbs/filters"
	"mybbs/models"
)

type LinkController struct {
	beego.Controller
}

func (c *LinkController) List() {
	c.Data["PageTitle"] = "友链列表"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	c.Data["FriendLinks"] = models.FindAllFriendLink()
	c.Layout = "layout/layout.tpl"
	c.TplName = "links/list.tpl"
}


