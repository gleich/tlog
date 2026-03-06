package tlog

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func log(fn func(string, ...any), o Op, msg string, v ...any) {
	fn(msg, append([]any{"operation", o}, v...)...)
}

func (o Op) Info(msg string, v ...any) {
	log(logger.Info, o, msg, v...)
}

func (o Op) Debug(msg string, v ...any) {
	log(logger.Debug, o, msg, v...)
}

func (o Op) Warn(msg string, v ...any) {
	log(logger.Warn, o, msg, v...)
}

func (o Op) Error(err error, msg string, v ...any) {
	log(logger.Error, o, msg, append([]any{"error", err}, v...)...)
}
