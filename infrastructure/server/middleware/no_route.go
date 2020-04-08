package middleware

import (
	"github.com/dmitrymatviets/myhood/infrastructure/constants"
	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) { ctx.Set(constants.KeyResponse, "404 unknown route") }
}
