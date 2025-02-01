package app

import (
	"github.com/opoccomaxao/tg-sharegallery/pkg/config"
	"github.com/opoccomaxao/tg-sharegallery/pkg/endpoints"
	"github.com/opoccomaxao/tg-sharegallery/pkg/server"
	"go.uber.org/fx"
)

func Run() error {
	app := fx.New(
		fx.Provide(NewCancelCause),
		config.Module(),
		server.Module(),
		endpoints.Invoke(),
	)

	app.Run()

	return nil
}
