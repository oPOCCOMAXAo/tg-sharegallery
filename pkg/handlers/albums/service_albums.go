package albums

import (
	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) Albums(ctx *router.Context) {
	update := ctx.Update()

	view := views.MenuAlbums{
		UserID:    update.Message.Chat.ID,
		MessageID: int64(update.Message.ID),
	}

	err := s.fillMenuAlbumsView(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.SendMessage(view.SendMessageParams()))
}
