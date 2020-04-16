package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dmitrymatviets/myhood/infrastructure"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	debugLevel = "debug"
	errorLevel = "error"
	infoLevel  = "info"
	warnLevel  = "warn"
)

const (
	timeEncoderEpoch   = "epoch"
	timeEncoderISO8601 = "ISO8601"
)

func New(cfg LoggerConfig) (*Logger, error) {
	config := cfg.withDefaults()
	builder, err := config.newBuilder()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	logger, err := builder.Build()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return wrap(logger.With(zap.String("app_name", cfg.AppName))), nil
}

func wrap(logger *zap.Logger) *Logger {
	return &Logger{logger: logger.Sugar()}
}

type Logger struct {
	logger *zap.SugaredLogger
}

func (log *Logger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.logger.Debugw(msg, log.addCtxFields(ctx, keysAndValues...)...)
}

func (log *Logger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.logger.Infow(msg, log.addCtxFields(ctx, keysAndValues...)...)
}

func (log *Logger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.logger.Warnw(msg, log.addCtxFields(ctx, keysAndValues...)...)
}

func (log *Logger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.logger.Errorw(msg, log.addCtxFields(ctx, keysAndValues...)...)
}

func (log *Logger) Fatal(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.logger.Fatalw(msg, log.addCtxFields(ctx, keysAndValues...)...)
}

func (log *Logger) addCtxFields(ctx context.Context, keysAndValues ...interface{}) []interface{} {
	result := make([]interface{}, 0)

	if requestId, ok := ctx.Value(infrastructure.CtxKeyRequestId).(string); ok {
		result = append(result, infrastructure.CtxKeyRequestId, requestId)
	}

	return append(result, keysAndValues...)
}

func ToString(o interface{}) string {
	err, ok := o.(error)
	if ok {
		return err.Error()
	}

	strVal, ok := o.(string)
	if ok {
		return strVal
	}

	buf, ok := o.([]byte)
	if ok {
		// json
		return string(buf)
	}

	if jsonBody, err := json.Marshal(o); err == nil {
		return string(jsonBody)
	}

	return fmt.Sprintf("%v", o)
}
