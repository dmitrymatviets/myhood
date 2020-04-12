package city

import (
	"context"
	"database/sql"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/infrastructure/database"
)

type MssqlCityRepository struct {
	db *database.Database
}

func NewMssqlCityRepository(db *database.Database) contract.ICityRepository {
	return &MssqlCityRepository{db: db}
}

func (cr *MssqlCityRepository) GetCities(ctx context.Context) ([]*model.City, error) {
	cities := make([]*model.City, 0)
	err := cr.db.SelectContext(ctx, &cities, `SELECT city_id
                                              , name 
                                           FROM cities`)
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (cr *MssqlCityRepository) GetById(ctx context.Context, id model.IntId) (*model.City, error) {
	var city model.City
	err := cr.db.GetContext(ctx, &city, `SELECT city_id
                                              , name 
                                       FROM cities
                                      WHERE city_id = ?`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &city, nil
}
