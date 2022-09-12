package models

type Domain struct {
	CommonModel
	Domain string
	Desc   string
	Users  []*User `gorm:"many2many:user_domain;"`
}
