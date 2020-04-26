package service

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
)

type CityService struct {
	cityRepo contract.ICityRepository
}

func NewCityService(cityRepo contract.ICityRepository) contract.ICityService {
	return &CityService{cityRepo}
}

func (c *CityService) GetCities(ctx context.Context) ([]*model.City, error) {
	return c.cityRepo.GetCities(ctx)
}
