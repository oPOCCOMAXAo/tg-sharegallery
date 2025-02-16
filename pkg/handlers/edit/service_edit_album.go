package edit

import (
	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/domain"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) EditAlbum(ctx *router.Context) {
	update := ctx.Update()

	view := views.MenuAlbum{
		UserID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: int64(update.CallbackQuery.Message.Message.ID),
	}

	query := ctx.Query()
	query.GetInt64Into("id", &view.AlbumID)

	err := s.domain.StartEditAlbum(ctx.Context(), domain.StartEditAlbumParams{
		UserTgID: view.UserID,
		AlbumID:  view.AlbumID,
	})
	if err != nil {
		ctx.Error(err)

		return
	}

	err = s.views.FillMenuAlbum(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.EditMessageText(view.EditMessageTextParams()))
}
