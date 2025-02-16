package views

import (
	"strconv"

	"github.com/go-telegram/bot"
	bmodels "github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-instrumentation/query"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
)

type PublicView struct {
	UserID      int64
	MessageID   int64
	AlbumID     int64
	CurrentPage int64
	HasNextPage bool
	HasPrevPage bool
	Album       *models.Album
	Image       *models.AlbumImage
}

func (v *PublicView) Text() string {
	return "Album\n" + v.Album.Title +
		"\n\nPage " + strconv.FormatInt(v.CurrentPage, 10)
}

func (v *PublicView) ReplyMarkup() bmodels.ReplyMarkup {
	var res bmodels.InlineKeyboardMarkup

	{
		row := []bmodels.InlineKeyboardButton{}

		if v.HasPrevPage {
			row = append(row, bmodels.InlineKeyboardButton{
				Text: "<< Prev",
				CallbackData: query.Command("view_album").
					WithParamInt64("id", v.AlbumID).
					WithParamInt64("page", v.CurrentPage-1).
					Encode(),
			})
		}

		if v.HasNextPage {
			row = append(row, bmodels.InlineKeyboardButton{
				Text: "Next >>",
				CallbackData: query.Command("view_album").
					WithParamInt64("id", v.AlbumID).
					WithParamInt64("page", v.CurrentPage+1).
					Encode(),
			})
		}

		if len(row) > 0 {
			res.InlineKeyboard = append(res.InlineKeyboard, row)
		}
	}

	res.InlineKeyboard = append(res.InlineKeyboard, []bmodels.InlineKeyboardButton{
		{
			Text:         "Close",
			CallbackData: "delete",
		},
	})

	return res
}

func (v *PublicView) EditMessageMediaParams() *bot.EditMessageMediaParams {
	return &bot.EditMessageMediaParams{
		ChatID:    v.UserID,
		MessageID: int(v.MessageID),
		Media: &bmodels.InputMediaPhoto{
			Media:     v.Image.TgFile,
			Caption:   v.Text(),
			ParseMode: bmodels.ParseModeHTML,
		},
		ReplyMarkup: v.ReplyMarkup(),
	}
}

func (v *PublicView) SendPhotoParams() *bot.SendPhotoParams {
	return &bot.SendPhotoParams{
		ChatID: v.UserID,
		Photo: &bmodels.InputFileString{
			Data: v.Image.TgFile,
		},
		Caption:        v.Text(),
		ParseMode:      bmodels.ParseModeHTML,
		ReplyMarkup:    v.ReplyMarkup(),
		ProtectContent: true,
	}
}
