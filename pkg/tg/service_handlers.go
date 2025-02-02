package tg

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/internal"
)

func (s *Service) telemetry(
	typ internal.HandlerType,
	pattern string,
) bot.Middleware {
	return func(handler bot.HandlerFunc) bot.HandlerFunc {
		return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
			var (
				userName    string
				userID      int64
				messageText string
			)

			switch {
			case update.Message != nil:
				userName = update.Message.From.Username
				userID = update.Message.From.ID
				messageText = update.Message.Text
			case update.CallbackQuery != nil:
				userName = update.CallbackQuery.From.Username
				userID = update.CallbackQuery.From.ID
				messageText = update.CallbackQuery.Data
			case update.InlineQuery != nil:
				userName = update.InlineQuery.From.Username
				userID = update.InlineQuery.From.ID
				messageText = update.InlineQuery.Query
			}

			s.logger.InfoContext(ctx, "request",
				slog.String("type", typ.String()),
				slog.String("pattern", pattern),
				slog.Int64("user_id", userID),
				slog.String("user_name", userName),
				slog.String("message_text", messageText),
			)

			handler(ctx, bot, update)
		}
	}
}
