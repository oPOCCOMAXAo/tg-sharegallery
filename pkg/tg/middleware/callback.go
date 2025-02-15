package middleware

import (
	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-instrumentation/router"
)

func RequiredCallbackMessage(ctx *router.Context) {
	update := ctx.Update()
	if update.CallbackQuery.Message.Message == nil {
		ctx.Abort()

		_, _ = ctx.AnswerCallbackQuery(&bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			Text:            "This command is not available for this message",
		})
	}
}
