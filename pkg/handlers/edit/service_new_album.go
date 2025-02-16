package edit

import (
	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) NewAlbumMessage(ctx *router.Context) {
	update := ctx.Update()

	view := views.MenuAlbum{
		UserID:    update.Message.From.ID,
		MessageID: int64(update.Message.ID),
		AlbumID:   0,
	}

	err := s.fillMenuAlbumView(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.SendMessage(view.SendMessageParams()))
}

func (s *Service) NewAlbum(ctx *router.Context) {
	update := ctx.Update()

	view := views.MenuAlbum{
		UserID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: int64(update.CallbackQuery.Message.Message.ID),
		AlbumID:   0,
	}

	err := s.fillMenuAlbumView(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.EditMessageText(view.EditMessageTextParams()))
}
