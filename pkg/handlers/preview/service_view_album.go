package preview

import (
	"github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) ViewAlbum(ctx *router.Context) {
	update := ctx.Update()

	view := views.PublicView{
		UserID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: int64(update.CallbackQuery.Message.Message.ID),
	}

	query := ctx.Query()
	query.GetInt64Into("id", &view.AlbumID)
	query.GetInt64Into("page", &view.CurrentPage)
	requiredNewMessage := query.Has("new")

	err := s.views.FillPublicView(ctx.Context(), &view)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.RespondCallbackText(""))

	if requiredNewMessage {
		ctx.LogError2(ctx.SendPhoto(view.SendPhotoParams()))
	} else {
		ctx.LogError2(ctx.EditMessageMedia(view.EditMessageMediaParams()))
	}
}
