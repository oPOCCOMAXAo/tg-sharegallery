package albums

import (
	"context"

	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) MenuAlbums(ctx *router.Context) {
	update := ctx.Update()

	view := views.MenuAlbums{
		UserID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: int64(update.CallbackQuery.Message.Message.ID),
	}

	err := s.fillMenuAlbumsView(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.EditMessageText(view.EditMessageTextParams()))
}

func (s *Service) fillMenuAlbumsView(
	ctx context.Context,
	view *views.MenuAlbums,
) error {
	stats, err := s.domain.GetUserAlbumStats(ctx, view.UserID)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	view.HasAlbums = stats.AlbumsCount > 0
	view.HasUnsaved = stats.HasUnsaved
	view.EditAlbumID = stats.CurrentAlbumID

	return nil
}
