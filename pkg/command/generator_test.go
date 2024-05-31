package command

import "testing"

func TestSetCommand(t *testing.T) {
	str := SetCommand("1", "key", "value", nil)

	if str != "1 SET key value" {
		t.Errorf("Expected command to be 1 SET key value, got %s", str)
	}
}

func TestSetCommand_WithTTL(t *testing.T) {
	ttl := uint64(3000)
	str := SetCommand("1", "key", "value", &ttl)

	if str != "1 SET key value 3000" {
		t.Errorf("Expected command to be 1 SET key value 3000, got %s", str)
	}
}

func TestGetCommand(t *testing.T) {
	str := GetCommand("1", "key")

	if str != "1 GET key" {
		t.Errorf("Expected command to be 1 GET key, got %s", str)
	}
}

func TestDeleteCommand(t *testing.T) {
	str := DeleteCommand("1", "key")

	if str != "1 DELETE key" {
		t.Errorf("Expected command to be 1 DELETE key, got %s", str)
	}
}

func TestPublishCommand(t *testing.T) {
	str := PublishCommand("1", "topic", "data")

	if str != "1 PUBLISH topic data" {
		t.Errorf("Expected command to be 1 PUBLISH topic data, got %s", str)
	}
}

func TestSubscribeCommand(t *testing.T) {
	str := SubscribeCommand("1", "topic")

	if str != "1 SUBSCRIBE topic" {
		t.Errorf("Expected command to be 1 SUBSCRIBE topic, got %s", str)
	}
}

func TestUnsubscribeCommand(t *testing.T) {
	str := UnsubscribeCommand("1", "topic")

	if str != "1 UNSUBSCRIBE topic" {
		t.Errorf("Expected command to be 1 UNSUBSCRIBE topic, got %s", str)
	}
}
