package controller

import (
	"github.com/gin-gonic/gin"
	"suxin2017.com/server/models"
	"suxin2017.com/server/service"
	"suxin2017.com/server/utils"
)

func HandleLogin(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err == nil {
		token, err := service.Login(&user, c)
		if err == nil {
			c.JSON(200, utils.Ok(token))
			return
		}
		c.JSON(200, utils.Error(err.Error()))
	} else {
		c.JSON(200, utils.ErrorWithTrack("解析json失败", err))
	}
}
