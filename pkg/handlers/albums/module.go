package albums

import (
	"github.com/opoccomaxao/tg-instrumentation/apimodels"
	pkgrouter "github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/middleware"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("handlers/albums",
		fx.Provide(fx.Private, NewService),
		fx.Invoke(RegisterHandlers),
	)
}

func RegisterHandlers(
	service *Service,
	router *pkgrouter.Router,
) error {
	router.
		Text("/albums",
			middleware.RequiredPrivateChat,
			service.AlbumsMessage,
		).
		WithDescription(apimodels.LCAll, apimodels.CSAllPrivateChats, "View my albums").
		WithDescription(apimodels.LCUk, apimodels.CSAllPrivateChats, "Переглянути мої альбоми")

	router.Callback("albums",
		pkgrouter.AutoAnswerCallbackQuery(),
		middleware.RequiredCallbackMessage,
		service.Albums,
	)

	router.Callback("list_albums",
		pkgrouter.AutoAnswerCallbackQuery(),
		middleware.RequiredCallbackMessage,
		service.ListAlbums,
	)

	return nil
}
