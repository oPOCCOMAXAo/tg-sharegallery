package edit

import (
	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) SaveAlbum(ctx *router.Context) {
	update := ctx.Update()
	query := ctx.Query()

	albumID, _ := query.GetInt64("id")
	if albumID == 0 {
		ctx.LogError2(ctx.RespondCallbackText("Error: Album not found"))
	} else {
		err := s.domain.SaveAlbum(
			ctx.Context(),
			update.CallbackQuery.From.ID,
			albumID,
		)
		if err != nil {
			ctx.LogError2(ctx.RespondCallbackText("Error: " + err.Error()))
			ctx.Error(err)

			return
		}
	}

	view := views.MenuAlbums{
		UserID:    update.CallbackQuery.From.ID,
		MessageID: int64(update.CallbackQuery.Message.Message.ID),
	}

	err := s.views.FillMenuAlbums(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.EditMessageText(view.EditMessageTextParams()))
}
