package models

import (
	"strconv"
	"github.com/astaxie/beego/orm"
	"time"
	"mybbs/utils"
)

type Topic struct {
	Id            int       `orm:"pk;auto"`
	Title         string    `orm:"unique"`
	Content       string    `orm:"type(text);null"`
	InTime        time.Time `orm:"auto_now_add;type(datetime)"`
	User          *User     `orm:"rel(fk)"`
	Section       *Section  `orm:"rel(fk)"`
	View          int       `orm:"default(0)"`
	ReplyCount    int       `orm:"default(0)"`
	LastReplyUser *User     `orm:"rel(fk);null"`
	LastReplyTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func SaveTopic(topic *Topic) int64 {
	o:= orm.NewOrm()
	id, _ := o.Insert(topic)
	return id
}

func FindTopicById(id int) Topic {
	o := orm.NewOrm()
	var topic Topic
	o.QueryTable(topic).RelatedSel().Filter("Id", id).One(&topic)
	return topic
}

func PageTopic(p int, size int, section *Section) utils.Page {
	o := orm.NewOrm()
	var topic Topic
	var list []Topic
	qs := o.QueryTable(topic)
	if section.Id > 0 {
		qs = qs.Filter("Section", section)
	}
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("InTime").Limit(size).Offset((p-1)*size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}
