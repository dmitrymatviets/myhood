package dto

import "github.com/dmitrymatviets/myhood/core/model"

type RemoveFriendRequest struct {
	Session  model.Session `json:"session"`
	FriendId model.IntId   `json:"friendId"`
}

type RemoveFriendResponse struct {
}
