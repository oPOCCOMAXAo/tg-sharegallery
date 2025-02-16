package albums

import (
	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) AlbumsMessage(ctx *router.Context) {
	update := ctx.Update()

	view := views.MenuAlbums{
		UserID:    update.Message.Chat.ID,
		MessageID: int64(update.Message.ID),
	}

	err := s.views.FillMenuAlbums(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.SendMessage(view.SendMessageParams()))
}

func (s *Service) Albums(ctx *router.Context) {
	update := ctx.Update()

	view := views.MenuAlbums{
		UserID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: int64(update.CallbackQuery.Message.Message.ID),
	}

	err := s.views.FillMenuAlbums(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.EditMessageText(view.EditMessageTextParams()))
}
