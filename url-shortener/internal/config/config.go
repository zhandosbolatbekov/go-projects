package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/ilyakaznacheev/cleanenv"
)

const (
	defaultHTTPPort = 8080
	EnvLocal        = "local"
	EnvProd         = "prod"
)

type Config struct {
	Environment string       `yaml:"environment"`
	Server      ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

func Init(configPath string) (*Config, error) {
	// check if file exists
	if _, err := os.Stat(configPath); err != nil {
		return nil, err
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
