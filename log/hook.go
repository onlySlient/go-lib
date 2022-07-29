package log

import (
	"github.com/rs/zerolog"
)

type ErrorHook struct{}

func (h ErrorHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	if level == zerolog.ErrorLevel {
		e = e.Caller(3)
	}
}
