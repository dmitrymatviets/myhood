package api

import (
	"github.com/dmitrymatviets/myhood/api/dto"
	"github.com/dmitrymatviets/myhood/core/contract"
	baseHTTP "github.com/dmitrymatviets/myhood/infrastructure/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthEndpoint struct {
	*Server
	authService contract.IAuthService
}

func NewAuthEndpoint(s *Server, authService contract.IAuthService) *AuthEndpoint {
	endpoint := &AuthEndpoint{s, authService}
	endpoint.Server = s

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/auth/signup",
		HandleFuncs: []gin.HandlerFunc{endpoint.SignupV1},
	})

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/auth/login",
		HandleFuncs: []gin.HandlerFunc{endpoint.LoginV1},
	})

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/auth/logout",
		HandleFuncs: []gin.HandlerFunc{endpoint.LogoutV1},
	})

	s.AddRoutes(&baseHTTP.Route{
		Method:      http.MethodPost,
		Path:        "v1/auth/checkSession",
		HandleFuncs: []gin.HandlerFunc{endpoint.CheckSessionV1},
	})

	return endpoint
}

func (e *AuthEndpoint) SignupV1(ctx *gin.Context) {
	var requestDto dto.SignupRequest
	e.ApiMethod(ctx, requestDto, func() (interface{}, error) {
		session, user, err := e.authService.SignUp(ctx, requestDto.SignupDto)
		if err != nil {
			return nil, err
		}
		return dto.SignupResponse{
			Session: session,
			User:    user,
		}, nil
	})
}

func (e *AuthEndpoint) LoginV1(ctx *gin.Context) {
	var requestDto dto.LoginRequest
	e.ApiMethod(ctx, requestDto, func() (interface{}, error) {
		session, user, err := e.authService.Login(ctx, requestDto.Credentials)
		if err != nil {
			return nil, err
		}
		return dto.SignupResponse{
			Session: session,
			User:    user,
		}, nil
	})
}

func (e *AuthEndpoint) LogoutV1(ctx *gin.Context) {

}

func (e *AuthEndpoint) CheckSessionV1(ctx *gin.Context) {

}
