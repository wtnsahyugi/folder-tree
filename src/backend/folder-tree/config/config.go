package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	ServerAddr         string `env:"ADDR,default=localhost:8081"`
	Host               string `env:"DATABASE_HOST,default=localhost"`
	Port               int    `env:"DATABASE_PORT,default=5432"`
	Username           string `env:"DATABASE_USERNAME,required"`
	Password           string `env:"DATABASE_PASSWORD,required"`
	Database           string `env:"DATABASE_NAME,required"`
	ConnectionRetry    int    `env:"DATABASE_CONNECTION_RETRY,default=3"`
	MaxConnection      int    `env:"DATABASE_MAX_CONNECTION,default=10"`
	MaxIdleConnection  int    `env:"DATABASE_MAX_IDLE_CONNECTION,default=10"`
	ConnectionLifeTime int    `env:"DATABASE_CONNECTION_LIFETIME,default=10"`
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	// just skip loading env files if it is not exists, env files only used in local dev
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
