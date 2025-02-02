package tg

import (
	"log/slog"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("tg",
		fx.Provide(newModule),
	)
}

type moduleParams struct {
	fx.In
	fx.Lifecycle

	Config Config
	Logger *slog.Logger
}

type moduleResult struct {
	fx.Out

	Service *Service
}

func newModule(
	params moduleParams,
) (moduleResult, error) {
	var (
		res moduleResult
		err error
	)

	res.Service, err = New(
		params.Config,
		params.Logger,
	)
	if err != nil {
		return res, err
	}

	params.Lifecycle.Append(fx.Hook{
		OnStart: res.Service.OnStart,
		OnStop:  res.Service.OnStop,
	})

	return res, nil
}
