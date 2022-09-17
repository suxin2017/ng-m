package models

type Domain struct {
	CommonModel
	Domain        string `validate:"required" json:"domain,omitempty"`
	Desc          string `json:"desc,omitempty"`
	CreatedUserID uint   `json:"createdUserId,omitempty"`
	CreatedUser   User   `json:"createdUser,omitempty"`
	UpdatedUserId uint   `json:"updatedUserId,omitempty"`
	UpdatedUser   User   `json:"updatedUser,omitempty"`
	Users         []User `gorm:"many2many:user_domain;" json:"users,omitempty"`
}

type Path struct {
}

type Upstream struct {
}
