package dto

import "github.com/dmitrymatviets/myhood/core/model"

type CheckSessionRequest struct {
	model.Session `json:"session"`
}

type CheckSessionResponse struct {
	User *model.User `json:"user"`
}
