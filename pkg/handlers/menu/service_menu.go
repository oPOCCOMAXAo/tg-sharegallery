package menu

import (
	"context"

	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

type ViewParams struct {
	Page      views.MenuPage
	ChatID    int64
	MessageID int64
	QueryID   string
}

func (s *Service) Menu(ctx *router.Context) {
	update := ctx.Update()

	view := views.Menu{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: int64(update.CallbackQuery.Message.Message.ID),
	}

	query := ctx.Query()
	query.GetInto("page", (*string)(&view.Page))

	err := s.fillMenuView(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.EditMessageText(view.EditMessageTextParams()))
}

func (s *Service) fillMenuView(
	_ context.Context,
	view *views.Menu,
) error {
	if view.Page == "" {
		view.Page = views.MenuPageMain
	}

	return nil
}
