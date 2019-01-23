package controllers

import (
	"mybbs/filters"
	"mybbs/models"
	"regexp"

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

func (c *UserController) Setting() {
	flash := beego.NewFlash()
	email, url, signature := c.Input().Get("email"), c.Input().Get("url"), c.Input().Get("signature")
	if len(email) > 0 {
		ok, _ := regexp.MatchString("^([a-z0-9A-Z]+[-|_|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$", email)
		if !ok {
			flash.Error("请输入正确的邮箱地址")
			flash.Store(&c.Controller)
			c.Redirect("/user/setting", 302)
			return
		}
	}

	if len(signature) > 1000 {
		flash.Error("个人签名长度不能超过1000字符")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}

	_, user := filters.IsLogin(c.Ctx)
	user.Email = email
	user.Url = url
	user.Signature = signature
	models.UpdateUser(&user)
	flash.Success("更新资料成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}
