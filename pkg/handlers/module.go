package handlers

import (
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers/menu"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("handlers",
		menu.Module(),
	)
}
