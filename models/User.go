package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int    `orm:"pk;auto"`
	Username  string `orm:"unique"`
	Password  string
	Token     string `orm:"unique"`
	Avatar    string
	Email     string    `orm:"null"`
	Url       string    `orm:"null"`
	Signature string    `orm:"null:size(1000)"`
	InTime    time.Time `orm:"auto_now_add;type(datetime)"`
	Roles     []*Role   `orm:"rel(m2m)"`
}

func FindUserById(id int) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id",id).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token",token).One(&user)
	return err != orm.ErrNoRows, user
}

func FindPermissionByUserIdAndPermissionName(userId int, name string) bool {
	o := orm.NewOrm()
	var permission Permission
	o.Raw("select p.* from permission p "+
	  "left join role_permissions rp on p.id = rp.permission_id "+
	  "left join role r on rp.role_id = r.id "+
	  "left join user_roles ur on r.id = ur.role_id "+
	  "left join user u on ur.user_id = u.id "+
	  "where u.id = ? and p.name = ?", userId, name).QueryRow(&permission)
	return permission.Id > 0
}