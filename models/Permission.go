package models

type Permission struct {
	Id               int `orm:"pk;auto"`
	Pid              int
	Url              string
	Name             string
	Description      string
	Roles            []*Role       `orm:"reverse(many)"`
	ChildPermissions []*Permission `orm:"-"`
}
