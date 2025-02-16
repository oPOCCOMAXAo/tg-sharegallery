package views

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("views",
		fx.Provide(NewService),
	)
}
