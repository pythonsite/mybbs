package models

type Role struct {
	Id          int           `orm:"pk;auto"`
	Name        string        `orm:"unique"`
	Users       []*User       `orm:"reverse(many)"`
	Permissions []*Permission `orm:"rel(m2m)"`
}
