package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Reply struct {
	Id      int       `orm:"pk;auto"`
	Topic   *Topic    `orm:"rel(fk)"`
	Content string    `orm:"type(text)"`
	User    *User     `orm:"rel(fk)"`
	Up      int       `orm:"defalut(0)"`
	InTime  time.Time `orm:"auto_now_add;type(datetime)"`
}

func FindReplyByUser(user *User, limit int) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replies []*Reply
	o.QueryTable(reply).RelatedSel("Topic", "User").OrderBy("-Intime").Limit(limit).All(&replies)
	return replies
}