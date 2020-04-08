package middleware

import (
	"github.com/dmitrymatviets/myhood/infrastructure/constants"
	"github.com/dmitrymatviets/myhood/infrastructure/server/protocol"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/dmitrymatviets/myhood/infrastructure/logger"
	"github.com/dmitrymatviets/myhood/infrastructure/tracing"
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
			ctx.Set(constants.KeyResponse, err)

			return
		}

		ctx.Set(constants.KeyMeta, req.Meta)
		ctx.Set(constants.KeyRequest, req.Data)

		requestString := req.Data

		logger.Info(ctx, "merchsheet incoming request",
			"url", ctx.Request.RequestURI,
			"headers", ctx.Request.Header,
			//TODO обрезка
			"body", requestString)

		tracing.SetValue(ctx, constants.KeyRequest, requestString)

		ctx.Next()
	}
}
