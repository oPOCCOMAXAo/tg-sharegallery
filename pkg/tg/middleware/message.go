package middleware

import (
	"github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-instrumentation/router"
)

func RequiredPrivateChat(ctx *router.Context) {
	update := ctx.Update()
	if update.Message.Chat.Type != models.ChatTypePrivate {
		ctx.Abort()
	}
}
