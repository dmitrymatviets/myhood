package city

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
)

type InMemCityRepository struct {
	cities map[model.IntId]model.City
}

func (i InMemCityRepository) GetCities(ctx context.Context) ([]model.City, error) {
	panic("implement me")
}

func NewInMemCityRepository() contract.ICityRepository {
	return &InMemCityRepository{
		cities: map[model.IntId]model.City{
			1: model.City{1, "Москва"},
			2: model.City{2, "Санкт-Петербург"},
			3: model.City{3, "Казань"},
			4: model.City{4, "Нижний Новгород"},
		}}
}
