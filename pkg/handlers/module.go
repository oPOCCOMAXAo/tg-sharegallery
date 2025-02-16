package handlers

import (
	"context"

	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers/albums"
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers/edit"
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers/menu"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("handlers",
		menu.Module(),
		albums.Module(),
		edit.Module(),
		fx.Invoke(FinishHandlers),
	)
}

func FinishHandlers(
	lc fx.Lifecycle,
	router *router.Router,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return router.UpdateCommandsDescription(ctx)
		},
	})
}
