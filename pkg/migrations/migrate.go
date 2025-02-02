package migrations

import (
	"context"

	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Migrate(
	ctx context.Context,
	dbOrig *gorm.DB,
) error {
	db := dbOrig.WithContext(ctx)

	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
