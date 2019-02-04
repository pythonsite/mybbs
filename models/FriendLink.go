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

func SaveLink(friendlink *FriendLink) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(friendlink)
	return id
}

func DeleteFriendLinkById(id int) {
	o := orm.NewOrm()
	o.Raw("delete from friend_link where id=?", id).Exec()
}