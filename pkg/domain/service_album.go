package domain

import (
	"context"
	"time"

	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/repo"
	"github.com/pkg/errors"
)

type UserAlbumStats struct {
	AlbumsCount    int64
	HasUnsaved     bool
	CurrentAlbumID int64 // album currently being edited. 0 if no album is being edited.
}

func (s *Service) GetUserAlbumStats(
	ctx context.Context,
	userTgID int64,
) (*UserAlbumStats, error) {
	var (
		res UserAlbumStats
		err error
	)

	stats, err := s.repo.GetUserAlbumCount(ctx, userTgID)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	for _, stat := range stats {
		res.AlbumsCount += stat.Count

		if !stat.Saved {
			res.HasUnsaved = true
		}
	}

	res.CurrentAlbumID, err = s.repo.GetUserCurrentAlbumID(ctx, userTgID)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return &res, nil
}

func (s *Service) GetOrCreateNewAlbumForUserByTgID(
	ctx context.Context,
	userTgID int64,
) (*models.AlbumDomain, error) {
	user, err := s.repo.GetOrCreateUserByTgID(ctx, userTgID)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	album, err := s.repo.GetOrCreateNewAlbumForUser(ctx, user.ID)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	if album.Title == "" {
		album.Title = "Untitled " + time.Unix(album.CreatedAt, 0).Format(FormatAlbumTime)

		err = s.repo.UpdateAlbumTitle(ctx, album)
		if err != nil {
			//nolint:wrapcheck
			return nil, err
		}
	}

	domain, err := s.repo.GetFullAlbums(ctx, repo.FullAlbumsParams{
		AlbumIDs: []int64{album.ID},
		TgUserID: userTgID,
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	for _, ad := range domain {
		if ad.ID == album.ID {
			return ad, nil
		}
	}

	return nil, errors.WithStack(models.ErrFailed)
}

func (s *Service) GetAlbumForUser(
	ctx context.Context,
	userTgID int64,
	albumID int64,
) (*models.AlbumDomain, error) {
	res, err := s.repo.GetFullAlbums(ctx, repo.FullAlbumsParams{
		AlbumIDs: []int64{albumID},
		TgUserID: userTgID,
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	for _, ad := range res {
		if ad.ID == albumID {
			return ad, nil
		}
	}

	return nil, errors.WithStack(models.ErrNotFound)
}
