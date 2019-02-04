package models

import (
	"github.com/astaxie/beego/orm"
)


type FriendLink struct {
	Id      int       `orm:"pk;auto"`
	LinkName string `orm:"unique"`
	LinkUrl string `orm:"unique"`
}

func FindAllFriendLink() []*FriendLink {
	o := orm.NewOrm()
	var friendLink FriendLink
	var friendLinks []*FriendLink
	o.QueryTable(friendLink).All(&friendLinks)
	return friendLinks
}