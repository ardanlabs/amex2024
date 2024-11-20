package testapp

import (
	"context"
	"math/rand/v2"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func test(ctx context.Context, r *http.Request) web.Encoder {
	if n := rand.IntN(100); n%2 == 0 {
		//return errs.Newf(errs.InvalidArgument, "you sent me bad stuff")
		panic("THIS IS A PANIC")
	}

	resp := status{
		Status: "OK",
	}

	return resp
}
