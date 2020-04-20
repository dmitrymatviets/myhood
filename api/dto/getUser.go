package dto

import "github.com/dmitrymatviets/myhood/core/model"

type GetUserRequest struct {
	UserId  model.IntId   `json:"userId"`
	Session model.Session `json:"session"`
}

type GetUserResponse struct {
	User *model.User `json:"user"`
}
