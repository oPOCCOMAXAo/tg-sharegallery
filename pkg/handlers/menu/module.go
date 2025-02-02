package menu

import (
	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers/internal"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/middleware"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("handlers/menu",
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

	StartTextHandler    internal.TextHandler     `group:"handlers"`
	HelpTextHandler     internal.TextHandler     `group:"handlers"`
	MenuCallbackHandler internal.CallbackHandler `group:"handlers"`
}

func Handlers(
	params HandlersParams,
) (HandlersResults, error) {
	var res HandlersResults

	res.StartTextHandler = internal.NewRawHandler(params.Service.Start).
		WithDescription(models.CSAllPrivateChats, models.LCEn, "Start").
		WithDescription(models.CSAllPrivateChats, models.LCUk, "Запуск").
		WithMiddleware(middleware.MustHaveMessage).
		Text("/start", bot.MatchTypeExact)

	res.HelpTextHandler = internal.NewRawHandler(params.Service.Help).
		WithDescription(models.CSAllPrivateChats, models.LCEn, "Help").
		WithDescription(models.CSAllPrivateChats, models.LCUk, "Допомога").
		WithMiddleware(middleware.MustHaveMessage).
		Text("/help", bot.MatchTypeExact)

	res.MenuCallbackHandler = internal.NewRawHandler(params.Service.Menu).
		WithMiddleware(middleware.MustHaveCallback).
		Callback("menu ", bot.MatchTypePrefix)

	return res, nil
}
