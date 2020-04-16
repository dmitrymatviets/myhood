package middleware

import (
	"github.com/dmitrymatviets/myhood/infrastructure"
	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) { ctx.Set(infrastructure.CtxKeyResponse, "404 unknown route") }
}
