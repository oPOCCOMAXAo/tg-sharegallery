package models

import "github.com/go-telegram/bot/models"

type CommandScope string

const (
	CSDefault               CommandScope = "default"
	CSAllPrivateChats       CommandScope = "all_private_chats"
	CSAllGroupChats         CommandScope = "all_group_chats"
	CSAllChatAdministrators CommandScope = "all_chat_administrators"
)

func (s CommandScope) BotCommandScope() models.BotCommandScope {
	switch s {
	case CSDefault:
		return &models.BotCommandScopeDefault{}
	case CSAllPrivateChats:
		return &models.BotCommandScopeAllPrivateChats{}
	case CSAllGroupChats:
		return &models.BotCommandScopeAllGroupChats{}
	case CSAllChatAdministrators:
		return &models.BotCommandScopeAllChatAdministrators{}
	}

	return nil
}
