package menu

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *Service) Help(
	ctx context.Context,
	router *bot.Bot,
	update *models.Update,
) {
	_, err := s.domain.GetCreateUserByTgID(ctx, update.Message.From.ID)
	if err != nil {
		s.logger.ErrorContext(ctx, "Help",
			slog.Any("error", err),
		)

		return
	}

	menu, err := s.getPageResult(ctx, MenuParams{
		Page: PageHelp,
	})
	if err != nil {
		s.logger.ErrorContext(ctx, "Help",
			slog.Any("error", err),
		)

		return
	}

	_, _ = router.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        menu.Text,
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: menu.ReplyMarkup,
	})
}
