package config

import (
	"context"
	"log"

	"github.com/audryus/miniature-octo-tribble/types"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Mongo `yaml:"mongo"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Mongo struct {
		URL      string `env-required:"true" yaml:"url" env:"MONGO_URL"`
		DATABASE string `env-required:"true" yaml:"database" env:"MONGO_DATABASE"`
	}
)

func New(ctx context.Context) context.Context {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return context.WithValue(ctx, types.Config, cfg)
}
func Get(ctx context.Context) *Config {
	return ctx.Value(types.Config).(*Config)
}
