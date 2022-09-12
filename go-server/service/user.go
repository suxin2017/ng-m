package service

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	. "suxin2017.com/server/models"
	"suxin2017.com/server/utils"
)

func Login(user *User, c *gin.Context) (string, error) {
	var hasUser User
	s := sha1.New()
	s.Write([]byte(user.Password))

	password := hex.EncodeToString(s.Sum(nil))
	DB.Where(&User{Email: user.Email, Password: password}).First(&hasUser)
	if hasUser.ID > 0 {
		token, err := utils.Generate(hasUser.ID)
		if err == nil {
			return token, nil
		}
		log.Fatalln("token 失败", err.Error())
	}
	DB.Where(&User{Email: user.Email}).First(&hasUser)
	if hasUser.ID > 0 {
		return "", fmt.Errorf("用户密码错误")
	}
	return "", fmt.Errorf("用户不存在")
}
