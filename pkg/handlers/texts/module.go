package texts

import (
	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers/internal"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/middleware"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("handlers/texts",
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

	StartTextHandler internal.TextHandler `group:"handlers"`
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

	return res, nil
}
