package oplog

import (
	"time"
)

func logSince(fn func(string, ...any), opt Operation, msg string, start time.Time, v ...any) {
	log(fn, opt, msg, append([]any{"duration", formatDuration(time.Since(start))}, v...)...)
}

func (o Operation) InfoSince(msg string, start time.Time, v ...any) {
	logSince(logger.Info, o, msg, start, v...)
}

func (o Operation) DebugSince(msg string, start time.Time, v ...any) {
	logSince(logger.Debug, o, msg, start, v...)
}

func (o Operation) WarnSince(msg string, start time.Time, v ...any) {
	logSince(logger.Warn, o, msg, start, v...)
}

func (o Operation) ErrorSince(msg string, start time.Time, v ...any) {
	logSince(logger.Error, o, msg, start, v...)
}
