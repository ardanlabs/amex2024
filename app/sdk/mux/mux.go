package mux

import (
	"context"

	"github.com/ardanlabs/service/app/domain/testapp"
	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func WebAPI(log *logger.Logger) *web.App {
	logger := func(ctx context.Context, msg string, args ...any) {
		log.Info(ctx, msg, args...)
	}

	app := web.NewApp(logger,
		mid.Logger(log),
		mid.Error(log),
		mid.Panics(),
	)

	testapp.Routes(app)

	return app
}
