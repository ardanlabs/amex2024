package web

import (
	"context"
	"fmt"
	"net/http"
)

type httpStatus interface {
	HTTPStatus() int
}

func Respond(ctx context.Context, w http.ResponseWriter, dataModel Encoder) error {
	data, ct, err := dataModel.Encode()
	if err != nil {
		return fmt.Errorf("respond: encode: %w", err)
	}

	w.Header().Set("Content-Type", ct)

	statusCode := http.StatusOK
	if v, ok := dataModel.(httpStatus); ok {
		statusCode = v.HTTPStatus()
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("respond: write: %w", err)
	}

	return nil
}
