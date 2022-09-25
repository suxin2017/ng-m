package controller

import (
	"fmt"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"suxin2017.com/server/constants"
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

type GetPathListParams struct {
	CommonPageParam
	DomainId uint `form:"domainId"`
}

func GetPathList(c *gin.Context) {
	var pathParam GetPathListParams = GetPathListParams{
		CommonPageParam: NewPageParam(),
	}
	if err := c.ShouldBind(&pathParam); err == nil {
		var targetDomain models.Domain
		var resultPageList []models.Location
		if pathParam.DomainId == 0 {
			c.Error(fmt.Errorf("domain id is empty"))
			return
		}

		models.DB.First(&targetDomain, pathParam.DomainId)
		if targetDomain.ID == 0 {
			c.Error(fmt.Errorf("domain is't existed"))
			return
		}
		DBPageChain(models.DB.Model(&targetDomain), pathParam.CommonPageParam).Association("Locations").Find(&resultPageList)
		total := models.DB.Model(&targetDomain).Association("Locations").Count()
		c.JSON(200, utils.Ok(gin.H{
			"total": total,
			"data":  resultPageList,
		}))

	} else {
		c.Error(err)
	}

}

type GetPathInfoParam struct {
	PathId uint `form:"pathId"`
}

func GetPathInfo(c *gin.Context) {
	var pathParam GetPathInfoParam
	if err := c.ShouldBind(&pathParam); err == nil {
		var targetLocation models.Location
		targetLocation.ID = pathParam.PathId
		models.DB.First(&targetLocation, targetLocation.ID)
		if targetLocation.ID == 0 {
			c.Error(fmt.Errorf("path is't existed"))
			return
		}
		c.JSON(200, utils.Ok(targetLocation))

	} else {
		c.Error(err)
	}

}

type AddStaticPathData struct {
	DomainId   uint   `form:"domainId"`
	PathId     uint   `form:"pathId"`
	Root       string `form:"root"`
	IP         string `form:"ip"`
	UpstreamId int    `from:"upstreamId"`
	Port       string `from:"port"`
	Type       int    `form:"type"`
}

func SaveStaticPathConfig(c *gin.Context) {
	user, exist := c.Get("currentUser")
	var pathParam AddStaticPathData
	if err := c.ShouldBind(&pathParam); err == nil && exist {
		var targetDomain models.Domain
		models.DB.First(&targetDomain, pathParam.DomainId)
		if targetDomain.ID == 0 {
			c.Error(fmt.Errorf("domain不存在"))
			c.Abort()
			return
		}
		var targetLocation models.Location

		models.DB.Model(&targetDomain).Association("Locations").Find(&targetLocation, pathParam.PathId)
		utils.DebugLog(targetLocation)

		if targetLocation.ID == 0 {
			c.Error(fmt.Errorf("路径不存在"))
			c.Abort()
			return
		}
		targetLocation.Type = pathParam.Type

		if pathParam.Type == 1 {
			targetLocation.Root = pathParam.Root
		} else if pathParam.Type == 2 {

			targetLocation.ServerPort = pathParam.Port
			targetLocation.ServerHost = pathParam.IP

		} else {
			c.Error(fmt.Errorf("不支持的类型"))
			return
		}

		if err := models.DB.Transaction(func(tx *gorm.DB) error {

			tx.Model(&targetLocation).Updates(targetLocation)

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
		models.DB.Preload("Locations").First(&targetDomain, pathParam.DomainId)

		nginxConfig := models.GetNginxConfigFromLocationAndDomain(targetDomain.Locations, targetDomain)
		fmt.Println(nginxConfig)

		domainNginxConfigPath := path.Join(constants.GetNginxConDDir(), fmt.Sprintf("%d.conf", targetDomain.ID))
		file, err := os.Create(domainNginxConfigPath)
		if err != nil {
			c.Error(fmt.Errorf("文件处理失败"))
			return
		}
		defer file.Close()

		file.WriteString(nginxConfig)
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
