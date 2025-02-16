package views

import (
	"strconv"

	"github.com/go-telegram/bot"
	bmodels "github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-instrumentation/query"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
)

type MenuListAlbums struct {
	UserID      int64
	MessageID   int64
	Albums      []*models.Album
	CurrentPage int64
	HasNextPage bool
	HasPrevPage bool
}

func (m *MenuListAlbums) Text() string {
	return "Your albums, page " + strconv.FormatInt(m.CurrentPage, 10)
}

func (m *MenuListAlbums) ReplyMarkup() bmodels.ReplyMarkup {
	var res bmodels.InlineKeyboardMarkup

	for _, album := range m.Albums {
		res.InlineKeyboard = append(res.InlineKeyboard, []bmodels.InlineKeyboardButton{
			{
				Text: album.Title,
				CallbackData: query.Command("edit_album").
					WithParamInt64("id", album.ID).
					Encode(),
			},
		})
	}

	if m.HasPrevPage || m.HasNextPage {
		row := []bmodels.InlineKeyboardButton{}

		if m.HasPrevPage {
			row = append(row, bmodels.InlineKeyboardButton{
				Text: "<< Prev",
				CallbackData: query.Command("list_albums").
					WithParamInt64("page", m.CurrentPage-1).
					Encode(),
			})
		}

		if m.HasNextPage {
			row = append(row, bmodels.InlineKeyboardButton{
				Text: "Next >>",
				CallbackData: query.Command("list_albums").
					WithParamInt64("page", m.CurrentPage+1).
					Encode(),
			})
		}

		res.InlineKeyboard = append(res.InlineKeyboard, row)
	}

	res.InlineKeyboard = append(res.InlineKeyboard, []bmodels.InlineKeyboardButton{
		{
			Text: "<< Back",
			CallbackData: query.Command("albums").
				Encode(),
		},
	})

	return res
}

func (m *MenuListAlbums) SendMessageParams() *bot.SendMessageParams {
	return &bot.SendMessageParams{
		ChatID:      m.UserID,
		Text:        m.Text(),
		ParseMode:   bmodels.ParseModeHTML,
		ReplyMarkup: m.ReplyMarkup(),
	}
}

func (m *MenuListAlbums) EditMessageTextParams() *bot.EditMessageTextParams {
	return &bot.EditMessageTextParams{
		ChatID:      m.UserID,
		MessageID:   int(m.MessageID),
		Text:        m.Text(),
		ParseMode:   bmodels.ParseModeHTML,
		ReplyMarkup: m.ReplyMarkup(),
	}
}
