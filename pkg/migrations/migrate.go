package migrations

import (
	"context"

	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	models.User
}

type Album struct {
	models.Album

	Owner *User `gorm:"foreignKey:OwnerID"`
}

type AlbumImage struct {
	models.AlbumImage

	Album *Album `gorm:"foreignKey:AlbumID"`
}

func Migrate(
	ctx context.Context,
	dbOrig *gorm.DB,
) error {
	db := dbOrig.WithContext(ctx)

	err := db.AutoMigrate(
		&User{},
		&Album{},
		&AlbumImage{},
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
