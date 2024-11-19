package web

import (
	"context"
)

type ctxKey int

const traceIDKey ctxKey = 1

func setTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(traceIDKey).(string)
	if !ok {
		return ""
	}

	return v
}
