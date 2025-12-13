package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

const timestamp_format = "2006-01-02-15-04-05"

func InitLogger(level slog.Level) {
	Log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					return slog.String(
						slog.TimeKey,
						a.Value.Time().Format(timestamp_format),
					)
				}
				return a
			},
		}),
	)
	slog.SetDefault(Log)
}
