package config

import (
	"github.com/dmitrymatviets/myhood/infrastructure/database"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

type Config struct {
	Database database.DatabaseConfig `envconfig:"db"`
	/*
		Server  config.ServerConfig   `envconfig:"server"`
		Options OptionsConfig         `envconfig:"options"`
	*/
}

type OptionsConfig struct{}

func Load() (dbCfg database.DatabaseConfig, err error) {
	godotenv.Load()

	log.Println("Env variables:")
	for _, environ := range os.Environ() {
		log.Println(environ)
	}

	cfg := &Config{}
	if err = envconfig.Process("myhood", cfg); err != nil {
		return
	}

	log.Printf("Config: %+v\n", cfg)

	return cfg.Database, nil
}
