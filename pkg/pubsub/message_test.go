package pubsub

import (
	"famcache/domain/command"
	"testing"
)

func TestNewPubsubMessage(t *testing.T) {
	topic := "topic"
	data := "data"
	to := "to"
	retried := uint(0)

	m := NewPubsubMessage(to, topic, data)

	if m.Topic() != topic {
		t.Errorf("Expected topic to be %s, got %s", topic, m.Topic())
	}

	if m.Data() != data {
		t.Errorf("Expected data to be %s, got %s", data, m.Data())
	}

	if m.Recipient() != to {
		t.Errorf("Expected recipient to be %s, got %s", to, m.Recipient())
	}

	if m.RetryCount() != retried {
		t.Errorf("Expected retry count to be %d, got %d", retried, m.RetryCount())
	}
}

func TestPubsubMessage_IncrementRetryCount(t *testing.T) {
	m := NewPubsubMessage("to", "topic", "data")
	m.IncrementRetryCount()

	if m.RetryCount() != 1 {
		t.Errorf("Expected retry count to be 1, got %d", m.RetryCount())
	}
}

func TestPubsubMessage_ToMessagingCommand(t *testing.T) {
	m := NewPubsubMessage("to", "topic", "data")
	cmd := m.ToMessagingCommand()

	if cmd == nil {
		t.Error("Command is nil")
	}

	if cmd.Type() != command.CommandPublish {
		t.Errorf("Expected type to be PUBLISH, got %s", cmd.Type())
	}

	if cmd.Data() != "data" {
		t.Errorf("Expected data to be data, got %s", cmd.Data())
	}

	if cmd.Topic() != "topic" {
		t.Errorf("Expected topic to be topic, got %s", cmd.Topic())
	}
}
