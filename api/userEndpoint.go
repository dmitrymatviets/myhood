package api

import (
	"github.com/dmitrymatviets/myhood/core/contract"
)

type UserEndpoint struct {
	*Server
	contract.IUserService
}

func NewUserEndpoint(s *Server, userService contract.IUserService) *UserEndpoint {
	endpoint := &UserEndpoint{s, userService}
	endpoint.Server = s

	return endpoint
}
