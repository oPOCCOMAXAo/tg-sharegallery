package domain

import "github.com/opoccomaxao/tg-sharegallery/pkg/repo"

type Service struct {
	repo *repo.Repo
}

func New(
	repo *repo.Repo,
) *Service {
	return &Service{
		repo: repo,
	}
}
