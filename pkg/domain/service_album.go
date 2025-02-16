package domain

import (
	"context"
	"time"

	"github.com/opoccomaxao/tg-instrumentation/query"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/repo"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (s *Service) GetFullAlbums(
	ctx context.Context,
	params repo.FullAlbumsParams,
) ([]*models.AlbumDomain, error) {
	albums, err := s.repo.GetFullAlbums(ctx, params)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	for _, album := range albums {
		if album.PublicID != nil {
			album.PublicLink, err = s.tg.CreateStartLink(query.Command("id").
				WithParamInt64("id", *album.PublicID).
				Encode())
			if err != nil {
				//nolint:wrapcheck
				return nil, err
			}
		}
	}

	return albums, nil
}

func (s *Service) GetAlbumByID(
	ctx context.Context,
	albumID int64,
) (*models.Album, error) {
	//nolint:wrapcheck
	return s.repo.GetAlbumByID(ctx, albumID)
}

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

func (s *Service) StartEditNewAlbumForUserByTgID(
	ctx context.Context,
	userTgID int64,
) (*models.Album, error) {
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

	err = s.repo.UpdateCurrentAlbumForUserByTgID(ctx, userTgID, album.ID)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return album, nil
}

type GetAlbumParams struct {
	UserTgID int64
	AlbumID  int64
}

func (s *Service) GetAlbumForUserByTgID(
	ctx context.Context,
	params GetAlbumParams,
) (*models.AlbumDomain, error) {
	res, err := s.GetFullAlbums(ctx, repo.FullAlbumsParams{
		AlbumIDs: []int64{params.AlbumID},
		UserTgID: params.UserTgID,
	})
	if err != nil {
		return nil, err
	}

	for _, ad := range res {
		if ad.ID == params.AlbumID {
			return ad, nil
		}
	}

	return nil, errors.WithStack(models.ErrNotFound)
}

func (s *Service) GetCurrentAlbumForUserByTgID(
	ctx context.Context,
	userTgID int64,
) (*models.AlbumDomain, error) {
	albumID, err := s.repo.GetUserCurrentAlbumID(ctx, userTgID)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	if albumID == 0 {
		return nil, errors.WithStack(models.ErrNotFound)
	}

	return s.GetAlbumForUserByTgID(ctx, GetAlbumParams{
		UserTgID: userTgID,
		AlbumID:  albumID,
	})
}

type ListAlbumsParams struct {
	UserTgID int64
	Limit    int64
	Offset   int64
}

type ListAlbumsResult struct {
	Albums []*models.AlbumDomain
	Total  int64
}

func (s *Service) ListAlbums(
	ctx context.Context,
	params ListAlbumsParams,
) (*ListAlbumsResult, error) {
	var res ListAlbumsResult

	list, err := s.repo.GetAlbumsList(ctx, repo.AlbumListParams{
		UserTgID: params.UserTgID,
		Limit:    params.Limit,
		Offset:   params.Offset,
	})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	albums, err := s.GetFullAlbums(ctx, repo.FullAlbumsParams{
		AlbumIDs: list.AlbumsIDs,
	})
	if err != nil {
		return nil, err
	}

	byID := lo.KeyBy(albums, func(a *models.AlbumDomain) int64 {
		return a.ID
	})

	res.Albums = lo.FilterMap(list.AlbumsIDs, func(id int64, _ int) (*models.AlbumDomain, bool) {
		album, ok := byID[id]

		return album, ok
	})

	res.Total = list.Total

	return &res, nil
}

func (s *Service) GetAlbumIDByPublicID(
	ctx context.Context,
	publicID int64,
) (int64, error) {
	albumID, err := s.repo.GetAlbumIDByPublicIDOrZero(ctx, publicID)
	if err != nil {
		//nolint:wrapcheck
		return 0, err
	}

	if albumID == 0 {
		return 0, errors.WithStack(models.ErrNotFound)
	}

	return albumID, nil
}
