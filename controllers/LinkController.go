package controllers

import (
	"github.com/astaxie/beego"
	"mybbs/filters"
	"mybbs/models"
	"fmt"
	"strconv"
)

type LinkController struct {
	beego.Controller
}

func (c *LinkController) List() {
	c.Data["PageTitle"] = "友链列表"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	c.Data["FriendLinks"] = models.FindAllFriendLink()
	c.Layout = "layout/layout.tpl"
	c.TplName = "link/list.tpl"
}

func (c *LinkController) Add() {
	c.Data["PageTitle"] = "添加友链"
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	c.Layout = "layout/layout.tpl"
	c.TplName = "link/add.tpl"
}

func (c *LinkController) Save() {
	flash := beego.NewFlash()
	name, url := c.GetString("name"), c.GetString("url")
	fmt.Println(name,url)
	if len(name) == 0 || len(url) == 0 {
		flash.Error("友链地址或友链名称不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/link/add", 302)
	} else {
		link := models.FriendLink{LinkName:name, LinkUrl:url}
		id := models.SaveLink(&link)
		if id > 0 {
			c.Redirect("/link/list", 302)
		} else {
			flash.Error("添加友链错误")
			flash.Store(&c.Controller)
			c.Redirect("/link/add", 302)
		}
	}
}

func (c *LinkController) Delete() {
	id,  _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		models.DeleteFriendLinkById(id)
		c.Redirect("/link/list", 302)
	} else {
		c.Ctx.WriteString("友链不存在")
	}
}


