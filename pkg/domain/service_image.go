package domain

import (
	"context"

	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
)

type ImagePageParams struct {
	AlbumID  int64
	UserTgID int64
	Offset   int64
}

type ImagePageResult struct {
	Image *models.AlbumImage
	Total int64
}

func (s *Service) GetImagePage(
	ctx context.Context,
	params ImagePageParams,
) (*ImagePageResult, error) {
	album, err := s.GetAlbumForUserByTgID(ctx, GetAlbumParams{
		UserTgID: params.UserTgID,
		AlbumID:  params.AlbumID,
	})
	if err != nil {
		return nil, err
	}

	var res ImagePageResult

	res.Image, err = s.repo.GetImagePage(ctx, album.ID, params.Offset)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	res.Total = album.ImagesCount

	return &res, nil
}
