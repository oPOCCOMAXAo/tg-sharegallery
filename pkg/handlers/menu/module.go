package menu

import (
	"github.com/opoccomaxao/tg-instrumentation/apimodels"
	pkgrouter "github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/middleware"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("handlers/menu",
		fx.Provide(fx.Private, NewService),
		fx.Invoke(RegisterHandlers),
	)
}

func RegisterHandlers(
	service *Service,
	router *pkgrouter.Router,
) error {
	router.
		Text("/start",
			middleware.RequiredPrivateChat,
			service.Start,
		).
		WithDescription(apimodels.LCEn, apimodels.CSAllPrivateChats, "Start").
		WithDescription(apimodels.LCUk, apimodels.CSAllPrivateChats, "Запуск")

	router.
		Text("/help",
			middleware.RequiredPrivateChat,
			service.Help,
		).
		WithDescription(apimodels.LCEn, apimodels.CSAllPrivateChats, "Help").
		WithDescription(apimodels.LCUk, apimodels.CSAllPrivateChats, "Допомога")

	router.Callback("menu",
		pkgrouter.AutoAnswerCallbackQuery(),
		middleware.RequiredCallbackMessage,
		service.Menu,
	)

	router.Callback("delete",
		pkgrouter.AutoAnswerCallbackQuery(),
		middleware.RequiredCallbackMessage,
		service.Delete,
	)

	return nil
}
