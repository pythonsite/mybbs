package utils

import (
	"github.com/astaxie/beego"
	"time"
	"github.com/xeonx/timeago"
	"github.com/russross/blackfriday"
	"mybbs/utils"
	"mybbs/models"
)

func FormatTime(time time.Time) string {
	return timeago.Chinese.Format(time)
}

func MarkDown(content string) string {
	return string(blackfriday.MarkdownCommon([]byte(utils.NoHtml(content))))
}

func HasPermission(userId int, name string) bool {
	return models.FindPermissionByUserIdAndPermissionName(userId, name)
}

func init() {
	beego.AddFuncMap("timeage", FormatTime)
	beego.AddFuncMap("markdown", MarkDown)
	beego.AddFuncMap("haspermission", HasPermission)
}