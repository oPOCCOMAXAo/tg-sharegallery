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
