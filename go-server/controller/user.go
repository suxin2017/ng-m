package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"suxin2017.com/server/models"
	"suxin2017.com/server/utils"
)

func CurrentUser(c *gin.Context) {
	user, exist := c.Get("currentUser")
	if exist {
		userAndDomains, _ := models.GetUserAndDomainById(user.(*models.User).ID)
		c.JSON(200, utils.Ok(userAndDomains))
	} else {
		c.String(403, "/login")
	}
}

type CommonPageParam struct {
	PageSize int `json:"pageSize,omitempty"`
	PageNum  int `json:"pageNum,omitempty"`
}

func NewPageParam() CommonPageParam {
	return CommonPageParam{
		PageSize: 10,
		PageNum:  1,
	}
}
func DBPageChain(db *gorm.DB, cp CommonPageParam) *gorm.DB {
	return db.Limit(cp.PageSize).Offset((cp.PageNum - 1) * cp.PageSize)
}

func UserList(c *gin.Context) {
	var commonPageParam CommonPageParam = NewPageParam()
	c.ShouldBindQuery(&commonPageParam)
	var users []models.User
	DBPageChain(models.DB, commonPageParam).Order("created_at desc").Find(&users)
	var total int64
	models.DB.Model(&models.User{}).Count(&total)
	c.JSON(200, utils.Ok(gin.H{
		"total": total,
		"data":  users,
	}))
}

func AddUser(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) == nil {
		if err := utils.CheckStruct(user); err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		result := models.DB.FirstOrCreate(&user, models.User{Email: user.Email})
		log.Println(result)
		if result.RowsAffected > 0 {
			c.JSON(200, utils.OkMessage())
			return
		} else {
			c.Error(fmt.Errorf("邮箱已经存在"))
			return
		}
	}
	c.Error(fmt.Errorf("绑定失败"))

}
func RemoveUser(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) == nil {
		if user.ID == 0 {
			c.Error(fmt.Errorf("id不能为空"))
			return
		}
		result := models.DB.Delete(&models.User{}, user.ID)
		if result.RowsAffected > 0 {
			c.JSON(200, utils.OkMessage())
			return
		}
	}
	c.Error(fmt.Errorf("操作失败"))
}

type DeleteUserBody struct {
	Id uint `json:"id,omitempty"`
}

func DeleteUser(c *gin.Context) {
	var deleteUserBody DeleteUserBody
	if c.ShouldBind(&deleteUserBody) == nil {
		if deleteUserBody.Id > 0 {
			models.DB.Delete(&models.User{}, deleteUserBody.Id)
			c.JSON(200, utils.OkMessage())
		} else {
			c.Error(fmt.Errorf("id不存在"))
		}
		return
	}
	c.Error(fmt.Errorf("绑定失败"))
}
