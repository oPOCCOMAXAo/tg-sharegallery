package internal

import (
	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-sharegallery/pkg/texts"
)

type Handler interface {
	HandlerFunc() bot.HandlerFunc
	Middleware() []bot.Middleware
}

type PatternMatcher interface {
	Pattern() string
	MatchType() bot.MatchType
}

type Describable interface {
	Description() []texts.CommandDescription
}

type Matcher interface {
	MatchFunc() bot.MatchFunc
}

type TextHandler interface {
	Handler
	PatternMatcher
	Describable
}

type CallbackHandler interface {
	Handler
	PatternMatcher
}

type CustomHandler interface {
	Handler
	Matcher
}
