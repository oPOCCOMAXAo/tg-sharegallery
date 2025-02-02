package handlers

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers/internal"
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers/texts"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("handlers",
		texts.Module(),
		fx.Invoke(RegisterHandlers),
	)
}

type RegisterParams struct {
	fx.In
	fx.Lifecycle

	Logger           *slog.Logger
	Tg               *tg.Service
	TextHandlers     []internal.TextHandler     `group:"handlers"`
	CallbackHandlers []internal.CallbackHandler `group:"handlers"`
	CustomHandlers   []internal.CustomHandler   `group:"handlers"`
}

func RegisterHandlers(
	params RegisterParams,
) error {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			for _, handler := range params.TextHandlers {
				if handler == nil {
					continue
				}

				params.Logger.Debug("text",
					slog.String("pattern", handler.Pattern()),
				)
				params.Tg.RegisterTextHandler(handler.Pattern(), handler.MatchType(), handler.HandlerFunc())
				if handler.MatchType() == bot.MatchTypeExact {
					params.Tg.AddCommandDescription(ctx, handler.Pattern(), handler.Description())
				}
			}

			for _, handler := range params.CallbackHandlers {
				if handler == nil {
					continue
				}

				params.Logger.Debug("callback",
					slog.String("pattern", handler.Pattern()),
				)
				params.Tg.RegisterCallbackHandler(handler.Pattern(), handler.MatchType(), handler.HandlerFunc())
			}

			for _, handler := range params.CustomHandlers {
				if handler == nil {
					continue
				}

				params.Logger.Debug("custom")
				params.Tg.RegisterCustomHandler(handler.MatchFunc(), handler.HandlerFunc())
			}

			return nil
		},
	})

	return nil
}
