package app

import (
	"github.com/opoccomaxao/tg-sharegallery/pkg/config"
	"github.com/opoccomaxao/tg-sharegallery/pkg/endpoints"
	"github.com/opoccomaxao/tg-sharegallery/pkg/server"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg"
	"go.uber.org/fx"
)

func Run() error {
	app := fx.New(
		fx.Provide(NewCancelCause),
		config.Module(),
		server.Module(),
		tg.Module(),
		endpoints.Invoke(),
		fx.Invoke(func(*tg.Service) {}),
	)

	app.Run()

	return nil
}
