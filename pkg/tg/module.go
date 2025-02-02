package tg

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("tg",
		fx.Provide(moduleProvide),
	)
}

type moduleParams struct {
	fx.In
	fx.Lifecycle

	Config Config
}

type moduleResult struct {
	fx.Out

	Service *Service
}

func moduleProvide(
	params moduleParams,
) (moduleResult, error) {
	var (
		res moduleResult
		err error
	)

	res.Service, err = New(params.Config)
	if err != nil {
		return res, err
	}

	params.Lifecycle.Append(fx.Hook{
		OnStart: res.Service.OnStart,
		OnStop:  res.Service.OnStop,
	})

	return res, nil
}
