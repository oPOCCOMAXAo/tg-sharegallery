package domain

import (
	"context"

	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
)

func (s *Service) GetOrCreateUserByTgID(
	ctx context.Context,
	tgID int64,
) (*models.User, error) {
	//nolint:wrapcheck
	return s.repo.GetOrCreateUserByTgID(ctx, tgID)
}
