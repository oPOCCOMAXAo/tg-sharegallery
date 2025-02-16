package views

import (
	"context"
	"errors"

	"github.com/opoccomaxao/tg-sharegallery/pkg/domain"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
)

type Service struct {
	domain *domain.Service
}

func NewService(
	domain *domain.Service,
) *Service {
	return &Service{
		domain: domain,
	}
}

func (s *Service) FillMenu(
	_ context.Context,
	view *Menu,
) error {
	if view.Page == "" {
		view.Page = MenuPageMain
	}

	return nil
}

func (s *Service) FillMenuAlbums(
	ctx context.Context,
	view *MenuAlbums,
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

func (s *Service) FillMenuAlbum(
	ctx context.Context,
	view *MenuAlbum,
) error {
	var err error

	if view.AlbumID == 0 {
		view.Album, err = s.domain.GetCurrentAlbumForUserByTgID(ctx, view.UserID)
		if err != nil && !errors.Is(err, models.ErrNotFound) {
			//nolint:wrapcheck
			return err
		}
	} else {
		view.Album, err = s.domain.GetAlbumForUserByTgID(ctx, domain.GetAlbumParams{
			UserTgID: view.UserID,
			AlbumID:  view.AlbumID,
		})
		if err != nil && !errors.Is(err, models.ErrNotFound) {
			//nolint:wrapcheck
			return err
		}
	}

	if view.Album != nil {
		view.AlbumID = view.Album.ID
	}

	return nil
}

func (s *Service) FillMenuListAlbums(
	ctx context.Context,
	view *MenuListAlbums,
) error {
	list, err := s.domain.ListAlbums(ctx, domain.ListAlbumsParams{
		UserTgID: view.UserID,
		Limit:    AlbumsPerPage,
		Offset:   view.CurrentPage * AlbumsPerPage,
	})
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	view.Albums = list.Albums
	view.HasPrevPage = view.CurrentPage > 1
	view.HasNextPage = list.Total > (view.CurrentPage+1)*AlbumsPerPage

	return nil
}

func (s *Service) FillPreview(
	ctx context.Context,
	view *Preview,
) error {
	page, err := s.domain.GetImagePage(ctx, domain.ImagePageParams{
		AlbumID:  view.AlbumID,
		UserTgID: view.UserID,
		Offset:   view.CurrentPage,
	})
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	view.Image = page.Image
	view.HasPrevPage = view.CurrentPage > 0
	view.HasNextPage = page.Total > view.CurrentPage+1

	return nil
}
