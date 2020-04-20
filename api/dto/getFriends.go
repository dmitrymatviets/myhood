package dto

import "github.com/dmitrymatviets/myhood/core/model"

type GetFriendsRequest struct {
	Session model.Session `json:"session"`
	UserId  model.IntId   `json:"userId"`
}

type GetFriendsResponse struct {
	Friends []*model.DisplayUserDto `json:"friends"`
}
