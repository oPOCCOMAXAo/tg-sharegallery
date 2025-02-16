package preview

import (
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

type Service struct {
	views *views.Service
}

func NewService(
	views *views.Service,
) *Service {
	return &Service{
		views: views,
	}
}
