package models

import "time"

type ReplyUpLog struct {
	Id     int       `orm:"pk;auto"`
	User   *User     `orm:"rel(fk)"`
	Reply  *Reply    `orm:"rel(fk)"`
	InTime time.Time `orm:"auto_now_add;type(datetime)"`
}
