package menu

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/opoccomaxao/tg-sharegallery/pkg/texts"
	"github.com/opoccomaxao/tg-sharegallery/pkg/views"
)

type ViewParams struct {
	Page    views.MenuPage
	ChatID  int64
	QueryID string
}

func (s *Service) Menu(
	ctx context.Context,
	router *bot.Bot,
	update *models.Update,
) {
	req := ViewParams{
		ChatID:  update.CallbackQuery.Message.Message.From.ID,
		QueryID: update.CallbackQuery.ID,
	}

	params := texts.DecodeQuery(update.CallbackQuery.Data)
	params.GetStringInto("page", (*string)(&req.Page))

	menu, cb, err := s.getMenuPageResult(ctx, req)
	if err != nil {
		s.logger.ErrorContext(ctx, "Menu",
			slog.Any("error", err),
		)

		return
	}

	if menu != nil {
		_, _ = router.EditMessageText(ctx, menu.EditMessageTextParams())
	}

	if cb != nil {
		_, _ = router.AnswerCallbackQuery(ctx, cb.AnswerCallbackQueryParams())
	}
}

//nolint:revive,unparam // these params will be used in the future.
func (s *Service) getMenuPageResult(
	ctx context.Context,
	params ViewParams,
) (*views.Menu, *views.Callback, error) {
	var (
		menu *views.Menu
		cb   *views.Callback
	)

	switch params.Page {
	case views.MenuPageMain:
		menu = &views.Menu{
			Text:       "Bot description",
			ChatID:     params.ChatID,
			ShowHelp:   true,
			ShowAlbums: true,
		}
	case views.MenuPageHelp:
		menu = &views.Menu{
			Text:       "Help",
			ChatID:     params.ChatID,
			ShowMain:   true,
			ShowAlbums: true,
		}
	default:
		cb = &views.Callback{
			Text:    "Unknown page",
			QueryID: params.QueryID,
		}
	}

	return menu, cb, nil
}
