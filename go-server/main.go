package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"suxin2017.com/server/constants"
	"suxin2017.com/server/controller"
	"suxin2017.com/server/middleware"
	"suxin2017.com/server/models"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	constants.InitEnvDir()
	models.ConnectDatabase()

	println(constants.GetGinLogPath())
	f, err := os.Create(constants.GetGinLogPath())
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Disable Console Color
	// gin.DisableConsoleColor()
	gin.ForceConsoleColor()
	r := gin.Default()
	r.Use(middleware.ErrorHandleMiddleware())
	r.Use(middleware.Auth())

	api := r.Group("/api")

	r.Static("/domain-preview", constants.GetNginxDomainConfigDir())

	// Ping test
	api.GET("/ping", func(c *gin.Context) {
		user, exist := c.Get("currentUser")
		log.Printf("user %#v,exist %v", user.(*models.User), exist)
		c.String(http.StatusOK, "pong")
	})

	api.POST("/login", controller.HandleLogin)
	userRouter := api.Group("/user")
	userRouter.GET("/info", controller.CurrentUser)
	userRouter.POST("/add", controller.AddUser)
	userRouter.POST("/delete", controller.DeleteUser)
	userRouter.GET("/list", controller.UserList)

	api.GET("/term", controller.HandleTerm)

	domainRouter := api.Group("/domain")
	domainRouter.GET("/list", controller.GetDominListWithLoginUser)
	domainRouter.POST("/add", controller.AddDomainByLoginUser)
	domainRouter.GET("/info", controller.GetDomainById)
	domainRouter.GET("/users", controller.GetUserListByDomain)
	domainRouter.POST("/user/add", controller.AddDomainToUser)

	domainRouter.GET("/resource", controller.GetDomainResourcePath)
	domainRouter.GET("/nginxConfig", controller.GetNginxConfig)
	domainRouter.POST("/resource/upload", controller.
		UplodaResourceToDomain)
	locationRouter := domainRouter.Group("/path")
	locationRouter.POST("/add", controller.AddPath)
	locationRouter.GET("/list", controller.GetPathList)
	locationRouter.POST("/add/root", controller.SaveStaticPathConfig)
	locationRouter.GET("/info", controller.GetPathInfo)

	nginxRouter := api.Group("nginx")
	nginxRouter.GET("/status", controller.GetNginxStatus)
	nginxRouter.GET("/stop", controller.StopNginx)
	nginxRouter.GET("/start", controller.StartNginx)
	nginxRouter.GET("/reload", controller.ReloadNginx)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	log.Println("listen server in http://localhost:8081")

	r.Run(":8081")

}
