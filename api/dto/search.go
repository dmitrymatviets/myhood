package dto

import "github.com/dmitrymatviets/myhood/core/model"

type SearchRequest struct {
	model.SearchDto
	Session model.Session `json:"session"`
}

type SearchResponse struct {
	Users []*model.DisplayUserDto `json:"users"`
}
