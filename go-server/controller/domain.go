package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"

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

func AddDomainByLoginUser(c *gin.Context) {
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

type AddDomainToUserData struct {
	DomainId uint `form:"domainId"`
	UserId   uint `form:"userId"`
}

func AddDomainToUser(c *gin.Context) {
	var addDomainToUserData AddDomainToUserData
	if c.ShouldBind(&addDomainToUserData) == nil {
		var targetDomain models.Domain
		var targetUser models.User
		models.DB.First(&targetDomain, addDomainToUserData.DomainId)
		models.DB.First(&targetUser, addDomainToUserData.UserId)

		if targetDomain.ID == 0 {
			c.Error(fmt.Errorf("域名不存在"))
			return
		}
		if targetUser.ID == 0 {
			c.Error(fmt.Errorf("用户不存在"))
			return
		}
		err := models.DB.Model(&targetDomain).Association("Users").Append(&targetUser)
		if err != nil {
			c.Error(fmt.Errorf("添加用户失败"))
			log.Println(err)
			return
		}
		c.JSON(200, utils.OkMessage())
	} else {
		c.Error(fmt.Errorf("参数解析错误"))
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
func UplodaResourceToDomain(c *gin.Context) {
	println("=============")
	// Multipart form
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	domainId := c.Request.FormValue("domainId")
	extractArchive := c.Request.FormValue("flag")
	log.Println("=============")
	log.Println(domainId)
	u64, err := strconv.ParseUint(domainId, 10, 32)
	if err != nil {
		c.Error(fmt.Errorf("domain is't a number"))
		return
	}
	var targetDomain models.Domain
	models.DB.First(&targetDomain, u64)
	if targetDomain.ID == 0 {
		c.Error(fmt.Errorf("domain 不存在"))
		return
	}

	dst := constants.GetNginxDomainResourceDir(uint(u64))

	// 如果是压缩包解压

	if extractArchive != "true" {
		fileExt := path.Ext(file.Filename)
		log.Printf("上传文件", fileExt)
		if fileExt == ".tar" || fileExt == ".zip" {
			fileName := file.Filename

			fileName = fmt.Sprintf("%v_%v_%v", time.Now().UnixMilli(), domainId, fileName)
			p := constants.GetNginxUploadDir()
			uploadPath := path.Join(p, fileName)
			targetPath := dst
			log.Printf("上传文件路径%v\n", uploadPath)

			if err := c.SaveUploadedFile(file, uploadPath); err != nil {
				c.Error(fmt.Errorf("上传文件失败"))
				log.Println(err)
				return
			}

			if fileExt == ".tar" {
				cmd := exec.Command("tar", "-xvf", uploadPath, "-C", targetPath)
				out, err := cmd.CombinedOutput()
				println(string(out[:]))

				if err != nil {
					log.Fatal(err)
				}

			}

			if fileExt == ".zip" {
				cmd := exec.Command("unzip", "-d", targetPath, "-o", uploadPath)
				out, err := cmd.CombinedOutput()

				println(string(out[:]))

				if err != nil {
					log.Fatal(err)
				}

			}

		} else {
			c.Error(fmt.Errorf("只支持.zip或.tar"))
			return
		}
	} else {
		log.Printf("上传文件路径%v\n", path.Join(dst, path.Clean(file.Filename)))
		// Upload the file to specific dst.
		if err := c.SaveUploadedFile(file, path.Join(dst, path.Clean(file.Filename))); err != nil {
			c.Error(fmt.Errorf("上传文件失败"))
			log.Println(err)
			return
		}
	}

	c.JSON(http.StatusOK, utils.OkMessage())

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

type GetNginxConfigParam struct {
	DomainId uint `form:"domainId"`
}

func GetNginxConfig(c *gin.Context) {

	var params GetNginxConfigParam
	if err := c.ShouldBind(&params); err == nil {
		if params.DomainId == 0 {
			c.Error(fmt.Errorf("DomainId is required"))
			c.Abort()
			return
		}

		var targetDomain models.Domain

		models.DB.Preload("Locations").First(&targetDomain, params.DomainId)
		utils.DebugLog(targetDomain)
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

		c.JSON(200, utils.Ok(nginxConfig))

	} else {
		c.Error(fmt.Errorf("something is error"))
		c.Abort()
		return
	}
}

func GetUserListByDomain(c *gin.Context) {

	var params GetNginxConfigParam
	if err := c.ShouldBind(&params); err == nil {
		var targetDomain models.Domain
		var userList []models.User

		models.DB.First(&targetDomain, params.DomainId)

		if targetDomain.ID == 0 {
			c.Error(fmt.Errorf("domain isn't exist"))
		}

		models.DB.Model(&targetDomain).Association("Users").Find(&userList)

		c.JSON(200, utils.Ok(gin.H{
			"total": len(userList),
			"data":  userList,
		}))

	} else {
		c.Error(fmt.Errorf("something is error"))
		c.Abort()
		return
	}
}
