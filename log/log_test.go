package log

import (
	"testing"

	"github.com/rs/zerolog"
)

func TestInfo(t *testing.T) {
	Info("hehe")
}

func TestDebug(t *testing.T) {
	Info("hehe")
}

func TestError(t *testing.T) {
	Error("hehe")
}

func TestTrace(t *testing.T) {
	Trace("hehe")
}

func TestSetGlobalLevel(t *testing.T) {
	SetGlobalLevel(zerolog.ErrorLevel)
	Debug("hehe")
	Error("hehe")
}
