package app

import (
	"github.com/opoccomaxao/tg-sharegallery/pkg/config"
	"github.com/opoccomaxao/tg-sharegallery/pkg/db"
	"github.com/opoccomaxao/tg-sharegallery/pkg/domain"
	"github.com/opoccomaxao/tg-sharegallery/pkg/endpoints"
	"github.com/opoccomaxao/tg-sharegallery/pkg/handlers"
	"github.com/opoccomaxao/tg-sharegallery/pkg/logger"
	"github.com/opoccomaxao/tg-sharegallery/pkg/repo"
	"github.com/opoccomaxao/tg-sharegallery/pkg/server"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg"
	"go.uber.org/fx"
)

func Run() error {
	app := fx.New(
		fx.Provide(NewCancelCause),
		fx.WithLogger(NewFxLogger),
		config.Module(),
		logger.Module(),
		server.Module(),
		tg.Module(),
		db.Module(),
		repo.Module(),
		domain.Module(),
		endpoints.Invoke(),
		handlers.Invoke(),
	)

	app.Run()

	return nil
}
