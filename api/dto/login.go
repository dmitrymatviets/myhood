package dto

import "github.com/dmitrymatviets/myhood/core/model"

type LoginRequest struct {
	model.Credentials
}

type LoginResponse struct {
	Session model.Session `json:"session"`
	User    *model.User   `json:"user"`
}
