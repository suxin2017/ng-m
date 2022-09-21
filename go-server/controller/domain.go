package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/gin-gonic/gin"
	"suxin2017.com/server/constants"
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
		var targetDomains []models.Domain
		DBPageChain(models.DB.Model(user), pageParam).Association("Domain").Find(&domains)

		for _, domain := range domains {
			models.DB.Joins("CreatedUser").Joins("UpdatedUser").Find(&domain)
			utils.DebugLog(domain)
			targetDomains = append(targetDomains, domain)
		}

		utils.DebugLog(domains)

		total := models.DB.Model(user).Association("Domain").Count()
		c.JSON(200, utils.Ok(gin.H{
			"total": total,
			"data":  targetDomains,
		}))
	} else {
		c.Error(fmt.Errorf("用户未登录"))
	}
}

type GetDomainByIdParams struct {
	ID uint `json:"id,omitempty" validate:"reqired" form:"id"`
}

func GetDomainById(c *gin.Context) {
	var domainId GetDomainByIdParams
	c.ShouldBind(&domainId)
	if domainId.ID > 0 {
		var domain models.Domain

		models.DB.Joins("CreatedUser").Joins("UpdatedUser").First(&domain, domainId.ID)
		if domain.ID > 0 {

			c.JSON(200, utils.Ok(domain))
			return
		}

		c.Error(fmt.Errorf("域名不存在"))
	} else {
		c.Error(fmt.Errorf("id未填写"))
	}
}

func AddDomainByUser(c *gin.Context) {
	user, exist := c.Get("currentUser")
	var addDomainByUserData models.Domain
	if c.ShouldBind(&addDomainByUserData) == nil && exist {
		addDomainByAddBody(c, addDomainByUserData, user.(*models.User))
	} else {
		utils.DebugLog(addDomainByUserData)
		println(exist)
		c.Error(fmt.Errorf("参数解析错误或者用户不存在"))
	}

}

func addDomainByAddBody(c *gin.Context, addBody models.Domain, user *models.User) {
	var targetUser models.User
	var targetDomain models.Domain
	if models.DB.Where("domain", addBody.Domain).First(&targetDomain) != nil && targetDomain.ID > 0 {

		c.Error(fmt.Errorf("域名已经存在"))
		c.Abort()
		return
	}
	if models.DB.Preload("Domain").First(&targetUser, user.ID).Error != nil {
		c.Error(fmt.Errorf("用户不存在"))
		c.Abort()
		return
	}
	// 域名存在
	if addBody.ID > 0 {
		// c.Error(fmt.Errorf("域名不存在"))
		targetDomain.ID = addBody.ID
		targetDomain.UpdatedUserId = user.ID
		targetDomain.Desc = addBody.Desc
		targetDomain.Domain = addBody.Domain
		if models.DB.Model(&targetDomain).Updates(targetDomain).RowsAffected == 0 {
			log.Println("没更新成功")

		}

	} else {
		// 新增域名
		targetDomain.CreatedUserID = user.ID
		targetDomain.UpdatedUserId = user.ID
		targetDomain.Desc = addBody.Desc
		targetDomain.Domain = addBody.Domain
		targetDomain.Users = []models.User{*user}
		models.DB.Create(&targetDomain)
		log.Println("init domain resource directory")
		constants.InitNotExistDir(constants.GetNginxDomainResourceDir(targetDomain.ID))
	}
	c.JSON(200, utils.OkMessage())

}

type GetDomainResourcePathParams struct {
	ID   uint   `json:"id,omitempty" form:"id"`
	Path string `json:"path,omitempty" form:"path"`
}

func GetDomainResourcePath(c *gin.Context) {
	var params GetDomainResourcePathParams
	if err := c.ShouldBind(&params); err == nil {
		if params.ID == 0 {
			c.Error(fmt.Errorf("id is required"))
			c.Abort()
			return
		}

		basePath := constants.GetNginxDomainResourceDir(params.ID)
		fileList := GetFileTreeFromDir(path.Join(basePath, params.Path))
		c.JSON(200, utils.Ok(fileList))

	} else {
		c.Error(fmt.Errorf("something is error"))
		c.Abort()
		return
	}

}

type FileNode struct {
	Path  string `json:"path,omitempty"`
	Size  int64  `json:"size,omitempty"`
	IsDir bool   `json:"isDir,omitempty"`
}

func GetFileTreeFromDir(dirPath string) []FileNode {
	fileRoot := []FileNode{}
	paths, _ := ioutil.ReadDir(dirPath)
	for _, path := range paths {
		fileRoot = append(fileRoot, FileNode{
			Path:  path.Name(),
			Size:  path.Size(),
			IsDir: path.IsDir(),
		})
	}

	return fileRoot

}
