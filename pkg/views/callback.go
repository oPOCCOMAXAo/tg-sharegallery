package views

import "github.com/go-telegram/bot"

type Callback struct {
	Text    string
	QueryID string
}

//nolint:mnd
func (c *Callback) AnswerCallbackQueryParams() *bot.AnswerCallbackQueryParams {
	return &bot.AnswerCallbackQueryParams{
		CallbackQueryID: c.QueryID,
		Text:            c.Text,
		ShowAlert:       true,
		CacheTime:       30,
	}
}
