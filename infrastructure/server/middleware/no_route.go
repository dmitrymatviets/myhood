package middleware

import (
	"github.com/dmitrymatviets/myhood/infrastructure/config"
	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) { ctx.Set(config.CtxKeyResponse, "404 unknown route") }
}
