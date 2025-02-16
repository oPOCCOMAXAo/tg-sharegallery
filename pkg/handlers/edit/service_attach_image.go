package edit

import (
	bmodels "github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-instrumentation/apimodels"
	"github.com/opoccomaxao/tg-instrumentation/router"
)

func (s *Service) attachImageMatcher(
	update *apimodels.Update,
) bool {
	return update.Message != nil &&
		update.Message.Chat.Type == bmodels.ChatTypePrivate &&
		update.Message.Photo != nil
}

func (s *Service) AttachImageMessage(ctx *router.Context) {
	update := ctx.Update()

	images := update.Message.Photo

	maxImage := images[0]
	for _, image := range images {
		if image.Width > maxImage.Width {
			maxImage = image
		}
	}

	err := s.domain.AttachImageToAlbumByUserTgID(
		ctx.Context(),
		update.Message.Chat.ID,
		maxImage.FileID,
	)
	if err != nil {
		ctx.Error(err)

		return
	}

	ctx.LogError2(ctx.RespondReactionEmoji("👍"))
}
