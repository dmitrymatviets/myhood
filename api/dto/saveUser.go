package dto

import "github.com/dmitrymatviets/myhood/core/model"

type SaveUserRequest struct {
	Session model.Session `json:"session"`
	User    *model.User   `json:"user"`
}

type SaveUserResponse struct {
	User *model.User `json:"user"`
}
