package logger

import (
	"context"
	"log/slog"
	"os"
)

type CtxKey struct{}

func CtxWithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	if logger == nil {
		return ctx
	}

	ctxLog, ok := ctx.Value(CtxKey{}).(*slog.Logger)
	if ok && ctxLog == logger {
		return ctx
	}

	return context.WithValue(ctx, CtxKey{}, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(CtxKey{}).(*slog.Logger)
	if ok {
		return logger
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
}
