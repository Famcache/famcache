package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig()

	if c == nil {
		t.Error("Expected config instance, got nil")
	}

	if _, ok := c.(*Config); !ok {
		t.Error("Expected Config instance, got different type")
	}

	t.Log("NewConfig() test passed")
}
