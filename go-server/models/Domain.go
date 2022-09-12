package models

type Domain struct {
	CommonModel
	Domain        string `validate:"required" json:"domain,omitempty"`
	Desc          string `json:"desc,omitempty"`
	CreatedUserID uint
	CreatedUser   User
	UpdatedUserId uint
	UpdatedUser   User
	Users         []*User `gorm:"many2many:user_domain;" json:"users,omitempty"`
}
