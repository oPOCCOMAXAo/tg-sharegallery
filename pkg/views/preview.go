package views

import (
	"strconv"

	"github.com/go-telegram/bot"
	bmodels "github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-instrumentation/query"
	"github.com/opoccomaxao/tg-sharegallery/pkg/models"
)

type Preview struct {
	UserID      int64
	MessageID   int64
	AlbumID     int64
	CurrentPage int64
	HasNextPage bool
	HasPrevPage bool
	Image       *models.AlbumImage
}

func (v *Preview) Text() string {
	return "Preview album, page " + strconv.FormatInt(v.CurrentPage, 10)
}

func (v *Preview) ReplyMarkup() bmodels.ReplyMarkup {
	var res bmodels.InlineKeyboardMarkup

	{
		row := []bmodels.InlineKeyboardButton{}

		if v.HasPrevPage {
			row = append(row, bmodels.InlineKeyboardButton{
				Text: "<< Prev",
				CallbackData: query.Command("preview_album").
					WithParamInt64("id", v.AlbumID).
					WithParamInt64("page", v.CurrentPage-1).
					Encode(),
			})
		}

		if v.HasNextPage {
			row = append(row, bmodels.InlineKeyboardButton{
				Text: "Next >>",
				CallbackData: query.Command("preview_album").
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

func (v *Preview) EditMessageMediaParams() *bot.EditMessageMediaParams {
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

func (v *Preview) SendPhotoParams() *bot.SendPhotoParams {
	return &bot.SendPhotoParams{
		ChatID: v.UserID,
		Photo: &bmodels.InputFileString{
			Data: v.Image.TgFile,
		},
		Caption:     v.Text(),
		ParseMode:   bmodels.ParseModeHTML,
		ReplyMarkup: v.ReplyMarkup(),
	}
}
