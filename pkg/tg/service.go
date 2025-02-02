package tg

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/logger"
	xmodels "github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/texts"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/internal"
	"github.com/pkg/errors"
)

type Service struct {
	config Config
	logger *slog.Logger
	client *bot.Bot
	runCtx context.Context //nolint:containedctx
	cancel context.CancelFunc

	describer *texts.CommandDescriber
}

type Config struct {
	Debug     bool   `env:"DEBUG"`
	NoInit    bool   `env:"NO_INIT"`
	Token     string `env:"TOKEN,required"`
	ServerURL string `env:"SERVER_URL"`
	HookURL   string `env:"HOOK_URL"`
}

func New(
	config Config,
	logger *slog.Logger,
) (*Service, error) {
	res := Service{
		config: config,
		logger: logger,

		describer: texts.NewCommandDescriber(),
	}

	err := res.initClient()
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *Service) initClient() error {
	var err error

	opts := []bot.Option{
		bot.WithSkipGetMe(),
		bot.WithDefaultHandler(s.telemetry(internal.HandlerTypeUnknown, "")(s.defaultHandler)),
		bot.WithDebugHandler(logger.AsPrintf(s.logger.Debug)),
		bot.WithErrorsHandler(s.ErrorHandler),
	}

	if s.config.Debug {
		opts = append(opts, bot.WithDebug())
	}

	if s.config.ServerURL != "" {
		opts = append(opts, bot.WithServerURL(s.config.ServerURL))
	}

	s.client, err = bot.New(s.config.Token, opts...)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Service) OnStart(ctx context.Context) error {
	if s.config.NoInit {
		return nil
	}

	s.setupCommands(ctx)

	ok, err := s.client.SetWebhook(ctx, &bot.SetWebhookParams{
		URL:         s.config.HookURL,
		SecretToken: "",
	})
	if err != nil {
		return errors.WithStack(err)
	}

	if !ok {
		return errors.Wrap(xmodels.ErrFailed, "start webhook")
	}

	s.runCtx, s.cancel = context.WithCancel(context.Background())

	//nolint:contextcheck
	go s.client.StartWebhook(s.runCtx)

	return nil
}

func (s *Service) OnStop(context.Context) error {
	if s.cancel != nil {
		s.cancel()
	}

	return nil
}

func (s *Service) WebhookHandler() http.HandlerFunc {
	return s.client.WebhookHandler()
}

func (s *Service) ErrorHandler(err error) {
	s.logger.Error("tg error", slog.Any("error", err))
}

func (s *Service) defaultHandler(
	_ context.Context,
	_ *bot.Bot,
	_ *models.Update) {
}
