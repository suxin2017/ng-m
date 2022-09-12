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

		total := models.DB.Model(&user).Association("domain").Count()
		c.JSON(200, utils.Ok(gin.H{
			"total": total,
			"data":  domains,
		}))
	}
}

type AddDomainByUserData struct {
	models.Domain
}

func AddDomainByUser(c *gin.Context) {
	// user, exist := c.Get("currentUser")
	// var addDomainByUserData AddDomainByUserData
	// if c.ShouldBind(&addDomainByUserData) != nil && exist {

	// AddDomainByAddBody(AddUserDomaianBody{UserId: user.(*models.User).ID})
	// }

}
