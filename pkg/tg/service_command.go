package tg

import (
	"context"

	"github.com/opoccomaxao/tg-sharegallery/pkg/texts"
)

func (s *Service) AddCommandDescription(
	ctx context.Context,
	command string,
	description []texts.CommandDescription,
) {
	s.describer.AddCommandDescription(command, description)
	s.setupCommands(ctx)
}

func (s *Service) setupCommands(ctx context.Context) {
	if s.client == nil {
		return
	}

	for _, value := range s.describer.ListCommandsParams() {
		_, _ = s.client.SetMyCommands(ctx, value)
	}
}
