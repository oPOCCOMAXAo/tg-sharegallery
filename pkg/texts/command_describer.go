package texts

import (
	"github.com/go-telegram/bot"
	bmodels "github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
)

type CommandDescription struct {
	Scope        models.CommandScope
	LanguageCode models.LanguageCode
	Description  string
}

type CommandDescriber struct {
	// map[scope]map[language]map[command]description
	data map[models.CommandScope]map[models.LanguageCode]map[string]string
}

func NewCommandDescriber() *CommandDescriber {
	return &CommandDescriber{}
}

func (s *CommandDescriber) AddCommandDescription(
	command string,
	description []CommandDescription,
) {
	for _, value := range description {
		s.addCommandDescriptionSingle(
			command,
			value.Description,
			value.Scope,
			value.LanguageCode,
		)
	}
}

func (s *CommandDescriber) addCommandDescriptionSingle(
	command string,
	description string,
	scope models.CommandScope,
	languageCode models.LanguageCode,
) {
	if description == "" {
		description = command
	}

	if s.data == nil {
		s.data = map[models.CommandScope]map[models.LanguageCode]map[string]string{}
	}

	if s.data[scope] == nil {
		s.data[scope] = map[models.LanguageCode]map[string]string{}
	}

	if s.data[scope][languageCode] == nil {
		s.data[scope][languageCode] = map[string]string{}
	}

	s.data[scope][languageCode][command] = description
}

func (s *CommandDescriber) ListCommandsParams() []*bot.SetMyCommandsParams {
	var res []*bot.SetMyCommandsParams

	for scope, next := range s.data {
		for lang, next := range next {
			data := make([]bmodels.BotCommand, 0, len(next))

			for command, description := range next {
				data = append(data, bmodels.BotCommand{
					Command:     command,
					Description: description,
				})
			}

			res = append(res, &bot.SetMyCommandsParams{
				Commands:     data,
				Scope:        scope.BotCommandScope(),
				LanguageCode: string(lang),
			})
		}
	}

	return res
}
