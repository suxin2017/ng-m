package models

import (
	"crypto/sha1"
	"encoding/hex"

	"gorm.io/gorm"
)

type User struct {
	CommonModel
	Name     string   `json:"name,omitempty" validate:"required"`
	Email    string   `json:"email,omitempty" validate:"required"`
	Password string   `json:"password,omitempty" validate:"required"`
	Avater   string   `json:"avater,omitempty"`
	Domain   []Domain `gorm:"many2many:user_domain;"`
}

func (u *User) BeforeSave(*gorm.DB) error {
	h := sha1.New()
	h.Write([]byte(u.Password))
	s := h.Sum(nil)
	u.Password = hex.EncodeToString(s)
	return nil
}

func GetUserAndDomainById(id uint) (User, error) {
	var user User
	err := DB.Model(&User{}).Preload("Domain").First(&user, id).Error
	return user, err
}

func AddDomainToUser(domain Domain, userId uint) {
}
