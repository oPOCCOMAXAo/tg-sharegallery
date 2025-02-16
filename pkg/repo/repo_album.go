package repo

import (
	"context"

	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type AlbumStats struct {
	UserID int64 `gorm:"column:user_id"`
	Count  int64 `gorm:"column:count"`
	Saved  bool  `gorm:"column:saved"`
}

func (r *Repo) GetUserAlbumCount(
	ctx context.Context,
	userTgID int64,
) ([]*AlbumStats, error) {
	var res []*AlbumStats

	err := r.db.WithContext(ctx).
		Table("albums").
		Select(
			"owner_id AS user_id",
			"COUNT(*) AS count",
			"saved",
		).
		Where("owner_id IN (SELECT id FROM users WHERE tg_id = ?)", userTgID).
		Group("owner_id, saved").
		Find(&res).
		Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

func (r *Repo) GetUserCurrentAlbumID(
	ctx context.Context,
	userTgID int64,
) (int64, error) {
	var res models.User

	err := r.db.WithContext(ctx).
		Table("users").
		Select("current_album_id").
		Where("tg_id = ?", userTgID).
		First(&res).
		Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.WithStack(err)
	}

	return lo.FromPtr(res.CurrentAlbumID), nil
}

func (r *Repo) GetOrCreateNewAlbumForUser(
	ctx context.Context,
	userID int64,
) (*models.Album, error) {
	var res models.Album

	err := r.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var existing models.Album

			err := tx.
				Where("owner_id = ?", userID).
				Where("saved = ?", false).
				First(&existing).
				Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.WithStack(err)
			}

			if existing.ID != 0 {
				res = existing

				return nil
			}

			res.OwnerID = userID
			res.Saved = false
			res.CreatedAt = r.Now().Unix()
			res.UpdatedAt = res.CreatedAt

			err = tx.
				Create(&res).
				Error
			if err != nil {
				return errors.WithStack(err)
			}

			err = tx.
				Model(&models.User{}).
				Where("id = ?", userID).
				Update("current_album_id", res.ID).
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

func (r *Repo) UpdateAlbumTitle(
	ctx context.Context,
	album *models.Album,
) error {
	album.UpdatedAt = r.Now().Unix()

	err := r.db.WithContext(ctx).
		Model(&models.Album{}).
		Select("title", "updated_at").
		Where("id = ?", album.ID).
		Updates(album).
		Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repo) UpdateAlbumTitleByUserTgID(
	ctx context.Context,
	userTgID int64,
	title string,
) error {
	res := r.db.WithContext(ctx).
		Model(&models.Album{}).
		Select("title", "updated_at").
		Where("id IN (SELECT current_album_id FROM users WHERE tg_id = ?)", userTgID).
		Updates(&models.Album{
			Title:     title,
			UpdatedAt: r.Now().Unix(),
		})

	if res.Error != nil {
		return errors.WithStack(res.Error)
	}

	if res.RowsAffected == 0 {
		return errors.WithStack(models.ErrNotFound)
	}

	return nil
}

func (r *Repo) UpdateAlbumSavedByUserTgID(
	ctx context.Context,
	albumID int64,
	userTgID int64,
) error {
	err := r.db.WithContext(ctx).
		Model(&models.Album{}).
		Select("saved", "updated_at").
		Where("id = ?", albumID).
		Where("owner_id IN (SELECT id FROM users WHERE tg_id = ?)", userTgID).
		Where(`EXISTS (SELECT id FROM album_images WHERE album_id = ?)`, albumID).
		Updates(&models.Album{
			Saved:     true,
			UpdatedAt: r.Now().Unix(),
		}).
		Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

type FullAlbumsParams struct {
	AlbumIDs []int64
	TgUserID int64
}

func (r *Repo) GetFullAlbums(
	ctx context.Context,
	params FullAlbumsParams,
) ([]*models.AlbumDomain, error) {
	if len(params.AlbumIDs) == 0 {
		return nil, nil
	}

	var res []*models.AlbumDomain

	query := r.db.WithContext(ctx).
		Select(
			"a.id",
			"a.public_id",
			"a.created_at",
			"a.updated_at",
			"a.deleted_at",
			"a.owner_id",
			"a.title",
			"a.saved",
			"COUNT(i.id) AS images_count",
		).
		Table("albums AS a").
		Joins("LEFT JOIN album_images AS i ON i.album_id = a.id").
		Group("a.id").
		Where("a.id IN (?)", params.AlbumIDs)

	if params.TgUserID != 0 {
		query = query.
			Where("a.owner_id IN (SELECT id FROM users WHERE tg_id = ?)", params.TgUserID)
	}

	err := query.
		Find(&res).
		Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

type AlbumListParams struct {
	UserTgID int64
	Limit    int64
	Offset   int64
}

type AlbumListResult struct {
	AlbumsIDs []int64
	Total     int64
}

func (r *Repo) GetAlbumsList(
	ctx context.Context,
	params AlbumListParams,
) (*AlbumListResult, error) {
	var res AlbumListResult

	query := r.db.WithContext(ctx).
		Table("albums").
		Where("owner_id IN (SELECT id FROM users WHERE tg_id = ?)", params.UserTgID)

	err := query.
		Count(&res.Total).
		Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = query.
		Select("id").
		Order("id DESC").
		Limit(int(params.Limit)).
		Offset(int(params.Offset)).
		Find(&res.AlbumsIDs).
		Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &res, nil
}
