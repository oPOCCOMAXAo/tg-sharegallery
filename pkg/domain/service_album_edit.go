package domain

import "context"

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
