package dto

import "github.com/dmitrymatviets/myhood/core/model"

type GetRecommendationsRequest struct {
	Session model.Session `json:"session"`
}

type GetRecommendationsResponse struct {
	Recommendations []*model.DisplayUserDto `json:"recommendations"`
}
