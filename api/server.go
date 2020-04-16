package api

import (
	"github.com/dmitrymatviets/myhood/core/contract"
	baseHTTP "github.com/dmitrymatviets/myhood/infrastructure/server"
)

type Server struct {
	*baseHTTP.HTTPServer
}

func NewServer(httpServer *baseHTTP.HTTPServer, authService contract.IAuthService, userService contract.IUserService) *Server {
	s := &Server{
		HTTPServer: httpServer,
	}

	return s
}
