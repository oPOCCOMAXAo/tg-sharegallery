package menu

import (
	"github.com/opoccomaxao/tg-instrumentation/apimodels"
	pkgrouter "github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/middleware"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("handlers/menu",
		fx.Provide(fx.Private, NewService),
		fx.Invoke(RegisterHandlers),
	)
}

func RegisterHandlers(
	service *Service,
	router *pkgrouter.Router,
) error {
	router.Text("/start", service.Start).
		WithDescription(apimodels.LCEn, apimodels.CSAllPrivateChats, "Start").
		WithDescription(apimodels.LCUk, apimodels.CSAllPrivateChats, "Запуск")

	router.Text("/help", service.Help).
		WithDescription(apimodels.LCEn, apimodels.CSAllPrivateChats, "Help").
		WithDescription(apimodels.LCUk, apimodels.CSAllPrivateChats, "Допомога")

	router.Callback("menu",
		pkgrouter.AutoAnswerCallbackQuery(),
		middleware.RequiredCallbackMessage,
		service.Menu,
	)

	return nil
}
