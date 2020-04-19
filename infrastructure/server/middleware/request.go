package middleware

import (
	"encoding/json"
	"github.com/dmitrymatviets/myhood/infrastructure"
	"github.com/dmitrymatviets/myhood/infrastructure/server/protocol"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"

	"github.com/dmitrymatviets/myhood/infrastructure/logger"
)

func RequestMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// before request

		var requestString json.RawMessage
		if ctx.Request.Method == http.MethodPost {
			var req protocol.Request
			err := ctx.ShouldBindJSON(&req)
			if err != nil {
				err = errors.Wrap(err, "request unmarshal error")
				logger.Error(ctx, err.Error())

				_ = ctx.Error(err)
				ctx.Set(infrastructure.CtxKeyResponse, err)

				return
			}

			ctx.Set(infrastructure.CtxKeyMeta, req.Meta)
			ctx.Set(infrastructure.CtxKeyRequest, req.Data)
			requestString = req.Data

		}

		reqId := uuid.New().String()
		ctx.Set(infrastructure.CtxKeyRequestId, reqId)

		logger.Info(ctx, "request",
			"url", ctx.Request.RequestURI,
			"requestId", reqId,
			"headers", ctx.Request.Header,
			"body", requestString)

		ctx.Next()
	}
}
