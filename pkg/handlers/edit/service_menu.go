package edit

import (
	"context"

	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

func (s *Service) fillMenuAlbumView(
	ctx context.Context,
	view *views.MenuAlbum,
) error {
	var err error

	if view.AlbumID == 0 {
		view.Album, err = s.domain.GetOrCreateNewAlbumForUserByTgID(ctx, view.UserID)
		if err != nil {
			//nolint:wrapcheck
			return err
		}
	} else {
		view.Album, err = s.domain.GetAlbumForUser(ctx, view.AlbumID, view.UserID)
		if err != nil {
			//nolint:wrapcheck
			return err
		}

		if view.Album != nil {
			view.AlbumID = view.Album.ID
		}
	}

	return nil
}
