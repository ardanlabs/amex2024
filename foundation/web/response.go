package web

import (
	"context"
	"fmt"
	"net/http"
)

func Respond(ctx context.Context, w http.ResponseWriter, dataModel Encoder) error {
	data, ct, err := dataModel.Encode()
	if err != nil {
		return fmt.Errorf("respond: encode: %w", err)
	}

	w.Header().Set("Content-Type", ct)
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("respond: write: %w", err)
	}

	return nil
}
