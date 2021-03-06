package main

import (
	_ "mybbs/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mybbs/models"
	_ "mybbs/utils"
	_ "github.com/go-sql-driver/mysql"
	_ "mybbs/templates"
)

func init() {
	url := beego.AppConfig.String("jdbc.url")
	port := beego.AppConfig.String("jdbc.port")
	username := beego.AppConfig.String("jdbc.username")
	password := beego.AppConfig.String("jdbc.password")
	orm.RegisterDataBase("default", "mysql", username+":"+password+"@tcp("+url+":"+port+")/mybbs-go?charset=utf8&parseTime=true&charset=utf8", 30)
	orm.RegisterModel(
		new(models.User),
		new(models.Topic),
		new(models.Section),
		new(models.Reply),
		new(models.ReplyUpLog),
		new(models.Role),
		new(models.Permission),
		new(models.FriendLink),
	)
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}

