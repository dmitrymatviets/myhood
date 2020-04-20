package dto

import "github.com/dmitrymatviets/myhood/core/model"

type AddFriendRequest struct {
	Session  model.Session `json:"session"`
	FriendId model.IntId   `json:"friendId"`
}

type AddFriendResponse struct {
}
