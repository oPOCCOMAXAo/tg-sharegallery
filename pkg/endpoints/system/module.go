package system

import (
	"net/http"

	"go.uber.org/fx"
)

func Invoke() fx.Option {
	return fx.Module("endpoints/system",
		fx.Provide(
			fx.Private,
			NewService,
		),
		fx.Invoke(RegisterEndpoints),
	)
}

func RegisterEndpoints(
	router *http.ServeMux,
	service *Service,
) error {
	router.HandleFunc("GET /health", service.Health)
	router.HandleFunc("PUT /shutdown", service.Shutdown)

	return nil
}
