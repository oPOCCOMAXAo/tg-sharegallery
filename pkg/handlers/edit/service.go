package edit

import (
	"github.com/opoccomaxao/tg-sharegallery/pkg/domain"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

type Service struct {
	domain *domain.Service
	views  *views.Service
}

func NewService(
	domain *domain.Service,
	views *views.Service,
) *Service {
	return &Service{
		domain: domain,
		views:  views,
	}
}
