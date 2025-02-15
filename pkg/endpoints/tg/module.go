package tg

import (
	"net/http"

	"github.com/opoccomaxao/tg-instrumentation/router"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("endpoints/tg",
		fx.Invoke(RegisterEndpoints),
	)
}

func RegisterEndpoints(
	router *http.ServeMux,
	tgRouter *router.Router,
) error {
	router.HandleFunc("POST /webhook", tgRouter.HandlerFunc)

	return nil
}
