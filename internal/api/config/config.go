package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	StandardConfig
}

type StandardConfig struct {
	GracefulShutdownTimeout time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT" default:"30s"`
	UIDirectory             string        `envconfig:"UI_DIR" default:"./ui/build"`
}

func ServerConfigFromEnv() ServerConfig {
	cfg := ServerConfig{}
	envconfig.MustProcess("", &cfg.StandardConfig)
	return cfg
}
