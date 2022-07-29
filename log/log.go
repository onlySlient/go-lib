package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	log = zerolog.New(os.Stdout).With().Timestamp().Logger()
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.DisableSampling(true)
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"
	zerolog.ErrorFieldName = "e"
	zerolog.CallerFieldName = "c"
	zerolog.TimeFieldFormat = time.RFC3339

	zerolog.DurationFieldUnit = time.Microsecond
	zerolog.DurationFieldInteger = false
	zerolog.ErrorHandler = func(err error) {
		// TODO emmm
	}
	zerolog.CallerSkipFrameCount = 3
}

func SetGlobalLevel(l zerolog.Level) {
	zerolog.SetGlobalLevel(l)
}

func EnableErrorCallerHook(enable bool) {
	if enable {
		log = log.Hook(ErrorHook{})
	}
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Infof(msg string, v ...interface{}) {
	log.Info().Msgf(msg, v...)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Debugf(msg string, v ...interface{}) {
	log.Debug().Msgf(msg, v...)
}

func Error(msg string) {
	log.Error().Msg(msg)
}

func Errorf(msg string, v ...interface{}) {
	log.Error().Msgf(msg, v...)
}

func Trace(msg string) {
	log.Trace().Msg(msg)
}

func Tracef(msg string, v ...interface{}) {
	log.Trace().Msgf(msg, v...)
}
