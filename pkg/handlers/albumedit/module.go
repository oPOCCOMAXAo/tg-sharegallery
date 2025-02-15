package albumedit

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("handlers/albumedit",
		fx.Provide(fx.Private, NewService),
		fx.Invoke(RegisterHandlers),
	)
}

func RegisterHandlers() error {
	return nil
}
