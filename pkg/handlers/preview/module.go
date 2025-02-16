package preview

import (
	pkgrouter "github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/middleware"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("handlers/preview",
		fx.Provide(fx.Private, NewService),
		fx.Invoke(RegisterHandlers),
	)
}

func RegisterHandlers(
	service *Service,
	router *pkgrouter.Router,
) error {
	router.Callback("preview_album",
		pkgrouter.AutoAnswerCallbackQuery(),
		middleware.RequiredCallbackMessage,
		service.PreviewAlbum,
	)

	router.Callback("view_album",
		pkgrouter.AutoAnswerCallbackQuery(),
		middleware.RequiredCallbackMessage,
		service.ViewAlbum,
	)

	return nil
}
