package logger

import "testing"

func TestNewLogger(t *testing.T) {
	l := NewLogger()

	if l == nil {
		t.Error("Expected logger instance, got nil")
	}

	if _, ok := l.(*Logger); !ok {
		t.Error("Expected Logger instance, got different type")
	}

	t.Log("NewLogger() test passed")
}
