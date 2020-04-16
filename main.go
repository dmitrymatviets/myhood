package main

import (
	"context"
	"github.com/dmitrymatviets/myhood/api"
	"github.com/dmitrymatviets/myhood/infrastructure/config"
	"github.com/dmitrymatviets/myhood/infrastructure/database"
	"github.com/dmitrymatviets/myhood/repository/city"
	"github.com/dmitrymatviets/myhood/repository/user"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.Load,
			database.NewDatabase,
			city.NewMssqlCityRepository,
			user.NewMssqlUserRepository,
			api.NewServer,
		),
		fx.Invoke(startApp),
	).Run()
}

func startApp(server *api.Server, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Stop()
		},
	})
}
