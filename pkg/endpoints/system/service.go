package system

import (
	"context"
)

type Service struct {
	cancel context.CancelCauseFunc
}

func NewService(
	cancel context.CancelCauseFunc,
) *Service {
	return &Service{
		cancel: cancel,
	}
}
