package testapp

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func test(ctx context.Context, r *http.Request) web.Encoder {
	resp := status{
		Status: "OK",
	}

	return resp
}
