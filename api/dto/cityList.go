package dto

import "github.com/dmitrymatviets/myhood/core/model"

type CityListRequest struct {
}

type CityListResponse struct {
	Cities []*model.City `json:"cities"`
}
