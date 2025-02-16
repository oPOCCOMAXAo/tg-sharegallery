package repo

import (
	"context"

	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (r *Repo) CreateAlbumImageByUserTgID(
	ctx context.Context,
	userTgID int64,
	imageFileID string,
) error {
	err := r.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var user models.User

			err := tx.
				Where("tg_id = ?", userTgID).
				First(&user).
				Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.WithStack(err)
			}

			if user.CurrentAlbumID == nil {
				return errors.WithStack(models.ErrNotFound)
			}

			image := models.AlbumImage{
				AlbumID:   *user.CurrentAlbumID,
				CreatedAt: r.Now().Unix(),
				TgFile:    imageFileID,
			}

			err = tx.
				Create(&image).
				Error
			if err != nil {
				return errors.WithStack(err)
			}

			return nil
		})
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	return nil
}
