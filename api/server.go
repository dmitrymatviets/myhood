package api

import (
	"encoding/json"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/infrastructure"
	baseHTTP "github.com/dmitrymatviets/myhood/infrastructure/server"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	*baseHTTP.HTTPServer
}

func NewServer(httpServer *baseHTTP.HTTPServer, authService contract.IAuthService, userService contract.IUserService) *Server {
	s := &Server{
		HTTPServer: httpServer,
	}

	NewAuthEndpoint(s, authService)
	NewUserEndpoint(s, userService)

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodGet,
		Path:        "/",
		HandleFuncs: []gin.HandlerFunc{static.ServeRoot("/", "ui")},
	})
	return s
}

func (s *Server) UnmarshalRequestData(ctx *gin.Context, to interface{}) error {
	return json.Unmarshal([]byte(s.RequestData(ctx)), to)
}

func (s *Server) RequestData(ctx *gin.Context) string {
	return ctx.GetString(infrastructure.CtxKeyRequest)
}

func (s *Server) RequestMeta(ctx *gin.Context) string {
	return ctx.GetString(infrastructure.CtxKeyMeta)
}

func (s *Server) ResponseError(ctx *gin.Context, err error) {
	_ = ctx.Error(err)
	ctx.Set(infrastructure.CtxKeyResponse, err)
}

func (s *Server) ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.Set(infrastructure.CtxKeyResponse, data)
}

func (e *Server) ApiMethod(ctx *gin.Context, requestDto interface{}, fn func() (interface{}, error)) {
	err := e.UnmarshalRequestData(ctx, &requestDto)
	if err != nil {
		e.ResponseError(ctx, err)
		return
	}

	result, err := fn()
	if err != nil {
		e.ResponseError(ctx, err)
		return
	}

	e.ResponseSuccess(ctx, result)
}
