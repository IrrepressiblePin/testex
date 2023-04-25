package configuration

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path/filepath"
)

type ApplicationConfig struct {
	ServerPort         int `env:"SERVER_PORT" env-default:"3000"`
	ServerWriteTimeout int `env:"SERVER_WRITE_TIMEOUT" env-default:"15"`
	ServerReadTimeout  int `env:"SERVER_READ_TIMEOUT" env-default:"15"`

	LimitRequest       int    `env:"LIMIT_REQUEST" env-required:"true"`
	LimitSecondRequest int    `env:"LIMIT_SECOND_REQUEST" env-required:"true"`
	IpPrefix           string `env:"IP_PREFIX" env-required:"true"`
}

var applicationConfig ApplicationConfig

func GetConfig() ApplicationConfig {
	return applicationConfig
}

func init() {
	if dir, err := os.Getwd(); err == nil {
		err := cleanenv.ReadConfig(filepath.Join(dir, ".env"), &applicationConfig)
		if err != nil {
			panic(err)
		}
	}
}
