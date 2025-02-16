package domain

import (
	"context"
)

func (s *Service) UpdateAlbumTitleByUserTgID(
	ctx context.Context,
	userTgID int64,
	title string,
) error {
	//nolint:wrapcheck
	return s.repo.UpdateAlbumTitleByUserTgID(ctx, userTgID, title)
}

func (s *Service) AttachImageToAlbumByUserTgID(
	ctx context.Context,
	userTgID int64,
	imageFileID string,
) error {
	//nolint:wrapcheck
	return s.repo.CreateAlbumImageByUserTgID(ctx, userTgID, imageFileID)
}

func (s *Service) StartEditAlbum(
	ctx context.Context,
	userTgID int64,
	albumID int64,
) error {
	_, err := s.GetAlbumForUserByTgID(ctx, userTgID, albumID)
	if err != nil {
		return err
	}

	err = s.repo.UpdateCurrentAlbumForUserByTgID(ctx, userTgID, albumID)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}

func (s *Service) SaveAlbum(
	ctx context.Context,
	userTgID int64,
	albumID int64,
) error {
	err := s.repo.UpdateAlbumSavedByUserTgID(ctx, albumID, userTgID)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = s.repo.RemoveCurrentAlbumFromUserByTgID(ctx, userTgID, albumID)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
