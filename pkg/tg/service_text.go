package tg

import (
	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/internal"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/middleware"
)

func (s *Service) RegisterTextHandler(
	pattern string,
	matchType bot.MatchType,
	handler bot.HandlerFunc,
	middlewares ...bot.Middleware,
) {
	middlewares = append([]bot.Middleware{
		s.telemetry(internal.HandlerTypeText, pattern),
		middleware.MustHaveMessage,
	}, middlewares...)

	_ = s.client.RegisterHandler(
		bot.HandlerTypeMessageText,
		pattern,
		matchType,
		handler,
		middlewares...,
	)
}
