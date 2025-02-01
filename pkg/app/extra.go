package app

import (
	"context"
	"log"

	"go.uber.org/fx"
)

func NewCancelCause(
	shutdowner fx.Shutdowner,
) context.CancelCauseFunc {
	return func(err error) {
		exitCode := 0

		if err != nil {
			exitCode = 1

			log.Printf("%+v", err)
		}

		err = shutdowner.Shutdown(
			fx.ExitCode(exitCode),
		)
		if err != nil {
			log.Printf("%+v", err)
		}
	}
}
