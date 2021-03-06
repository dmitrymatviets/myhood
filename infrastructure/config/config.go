package config

import (
	"github.com/dmitrymatviets/myhood/infrastructure/database"
	"github.com/dmitrymatviets/myhood/infrastructure/logger"
	"github.com/dmitrymatviets/myhood/infrastructure/server/config"
	"github.com/dmitrymatviets/myhood/util"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

type Config struct {
	Database database.DatabaseConfig `envconfig:"db"`
	Server   config.ServerConfig     `envconfig:"server"`
	Logger   logger.LoggerConfig     `envconfig:"logger"`
}

func Load() (dbCfg database.DatabaseConfig, serverCfg config.ServerConfig, loggerCfg logger.LoggerConfig, err error) {
	isTestMode := util.IsTestMode()
	envPath := util.RelPathToAbs(".env")
	godotenv.Load(envPath)

	if !isTestMode {
		log.Println("Env variables:")
		for _, environ := range os.Environ() {
			log.Println(environ)
		}
	}

	cfg := &Config{}
	if err = envconfig.Process("myhood", cfg); err != nil {
		return
	}

	if !isTestMode {
		log.Printf("Config: %+v\n", cfg)
	}

	return cfg.Database, cfg.Server, cfg.Logger, nil
}
