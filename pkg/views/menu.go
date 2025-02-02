package views

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/texts"
)

type MenuPage string

const (
	MenuPageMain MenuPage = "main"
	MenuPageHelp MenuPage = "help"
)

type Menu struct {
	Text       string
	ChatID     int64
	ShowHelp   bool
	ShowMain   bool
	ShowAlbums bool
}

func (m *Menu) ReplyMarkup() models.ReplyMarkup {
	if !m.ShowHelp && !m.ShowMain && !m.ShowAlbums {
		return nil
	}

	var res models.InlineKeyboardMarkup

	if m.ShowHelp {
		res.InlineKeyboard = append(res.InlineKeyboard, []models.InlineKeyboardButton{{
			Text: "Help",
			CallbackData: texts.QueryCommand("menu").
				AddParam("page", string(MenuPageHelp)).
				Encode(),
		}})
	}

	if m.ShowMain {
		res.InlineKeyboard = append(res.InlineKeyboard, []models.InlineKeyboardButton{{
			Text: "Main",
			CallbackData: texts.QueryCommand("menu").
				AddParam("page", string(MenuPageMain)).
				Encode(),
		}})
	}

	if m.ShowAlbums {
		res.InlineKeyboard = append(res.InlineKeyboard, []models.InlineKeyboardButton{
			{
				Text:         "New album",
				CallbackData: "newalbum",
			},
			{
				Text:         "My albums",
				CallbackData: "myalbums",
			},
		})
	}

	if len(res.InlineKeyboard) == 0 {
		return nil
	}

	return res
}

func (m *Menu) SendMessageParams() *bot.SendMessageParams {
	return &bot.SendMessageParams{
		ChatID:      m.ChatID,
		Text:        m.Text,
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: m.ReplyMarkup(),
	}
}

func (m *Menu) EditMessageTextParams() *bot.EditMessageTextParams {
	return &bot.EditMessageTextParams{
		ChatID:      m.ChatID,
		Text:        m.Text,
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: m.ReplyMarkup(),
	}
}
