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
	domainRouter.POST("/add", controller.AddDomainByUser)
	domainRouter.GET("/info", controller.GetDomainById)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	log.Println("listen server in http://localhost:8081")

	r.Run(":8081")

}
