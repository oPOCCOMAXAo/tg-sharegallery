package domain

import (
	snowflake "github.com/opoccomaxao/go-snowflake"
	"github.com/opoccomaxao/tg-sharegallery/pkg/repo"
	"github.com/opoccomaxao/tg-sharegallery/pkg/tg"
)

type Service struct {
	repo      *repo.Repo
	tg        *tg.Service
	snowflake *snowflake.Generator
}

func New(
	repo *repo.Repo,
	tg *tg.Service,
	snowflake *snowflake.Generator,
) *Service {
	return &Service{
		repo:      repo,
		tg:        tg,
		snowflake: snowflake,
	}
}
