package views

import (
	"strings"

	"github.com/go-telegram/bot"
	bmodels "github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-instrumentation/query"
	"github.com/samber/lo"
)

type MenuAlbums struct {
	UserID      int64
	MessageID   int64
	EditAlbumID int64
	HasAlbums   bool
	HasUnsaved  bool
}

func (m *MenuAlbums) Text() string {
	lines := []string{
		"Albums",
	}

	if m.HasUnsaved {
		lines = append(lines, "\nYou have unsaved album")
	}

	if !m.HasAlbums {
		lines = append(lines, "\nYou don't have any albums yet")
	}

	return strings.Join(lines, "\n")
}

func (m *MenuAlbums) ReplyMarkup() bmodels.ReplyMarkup {
	var res bmodels.InlineKeyboardMarkup

	res.InlineKeyboard = append(res.InlineKeyboard, []bmodels.InlineKeyboardButton{
		{
			Text:         lo.Ternary(m.HasUnsaved, "View unsaved album", "Create new album"),
			CallbackData: "new_album",
		},
	})

	if m.EditAlbumID != 0 {
		res.InlineKeyboard = append(res.InlineKeyboard, []bmodels.InlineKeyboardButton{
			{
				Text: "View opened album",
				CallbackData: query.Command("edit_album").
					WithParamInt64("id", m.EditAlbumID).
					Encode(),
			},
		})
	}

	if m.HasAlbums {
		res.InlineKeyboard = append(res.InlineKeyboard, []bmodels.InlineKeyboardButton{
			{
				Text:         "View my albums",
				CallbackData: "list_albums",
			},
		})
	}

	res.InlineKeyboard = append(res.InlineKeyboard, []bmodels.InlineKeyboardButton{
		{
			Text:         "Menu",
			CallbackData: "menu",
		},
	})

	return res
}

func (m *MenuAlbums) SendMessageParams() *bot.SendMessageParams {
	return &bot.SendMessageParams{
		ChatID:      m.UserID,
		Text:        m.Text(),
		ParseMode:   bmodels.ParseModeHTML,
		ReplyMarkup: m.ReplyMarkup(),
	}
}

func (m *MenuAlbums) EditMessageTextParams() *bot.EditMessageTextParams {
	return &bot.EditMessageTextParams{
		ChatID:      m.UserID,
		MessageID:   int(m.MessageID),
		Text:        m.Text(),
		ParseMode:   bmodels.ParseModeHTML,
		ReplyMarkup: m.ReplyMarkup(),
	}
}
