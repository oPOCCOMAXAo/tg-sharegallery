package menu

import (
	"log/slog"

	"github.com/opoccomaxao/tg-sharegallery/pkg/domain"
)

type Service struct {
	logger *slog.Logger
	domain *domain.Service
}

func NewService(
	logger *slog.Logger,
	domain *domain.Service,
) *Service {
	return &Service{
		logger: logger,
		domain: domain,
	}
}
