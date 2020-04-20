package dto

import "github.com/dmitrymatviets/myhood/core/model"

type LogoutRequest struct {
	model.Session `json:"session"`
}

type LogoutResponse struct {
	User *model.User `json:"user"`
}
