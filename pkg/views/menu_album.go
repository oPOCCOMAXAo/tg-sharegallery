package views

import (
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	bmodels "github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-instrumentation/query"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
	"github.com/samber/lo"
)

type MenuAlbum struct {
	UserID    int64
	MessageID int64
	AlbumID   int64
	Album     *models.AlbumDomain
}

func (m *MenuAlbum) Text() string {
	lines := []string{}

	if m.Album == nil {
		lines = append(lines,
			"Album not found",
		)
	} else {
		lines = append(lines,
			"Album",
			"<b>Title</b>:\n"+m.Album.Title,
			"<b>Images</b>: "+strconv.FormatInt(m.Album.ImagesCount, 10),
		)

		if m.Album.PublicLink != "" {
			lines = append(lines,
				"<b>Public link</b>: "+m.Album.PublicLink,
				"<b>Click to copy</b>: <code>"+m.Album.PublicLink+"</code>",
			)
		} else {
			lines = append(lines,
				"<b>Public link</b>: not published. Use [Publish] button from the menu.",
			)
		}

		lines = append(lines,
			"\nUpload images to this album by sending them to this chat.",
			"\nSend text message to this chat to set the title of the album.",
			"\nOnce you have uploaded images, you can save the album and exit.",
			"\nYou can also publish the album to get a public link.",
		)
	}

	return strings.Join(lines, "\n")
}

func (m *MenuAlbum) ReplyMarkup() bmodels.ReplyMarkup {
	var res bmodels.InlineKeyboardMarkup

	res.InlineKeyboard = append(res.InlineKeyboard, []bmodels.InlineKeyboardButton{
		{
			Text: "Update page",
			CallbackData: query.Command("edit_album").
				WithParamInt64("id", m.AlbumID).
				Encode(),
		},
	})

	if m.Album != nil {
		row := []bmodels.InlineKeyboardButton{}

		canSave := m.Album.Saved || m.Album.ImagesCount > 0

		row = append(row, bmodels.InlineKeyboardButton{
			Text: lo.Ternary(canSave, "Save and exit", "Exit"),
			CallbackData: query.Command("save_album").
				WithParamInt64("id", m.Album.ID).
				Encode(),
		})

		if m.Album.PublicLink == "" && canSave {
			row = append(row, bmodels.InlineKeyboardButton{
				Text: "Publish",
				CallbackData: query.Command("publish_album").
					WithParamInt64("id", m.Album.ID).
					Encode(),
			})
		}

		if m.Album.ImagesCount > 0 {
			row = append(row, bmodels.InlineKeyboardButton{
				Text: "Preview",
				CallbackData: query.Command("preview_album").
					WithParamInt64("id", m.Album.ID).
					WithParamEmpty("new").
					Encode(),
			})
		}

		if len(row) > 0 {
			res.InlineKeyboard = append(res.InlineKeyboard, row)
		}
	}

	res.InlineKeyboard = append(res.InlineKeyboard, []bmodels.InlineKeyboardButton{
		{
			Text:         "Albums menu",
			CallbackData: "albums",
		},
	})

	return res
}

func (m *MenuAlbum) SendMessageParams() *bot.SendMessageParams {
	return &bot.SendMessageParams{
		ChatID:      m.UserID,
		Text:        m.Text(),
		ParseMode:   bmodels.ParseModeHTML,
		ReplyMarkup: m.ReplyMarkup(),
		ReplyParameters: &bmodels.ReplyParameters{
			MessageID: int(m.MessageID),
			ChatID:    m.UserID,
		},
	}
}

func (m *MenuAlbum) EditMessageTextParams() *bot.EditMessageTextParams {
	return &bot.EditMessageTextParams{
		ChatID:      m.UserID,
		MessageID:   int(m.MessageID),
		Text:        m.Text(),
		ParseMode:   bmodels.ParseModeHTML,
		ReplyMarkup: m.ReplyMarkup(),
	}
}

func (m *MenuAlbum) SendMessageNotFound() *bot.SendMessageParams {
	return &bot.SendMessageParams{
		ChatID:    m.UserID,
		Text:      "Album not found. Please, create a new album.",
		ParseMode: bmodels.ParseModeHTML,
		ReplyParameters: &bmodels.ReplyParameters{
			MessageID: int(m.MessageID),
			ChatID:    m.UserID,
		},
	}
}
