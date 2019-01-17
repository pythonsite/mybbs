package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"mybbs/models"
)

func IsLogin(ctx *context.Context) (bool, models.User) {
	token, flag := ctx.GetSecureCookie(beego.AppConfig.String("cooke.secure"),beego.AppConfig.String("cookie.token"))
	var user models.User
	if flag {
		flag, user = models.FindUserByToken(token)
	}
	return flag, user
}

