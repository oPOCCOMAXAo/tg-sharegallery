package texts

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *Service) Start(
	ctx context.Context,
	_ *bot.Bot,
	update *models.Update,
) {
	_, err := s.domain.GetCreateUserByTgID(ctx, update.Message.From.ID)
	if err != nil {
		s.logger.ErrorContext(ctx, "GetCreateUserByTgID",
			slog.Any("tg_id", update.Message.From.ID),
			slog.Any("error", err),
		)

		return
	}
}
