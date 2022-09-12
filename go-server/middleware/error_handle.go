package middleware

import (
	"github.com/gin-gonic/gin"
	"suxin2017.com/server/utils"
)

func ErrorHandleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()

		for _, v := range ctx.Errors {
			ctx.JSON(200, utils.Error(v.Error()))
			ctx.Abort()
			return
		}
	}
}
