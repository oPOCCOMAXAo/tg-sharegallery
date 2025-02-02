package menu

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) Start(
	ctx context.Context,
	router *bot.Bot,
	update *models.Update,
) {
	_, err := s.domain.GetCreateUserByTgID(ctx, update.Message.From.ID)
	if err != nil {
		s.logger.ErrorContext(ctx, "Start",
			slog.Any("error", err),
		)

		return
	}

	menu, _, err := s.getMenuPageResult(ctx, ViewParams{
		Page:   views.MenuPageHelp,
		ChatID: update.Message.Chat.ID,
	})
	if err != nil {
		s.logger.ErrorContext(ctx, "Start",
			slog.Any("error", err),
		)

		return
	}

	if menu != nil {
		_, _ = router.SendMessage(ctx, menu.SendMessageParams())
	}
}
