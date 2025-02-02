package albumedit

import (
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers/internal"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("handlers/albumedit",
		fx.Provide(fx.Private, NewService),
		fx.Provide(module),
	)
}

type moduleParams struct {
	fx.In

	Service *Service
}

type moduleResults struct {
	fx.Out

	NewAlbum   internal.TextHandler     `group:"handlers"`
	NewAlbumCB internal.CallbackHandler `group:"handlers"`
}

func module(
	params moduleParams,
) (moduleResults, error) {
	var res moduleResults

	return res, nil
}
