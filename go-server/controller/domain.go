package controller

import (
	"github.com/gin-gonic/gin"
	"suxin2017.com/server/models"
	"suxin2017.com/server/utils"
)

func GetDominListWithLoginUser(c *gin.Context) {
	user, exist := c.Get("currentUser")
	pageParam := NewPageParam()
	c.ShouldBind(pageParam)
	if exist {
		user = user.(*models.User)
		var domains []models.Domain
		DBPageChain(models.DB.Model(&user), pageParam).Order("created_at desc").Association("domain").Find(&domains)

		var total int64
		models.DB.Model(&models.Domain{}).Count(&total)
		c.JSON(200, utils.Ok(gin.H{
			"total": total,
			"data":  domains,
		}))
	}
}
