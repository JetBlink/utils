package log

import (
	"testing"
)

func TestLogMessages(t *testing.T) {
	_ = New(true)
	defer Sync()
	Info("this is a test message")
}
