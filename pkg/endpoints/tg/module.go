package tg

import (
	"net/http"

	"github.com/opoccomaxao/tg-sharegallery/pkg/tg"
	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("endpoints/tg",
		fx.Invoke(RegisterEndpoints),
	)
}

func RegisterEndpoints(
	router *http.ServeMux,
	service *tg.Service,
) error {
	router.HandleFunc("POST /webhook", service.WebhookHandler())

	return nil
}
