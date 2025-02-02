package domain

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("domain",
		fx.Provide(New),
	)
}
