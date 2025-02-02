package menu

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/texts"
	"github.com/pkg/errors"
)

type MenuParams struct {
	Page Page
}

type MenuResults struct {
	Text             string
	CallbackResponse string
	ReplyMarkup      *models.InlineKeyboardMarkup
}

func (s *Service) Menu(
	ctx context.Context,
	router *bot.Bot,
	update *models.Update,
) {
	var req MenuParams

	params := texts.DecodeQuery(update.CallbackQuery.Data)
	params.GetStringInto("page", (*string)(&req.Page))

	res, err := s.getPageResult(ctx, req)
	if err != nil {
		s.logger.ErrorContext(ctx, "Menu",
			slog.Any("error", err),
		)

		return
	}

	_, err = router.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.Message.ID,
		Text:        res.Text,
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: res.ReplyMarkup,
	})
	if err != nil {
		s.logger.ErrorContext(ctx, "Menu",
			slog.Any("error", errors.WithStack(err)),
		)

		return
	}

	if res.CallbackResponse != "" {
		_, err = router.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			Text:            res.CallbackResponse,
			ShowAlert:       true,
			CacheTime:       300,
		})
		if err != nil {
			s.logger.ErrorContext(ctx, "Menu",
				slog.Any("error", errors.WithStack(err)),
			)
		}
	}
}

func (s *Service) getPageResult(
	ctx context.Context,
	params MenuParams,
) (*MenuResults, error) {
	var res MenuResults

	var btnHelp, btnMain, btnGallery bool

	var keyboard [][]models.InlineKeyboardButton

	switch params.Page {
	case PageMain:
		res.Text = "Bot description"
		btnHelp = true
		btnGallery = true
	case PageHelp:
		res.Text = "Help"
		btnMain = true
		btnGallery = true
	default:
		res.CallbackResponse = "Unknown page"
	}

	if btnHelp {
		keyboard = append(keyboard, []models.InlineKeyboardButton{{
			Text: "Help",
			CallbackData: texts.QueryCommand("menu").
				AddParam("page", string(PageHelp)).
				Encode(),
		}})
	}

	if btnMain {
		keyboard = append(keyboard, []models.InlineKeyboardButton{{
			Text: "Main",
			CallbackData: texts.QueryCommand("menu").
				AddParam("page", string(PageMain)).
				Encode(),
		}})
	}

	if btnGallery {
		keyboard = append(keyboard, []models.InlineKeyboardButton{
			{
				Text:         "New gallery",
				CallbackData: "newgallery",
			},
			{
				Text:         "My galleries",
				CallbackData: "mygalleries",
			},
		})
	}

	if len(keyboard) > 0 {
		res.ReplyMarkup = &models.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		}
	}

	return &res, nil
}
