package main

import (
	"context"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/core/model"
	"github.com/dmitrymatviets/myhood/infrastructure/config"
	"github.com/dmitrymatviets/myhood/infrastructure/database"
	"github.com/dmitrymatviets/myhood/repository/city"
	"github.com/dmitrymatviets/myhood/repository/user"
	"go.uber.org/fx"
	"log"
	"time"
)

func main() {
	fx.New(
		fx.Provide(
			config.Load,
			database.NewDatabase,
			city.NewMssqlCityRepository,
			user.NewMssqlUserRepository,
		),
		fx.Invoke(startApp),
	).Run()
}

func startApp( /*server *api.Server, */ lc fx.Lifecycle, userRepo contract.IUserRepository) {

	session, user, err := userRepo.SignUp(context.Background(), &model.UserWithPassword{
		User: &model.User{
			Email:       "asdads@dffdf.ru",
			Name:        "Петя",
			Surname:     "Иванов",
			DateOfBirth: time.Now(),
			Gender:      "М",
			Interests:   []string{"программирование, твин пикс"},
			CityId:      1,
			Page: model.Page{
				Slug: "44423",
			},
		},
		Password: "1234",
	})

	log.Println(session, user, err)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			//	server.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			//	return server.Stop()
			return nil
		},
	})
}
