package middleware

import (
	"log/slog"
	"time"

	"github.com/opoccomaxao/tg-instrumentation/router"
)

type telemetry struct {
	logger *slog.Logger
}

func Telemetry(logger *slog.Logger) router.Handler {
	res := telemetry{
		logger: logger,
	}

	return res.Handle
}

func (h *telemetry) Handle(ctx *router.Context) {
	h.CaptureRequest(ctx)
	defer h.CaptureResponse(ctx)

	ctx.Next()
}

func (h *telemetry) CaptureRequest(ctx *router.Context) {
	update := ctx.Update()

	args := []any{
		slog.Int64("id", update.ID),
		slog.String("pattern", ctx.Pattern()),
	}

	{
		raw := ctx.RawDebug()
		if len(raw) > 0 {
			args = append(args, slog.String("raw", string(raw)))
		}
	}

	switch {
	case update.Message != nil:
		args = append(args,
			slog.String("update_type", "message"),
			slog.Int64("user_id", update.Message.From.ID),
			slog.String("user_name", update.Message.From.Username),
			slog.String("message_text", update.Message.Text),
			slog.Time("message_date", time.Unix(int64(update.Message.Date), 0)),
		)
	case update.CallbackQuery != nil:
		args = append(args,
			slog.String("update_type", "callback"),
			slog.Int64("user_id", update.CallbackQuery.From.ID),
			slog.String("user_name", update.CallbackQuery.From.Username),
			slog.String("data", update.CallbackQuery.Data),
		)
	case update.InlineQuery != nil:
		args = append(args,
			slog.String("update_type", "inline"),
			slog.Int64("user_id", update.InlineQuery.From.ID),
			slog.String("user_name", update.InlineQuery.From.Username),
			slog.String("query", update.InlineQuery.Query),
		)
	}

	h.logger.InfoContext(ctx.Context(), "request", args...)
}

func (h *telemetry) CaptureResponse(ctx *router.Context) {
	update := ctx.Update()

	args := []any{
		slog.Int64("id", update.ID),
		slog.String("pattern", ctx.Pattern()),
		slog.Bool("accepted", ctx.IsAccepted()),
		slog.Bool("aborted", ctx.IsAborted()),
	}

	errs := ctx.Errors()
	if len(errs) > 0 {
		for _, err := range errs {
			args = append(args, slog.Any("error", err))
		}

		h.logger.ErrorContext(ctx.Context(), "response", args...)
	} else {
		h.logger.InfoContext(ctx.Context(), "response", args...)
	}
}
