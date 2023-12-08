package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/italorfeitosa/transactions-service/pkg/requestid"
)

var availableLogLevels = []slog.Level{
	slog.LevelInfo,
	slog.LevelWarn,
	slog.LevelError,
	slog.LevelDebug,
}

func Init(level string) {
	slog.SetDefault(slog.New(newCustomHandler(level)))
}

func getLevel(level string) slog.Level {
	for _, v := range availableLogLevels {
		if v.String() == level {
			return v
		}
	}

	return slog.LevelInfo
}

type customHandler struct {
	slog.Handler
}

func newCustomHandler(level string) slog.Handler {
	return &customHandler{slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: getLevel(strings.ToUpper(level)),
	})}
}

func (h *customHandler) Handle(ctx context.Context, r slog.Record) error {
	if v := requestid.FromContext(ctx); v != "" {
		r.Add("request_id", v)
	}

	return h.Handler.Handle(ctx, r)
}
