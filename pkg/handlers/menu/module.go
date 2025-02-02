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
		fx.Provide(module),
	)
}

type moduleParams struct {
	fx.In

	Service *Service
}

type moduleResults struct {
	fx.Out

	StartTextHandler    internal.TextHandler     `group:"handlers"`
	HelpTextHandler     internal.TextHandler     `group:"handlers"`
	TestTextHandler     internal.TextHandler     `group:"handlers"`
	MenuCallbackHandler internal.CallbackHandler `group:"handlers"`
}

func module(
	params moduleParams,
) (moduleResults, error) {
	var res moduleResults

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

	res.TestTextHandler = internal.NewRawHandler(params.Service.Test).
		WithMiddleware(middleware.MustHaveMessage).
		Text("/test", bot.MatchTypePrefix)

	res.MenuCallbackHandler = internal.NewRawHandler(params.Service.Menu).
		WithMiddleware(middleware.MustHaveCallback).
		Callback("menu ", bot.MatchTypePrefix)

	return res, nil
}
