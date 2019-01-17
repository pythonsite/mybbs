package models

import "time"

type Reply struct {
	Id      int       `orm:"pk;auto"`
	Topic   *Topic    `orm:"rel(fk)"`
	Content string    `orm:"type(text)"`
	User    *User     `orm:"rel(fk)"`
	Up      int       `orm:"defalut(0)"`
	InTime  time.Time `orm:"auto_now_add;type(datetime)"`
}
