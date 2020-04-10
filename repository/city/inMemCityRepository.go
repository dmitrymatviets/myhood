package city

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
)

type InMemCityRepository struct {
	cities map[model.IntId]model.City
}

func (i *InMemCityRepository) GetCities(ctx context.Context) ([]model.City, error) {
	result := make([]model.City, 0, 4)
	for _, city := range i.cities {
		result = append(result, city)
	}
	return result, nil
}

func (i *InMemCityRepository) GetById(ctx context.Context, id model.IntId) (*model.City, error) {
	if city, ok := i.cities[id]; ok {
		return &city, nil
	}
	return nil, nil
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
