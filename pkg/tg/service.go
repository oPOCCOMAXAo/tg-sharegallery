package tg

import (
	"context"
	"log"
	"net/http"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	xmodels "github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/pkg/errors"
)

type Service struct {
	config Config
	client *bot.Bot
	runCtx context.Context //nolint:containedctx
	cancel context.CancelFunc
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
) (*Service, error) {
	res := Service{
		config: config,
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
		bot.WithDefaultHandler(s.defaultHandler),
		bot.WithSkipGetMe(),
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

func (s *Service) defaultHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {
	if update.Message != nil {
		s.logMessage(ctx, update.Message)
	}

	if update.CallbackQuery != nil {
		s.logCallback(ctx, update.CallbackQuery)
	}
}

func (s *Service) logMessage(_ context.Context, msg *models.Message) {
	log.Printf("%d @%s: %s",
		msg.From.ID,
		msg.From.Username,
		msg.Text,
	)
}

func (s *Service) logCallback(_ context.Context, cb *models.CallbackQuery) {
	log.Printf("%d @%s: %s",
		cb.From.ID,
		cb.From.Username,
		cb.Data,
	)
}
