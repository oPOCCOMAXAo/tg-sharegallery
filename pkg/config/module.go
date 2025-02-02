package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/opoccomaxao/tg-sharegallery/pkg/server"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg"
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

	Server server.Config `envPrefix:"SERVER_"`
	TG     tg.Config     `envPrefix:"TG_"`
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
