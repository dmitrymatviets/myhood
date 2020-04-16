package middleware

import (
	"github.com/dmitrymatviets/myhood/infrastructure/config"
	"github.com/dmitrymatviets/myhood/infrastructure/server/protocol"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/dmitrymatviets/myhood/infrastructure/logger"
)

func RequestMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// before request

		var req protocol.Request
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			err = errors.Wrap(err, "request unmarshal error")
			logger.Error(ctx, err.Error())

			_ = ctx.Error(err)
			ctx.Set(config.CtxKeyResponse, err)

			return
		}

		ctx.Set(config.CtxKeyMeta, req.Meta)
		ctx.Set(config.CtxKeyRequestId, req.Data)

		requestString := req.Data

		logger.Info(ctx, "request",
			"url", ctx.Request.RequestURI,
			"headers", ctx.Request.Header,
			"body", requestString)

		ctx.Next()
	}
}
