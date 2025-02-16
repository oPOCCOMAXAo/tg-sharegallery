package menu

import (
	"log/slog"

	"github.com/opoccomaxao/tg-sharegallery/pkg/domain"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

type Service struct {
	logger *slog.Logger
	domain *domain.Service
	views  *views.Service
}

func NewService(
	logger *slog.Logger,
	domain *domain.Service,
	views *views.Service,
) *Service {
	return &Service{
		logger: logger,
		domain: domain,
		views:  views,
	}
}
