package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// AdminConfig represents configuration for an admin account.
type AdminConfig struct {
	// Password is the password for the admin account.
	Password string `envconfig:"ADMIN_ACCOUNT_PASSWORD" required:"true"`
	// TokenSigningKey is the key used to sign ID tokens for the admin account.
	TokenSigningKey []byte `envconfig:"TOKEN_SIGNING_KEY" required:"true"`
}

// AdminConfigFromEnv returns an AdminConfig populated from environment
// variables.
func AdminConfigFromEnv() AdminConfig {
	cfg := AdminConfig{}
	envconfig.MustProcess("", &cfg)
	return cfg
}

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
