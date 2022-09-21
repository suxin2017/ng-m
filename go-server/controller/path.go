package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"suxin2017.com/server/models"
	"suxin2017.com/server/utils"
)

type PathParam struct {
	DomainId  uint   `form:"domainId"`
	MatchRule string `form:"matchRule"`
	Path      string `form:"path"`
}

func AddPath(c *gin.Context) {
	user, exist := c.Get("currentUser")
	var pathParam PathParam
	if err := c.ShouldBind(&pathParam); err == nil && exist {
		var targetDomain models.Domain
		models.DB.First(&targetDomain, pathParam.DomainId)
		if targetDomain.ID == 0 {
			c.Error(fmt.Errorf("domain不存在"))
			c.Abort()
			return
		}
		var targetLocation models.Location
		models.DB.Model(&targetDomain).Where("path", pathParam.Path).Association("Locations").Find(&targetLocation)
		utils.DebugLog(targetLocation)

		if targetLocation.ID > 0 {
			c.Error(fmt.Errorf("路径已经存在"))
			c.Abort()
			return
		}
		if err := models.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&targetDomain).Association("Locations").Append(&models.Location{
				Path:      pathParam.Path,
				MatchRule: pathParam.MatchRule,
			}); err != nil {
				return err
			}
			if err := tx.Model(&targetDomain).Association("UpdatedUser").Replace(&models.User{
				CommonModel: models.CommonModel{
					ID: user.(*models.User).ID,
				},
			}); err != nil {
				return err
			}
			return nil
		}); err != nil {
			c.Error(fmt.Errorf("操作失败"))
			c.Abort()
			return
		}
		c.JSON(200, utils.OkMessage())

	} else {
		println(exist)
		if !exist {

			c.Error(fmt.Errorf("用户不存在"))
		} else {
			c.Error(err)
		}
	}

}
