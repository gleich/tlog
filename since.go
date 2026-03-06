package tlog

import (
	"time"
)

func logSince(fn func(string, ...any), opt Op, msg string, start time.Time, v ...any) {
	log(fn, opt, msg, append([]any{"duration", formatDuration(time.Since(start))}, v...)...)
}

func (o Op) InfoSince(msg string, start time.Time, v ...any) {
	logSince(logger.Info, o, msg, start, v...)
}

func (o Op) DebugSince(msg string, start time.Time, v ...any) {
	logSince(logger.Debug, o, msg, start, v...)
}

func (o Op) WarnSince(msg string, start time.Time, v ...any) {
	logSince(logger.Warn, o, msg, start, v...)
}

func (o Op) ErrorSince(msg string, start time.Time, v ...any) {
	logSince(logger.Error, o, msg, start, v...)
}
