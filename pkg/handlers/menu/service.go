package menu

import (
	"github.com/go-telegram/bot"
	"github.com/opoccomaxao/tg-sharegallery/pkg/domain"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

type Service struct {
	domain *domain.Service
	views  *views.Service
	client *bot.Bot
}

func NewService(
	domain *domain.Service,
	views *views.Service,
	client *bot.Bot,
) *Service {
	return &Service{
		domain: domain,
		views:  views,
		client: client,
	}
}
