package dto

import "github.com/dmitrymatviets/myhood/core/model"

type LoginRequest struct {
	model.Credentials `json:"session"`
}

type LoginResponse struct {
	Session model.Session `json:"session"`
	User    *model.User   `json:"user"`
}
