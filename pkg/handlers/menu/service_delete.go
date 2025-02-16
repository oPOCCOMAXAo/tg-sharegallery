package menu

import (
	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-instrumentation/router"
)

func (s *Service) Delete(ctx *router.Context) {
	update := ctx.Update()

	ctx.LogError2(s.client.DeleteMessage(ctx.Context(), &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
	}))
}
