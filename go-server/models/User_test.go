package models

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ConnectDatabase()
	os.Exit(m.Run())
}

func TestAddUserAndDomains(t *testing.T) {
	user := User{
		Name:     "jinzhu",
		Password: "123",
		Email:    "fff",
		Avater:   "ddd",
		Domain: []Domain{
			{
				CommonModel: CommonModel{
					ID: 2,
				},
				Domain: "123",
			},
		},
	}
	DB.Create(&user)
}
func TestUser(t *testing.T) {
	u, _ := GetUserAndDomainById(6)
	empJSON, _ := json.MarshalIndent(u, "", "  ")
	fmt.Printf("%s", empJSON)
}
