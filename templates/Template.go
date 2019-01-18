package utils

import (
	"mybbs/models"
	"mybbs/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"github.com/xeonx/timeago"
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
	beego.AddFuncMap("timeago", FormatTime)
	beego.AddFuncMap("markdown", MarkDown)
	beego.AddFuncMap("haspermission", HasPermission)
}
