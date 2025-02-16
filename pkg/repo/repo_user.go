package repo

import (
	"context"

	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (r *Repo) GetOrCreateUserByTgID(
	ctx context.Context,
	tgID int64,
) (*models.User, error) {
	var res models.User

	err := r.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			err := tx.
				Where("tg_id = ?", tgID).
				Take(&res).
				Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.WithStack(err)
			}

			if res.ID != 0 {
				return nil
			}

			res.TgID = tgID
			res.CreatedAt = r.Now().Unix()
			res.UpdatedAt = res.CreatedAt

			err = tx.
				Create(&res).
				Error
			if err != nil {
				return errors.WithStack(err)
			}

			return nil
		})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return &res, nil
}

func (r *Repo) RemoveCurrentAlbumFromUserByTgID(
	ctx context.Context,
	userTgID int64,
	albumID int64,
) error {
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("tg_id = ?", userTgID).
		Where("current_album_id = ?", albumID).
		Select("current_album_id", "updated_at").
		Updates(&models.User{
			UpdatedAt:      r.Now().Unix(),
			CurrentAlbumID: nil,
		}).
		Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repo) UpdateCurrentAlbumForUserByTgID(
	ctx context.Context,
	userTgID int64,
	albumID int64,
) error {
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("tg_id = ?", userTgID).
		Select("current_album_id", "updated_at").
		Updates(&models.User{
			UpdatedAt:      r.Now().Unix(),
			CurrentAlbumID: &albumID,
		}).
		Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
