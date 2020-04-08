package middleware

import (
	"github.com/dmitrymatviets/myhood/infrastructure/constants"
	"github.com/dmitrymatviets/myhood/infrastructure/tracing"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

const NoJaeger string = "no_jaeger"

func TracerMiddleware(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(constants.KeyRequestId, uuid.New().String())

		var opts []opentracing.StartSpanOption
		spanContext, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		if err == nil {
			opts = append(opts, opentracing.ChildOf(spanContext))
		}

		operationName := ctx.Request.URL.Path

		span := tracer.StartSpan(operationName, opts...)
		defer span.Finish()

		ctx.Set(tracing.SpanContextKey, span)

		ext.HTTPMethod.Set(span, ctx.Request.Method)
		ext.HTTPUrl.Set(span, ctx.Request.URL.String())

		tracing.SetValue(ctx, constants.KeyRequestId, ctx.MustGet(constants.KeyRequestId).(string))

		jaegerSpanContext, ok := span.Context().(jaeger.SpanContext)
		if ok {
			ctx.Set(constants.KeyTraceId, jaegerSpanContext.TraceID().String())
		} else {
			ctx.Set(constants.KeyTraceId, NoJaeger)
		}

		ctx.Next()

		ext.HTTPStatusCode.Set(span, uint16(ctx.Writer.Status()))
	}
}
