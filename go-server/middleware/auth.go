package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"suxin2017.com/server/models"
	"suxin2017.com/server/utils"
)

var whitePaths = []string{"/api/login", "/api/term", "/domain"}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(ctx.Request.URL.Path)
		for _, path := range whitePaths {
			// if ctx.Request.URL.Path
			if strings.HasPrefix(ctx.Request.URL.Path, path) {
				ctx.Next()
				return
			}

		}
		authorization := ctx.Request.Header.Get("Authorization")
		tokens := strings.Split(authorization, "Bearer ")
		if len(tokens) == 2 {

			token := strings.Trim(tokens[1], " ")
			log.Printf("====%v====", token)
			user, err := utils.ParseToken(token)

			if err == nil {
				var currentUser models.User
				models.DB.First(&currentUser, user.UserId)
				if currentUser.ID > 0 {
					ctx.Set("currentUser", &currentUser)
				}
				ctx.Next()
				return
			} else {
				ctx.String(403, "/login")
				ctx.Abort()
				log.Println("token 解析失败")
				log.Println(err)
				ctx.Abort()
				return
			}

		}
		ctx.String(403, "/login")
		ctx.Abort()

	}
}
