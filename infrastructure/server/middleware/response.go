package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/dmitrymatviets/myhood/infrastructure/constants"
	"github.com/dmitrymatviets/myhood/infrastructure/logger"
	"github.com/dmitrymatviets/myhood/infrastructure/server/config"
	"github.com/dmitrymatviets/myhood/infrastructure/server/protocol"
	"github.com/dmitrymatviets/myhood/infrastructure/tracing"
	"github.com/dmitrymatviets/myhood/tools/errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ResponseMiddleware(cfg *config.ServerConfig, logger *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		// after request

		if file, ok := ctx.Get(constants.KeyResponseFile); ok {
			if fileDto, ok := file.(protocol.FileResponseDto); ok {
				sendFileResponse(ctx, fileDto, logger)
				return
			}
			panic("incorrect file response")
		}
		sendDecoratedJsonResponse(ctx, cfg, logger)
	}
}

func sendDecoratedJsonResponse(ctx *gin.Context, cfg *config.ServerConfig, logger *logger.Logger) {
	if _, ok := ctx.Get(constants.KeyPlainTextResponse); ok {
		return
	}

	meta := map[string]interface{}{
		"_requestId":  fmt.Sprint(ctx.MustGet(constants.KeyRequestId)),
		"_appVersion": fmt.Sprint(cfg.Version),
	}

	metaJson, _ := json.Marshal(meta)

	result, _ := json.Marshal(ctx.MustGet(constants.KeyResponse))
	var response interface{}

	isError := len(ctx.Errors) > 0
	if isError {
		err := errors.ToDomainError(ctx.MustGet(constants.KeyResponse))
		// залогируем в т.ч. стектрейсы ошибок
		logger.Error(ctx, err.Error(), "error", fmt.Sprintf("%+v", err.Extra["error"]))

		response = protocol.ResponseError{
			Success:  0,
			Envelope: protocol.Envelope{Meta: metaJson},
			Error: protocol.RError{
				Message:     err.Message,
				Code:        err.Code,
				Description: err.Description,
			},
		}
	} else {
		response = protocol.ResponseSuccess{
			Success:  1,
			Envelope: protocol.Envelope{Meta: metaJson},
			Data:     result,
		}
	}

	logger.Info(ctx, "merchsheet outcoming response",
		"url", ctx.Request.RequestURI,
		//TODO обрезка
		"body", response)

	tracing.SetValue(ctx, constants.KeyResponse, response)

	ctx.JSON(http.StatusOK, response)
}

func sendFileResponse(ctx *gin.Context, attachment protocol.FileResponseDto, logger *logger.Logger) {
	ctx.FileAttachment(attachment.Path, attachment.Name)
	stat, _ := os.Stat(attachment.Path)
	textForLog := fmt.Sprintf("file: %v, %v bytes", attachment.Path, stat.Size())
	logger.Info(ctx, "merchsheet outcoming response",
		"url", ctx.Request.RequestURI,
		"body", textForLog)
	tracing.SetValue(ctx, constants.KeyResponse, textForLog)
	_ = os.Remove(attachment.Path)
}
