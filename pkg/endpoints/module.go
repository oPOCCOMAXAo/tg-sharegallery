package endpoints

import (
	"github.com/opoccomaxao/tg-sharegallery/pkg/endpoints/system"
	"github.com/opoccomaxao/tg-sharegallery/pkg/endpoints/tg"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Options(
		system.Invoke(),
		tg.Invoke(),
	)
}
