package mid

import (
	"context"
	"net/http"
	"runtime/debug"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

func bill() (err error) {
	return
}

func Panics() web.MidFunc {
	m := func(next web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, r *http.Request) (resp web.Encoder) {
			defer func() {
				if rec := recover(); rec != nil {
					trace := debug.Stack()
					resp = errs.Newf(errs.InternalOnlyLog, "PANIC [%v] TRACE[%s]", rec, string(trace))
				}
			}()

			return next(ctx, r)
		}

		return h
	}

	return m
}
