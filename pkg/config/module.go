package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/opoccomaxao/tg-sharegallery/pkg/server"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("config",
		fx.Provide(New),
	)
}

type Config struct {
	fx.Out

	server.Config `envPrefix:"SERVER_"`
}

func New() (Config, error) {
	var res Config

	err := env.ParseWithOptions(&res, env.Options{
		UseFieldNameByDefault: false,
		RequiredIfNoDef:       false,
	})
	if err != nil {
		return res, errors.WithStack(err)
	}

	return res, nil
}
