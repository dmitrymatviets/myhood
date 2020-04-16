package dto

import "github.com/dmitrymatviets/myhood/core/model"

type SignupRequest struct {
	model.SignupDto
}

type SignupResponse struct {
	Session model.Session `json:"session"`
	User    *model.User
}
