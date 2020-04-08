package config

import "time"

type ServerConfig struct {
	Version         string        `envconfig:"version"`
	Host            string        `envconfig:"host"`
	Port            int           `envconfig:"port"`
	ReadTimeout     time.Duration `envconfig:"read_timeout"`
	WriteTimeout    time.Duration `envconfig:"write_timeout"`
	ShutdownTimeout time.Duration `envconfig:"shutdown_timeout"`
}
