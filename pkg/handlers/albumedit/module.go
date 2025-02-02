package albumedit

import (
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers/internal"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("handlers/albumedit",
		fx.Provide(fx.Private, NewService),
		fx.Provide(Handlers),
	)
}

type HandlersParams struct {
	fx.In

	Service *Service
}

type HandlersResults struct {
	fx.Out

	NewAlbum   internal.TextHandler     `group:"handlers"`
	NewAlbumCB internal.CallbackHandler `group:"handlers"`
}

func Handlers(
	params HandlersParams,
) (HandlersResults, error) {
	var res HandlersResults

	return res, nil
}
