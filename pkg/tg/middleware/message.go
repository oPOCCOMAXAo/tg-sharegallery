package middleware

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// MustHaveCallback checks if the update has a callback query and calls the next handler if it does.
//
// This middleware is called by default in the RegisterTextHandler method.
func MustHaveMessage(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, router *bot.Bot, update *models.Update) {
		if update.Message != nil {
			next(ctx, router, update)

			return
		}
	}
}
