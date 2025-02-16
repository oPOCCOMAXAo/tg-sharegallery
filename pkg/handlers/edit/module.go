package edit

import (
	"github.com/opoccomaxao/tg-instrumentation/apimodels"
	pkgrouter "github.com/opoccomaxao/tg-instrumentation/router"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg/middleware"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("handlers/edit",
		fx.Provide(fx.Private, NewService),
		fx.Invoke(RegisterHandlers),
	)
}

func RegisterHandlers(
	service *Service,
	router *pkgrouter.Router,
) error {
	router.
		Text("/newalbum",
			middleware.RequiredPrivateChat,
			service.NewAlbumMessage,
		).
		WithDescription(apimodels.LCAll, apimodels.CSAllPrivateChats, "Create new album").
		WithDescription(apimodels.LCUk, apimodels.CSAllPrivateChats, "Створити новий альбом")

	router.Callback("new_album",
		pkgrouter.AutoAnswerCallbackQuery(),
		middleware.RequiredCallbackMessage,
		service.NewAlbum,
	)

	router.Callback("edit_album",
		pkgrouter.AutoAnswerCallbackQuery(),
		middleware.RequiredCallbackMessage,
		service.EditAlbum,
	)

	router.Callback("save_album",
		pkgrouter.AutoAnswerCallbackQuery(),
		middleware.RequiredCallbackMessage,
		service.SaveAlbum,
	)

	router.Custom(service.editTitleMatcher, service.EditTitleMessage)

	router.Custom(service.attachImageMatcher, service.AttachImageMessage)

	return nil
}
