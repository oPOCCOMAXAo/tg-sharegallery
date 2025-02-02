package middleware

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// MustHaveCallback checks if the update has a callback query and calls the next handler if it does.
//
// This middleware is called by default in the RegisterCallbackHandler method.
func MustHaveCallback(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, router *bot.Bot, update *models.Update) {
		if update.CallbackQuery != nil {
			next(ctx, router, update)

			return
		}
	}
}

func MustHaveCallbackMessage(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, router *bot.Bot, update *models.Update) {
		if update.CallbackQuery.Message.Message != nil {
			next(ctx, router, update)

			return
		}

		_, _ = router.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			Text:            "This command is not available for this message",
		})
	}
}

func AutoFinishCallback(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, router *bot.Bot, update *models.Update) {
		next(ctx, router, update)

		if update.CallbackQuery != nil {
			_, _ = router.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
				CallbackQueryID: update.CallbackQuery.ID,
			})
		}
	}
}
