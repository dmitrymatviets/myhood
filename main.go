package main

import (
	"context"
	"flag"
	"github.com/dmitrymatviets/myhood/api"
	"github.com/dmitrymatviets/myhood/core/contract"
	"github.com/dmitrymatviets/myhood/infrastructure/config"
	"github.com/dmitrymatviets/myhood/infrastructure/database"
	"github.com/dmitrymatviets/myhood/infrastructure/logger"
	"github.com/dmitrymatviets/myhood/infrastructure/server"
	"github.com/dmitrymatviets/myhood/infrastructure/validator"
	"github.com/dmitrymatviets/myhood/repository/city"
	"github.com/dmitrymatviets/myhood/repository/user"
	"github.com/dmitrymatviets/myhood/service"
	"go.uber.org/fx"
	"os"
)

const countGenerateProfiles = 100000

func main() {
	fx.New(
		fx.Provide(
			config.Load,
			database.NewDatabase,
			server.NewHTTPServer,
			logger.New,
			validator.NewValidator,
			city.NewMssqlCityRepository,
			user.NewMssqlUserRepository,
			service.NewAuthService,
			service.NewUserService,
			service.NewCityService,
			api.NewServer,
		),
		fx.Invoke(getCallback()),
	).Run()
}

func getCallback() interface{} {
	isGenerateProfilesMode := flag.Bool("generate", false, "help message for flagname")
	flag.Parse()

	if *isGenerateProfilesMode {
		return startProfilesGeneration
	}

	return startApp
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

func startProfilesGeneration(ur contract.IUserRepository) {
	// hack for faster generation
	if urImpl, ok := ur.(*user.MssqlUserRepository); ok {
		urImpl.DisableFastGenerationMode = false
	}

	err := service.GenerateProfiles(ur, "123456", countGenerateProfiles, 500)
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
