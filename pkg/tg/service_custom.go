package tg

import (
	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/internal"
)

func (s *Service) RegisterCustomHandler(
	matcher bot.MatchFunc,
	handler bot.HandlerFunc,
	middlewares ...bot.Middleware,
) {
	middlewares = append([]bot.Middleware{
		s.telemetry(internal.HandlerTypeUnknown, ""),
	}, middlewares...)

	_ = s.client.RegisterHandlerMatchFunc(
		matcher,
		handler,
		middlewares...,
	)
}
