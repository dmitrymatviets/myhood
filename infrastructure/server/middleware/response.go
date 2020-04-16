package middleware

import (
	"encoding/json"
	"fmt"
	config2 "github.com/dmitrymatviets/myhood/infrastructure/config"
	"github.com/dmitrymatviets/myhood/infrastructure/logger"
	"github.com/dmitrymatviets/myhood/infrastructure/server/config"
	"github.com/dmitrymatviets/myhood/infrastructure/server/protocol"
	"github.com/dmitrymatviets/myhood/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	defaultErrMessage = "Что-то пошло не так..."
	defaultErrCode    = "INTERNAL_ERROR"
)

func ResponseMiddleware(cfg *config.ServerConfig, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		// after request
		sendDecoratedJsonResponse(ctx, cfg, logger)
	}
}

func sendDecoratedJsonResponse(ctx *gin.Context, cfg *config.ServerConfig, logger *logger.Logger) {
	meta := map[string]interface{}{
		"_requestId":  fmt.Sprint(ctx.MustGet(config2.CtxKeyRequestId)),
		"_appVersion": fmt.Sprint(cfg.Version),
	}

	metaJson, _ := json.Marshal(meta)

	var response interface{}

	isError := len(ctx.Errors) > 0
	if isError {
		errMessage := defaultErrMessage
		errCode := defaultErrCode

		err := ctx.MustGet(config2.CtxKeyResponse)
		if realErr, ok := err.(error); ok {
			// залогируем ошибку + внутренности
			logger.Error(ctx, realErr.Error(), "error", fmt.Sprintf("%+v", err))
			if publicError, ok := realErr.(*pkg.PublicError); ok {
				errMessage = publicError.Message
				errCode = string(publicError.Code)
			}
		}

		response = protocol.ResponseError{
			Success:  0,
			Envelope: protocol.Envelope{Meta: metaJson},
			Error: protocol.RError{
				Message: errMessage,
				Code:    errCode,
			},
		}
	} else {
		result, _ := json.Marshal(ctx.MustGet(config2.CtxKeyResponse))
		response = protocol.ResponseSuccess{
			Success:  1,
			Envelope: protocol.Envelope{Meta: metaJson},
			Data:     result,
		}
	}

	logger.Info(ctx, "response",
		"url", ctx.Request.RequestURI,
		"body", response)

	ctx.JSON(http.StatusOK, response)
}
