package endpoints

import (
	"github.com/opoccomaxao/tg-sharegallery/pkg/endpoints/system"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Options(
		system.Invoke(),
	)
}
