package views

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-instrumentation/query"
)

type MenuPage string

const (
	MenuPageMain MenuPage = "main"
	MenuPageHelp MenuPage = "help"
)

type Menu struct {
	Page      MenuPage
	ChatID    int64
	MessageID int64
}

func (m *Menu) Text() string {
	switch m.Page {
	case MenuPageMain:
		return "This bot helps you to share your images with others by link."
	default:
		return "Page not found."
	}
}

func (m *Menu) ReplyMarkup() models.ReplyMarkup {
	var res models.InlineKeyboardMarkup

	if m.Page != MenuPageHelp {
		res.InlineKeyboard = append(res.InlineKeyboard, []models.InlineKeyboardButton{{
			Text: "Help",
			CallbackData: query.Command("menu").
				WithParam("page", string(MenuPageHelp)).
				Encode(),
		}})
	}

	if m.Page != MenuPageMain {
		res.InlineKeyboard = append(res.InlineKeyboard, []models.InlineKeyboardButton{{
			Text: "Main",
			CallbackData: query.Command("menu").
				WithParam("page", string(MenuPageMain)).
				Encode(),
		}})
	}

	res.InlineKeyboard = append(res.InlineKeyboard, []models.InlineKeyboardButton{
		{
			Text:         "Albums >>",
			CallbackData: "albums",
		},
	})

	if len(res.InlineKeyboard) == 0 {
		return nil
	}

	return res
}

func (m *Menu) SendMessageParams() *bot.SendMessageParams {
	return &bot.SendMessageParams{
		ChatID:      m.ChatID,
		Text:        m.Text(),
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: m.ReplyMarkup(),
	}
}

func (m *Menu) EditMessageTextParams() *bot.EditMessageTextParams {
	return &bot.EditMessageTextParams{
		ChatID:      m.ChatID,
		MessageID:   int(m.MessageID),
		Text:        m.Text(),
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: m.ReplyMarkup(),
	}
}
