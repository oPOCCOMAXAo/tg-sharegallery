package tg

import (
	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/internal"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/middleware"
)

func (s *Service) RegisterCallbackHandler(
	pattern string,
	matchType bot.MatchType,
	handler bot.HandlerFunc,
	middlewares ...bot.Middleware,
) {
	middlewares = append([]bot.Middleware{
		s.telemetry(internal.HandlerTypeCallback, pattern),
		middleware.AutoFinishCallback,
		middleware.MustHaveCallback,
	}, middlewares...)

	_ = s.client.RegisterHandler(
		bot.HandlerTypeCallbackQueryData,
		pattern,
		matchType,
		handler,
		middlewares...,
	)
}
