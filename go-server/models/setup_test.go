package models_test

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"testing"

	"suxin2017.com/server/models"
	. "suxin2017.com/server/models"
	models2 "suxin2017.com/server/models"
)

func TestConnect(t *testing.T) {
	models2.ConnectDatabase()
	user := &models.User{Name: "abc", Email: "abc@123.com", Password: "123123"}
	var hasUser models.User
	s := sha1.New()
	s.Write([]byte(user.Password))

	password := hex.EncodeToString(s.Sum(nil))
	DB.Where(&User{Email: user.Email, Password: password}).First(&hasUser)

	fmt.Printf("%+v", hasUser)
}
