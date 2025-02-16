package menu

import (
	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) Start(ctx *router.Context) {
	update := ctx.Update()

	_, err := s.domain.GetOrCreateUserByTgID(ctx.Context(), update.Message.From.ID)
	if err != nil {
		ctx.Error(err)

		return
	}

	view := views.Menu{
		UserID:    update.Message.From.ID,
		MessageID: int64(update.Message.ID),
		Page:      views.MenuPageMain,
	}

	err = s.views.FillMenu(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.SendMessage(view.SendMessageParams()))
}
