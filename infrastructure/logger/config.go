package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defaultLevel       = infoLevel
	defaultTimeEncoder = timeEncoderEpoch
)

type LoggerConfig struct {
	Debug       bool     `envconfig:"debug"`
	Level       string   `envconfig:"level"`
	Output      []string `envconfig:"output"`
	TimeEncoder string   `envconfig:"time_encoder"`
	AppName     string   `envconfig:"app_name"`
}

func (c *LoggerConfig) newBuilder() (*zap.Config, error) {
	var config zap.Config
	if c.Debug {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.Level = c.getLevel()

	if len(c.Output) != 0 {
		config.OutputPaths = c.Output
	}

	return &config, nil
}

func (c *LoggerConfig) getEncoder() (zapcore.TimeEncoder, error) {
	switch c.TimeEncoder {
	case timeEncoderEpoch:
		return zapcore.EpochTimeEncoder, nil
	case timeEncoderISO8601:
		return zapcore.ISO8601TimeEncoder, nil
	}

	return nil, errors.New("invalid time encoder")
}

func (c *LoggerConfig) getLevel() zap.AtomicLevel {
	switch c.Level {
	case debugLevel:
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case errorLevel:
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case infoLevel:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case warnLevel:
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	default:
		panic("unknown log level")
	}
}

func (c *LoggerConfig) withDefaults() (config LoggerConfig) {
	if c != nil {
		config = *c
	}

	if len(config.Level) == 0 {
		config.Level = defaultLevel
	}

	if len(config.TimeEncoder) == 0 {
		config.TimeEncoder = defaultTimeEncoder
	}

	return
}
