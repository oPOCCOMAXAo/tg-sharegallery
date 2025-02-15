package menu

import (
	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) Help(ctx *router.Context) {
	update := ctx.Update()

	_, err := s.domain.GetCreateUserByTgID(ctx.Context(), update.Message.From.ID)
	if err != nil {
		ctx.Error(err)

		return
	}

	view := views.Menu{
		ChatID:    update.Message.Chat.ID,
		MessageID: int64(update.Message.ID),
		Page:      views.MenuPageHelp,
	}

	err = s.fillMenuView(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.SendMessage(view.SendMessageParams()))
}
