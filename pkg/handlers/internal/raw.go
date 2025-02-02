package internal

import (
	"slices"

	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/texts"
)

type RawHandler struct {
	pattern        string
	matchType      bot.MatchType
	handlerFunc    bot.HandlerFunc
	middleware     []bot.Middleware
	cmdDescription []texts.CommandDescription
	matchFunc      bot.MatchFunc
}

func NewRawHandler(
	handler bot.HandlerFunc,
) *RawHandler {
	return &RawHandler{
		handlerFunc: handler,
	}
}

func (h *RawHandler) clone() *RawHandler {
	res := *h
	res.cmdDescription = slices.Clone(h.cmdDescription)
	res.middleware = slices.Clone(h.middleware)

	return &res
}

func (h *RawHandler) WithDescription(
	scope models.CommandScope,
	languageCode models.LanguageCode,
	description string,
) *RawHandler {
	res := h.clone()
	res.cmdDescription = append(res.cmdDescription, texts.CommandDescription{
		Scope:        scope,
		LanguageCode: languageCode,
		Description:  description,
	})

	return res
}

func (h *RawHandler) WithMiddleware(
	middleware ...bot.Middleware,
) *RawHandler {
	res := h.clone()
	res.middleware = append(res.middleware, middleware...)

	return res
}

func (h *RawHandler) Pattern() string {
	return h.pattern
}

func (h *RawHandler) MatchType() bot.MatchType {
	return h.matchType
}

func (h *RawHandler) HandlerFunc() bot.HandlerFunc {
	return h.handlerFunc
}

func (h *RawHandler) MatchFunc() bot.MatchFunc {
	return h.matchFunc
}

func (h *RawHandler) Middleware() []bot.Middleware {
	return h.middleware
}

func (h *RawHandler) Description() []texts.CommandDescription {
	return h.cmdDescription
}

func (h *RawHandler) Text(
	pattern string,
	matchType bot.MatchType,
) TextHandler {
	res := h.clone()
	res.pattern = pattern
	res.matchType = matchType

	return res
}

func (h *RawHandler) Callback(
	pattern string,
	matchType bot.MatchType,
) CallbackHandler {
	res := h.clone()
	res.pattern = pattern
	res.matchType = matchType

	return res
}

func (h *RawHandler) Custom(
	matchFunc bot.MatchFunc,
) CustomHandler {
	res := *h
	res.matchFunc = matchFunc

	return &res
}
