package edit

import (
	"errors"

	bmodels "github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-instrumentation/apimodels"
	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) editTitleMatcher(
	update *apimodels.Update,
) bool {
	return update.Message != nil &&
		update.Message.Chat.Type == bmodels.ChatTypePrivate &&
		update.Message.Text != ""
}

func (s *Service) EditTitleMessage(ctx *router.Context) {
	update := ctx.Update()

	view := views.MenuAlbum{
		UserID:    update.Message.Chat.ID,
		MessageID: int64(update.Message.ID),
		AlbumID:   0,
	}

	err := s.domain.UpdateAlbumTitleByUserTgID(
		ctx.Context(),
		view.UserID,
		update.Message.Text,
	)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			ctx.LogError2(ctx.SendMessage(view.SendMessageNotFound()))
		} else {
			ctx.Error(err)
		}

		return
	}

	err = s.fillMenuAlbumView(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.SendMessage(view.SendMessageParams()))
}
